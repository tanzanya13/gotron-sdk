// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Endpoint struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	NodeId  []byte `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To        *Endpoint `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Version   int32     `protobuf:"varint,3,opt,name=version" json:"version,omitempty"`
	Timestamp int64     `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PingMessage) Reset()                    { *m = PingMessage{} }
func (m *PingMessage) String() string            { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()               {}
func (*PingMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Echo      int32     `protobuf:"varint,2,opt,name=echo" json:"echo,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PongMessage) Reset()                    { *m = PongMessage{} }
func (m *PongMessage) String() string            { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()               {}
func (*PongMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From      *Endpoint `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	TargetId  []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp int64     `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *FindNeighbours) Reset()                    { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string            { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()               {}
func (*FindNeighbours) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From       *Endpoint   `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Neighbours []*Endpoint `protobuf:"bytes,2,rep,name=neighbours" json:"neighbours,omitempty"`
	Timestamp  int64       `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Neighbours) Reset()                    { *m = Neighbours{} }
func (m *Neighbours) String() string            { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()               {}
func (*Neighbours) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcb, 0x6a, 0xf3, 0x30,
	0x10, 0x85, 0xf1, 0x25, 0xf9, 0xfd, 0x4f, 0x42, 0x0b, 0x2a, 0x14, 0x53, 0xba, 0x30, 0x5e, 0x94,
	0x50, 0x88, 0x0d, 0xe9, 0x1b, 0x84, 0x5e, 0xc8, 0xa2, 0x25, 0x78, 0xd9, 0x9d, 0x23, 0xa9, 0xb2,
	0x20, 0xd6, 0x18, 0x49, 0xc9, 0x23, 0x74, 0xdf, 0x37, 0x2e, 0x56, 0xab, 0xf4, 0x42, 0xef, 0x2b,
	0xcd, 0xd1, 0x1c, 0xcd, 0x77, 0x04, 0x03, 0x07, 0x14, 0x35, 0x2f, 0xcf, 0xa5, 0xa1, 0xb8, 0xe5,
	0xba, 0xe8, 0x34, 0x5a, 0x24, 0x89, 0x3b, 0x28, 0xae, 0xf3, 0x25, 0x24, 0x17, 0x8a, 0x75, 0x28,
	0x95, 0x25, 0x29, 0xfc, 0xab, 0x19, 0xd3, 0xdc, 0x98, 0x34, 0xc8, 0x82, 0xc9, 0xb8, 0xf2, 0x92,
	0x10, 0x88, 0x3b, 0xd4, 0x36, 0x0d, 0xb3, 0x60, 0x32, 0xa8, 0x5c, 0x4d, 0x0e, 0x61, 0xa8, 0x90,
	0xf1, 0x05, 0x4b, 0x23, 0x67, 0x7e, 0x56, 0xf9, 0x43, 0x00, 0xa3, 0xa5, 0x54, 0xe2, 0x9a, 0x1b,
	0x53, 0x0b, 0x4e, 0x4e, 0x20, 0xbe, 0xd3, 0xd8, 0xba, 0x91, 0xa3, 0x19, 0x29, 0x3c, 0xba, 0xf0,
	0xdc, 0xca, 0xf5, 0x49, 0x0e, 0xa1, 0x45, 0x47, 0xf8, 0xd8, 0x15, 0x5a, 0xec, 0x13, 0x6e, 0xb9,
	0x36, 0x12, 0x95, 0x83, 0x0e, 0x2a, 0x2f, 0xc9, 0x31, 0xfc, 0xb7, 0xb2, 0xe5, 0xc6, 0xd6, 0x6d,
	0x97, 0xc6, 0x59, 0x30, 0x89, 0xaa, 0x97, 0x8b, 0x5c, 0xc0, 0x68, 0x89, 0xbf, 0x8f, 0x44, 0x20,
	0xe6, 0xb4, 0x41, 0xff, 0xed, 0xbe, 0x7e, 0x0b, 0x8a, 0xde, 0x83, 0x34, 0xec, 0x5d, 0x4a, 0xc5,
	0x6e, 0xb8, 0x14, 0xcd, 0x0a, 0x37, 0xda, 0xfc, 0x98, 0x75, 0x04, 0x89, 0xad, 0xb5, 0xe0, 0x76,
	0xc1, 0x1c, 0x6f, 0x5c, 0xed, 0xf4, 0x37, 0xcc, 0xfb, 0x00, 0xe0, 0x0f, 0xc0, 0x19, 0x80, 0xda,
	0xbd, 0x4a, 0xc3, 0x2c, 0xfa, 0xc4, 0xfd, 0xca, 0xf5, 0x75, 0x90, 0xf9, 0x15, 0xec, 0xa3, 0x16,
	0x85, 0xd5, 0xa8, 0x9e, 0xe6, 0x98, 0x79, 0xe2, 0x17, 0xef, 0xf6, 0x54, 0x48, 0xdb, 0x6c, 0x56,
	0x05, 0xc5, 0xb6, 0xec, 0x1d, 0x1e, 0x54, 0x0a, 0x9c, 0xd2, 0xb5, 0xe4, 0xca, 0x4e, 0xeb, 0x4e,
	0x96, 0xfd, 0xb6, 0xae, 0x86, 0xae, 0x79, 0xf6, 0x18, 0x00, 0x00, 0xff, 0xff, 0x86, 0xff, 0x02,
	0x81, 0xbc, 0x02, 0x00, 0x00,
}