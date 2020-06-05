package node

import (
	"context"

	"github.com/gogo/protobuf/proto"

	"github.com/decanus/bureka/pb"
)

type handlerFunc func(ctx context.Context, message *pb.Message) *pb.Message

func (n *Node) handler(t pb.Message_Type) handlerFunc {
	switch t {
	case pb.Message_MESSAGE:
		return n.onMessage
	case pb.Message_NODE_JOIN:
		return n.onNodeJoin
	case pb.Message_NODE_ANNOUNCE:
		return n.onNodeAnnounce
	case pb.Message_NODE_EXIT:
		return n.onNodeExit
	case pb.Message_HEARTBEAT:
		return n.onHeartbeat
	case pb.Message_REPAIR_REQUEST:
		return n.onRepairRequest
	case pb.Message_STATE_REQUEST:
		return n.onStateRequest
	case pb.Message_STATE_RESPONSE:
		return n.onStateRequest
	}

	return nil
}

func (n *Node) onMessage(ctx context.Context, message *pb.Message) *pb.Message {
	err := n.Send(ctx, message)
	if err != nil {
		// @todo
	}

	return nil
}

func (n *Node) onNodeJoin(ctx context.Context, message *pb.Message) *pb.Message {
	// @TODO THIS IS QUESTIONABLE CAUSE IT MAY BE HANDLED THROUGH ANOTHER PATH ALREADY
	return nil
}

func (n *Node) onNodeAnnounce(ctx context.Context, message *pb.Message) *pb.Message {
	return nil
}

func (n *Node) onNodeExit(ctx context.Context, message *pb.Message) *pb.Message {
	n.dht.RemovePeer(message.Sender)
	return nil
}

func (n *Node) onHeartbeat(_ context.Context, message *pb.Message) *pb.Message {
	n.dht.Heartbeat(message.Sender)
	return nil
}

func (n *Node) onRepairRequest(ctx context.Context, message *pb.Message) *pb.Message {
	return nil
}

func (n *Node) onStateRequest(ctx context.Context, message *pb.Message) *pb.Message {
	return nil
}

func (n *Node) onStateResponse(ctx context.Context, message *pb.Message) *pb.Message {
	req := &pb.State{}
	err := proto.Unmarshal(message.Data, req)
	if err != nil {
		// @todo
		return nil
	}

	n.dht.Lock()
	defer n.dht.Unlock()

	for _, peer := range req.RoutingTable {
		n.dht.RoutingTable = n.dht.RoutingTable.Insert(n.dht.ID, peer)
	}

	for _, peer := range req.Neighborhood {
		n.dht.NeighborhoodSet = n.dht.NeighborhoodSet.Insert(peer)
	}

	for _, peer := range req.Leafset {
		n.dht.LeafSet.Insert(peer)
	}

	return nil
}
