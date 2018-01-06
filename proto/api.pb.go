// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package hostapd is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PingRequest
	PongResponse
	ListRequest
	Client
	ListResponse
*/
package hostapd

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PongResponse struct {
}

func (m *PongResponse) Reset()                    { *m = PongResponse{} }
func (m *PongResponse) String() string            { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()               {}
func (*PongResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Client struct {
	Addr          string   `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
	Flag          []string `protobuf:"bytes,2,rep,name=flag" json:"flag,omitempty"`
	ConnectedTime uint32   `protobuf:"varint,3,opt,name=connected_time,json=connectedTime" json:"connected_time,omitempty"`
	IdleMsec      uint32   `protobuf:"varint,4,opt,name=idle_msec,json=idleMsec" json:"idle_msec,omitempty"`
	RxPackets     uint32   `protobuf:"varint,5,opt,name=rx_packets,json=rxPackets" json:"rx_packets,omitempty"`
	TxPackets     uint32   `protobuf:"varint,6,opt,name=tx_packets,json=txPackets" json:"tx_packets,omitempty"`
	RxBytes       uint32   `protobuf:"varint,7,opt,name=rx_bytes,json=rxBytes" json:"rx_bytes,omitempty"`
	TxBytes       uint32   `protobuf:"varint,8,opt,name=tx_bytes,json=txBytes" json:"tx_bytes,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Client) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Client) GetFlag() []string {
	if m != nil {
		return m.Flag
	}
	return nil
}

func (m *Client) GetConnectedTime() uint32 {
	if m != nil {
		return m.ConnectedTime
	}
	return 0
}

func (m *Client) GetIdleMsec() uint32 {
	if m != nil {
		return m.IdleMsec
	}
	return 0
}

func (m *Client) GetRxPackets() uint32 {
	if m != nil {
		return m.RxPackets
	}
	return 0
}

func (m *Client) GetTxPackets() uint32 {
	if m != nil {
		return m.TxPackets
	}
	return 0
}

func (m *Client) GetRxBytes() uint32 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *Client) GetTxBytes() uint32 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

type ListResponse struct {
	Client []*Client `protobuf:"bytes,1,rep,name=client" json:"client,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ListResponse) GetClient() []*Client {
	if m != nil {
		return m.Client
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "hostapd.PingRequest")
	proto.RegisterType((*PongResponse)(nil), "hostapd.PongResponse")
	proto.RegisterType((*ListRequest)(nil), "hostapd.ListRequest")
	proto.RegisterType((*Client)(nil), "hostapd.Client")
	proto.RegisterType((*ListResponse)(nil), "hostapd.ListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HostapdControl service

type HostapdControlClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error)
	ListClients(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type hostapdControlClient struct {
	cc *grpc.ClientConn
}

func NewHostapdControlClient(cc *grpc.ClientConn) HostapdControlClient {
	return &hostapdControlClient{cc}
}

func (c *hostapdControlClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := grpc.Invoke(ctx, "/hostapd.HostapdControl/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostapdControlClient) ListClients(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/hostapd.HostapdControl/ListClients", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HostapdControl service

type HostapdControlServer interface {
	Ping(context.Context, *PingRequest) (*PongResponse, error)
	ListClients(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterHostapdControlServer(s *grpc.Server, srv HostapdControlServer) {
	s.RegisterService(&_HostapdControl_serviceDesc, srv)
}

func _HostapdControl_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostapdControlServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hostapd.HostapdControl/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostapdControlServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HostapdControl_ListClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostapdControlServer).ListClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hostapd.HostapdControl/ListClients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostapdControlServer).ListClients(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostapdControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hostapd.HostapdControl",
	HandlerType: (*HostapdControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HostapdControl_Ping_Handler,
		},
		{
			MethodName: "ListClients",
			Handler:    _HostapdControl_ListClients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0x87, 0x9b, 0xb7, 0x7d, 0xd3, 0x66, 0xfa, 0x47, 0x58, 0x14, 0x56, 0x45, 0x08, 0x01, 0xb1,
	0xa7, 0x1e, 0x2a, 0xe2, 0xc5, 0x93, 0xbd, 0x78, 0x50, 0x28, 0xc1, 0x7b, 0x48, 0x93, 0xb1, 0x2e,
	0xa6, 0xd9, 0xb8, 0x3b, 0x42, 0xfc, 0x00, 0x7e, 0x57, 0x3f, 0x86, 0xec, 0x6e, 0x9b, 0xee, 0x6d,
	0xe6, 0xf7, 0x64, 0xc2, 0xce, 0x33, 0x10, 0xe5, 0x8d, 0x58, 0x34, 0x4a, 0x92, 0x64, 0xc3, 0x77,
	0xa9, 0x29, 0x6f, 0xca, 0x64, 0x0a, 0xe3, 0xb5, 0xa8, 0xb7, 0x29, 0x7e, 0x7e, 0xa1, 0xa6, 0x64,
	0x06, 0x93, 0xb5, 0x34, 0xad, 0x6e, 0x64, 0xad, 0xd1, 0xe0, 0x67, 0xa1, 0xe9, 0x80, 0x7f, 0x03,
	0x08, 0x57, 0x95, 0xc0, 0x9a, 0x18, 0x83, 0x41, 0x5e, 0x96, 0x8a, 0x07, 0x71, 0x30, 0x8f, 0x52,
	0x5b, 0x9b, 0xec, 0xad, 0xca, 0xb7, 0xfc, 0x5f, 0xdc, 0x37, 0x99, 0xa9, 0xd9, 0x35, 0xcc, 0x0a,
	0x59, 0xd7, 0x58, 0x10, 0x96, 0x19, 0x89, 0x1d, 0xf2, 0x7e, 0x1c, 0xcc, 0xa7, 0xe9, 0xb4, 0x4b,
	0x5f, 0xc5, 0x0e, 0xd9, 0x25, 0x44, 0xa2, 0xac, 0x30, 0xdb, 0x69, 0x2c, 0xf8, 0xc0, 0x7e, 0x31,
	0x32, 0xc1, 0x8b, 0xc6, 0x82, 0x5d, 0x01, 0xa8, 0x36, 0x6b, 0xf2, 0xe2, 0x03, 0x49, 0xf3, 0xff,
	0x96, 0x46, 0xaa, 0x5d, 0xbb, 0xc0, 0x60, 0x3a, 0xe2, 0xd0, 0x61, 0xea, 0xf0, 0x39, 0x8c, 0x54,
	0x9b, 0x6d, 0xbe, 0x09, 0x35, 0x1f, 0x5a, 0x38, 0x54, 0xed, 0xa3, 0x69, 0x0d, 0xa2, 0x03, 0x1a,
	0x39, 0x44, 0x0e, 0x25, 0xf7, 0x30, 0x71, 0x9b, 0x3b, 0x13, 0xec, 0x06, 0xc2, 0xc2, 0x6e, 0xce,
	0x83, 0xb8, 0x3f, 0x1f, 0x2f, 0x4f, 0x16, 0x7b, 0x85, 0x0b, 0x27, 0x24, 0xdd, 0xe3, 0xe5, 0x4f,
	0x00, 0xb3, 0x27, 0x87, 0x56, 0xb2, 0x26, 0x25, 0x2b, 0x76, 0x07, 0x03, 0x23, 0x99, 0x9d, 0x76,
	0x33, 0x9e, 0xf3, 0x8b, 0xb3, 0x63, 0xea, 0xab, 0xef, 0xb1, 0x07, 0x27, 0xdf, 0xfd, 0x5f, 0x7b,
	0xd3, 0xde, 0x49, 0xbc, 0x69, 0xff, 0xb9, 0x49, 0x6f, 0x13, 0xda, 0x4b, 0xdf, 0xfe, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xa7, 0x01, 0x0e, 0xf6, 0x01, 0x00, 0x00,
}
