// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectorEvent.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventMessage struct {
	Tenant               string   `protobuf:"bytes,1,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Timeout              string   `protobuf:"bytes,4,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Timestamp            string   `protobuf:"bytes,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	UUID                 string   `protobuf:"bytes,6,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Event                string   `protobuf:"bytes,7,opt,name=Event,proto3" json:"Event,omitempty"`
	Payload              string   `protobuf:"bytes,8,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessage) Reset()         { *m = EventMessage{} }
func (m *EventMessage) String() string { return proto.CompactTextString(m) }
func (*EventMessage) ProtoMessage()    {}
func (*EventMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{0}
}

func (m *EventMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessage.Unmarshal(m, b)
}
func (m *EventMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessage.Marshal(b, m, deterministic)
}
func (m *EventMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessage.Merge(m, src)
}
func (m *EventMessage) XXX_Size() int {
	return xxx_messageInfo_EventMessage.Size(m)
}
func (m *EventMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessage proto.InternalMessageInfo

func (m *EventMessage) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *EventMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EventMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventMessage) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *EventMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *EventMessage) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *EventMessage) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type EventMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=Event,proto3" json:"Event,omitempty"`
	Topic                string   `protobuf:"bytes,3,opt,name=Topic,proto3" json:"Topic,omitempty"`
	IteratorId           string   `protobuf:"bytes,4,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventMessageWait) Reset()         { *m = EventMessageWait{} }
func (m *EventMessageWait) String() string { return proto.CompactTextString(m) }
func (*EventMessageWait) ProtoMessage()    {}
func (*EventMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{1}
}

func (m *EventMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventMessageWait.Unmarshal(m, b)
}
func (m *EventMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventMessageWait.Marshal(b, m, deterministic)
}
func (m *EventMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventMessageWait.Merge(m, src)
}
func (m *EventMessageWait) XXX_Size() int {
	return xxx_messageInfo_EventMessageWait.Size(m)
}
func (m *EventMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_EventMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_EventMessageWait proto.InternalMessageInfo

func (m *EventMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *EventMessageWait) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventMessageWait) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *EventMessageWait) GetIteratorId() string {
	if m != nil {
		return m.IteratorId
	}
	return ""
}

type TopicMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	IteratorId           string   `protobuf:"bytes,3,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicMessageWait) Reset()         { *m = TopicMessageWait{} }
func (m *TopicMessageWait) String() string { return proto.CompactTextString(m) }
func (*TopicMessageWait) ProtoMessage()    {}
func (*TopicMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_a15ac99c650ec24e, []int{2}
}

func (m *TopicMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMessageWait.Unmarshal(m, b)
}
func (m *TopicMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMessageWait.Marshal(b, m, deterministic)
}
func (m *TopicMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessageWait.Merge(m, src)
}
func (m *TopicMessageWait) XXX_Size() int {
	return xxx_messageInfo_TopicMessageWait.Size(m)
}
func (m *TopicMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessageWait proto.InternalMessageInfo

func (m *TopicMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *TopicMessageWait) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *TopicMessageWait) GetIteratorId() string {
	if m != nil {
		return m.IteratorId
	}
	return ""
}

func init() {
	proto.RegisterType((*EventMessage)(nil), "grpc.EventMessage")
	proto.RegisterType((*EventMessageWait)(nil), "grpc.EventMessageWait")
	proto.RegisterType((*TopicMessageWait)(nil), "grpc.TopicMessageWait")
}

func init() {
	proto.RegisterFile("connectorEvent.proto", fileDescriptor_a15ac99c650ec24e)
}

var fileDescriptor_a15ac99c650ec24e = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x6e, 0xf2, 0x30,
	0x14, 0xfd, 0x12, 0xfe, 0x3e, 0x6e, 0x51, 0x1b, 0x19, 0x5a, 0x59, 0x11, 0xaa, 0xaa, 0x4c, 0x9d,
	0x32, 0xb4, 0x62, 0xea, 0x52, 0x41, 0x3b, 0x30, 0x54, 0x42, 0x05, 0xc4, 0xec, 0x26, 0x2e, 0x4a,
	0x21, 0x76, 0x64, 0x0c, 0x12, 0x6b, 0xa7, 0x3e, 0x5d, 0x9f, 0xa9, 0xd8, 0x4e, 0x48, 0x42, 0x61,
	0xe9, 0xe6, 0x73, 0xee, 0xcf, 0xb9, 0xe7, 0xda, 0x86, 0x4e, 0xc0, 0x19, 0xa3, 0x81, 0xe4, 0xe2,
	0x79, 0x43, 0x99, 0xf4, 0x13, 0xc1, 0x25, 0x47, 0xd5, 0xb9, 0x48, 0x02, 0xf7, 0x62, 0x1f, 0x33,
	0xb4, 0xf7, 0x6d, 0x41, 0x4b, 0xa7, 0xbd, 0xd0, 0xd5, 0x8a, 0xcc, 0x29, 0xba, 0x82, 0xfa, 0x84,
	0x32, 0xc2, 0x24, 0xb6, 0x6e, 0xac, 0xdb, 0xe6, 0x6b, 0x8a, 0x50, 0x07, 0x6a, 0x13, 0xbe, 0xa0,
	0x0c, 0xdb, 0x9a, 0x36, 0xc0, 0xb0, 0x49, 0x14, 0xe0, 0x4a, 0xc6, 0xee, 0x00, 0xc2, 0xd0, 0x98,
	0x44, 0x31, 0xe5, 0x6b, 0x89, 0xab, 0x9a, 0xcf, 0x20, 0xea, 0x42, 0x53, 0x1d, 0x57, 0x92, 0xc4,
	0x09, 0xae, 0xe9, 0x58, 0x4e, 0x20, 0x04, 0xd5, 0xe9, 0x74, 0xf8, 0x84, 0xeb, 0x3a, 0xa0, 0xcf,
	0x4a, 0x41, 0xcf, 0x87, 0x1b, 0x46, 0x41, 0x03, 0xa5, 0x30, 0x22, 0xdb, 0x25, 0x27, 0x21, 0xfe,
	0x6f, 0x14, 0x52, 0xe8, 0x7d, 0x5a, 0xe0, 0x14, 0x0d, 0xcd, 0x48, 0x24, 0x91, 0x07, 0xad, 0x19,
	0x17, 0x0b, 0x2a, 0xc6, 0x7c, 0x2d, 0x02, 0x9a, 0x5a, 0x2b, 0x71, 0xb9, 0x90, 0x5d, 0x14, 0x3a,
	0x6e, 0xf0, 0x1a, 0x60, 0x28, 0xa9, 0x20, 0xbb, 0x3d, 0x0e, 0xc3, 0xd4, 0x63, 0x81, 0xf1, 0x96,
	0xe0, 0xe8, 0xc4, 0x3f, 0xcc, 0x60, 0xd4, 0xec, 0xd3, 0x6a, 0x95, 0x43, 0xb5, 0xbb, 0x2f, 0x1b,
	0xce, 0x07, 0xa5, 0x3b, 0x47, 0x3d, 0x70, 0xc6, 0x94, 0x85, 0xa5, 0x9b, 0x45, 0xbe, 0x7a, 0x02,
	0x7e, 0x91, 0x73, 0xcf, 0x52, 0x2e, 0x4e, 0xe4, 0xd6, 0xfb, 0x87, 0x1e, 0xc1, 0x51, 0xb3, 0x96,
	0x1f, 0xc4, 0xef, 0x32, 0x95, 0xe3, 0x1e, 0x69, 0x97, 0x77, 0x28, 0xba, 0xcf, 0x3a, 0x1c, 0x6e,
	0xe4, 0x44, 0x87, 0x07, 0x68, 0x0f, 0x04, 0x25, 0x92, 0x66, 0x0e, 0x8d, 0xa3, 0xe2, 0xa4, 0xee,
	0xa5, 0x01, 0x59, 0xc6, 0xbe, 0xb8, 0xdf, 0x83, 0x6e, 0xc0, 0x63, 0x3f, 0x8c, 0xa4, 0x88, 0xa4,
	0x3f, 0x27, 0x2c, 0x24, 0xcb, 0x77, 0xff, 0x83, 0x6c, 0x88, 0xae, 0xe8, 0xb7, 0xcb, 0x7b, 0x1a,
	0xa9, 0x3f, 0x30, 0xb2, 0xde, 0xea, 0xfa, 0x33, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x1c, 0xd8, 0xde, 0x3b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectorEventClient is the client API for ConnectorEvent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectorEventClient interface {
	SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error)
	WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error)
	CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error)
}

type connectorEventClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorEventClient(cc grpc.ClientConnInterface) ConnectorEventClient {
	return &connectorEventClient{cc}
}

func (c *connectorEventClient) SendEventMessage(ctx context.Context, in *EventMessage, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/SendEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitEventMessage(ctx context.Context, in *EventMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/WaitEventMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) WaitTopicMessage(ctx context.Context, in *TopicMessageWait, opts ...grpc.CallOption) (*EventMessage, error) {
	out := new(EventMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/WaitTopicMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorEventClient) CreateIteratorEvent(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error) {
	out := new(IteratorMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorEvent/CreateIteratorEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorEventServer is the server API for ConnectorEvent service.
type ConnectorEventServer interface {
	SendEventMessage(context.Context, *EventMessage) (*Empty, error)
	WaitEventMessage(context.Context, *EventMessageWait) (*EventMessage, error)
	WaitTopicMessage(context.Context, *TopicMessageWait) (*EventMessage, error)
	CreateIteratorEvent(context.Context, *Empty) (*IteratorMessage, error)
}

// UnimplementedConnectorEventServer can be embedded to have forward compatible implementations.
type UnimplementedConnectorEventServer struct {
}

func (*UnimplementedConnectorEventServer) SendEventMessage(ctx context.Context, req *EventMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEventMessage not implemented")
}
func (*UnimplementedConnectorEventServer) WaitEventMessage(ctx context.Context, req *EventMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitEventMessage not implemented")
}
func (*UnimplementedConnectorEventServer) WaitTopicMessage(ctx context.Context, req *TopicMessageWait) (*EventMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitTopicMessage not implemented")
}
func (*UnimplementedConnectorEventServer) CreateIteratorEvent(ctx context.Context, req *Empty) (*IteratorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIteratorEvent not implemented")
}

func RegisterConnectorEventServer(s *grpc.Server, srv ConnectorEventServer) {
	s.RegisterService(&_ConnectorEvent_serviceDesc, srv)
}

func _ConnectorEvent_SendEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/SendEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).SendEventMessage(ctx, req.(*EventMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitEventMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/WaitEventMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitEventMessage(ctx, req.(*EventMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_WaitTopicMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/WaitTopicMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).WaitTopicMessage(ctx, req.(*TopicMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorEvent_CreateIteratorEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorEvent/CreateIteratorEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorEventServer).CreateIteratorEvent(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorEvent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ConnectorEvent",
	HandlerType: (*ConnectorEventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEventMessage",
			Handler:    _ConnectorEvent_SendEventMessage_Handler,
		},
		{
			MethodName: "WaitEventMessage",
			Handler:    _ConnectorEvent_WaitEventMessage_Handler,
		},
		{
			MethodName: "WaitTopicMessage",
			Handler:    _ConnectorEvent_WaitTopicMessage_Handler,
		},
		{
			MethodName: "CreateIteratorEvent",
			Handler:    _ConnectorEvent_CreateIteratorEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorEvent.proto",
}
