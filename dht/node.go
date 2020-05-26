package dht

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"

	"github.com/decanus/bureka/dht/state"
	"github.com/decanus/bureka/pb"
)

var logger = logging.Logger("dht")

const pastry protocol.ID = "/pastry/1.0/proto"

// ApplicationID represents a unique identifier for the application.
type ApplicationID string

// Application represents a pastry application
type Application interface {
	Deliver(msg pb.Message)
	Forward(msg pb.Message, target state.Peer) bool
	Heartbeat(id state.Peer)
}

// Node is a pastry node.
type Node struct {
	ctx context.Context

	sync.RWMutex

	LeafSet         state.LeafSet
	NeighborhoodSet state.Set
	RoutingTable    state.RoutingTable

	Host host.Host

	applications map[ApplicationID]Application

	writers map[string]chan<- pb.Message
}

// Guarantee that we implement interfaces.
var _ routing.PeerRouting = (*Node)(nil)

func New(ctx context.Context, host host.Host) *Node {
	id, _ := host.ID().MarshalBinary()

	n := &Node{
		ctx:             ctx,
		LeafSet:         state.NewLeafSet(id),
		NeighborhoodSet: make(state.Set, 0),
		applications:    make(map[ApplicationID]Application),
		Host:            host,
	}

	n.Host.SetStreamHandler(pastry, n.streamHandler)

	return n
}

// AddApplication adds an application as a message receiver.
func (n *Node) AddApplication(aid ApplicationID, app Application) {
	n.Lock()
	defer n.Unlock()

	n.applications[aid] = app
}

// RemoveApplication removes an application from the set.
func (n *Node) RemoveApplication(aid ApplicationID) {
	n.Lock()
	defer n.Unlock()

	delete(n.applications, aid)
}

// Send sends a message to the target or the next closest peer.
func (n *Node) Send(ctx context.Context, msg pb.Message) error {
	key := msg.Key

	if bytes.Equal(key, n.ID()) {
		n.deliver(msg) // @todo we may need to do this for more than just message types, like when the routing table is updated.
		return nil
	}

	target := n.route(key)
	if target == nil {
		// no target to be found, delivering to self
		return nil
	}

	forward := n.forward(msg, target)
	if !forward {
		return nil
	}

	err := n.send(msg, target)
	if err != nil {
		return err
	}

	return nil
}

// ID returns a nodes ID, mainly for testing purposes.
func (n *Node) ID() state.Peer {
	id, _ := n.Host.ID().MarshalBinary()
	return id
}

func (n *Node) FindPeer(ctx context.Context, id peer.ID) (peer.AddrInfo, error) {
	if err := id.Validate(); err != nil {
		return peer.AddrInfo{}, err
	}

	logger.Debug("finding peer", "peer", id)

	b, _ := id.MarshalBinary()

	local := n.route(b)
	if local != nil {
		id, err := peer.IDFromBytes(local)
		if err != nil {
			return peer.AddrInfo{}, err
		}

		return n.Host.Peerstore().PeerInfo(id), nil
	}

	return peer.AddrInfo{}, nil
}

// @todo probably want to return error if not found
func (n *Node) route(to state.Peer) state.Peer {
	if n.LeafSet.IsInRange(to) {
		id := n.LeafSet.Closest(to)
		if id != nil {
			return id
		}
	}

	// @todo this is flimsy but will fix later
	id := n.RoutingTable.Route(n.ID(), to)
	if id != nil {
		return id
	}

	return nil
}

// deliver sends the message to all connected applications.
func (n *Node) deliver(msg pb.Message) {
	n.RLock()
	defer n.RUnlock()

	for _, app := range n.applications {
		app.Deliver(msg)
	}
}

// forward asks all applications whether a message should be forwarded to a peer or not.
func (n *Node) forward(msg pb.Message, target state.Peer) bool {
	n.RLock()
	defer n.RUnlock()

	// @todo need to run over this logic
	forward := true
	for _, app := range n.applications {
		f := app.Forward(msg, target)
		if forward {
			forward = f
		}
	}

	return forward
}

func (n *Node) send(msg pb.Message, target state.Peer) error {
	out, ok := n.writers[string(target)]
	if !ok {
		return fmt.Errorf("peer %s not found", string(target))
	}

	out <- msg
	return nil
}

func (n *Node) createWriter(target peer.ID) chan pb.Message {
	n.Lock()
	defer n.Unlock()

	c := make(chan pb.Message) // @todo buffer size
	n.writers[string(target)] = c
	return c
}

func (n *Node) addPeer(id state.Peer) {
	n.Lock()
	defer n.Unlock()

	n.LeafSet.Insert(id)
	n.NeighborhoodSet.Insert(id)
	n.RoutingTable = n.RoutingTable.Insert(n.ID(), id)
}
