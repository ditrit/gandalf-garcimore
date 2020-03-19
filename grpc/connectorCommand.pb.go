// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectorCommand.proto

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

type CommandMessage struct {
	SourceAggregator      string   `protobuf:"bytes,1,opt,name=SourceAggregator,proto3" json:"SourceAggregator,omitempty"`
	SourceConnector       string   `protobuf:"bytes,2,opt,name=SourceConnector,proto3" json:"SourceConnector,omitempty"`
	SourceWorker          string   `protobuf:"bytes,3,opt,name=SourceWorker,proto3" json:"SourceWorker,omitempty"`
	DestinationAggregator string   `protobuf:"bytes,4,opt,name=DestinationAggregator,proto3" json:"DestinationAggregator,omitempty"`
	DestinationConnector  string   `protobuf:"bytes,5,opt,name=DestinationConnector,proto3" json:"DestinationConnector,omitempty"`
	DestinationWorker     string   `protobuf:"bytes,6,opt,name=DestinationWorker,proto3" json:"DestinationWorker,omitempty"`
	Tenant                string   `protobuf:"bytes,7,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	Token                 string   `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
	Context               string   `protobuf:"bytes,9,opt,name=Context,proto3" json:"Context,omitempty"`
	Timeout               string   `protobuf:"bytes,10,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Timestamp             string   `protobuf:"bytes,11,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Major                 string   `protobuf:"bytes,12,opt,name=Major,proto3" json:"Major,omitempty"`
	Minor                 string   `protobuf:"bytes,13,opt,name=Minor,proto3" json:"Minor,omitempty"`
	UUID                  string   `protobuf:"bytes,14,opt,name=UUID,proto3" json:"UUID,omitempty"`
	ConnectorType         string   `protobuf:"bytes,15,opt,name=ConnectorType,proto3" json:"ConnectorType,omitempty"`
	CommandType           string   `protobuf:"bytes,16,opt,name=CommandType,proto3" json:"CommandType,omitempty"`
	Command               string   `protobuf:"bytes,17,opt,name=Command,proto3" json:"Command,omitempty"`
	Payload               string   `protobuf:"bytes,18,opt,name=Payload,proto3" json:"Payload,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *CommandMessage) Reset()         { *m = CommandMessage{} }
func (m *CommandMessage) String() string { return proto.CompactTextString(m) }
func (*CommandMessage) ProtoMessage()    {}
func (*CommandMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ea62e0fbc109bd, []int{0}
}

func (m *CommandMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandMessage.Unmarshal(m, b)
}
func (m *CommandMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandMessage.Marshal(b, m, deterministic)
}
func (m *CommandMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMessage.Merge(m, src)
}
func (m *CommandMessage) XXX_Size() int {
	return xxx_messageInfo_CommandMessage.Size(m)
}
func (m *CommandMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMessage proto.InternalMessageInfo

func (m *CommandMessage) GetSourceAggregator() string {
	if m != nil {
		return m.SourceAggregator
	}
	return ""
}

func (m *CommandMessage) GetSourceConnector() string {
	if m != nil {
		return m.SourceConnector
	}
	return ""
}

func (m *CommandMessage) GetSourceWorker() string {
	if m != nil {
		return m.SourceWorker
	}
	return ""
}

func (m *CommandMessage) GetDestinationAggregator() string {
	if m != nil {
		return m.DestinationAggregator
	}
	return ""
}

func (m *CommandMessage) GetDestinationConnector() string {
	if m != nil {
		return m.DestinationConnector
	}
	return ""
}

func (m *CommandMessage) GetDestinationWorker() string {
	if m != nil {
		return m.DestinationWorker
	}
	return ""
}

func (m *CommandMessage) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *CommandMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CommandMessage) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *CommandMessage) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *CommandMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *CommandMessage) GetMajor() string {
	if m != nil {
		return m.Major
	}
	return ""
}

func (m *CommandMessage) GetMinor() string {
	if m != nil {
		return m.Minor
	}
	return ""
}

func (m *CommandMessage) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *CommandMessage) GetConnectorType() string {
	if m != nil {
		return m.ConnectorType
	}
	return ""
}

func (m *CommandMessage) GetCommandType() string {
	if m != nil {
		return m.CommandType
	}
	return ""
}

func (m *CommandMessage) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandMessage) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type CommandMessageUUID struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandMessageUUID) Reset()         { *m = CommandMessageUUID{} }
func (m *CommandMessageUUID) String() string { return proto.CompactTextString(m) }
func (*CommandMessageUUID) ProtoMessage()    {}
func (*CommandMessageUUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ea62e0fbc109bd, []int{1}
}

func (m *CommandMessageUUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandMessageUUID.Unmarshal(m, b)
}
func (m *CommandMessageUUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandMessageUUID.Marshal(b, m, deterministic)
}
func (m *CommandMessageUUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMessageUUID.Merge(m, src)
}
func (m *CommandMessageUUID) XXX_Size() int {
	return xxx_messageInfo_CommandMessageUUID.Size(m)
}
func (m *CommandMessageUUID) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMessageUUID.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMessageUUID proto.InternalMessageInfo

func (m *CommandMessageUUID) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type CommandMessageWait struct {
	WorkerSource         string   `protobuf:"bytes,1,opt,name=WorkerSource,proto3" json:"WorkerSource,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	IteratorId           string   `protobuf:"bytes,3,opt,name=IteratorId,proto3" json:"IteratorId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandMessageWait) Reset()         { *m = CommandMessageWait{} }
func (m *CommandMessageWait) String() string { return proto.CompactTextString(m) }
func (*CommandMessageWait) ProtoMessage()    {}
func (*CommandMessageWait) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ea62e0fbc109bd, []int{2}
}

func (m *CommandMessageWait) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandMessageWait.Unmarshal(m, b)
}
func (m *CommandMessageWait) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandMessageWait.Marshal(b, m, deterministic)
}
func (m *CommandMessageWait) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMessageWait.Merge(m, src)
}
func (m *CommandMessageWait) XXX_Size() int {
	return xxx_messageInfo_CommandMessageWait.Size(m)
}
func (m *CommandMessageWait) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMessageWait.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMessageWait proto.InternalMessageInfo

func (m *CommandMessageWait) GetWorkerSource() string {
	if m != nil {
		return m.WorkerSource
	}
	return ""
}

func (m *CommandMessageWait) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *CommandMessageWait) GetIteratorId() string {
	if m != nil {
		return m.IteratorId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommandMessage)(nil), "grpc.CommandMessage")
	proto.RegisterType((*CommandMessageUUID)(nil), "grpc.CommandMessageUUID")
	proto.RegisterType((*CommandMessageWait)(nil), "grpc.CommandMessageWait")
}

func init() {
	proto.RegisterFile("connectorCommand.proto", fileDescriptor_84ea62e0fbc109bd)
}

var fileDescriptor_84ea62e0fbc109bd = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x09, 0x74, 0x1b, 0x7d, 0xdd, 0xd6, 0xee, 0xa9, 0x9d, 0xac, 0x6a, 0x42, 0x53, 0xc5,
	0xa1, 0x42, 0x28, 0x87, 0x81, 0xc4, 0x89, 0x03, 0xeb, 0x40, 0xea, 0x61, 0x52, 0xc5, 0x3a, 0x76,
	0x36, 0x89, 0x89, 0xb2, 0x35, 0x76, 0xe4, 0xba, 0x88, 0xfe, 0xb3, 0x48, 0xfc, 0x27, 0xc4, 0xcf,
	0x4e, 0x9b, 0x64, 0xb9, 0xe5, 0xfb, 0x7d, 0xdf, 0xb3, 0x9f, 0x93, 0xe7, 0xc0, 0x79, 0xa4, 0xa4,
	0x14, 0x91, 0x51, 0x7a, 0xa6, 0xb2, 0x8c, 0xcb, 0x38, 0xcc, 0xb5, 0x32, 0x0a, 0x3b, 0x89, 0xce,
	0xa3, 0x71, 0x7f, 0xe7, 0x3a, 0x3c, 0xf9, 0xdb, 0x81, 0x53, 0x1f, 0xbc, 0x15, 0xeb, 0x35, 0x4f,
	0x04, 0xbe, 0x83, 0xc1, 0x9d, 0xda, 0xe8, 0x48, 0x7c, 0x49, 0x12, 0x2d, 0x12, 0x5e, 0x84, 0x59,
	0x70, 0x19, 0x4c, 0xbb, 0xdf, 0x9f, 0x71, 0x9c, 0x42, 0xdf, 0xb1, 0x59, 0xb9, 0x2e, 0x7b, 0x49,
	0xd1, 0x26, 0xc6, 0x09, 0x1c, 0x3b, 0xf4, 0xa0, 0xf4, 0x93, 0xd0, 0xec, 0x15, 0xc5, 0x6a, 0x0c,
	0x3f, 0xc2, 0xe8, 0x46, 0xac, 0x4d, 0x2a, 0xb9, 0x49, 0x95, 0xac, 0x6c, 0xdf, 0xa1, 0x70, 0xbb,
	0x89, 0x57, 0x30, 0xac, 0x18, 0xfb, 0x46, 0x0e, 0xa8, 0xa8, 0xd5, 0xc3, 0xf7, 0x70, 0x56, 0xe1,
	0xbe, 0xa5, 0x43, 0x2a, 0x78, 0x6e, 0xe0, 0x39, 0x1c, 0x2e, 0x85, 0xe4, 0xd2, 0xb0, 0x23, 0x8a,
	0x78, 0x85, 0x43, 0x38, 0x58, 0xaa, 0x27, 0x21, 0xd9, 0x6b, 0xc2, 0x4e, 0x20, 0x83, 0xa3, 0x62,
	0x23, 0x23, 0xfe, 0x18, 0xd6, 0x25, 0x5e, 0x4a, 0xeb, 0x2c, 0xd3, 0x4c, 0xa8, 0x8d, 0x61, 0xe0,
	0x1c, 0x2f, 0xf1, 0x02, 0xba, 0xf6, 0x71, 0x6d, 0x78, 0x96, 0xb3, 0x1e, 0x79, 0x7b, 0x60, 0xf7,
	0xb9, 0xe5, 0x8f, 0xc5, 0x91, 0x8e, 0xdd, 0x3e, 0x24, 0x88, 0xa6, 0xb2, 0xa0, 0x27, 0x9e, 0x5a,
	0x81, 0x08, 0x9d, 0xfb, 0xfb, 0xf9, 0x0d, 0x3b, 0x25, 0x48, 0xcf, 0xf8, 0x16, 0x4e, 0x76, 0x47,
	0x5f, 0x6e, 0x73, 0xc1, 0xfa, 0x64, 0xd6, 0x21, 0x5e, 0x42, 0xcf, 0x4f, 0x02, 0x65, 0x06, 0x94,
	0xa9, 0x22, 0x77, 0x32, 0x92, 0xec, 0xac, 0x3c, 0x19, 0x49, 0xeb, 0x2c, 0xf8, 0x76, 0xa5, 0x78,
	0xcc, 0xd0, 0x39, 0x5e, 0x4e, 0xa6, 0x80, 0xf5, 0xf9, 0xa2, 0x8e, 0xca, 0x2e, 0x83, 0x7d, 0x97,
	0x13, 0xd9, 0x4c, 0x3e, 0xf0, 0xd4, 0xd8, 0xb9, 0x71, 0x5f, 0xc1, 0x4d, 0x8a, 0xaf, 0xa8, 0x31,
	0xfb, 0x26, 0x7e, 0xf0, 0xd5, 0x46, 0xf8, 0xd9, 0x73, 0x02, 0xdf, 0x00, 0xcc, 0x8d, 0xd0, 0x76,
	0x46, 0xe6, 0xb1, 0x9f, 0xb7, 0x0a, 0xb9, 0xfa, 0x17, 0xc0, 0x60, 0xd6, 0xb8, 0x2c, 0xf8, 0x0d,
	0xf0, 0x4e, 0xc8, 0xb8, 0x71, 0x25, 0x86, 0xa1, 0xbd, 0x3d, 0x61, 0x9d, 0x8e, 0x59, 0x1b, 0xa5,
	0xa3, 0xbc, 0xb0, 0xeb, 0xd8, 0xf6, 0x1b, 0xeb, 0xb4, 0x56, 0xd8, 0xdc, 0xb8, 0x75, 0x87, 0x62,
	0x9d, 0xcf, 0x30, 0x9a, 0x69, 0xc1, 0x8d, 0x28, 0x1b, 0x2f, 0x1b, 0xed, 0xb9, 0x82, 0xaf, 0x59,
	0x6e, 0xb6, 0xe3, 0x91, 0x13, 0x65, 0x66, 0x57, 0x7e, 0xfd, 0x09, 0x2e, 0x22, 0x95, 0x85, 0x71,
	0x6a, 0x74, 0x6a, 0xc2, 0xa4, 0x28, 0xe4, 0xab, 0x5f, 0xe1, 0x23, 0xff, 0xcd, 0xa9, 0xe2, 0x7a,
	0xd4, 0x7c, 0x01, 0x0b, 0xfb, 0x57, 0x58, 0x04, 0x3f, 0x0f, 0xe9, 0xf7, 0xf0, 0xe1, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xca, 0xe9, 0x3a, 0xd4, 0x4f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnectorCommandClient is the client API for ConnectorCommand service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnectorCommandClient interface {
	SendCommandMessage(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandMessageUUID, error)
	WaitCommandMessage(ctx context.Context, in *CommandMessageWait, opts ...grpc.CallOption) (*CommandMessage, error)
	CreateIteratorCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error)
}

type connectorCommandClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorCommandClient(cc grpc.ClientConnInterface) ConnectorCommandClient {
	return &connectorCommandClient{cc}
}

func (c *connectorCommandClient) SendCommandMessage(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandMessageUUID, error) {
	out := new(CommandMessageUUID)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorCommand/SendCommandMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorCommandClient) WaitCommandMessage(ctx context.Context, in *CommandMessageWait, opts ...grpc.CallOption) (*CommandMessage, error) {
	out := new(CommandMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorCommand/WaitCommandMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorCommandClient) CreateIteratorCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*IteratorMessage, error) {
	out := new(IteratorMessage)
	err := c.cc.Invoke(ctx, "/grpc.ConnectorCommand/CreateIteratorCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorCommandServer is the server API for ConnectorCommand service.
type ConnectorCommandServer interface {
	SendCommandMessage(context.Context, *CommandMessage) (*CommandMessageUUID, error)
	WaitCommandMessage(context.Context, *CommandMessageWait) (*CommandMessage, error)
	CreateIteratorCommand(context.Context, *Empty) (*IteratorMessage, error)
}

// UnimplementedConnectorCommandServer can be embedded to have forward compatible implementations.
type UnimplementedConnectorCommandServer struct {
}

func (*UnimplementedConnectorCommandServer) SendCommandMessage(ctx context.Context, req *CommandMessage) (*CommandMessageUUID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandMessage not implemented")
}
func (*UnimplementedConnectorCommandServer) WaitCommandMessage(ctx context.Context, req *CommandMessageWait) (*CommandMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitCommandMessage not implemented")
}
func (*UnimplementedConnectorCommandServer) CreateIteratorCommand(ctx context.Context, req *Empty) (*IteratorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIteratorCommand not implemented")
}

func RegisterConnectorCommandServer(s *grpc.Server, srv ConnectorCommandServer) {
	s.RegisterService(&_ConnectorCommand_serviceDesc, srv)
}

func _ConnectorCommand_SendCommandMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).SendCommandMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorCommand/SendCommandMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).SendCommandMessage(ctx, req.(*CommandMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorCommand_WaitCommandMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessageWait)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).WaitCommandMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorCommand/WaitCommandMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).WaitCommandMessage(ctx, req.(*CommandMessageWait))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorCommand_CreateIteratorCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorCommandServer).CreateIteratorCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ConnectorCommand/CreateIteratorCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorCommandServer).CreateIteratorCommand(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConnectorCommand_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.ConnectorCommand",
	HandlerType: (*ConnectorCommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCommandMessage",
			Handler:    _ConnectorCommand_SendCommandMessage_Handler,
		},
		{
			MethodName: "WaitCommandMessage",
			Handler:    _ConnectorCommand_WaitCommandMessage_Handler,
		},
		{
			MethodName: "CreateIteratorCommand",
			Handler:    _ConnectorCommand_CreateIteratorCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connectorCommand.proto",
}
