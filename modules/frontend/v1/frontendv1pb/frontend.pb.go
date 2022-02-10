// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/frontend/v1/frontendv1pb/frontend.proto

// Protobuf package should not be changed when moving around go packages
// in order to not break backward compatibility.

package frontendv1pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	stats "github.com/grafana/tempo/modules/querier/stats"
	httpgrpc "github.com/weaveworks/common/httpgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Type int32

const (
	Type_HTTP_REQUEST Type = 0
	Type_GET_ID       Type = 1
)

var Type_name = map[int32]string{
	0: "HTTP_REQUEST",
	1: "GET_ID",
}

var Type_value = map[string]int32{
	"HTTP_REQUEST": 0,
	"GET_ID":       1,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8e6c94795ed772cd, []int{0}
}

type FrontendToClient struct {
	HttpRequest *httpgrpc.HTTPRequest `protobuf:"bytes,1,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	Type        Type                  `protobuf:"varint,2,opt,name=type,proto3,enum=frontend.Type" json:"type,omitempty"`
	// Whether query statistics tracking should be enabled. The response will include
	// statistics only when this option is enabled.
	StatsEnabled bool `protobuf:"varint,3,opt,name=statsEnabled,proto3" json:"statsEnabled,omitempty"`
}

func (m *FrontendToClient) Reset()         { *m = FrontendToClient{} }
func (m *FrontendToClient) String() string { return proto.CompactTextString(m) }
func (*FrontendToClient) ProtoMessage()    {}
func (*FrontendToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c94795ed772cd, []int{0}
}
func (m *FrontendToClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FrontendToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FrontendToClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FrontendToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrontendToClient.Merge(m, src)
}
func (m *FrontendToClient) XXX_Size() int {
	return m.Size()
}
func (m *FrontendToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_FrontendToClient.DiscardUnknown(m)
}

var xxx_messageInfo_FrontendToClient proto.InternalMessageInfo

func (m *FrontendToClient) GetHttpRequest() *httpgrpc.HTTPRequest {
	if m != nil {
		return m.HttpRequest
	}
	return nil
}

func (m *FrontendToClient) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_HTTP_REQUEST
}

func (m *FrontendToClient) GetStatsEnabled() bool {
	if m != nil {
		return m.StatsEnabled
	}
	return false
}

type ClientToFrontend struct {
	HttpResponse *httpgrpc.HTTPResponse `protobuf:"bytes,1,opt,name=httpResponse,proto3" json:"httpResponse,omitempty"`
	ClientID     string                 `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Stats        *stats.Stats           `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *ClientToFrontend) Reset()         { *m = ClientToFrontend{} }
func (m *ClientToFrontend) String() string { return proto.CompactTextString(m) }
func (*ClientToFrontend) ProtoMessage()    {}
func (*ClientToFrontend) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c94795ed772cd, []int{1}
}
func (m *ClientToFrontend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientToFrontend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientToFrontend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientToFrontend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientToFrontend.Merge(m, src)
}
func (m *ClientToFrontend) XXX_Size() int {
	return m.Size()
}
func (m *ClientToFrontend) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientToFrontend.DiscardUnknown(m)
}

var xxx_messageInfo_ClientToFrontend proto.InternalMessageInfo

func (m *ClientToFrontend) GetHttpResponse() *httpgrpc.HTTPResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *ClientToFrontend) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *ClientToFrontend) GetStats() *stats.Stats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type NotifyClientShutdownRequest struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (m *NotifyClientShutdownRequest) Reset()         { *m = NotifyClientShutdownRequest{} }
func (m *NotifyClientShutdownRequest) String() string { return proto.CompactTextString(m) }
func (*NotifyClientShutdownRequest) ProtoMessage()    {}
func (*NotifyClientShutdownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c94795ed772cd, []int{2}
}
func (m *NotifyClientShutdownRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyClientShutdownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyClientShutdownRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyClientShutdownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyClientShutdownRequest.Merge(m, src)
}
func (m *NotifyClientShutdownRequest) XXX_Size() int {
	return m.Size()
}
func (m *NotifyClientShutdownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyClientShutdownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyClientShutdownRequest proto.InternalMessageInfo

func (m *NotifyClientShutdownRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

type NotifyClientShutdownResponse struct {
}

func (m *NotifyClientShutdownResponse) Reset()         { *m = NotifyClientShutdownResponse{} }
func (m *NotifyClientShutdownResponse) String() string { return proto.CompactTextString(m) }
func (*NotifyClientShutdownResponse) ProtoMessage()    {}
func (*NotifyClientShutdownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e6c94795ed772cd, []int{3}
}
func (m *NotifyClientShutdownResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NotifyClientShutdownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotifyClientShutdownResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NotifyClientShutdownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyClientShutdownResponse.Merge(m, src)
}
func (m *NotifyClientShutdownResponse) XXX_Size() int {
	return m.Size()
}
func (m *NotifyClientShutdownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyClientShutdownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyClientShutdownResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("frontend.Type", Type_name, Type_value)
	proto.RegisterType((*FrontendToClient)(nil), "frontend.FrontendToClient")
	proto.RegisterType((*ClientToFrontend)(nil), "frontend.ClientToFrontend")
	proto.RegisterType((*NotifyClientShutdownRequest)(nil), "frontend.NotifyClientShutdownRequest")
	proto.RegisterType((*NotifyClientShutdownResponse)(nil), "frontend.NotifyClientShutdownResponse")
}

func init() {
	proto.RegisterFile("modules/frontend/v1/frontendv1pb/frontend.proto", fileDescriptor_8e6c94795ed772cd)
}

var fileDescriptor_8e6c94795ed772cd = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xbd, 0x50, 0x4a, 0x98, 0x58, 0x95, 0xb5, 0x02, 0x14, 0x19, 0x64, 0x05, 0x0b, 0xaa,
	0x08, 0x09, 0x2f, 0x0d, 0x48, 0x08, 0x8e, 0xd0, 0x50, 0x7a, 0x41, 0x65, 0x63, 0x2e, 0x5c, 0xaa,
	0xd8, 0xd9, 0xfc, 0x11, 0x89, 0xc7, 0xf5, 0xae, 0x13, 0xe5, 0x25, 0x10, 0x12, 0x6f, 0xc3, 0x13,
	0x70, 0xec, 0x91, 0x23, 0x4a, 0x5e, 0x04, 0x65, 0x37, 0x76, 0x9d, 0xa8, 0x82, 0xcb, 0x6a, 0xc6,
	0xf3, 0xcd, 0xec, 0x6f, 0x3e, 0xdb, 0xc0, 0xa6, 0xd8, 0xcf, 0x27, 0x42, 0xb2, 0x41, 0x86, 0x89,
	0x12, 0x49, 0x9f, 0xcd, 0x8e, 0xca, 0x78, 0x76, 0x94, 0x46, 0x65, 0x12, 0xa4, 0x19, 0x2a, 0xa4,
	0xb5, 0x22, 0x77, 0x9f, 0x0d, 0xc7, 0x6a, 0x94, 0x47, 0x41, 0x8c, 0x53, 0x36, 0xc4, 0x21, 0x32,
	0x2d, 0x88, 0xf2, 0x81, 0xce, 0x74, 0xa2, 0x23, 0xd3, 0xe8, 0xbe, 0xac, 0xc8, 0xe7, 0xa2, 0x37,
	0x13, 0x73, 0xcc, 0xbe, 0x4a, 0x16, 0xe3, 0x74, 0x8a, 0x09, 0x1b, 0x29, 0x95, 0x0e, 0xb3, 0x34,
	0x2e, 0x83, 0x4d, 0xd7, 0xa3, 0x82, 0xef, 0x22, 0x17, 0xd9, 0x58, 0x64, 0x4c, 0xaa, 0x9e, 0x92,
	0xe6, 0x34, 0x12, 0xff, 0x07, 0x01, 0xe7, 0xfd, 0x06, 0x2a, 0xc4, 0x77, 0x93, 0xb1, 0x48, 0x14,
	0x7d, 0x05, 0xf5, 0xf5, 0x24, 0x2e, 0x2e, 0x72, 0x21, 0x55, 0x83, 0x34, 0x49, 0xab, 0xde, 0xbe,
	0x17, 0x94, 0xd3, 0x3f, 0x84, 0xe1, 0xd9, 0xa6, 0xc8, 0xab, 0x4a, 0xea, 0xc3, 0x9e, 0x5a, 0xa4,
	0xa2, 0x71, 0xa3, 0x49, 0x5a, 0x07, 0xed, 0x83, 0xa0, 0x5c, 0x3f, 0x5c, 0xa4, 0x82, 0xeb, 0x1a,
	0xf5, 0xc1, 0xd6, 0x00, 0x9d, 0xa4, 0x17, 0x4d, 0x44, 0xbf, 0x71, 0xb3, 0x49, 0x5a, 0x35, 0xbe,
	0xf5, 0xcc, 0xff, 0x46, 0xc0, 0x31, 0x2c, 0x21, 0x16, 0x74, 0xf4, 0x0d, 0xd8, 0xe6, 0x2e, 0x99,
	0x62, 0x22, 0xc5, 0x06, 0xeb, 0xfe, 0x2e, 0x96, 0xa9, 0xf2, 0x2d, 0x2d, 0x75, 0xa1, 0x16, 0xeb,
	0x79, 0xa7, 0xc7, 0x1a, 0xee, 0x0e, 0x2f, 0x73, 0xea, 0xc3, 0x2d, 0x7d, 0xb9, 0x26, 0xa9, 0xb7,
	0xed, 0xc0, 0xf8, 0xd3, 0x5d, 0x9f, 0xdc, 0x94, 0xfc, 0xd7, 0xf0, 0xe0, 0x23, 0xaa, 0xf1, 0x60,
	0x61, 0xa8, 0xba, 0xa3, 0x5c, 0xf5, 0x71, 0x9e, 0x14, 0x7b, 0x57, 0xc7, 0x93, 0xed, 0xf1, 0xbe,
	0x07, 0x0f, 0xaf, 0x6f, 0x35, 0x68, 0x4f, 0x1f, 0xc3, 0xde, 0xda, 0x1d, 0xea, 0x80, 0xbd, 0x5e,
	0xe0, 0x9c, 0x77, 0x3e, 0x7d, 0xee, 0x74, 0x43, 0xc7, 0xa2, 0x00, 0xfb, 0x27, 0x9d, 0xf0, 0xfc,
	0xf4, 0xd8, 0x21, 0xed, 0x9f, 0x04, 0x6a, 0xa5, 0x13, 0x27, 0x70, 0xfb, 0x2c, 0xc3, 0x58, 0x48,
	0x49, 0xdd, 0x2b, 0x8f, 0x77, 0x0d, 0x73, 0x2b, 0xb5, 0xdd, 0x57, 0xec, 0x5b, 0x2d, 0xf2, 0x9c,
	0x50, 0x01, 0x77, 0xaf, 0x63, 0xa3, 0x4f, 0xae, 0x3a, 0xff, 0xb1, 0xb6, 0x7b, 0xf8, 0x3f, 0x99,
	0x59, 0xf1, 0xed, 0xe1, 0xaf, 0xa5, 0x47, 0x2e, 0x97, 0x1e, 0xf9, 0xb3, 0xf4, 0xc8, 0xf7, 0x95,
	0x67, 0x5d, 0xae, 0x3c, 0xeb, 0xf7, 0xca, 0xb3, 0xbe, 0xd8, 0xd5, 0xbf, 0x25, 0xda, 0xd7, 0xdf,
	0xe4, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0x88, 0x53, 0x7e, 0x58, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FrontendClient interface {
	// After calling this method, client enters a loop, in which it waits for
	// a "FrontendToClient" message and replies with single "ClientToFrontend" message.
	Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error)
	// The client notifies the query-frontend that it started a graceful shutdown.
	NotifyClientShutdown(ctx context.Context, in *NotifyClientShutdownRequest, opts ...grpc.CallOption) (*NotifyClientShutdownResponse, error)
}

type frontendClient struct {
	cc *grpc.ClientConn
}

func NewFrontendClient(cc *grpc.ClientConn) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Process(ctx context.Context, opts ...grpc.CallOption) (Frontend_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontend_serviceDesc.Streams[0], "/frontend.Frontend/Process", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontendProcessClient{stream}
	return x, nil
}

type Frontend_ProcessClient interface {
	Send(*ClientToFrontend) error
	Recv() (*FrontendToClient, error)
	grpc.ClientStream
}

type frontendProcessClient struct {
	grpc.ClientStream
}

func (x *frontendProcessClient) Send(m *ClientToFrontend) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontendProcessClient) Recv() (*FrontendToClient, error) {
	m := new(FrontendToClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *frontendClient) NotifyClientShutdown(ctx context.Context, in *NotifyClientShutdownRequest, opts ...grpc.CallOption) (*NotifyClientShutdownResponse, error) {
	out := new(NotifyClientShutdownResponse)
	err := c.cc.Invoke(ctx, "/frontend.Frontend/NotifyClientShutdown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServer is the server API for Frontend service.
type FrontendServer interface {
	// After calling this method, client enters a loop, in which it waits for
	// a "FrontendToClient" message and replies with single "ClientToFrontend" message.
	Process(Frontend_ProcessServer) error
	// The client notifies the query-frontend that it started a graceful shutdown.
	NotifyClientShutdown(context.Context, *NotifyClientShutdownRequest) (*NotifyClientShutdownResponse, error)
}

// UnimplementedFrontendServer can be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (*UnimplementedFrontendServer) Process(srv Frontend_ProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method Process not implemented")
}
func (*UnimplementedFrontendServer) NotifyClientShutdown(ctx context.Context, req *NotifyClientShutdownRequest) (*NotifyClientShutdownResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyClientShutdown not implemented")
}

func RegisterFrontendServer(s *grpc.Server, srv FrontendServer) {
	s.RegisterService(&_Frontend_serviceDesc, srv)
}

func _Frontend_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontendServer).Process(&frontendProcessServer{stream})
}

type Frontend_ProcessServer interface {
	Send(*FrontendToClient) error
	Recv() (*ClientToFrontend, error)
	grpc.ServerStream
}

type frontendProcessServer struct {
	grpc.ServerStream
}

func (x *frontendProcessServer) Send(m *FrontendToClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontendProcessServer) Recv() (*ClientToFrontend, error) {
	m := new(ClientToFrontend)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Frontend_NotifyClientShutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyClientShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).NotifyClientShutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/frontend.Frontend/NotifyClientShutdown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).NotifyClientShutdown(ctx, req.(*NotifyClientShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Frontend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "frontend.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyClientShutdown",
			Handler:    _Frontend_NotifyClientShutdown_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _Frontend_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "modules/frontend/v1/frontendv1pb/frontend.proto",
}

func (m *FrontendToClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FrontendToClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FrontendToClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.StatsEnabled {
		i--
		if m.StatsEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Type != 0 {
		i = encodeVarintFrontend(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if m.HttpRequest != nil {
		{
			size, err := m.HttpRequest.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClientToFrontend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientToFrontend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientToFrontend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintFrontend(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0x12
	}
	if m.HttpResponse != nil {
		{
			size, err := m.HttpResponse.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFrontend(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyClientShutdownRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyClientShutdownRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyClientShutdownRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintFrontend(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NotifyClientShutdownResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotifyClientShutdownResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotifyClientShutdownResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintFrontend(dAtA []byte, offset int, v uint64) int {
	offset -= sovFrontend(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FrontendToClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpRequest != nil {
		l = m.HttpRequest.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovFrontend(uint64(m.Type))
	}
	if m.StatsEnabled {
		n += 2
	}
	return n
}

func (m *ClientToFrontend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HttpResponse != nil {
		l = m.HttpResponse.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovFrontend(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *NotifyClientShutdownRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovFrontend(uint64(l))
	}
	return n
}

func (m *NotifyClientShutdownResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovFrontend(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFrontend(x uint64) (n int) {
	return sovFrontend(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FrontendToClient) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FrontendToClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FrontendToClient: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpRequest == nil {
				m.HttpRequest = &httpgrpc.HTTPRequest{}
			}
			if err := m.HttpRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.StatsEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClientToFrontend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientToFrontend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientToFrontend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpResponse", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HttpResponse == nil {
				m.HttpResponse = &httpgrpc.HTTPResponse{}
			}
			if err := m.HttpResponse.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &stats.Stats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyClientShutdownRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyClientShutdownRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyClientShutdownRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFrontend
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFrontend
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NotifyClientShutdownResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NotifyClientShutdownResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotifyClientShutdownResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipFrontend(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFrontend
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFrontend(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFrontend
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFrontend
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFrontend
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFrontend
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFrontend
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFrontend        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFrontend          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFrontend = fmt.Errorf("proto: unexpected end of group")
)