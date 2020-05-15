// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pastry.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Message_Type int32

const (
	Message_MESSAGE        Message_Type = 0
	Message_NODE_JOIN      Message_Type = 1
	Message_NODE_ANNOUNCE  Message_Type = 2
	Message_NODE_EXIT      Message_Type = 3
	Message_HEARTBEAT      Message_Type = 4
	Message_REPAIR_REQUEST Message_Type = 5
	Message_STATE_REQUEST  Message_Type = 6
	Message_STATE_RESPONSE Message_Type = 7
)

var Message_Type_name = map[int32]string{
	0: "MESSAGE",
	1: "NODE_JOIN",
	2: "NODE_ANNOUNCE",
	3: "NODE_EXIT",
	4: "HEARTBEAT",
	5: "REPAIR_REQUEST",
	6: "STATE_REQUEST",
	7: "STATE_RESPONSE",
}

var Message_Type_value = map[string]int32{
	"MESSAGE":        0,
	"NODE_JOIN":      1,
	"NODE_ANNOUNCE":  2,
	"NODE_EXIT":      3,
	"HEARTBEAT":      4,
	"REPAIR_REQUEST": 5,
	"STATE_REQUEST":  6,
	"STATE_RESPONSE": 7,
}

func (x Message_Type) String() string {
	return proto.EnumName(Message_Type_name, int32(x))
}

func (Message_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_85373c68ed8b9855, []int{0, 0}
}

type Message struct {
	Type                 Message_Type `protobuf:"varint,1,opt,name=type,proto3,enum=pb.Message_Type" json:"type,omitempty"`
	Sender               string       `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Key                  string       `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Hops                 uint64       `protobuf:"varint,4,opt,name=hops,proto3" json:"hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_85373c68ed8b9855, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetType() Message_Type {
	if m != nil {
		return m.Type
	}
	return Message_MESSAGE
}

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Message) GetHops() uint64 {
	if m != nil {
		return m.Hops
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Message_Type", Message_Type_name, Message_Type_value)
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("pastry.proto", fileDescriptor_85373c68ed8b9855) }

var fileDescriptor_85373c68ed8b9855 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x4d, 0x1b, 0x5b, 0x76, 0x74, 0x97, 0x38, 0x07, 0xe9, 0xb1, 0x2c, 0x1e, 0x7a, 0xea,
	0x41, 0x9f, 0x20, 0xea, 0xa0, 0x15, 0x36, 0x5d, 0x93, 0x2c, 0x78, 0x5b, 0xb6, 0x18, 0x14, 0x04,
	0x1b, 0xda, 0x5e, 0xfa, 0x12, 0x3e, 0xb2, 0x48, 0xc2, 0xba, 0xb7, 0xf9, 0xbe, 0xff, 0xe7, 0x3f,
	0x0c, 0x5c, 0xfa, 0xc3, 0x38, 0x0d, 0x73, 0xed, 0x87, 0x7e, 0xea, 0x31, 0xf1, 0xdd, 0xfa, 0x97,
	0x41, 0xbe, 0x71, 0xe3, 0x78, 0xf8, 0x70, 0x78, 0x03, 0x7c, 0x9a, 0xbd, 0x2b, 0x58, 0xc9, 0xaa,
	0xd5, 0xad, 0xa8, 0x7d, 0x57, 0x1f, 0xa3, 0xda, 0xce, 0xde, 0xe9, 0x98, 0xe2, 0x35, 0x64, 0xa3,
	0xfb, 0x7e, 0x77, 0x43, 0x91, 0x94, 0xac, 0x5a, 0xe8, 0x23, 0xa1, 0x80, 0xf4, 0xcb, 0xcd, 0x45,
	0x1a, 0x65, 0x38, 0x11, 0x81, 0x7f, 0xf6, 0x7e, 0x2c, 0x78, 0xc9, 0x2a, 0xae, 0xe3, 0xbd, 0xfe,
	0x61, 0xc0, 0xc3, 0x18, 0x5e, 0x40, 0xbe, 0x21, 0x63, 0xe4, 0x13, 0x89, 0x33, 0x5c, 0xc2, 0x42,
	0xb5, 0x8f, 0xb4, 0x7f, 0x69, 0x1b, 0x25, 0x18, 0x5e, 0xc1, 0x32, 0xa2, 0x54, 0xaa, 0xdd, 0xa9,
	0x07, 0x12, 0xc9, 0xa9, 0x41, 0x6f, 0x8d, 0x15, 0x69, 0xc0, 0x67, 0x92, 0xda, 0xde, 0x93, 0xb4,
	0x82, 0x23, 0xc2, 0x4a, 0xd3, 0x56, 0x36, 0x7a, 0xaf, 0xe9, 0x75, 0x47, 0xc6, 0x8a, 0xf3, 0x30,
	0x62, 0xac, 0xb4, 0x74, 0x52, 0x59, 0xa8, 0xfd, 0x2b, 0xb3, 0x6d, 0x95, 0x21, 0x91, 0x77, 0x59,
	0xfc, 0xc5, 0xdd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xcc, 0x6f, 0x0e, 0x1b, 0x01, 0x00,
	0x00,
}