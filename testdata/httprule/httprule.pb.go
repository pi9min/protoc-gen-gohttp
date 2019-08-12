// Code generated by protoc-gen-go. DO NOT EDIT.
// source: httprule/httprule.proto

package httprulepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetMessageRequest struct {
	MessageId            string                        `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Sub                  *GetMessageRequest_SubMessage `protobuf:"bytes,2,opt,name=sub,proto3" json:"sub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *GetMessageRequest) Reset()         { *m = GetMessageRequest{} }
func (m *GetMessageRequest) String() string { return proto.CompactTextString(m) }
func (*GetMessageRequest) ProtoMessage()    {}
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{0}
}

func (m *GetMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageRequest.Unmarshal(m, b)
}
func (m *GetMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageRequest.Marshal(b, m, deterministic)
}
func (m *GetMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageRequest.Merge(m, src)
}
func (m *GetMessageRequest) XXX_Size() int {
	return xxx_messageInfo_GetMessageRequest.Size(m)
}
func (m *GetMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageRequest proto.InternalMessageInfo

func (m *GetMessageRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *GetMessageRequest) GetSub() *GetMessageRequest_SubMessage {
	if m != nil {
		return m.Sub
	}
	return nil
}

type GetMessageRequest_SubMessage struct {
	Subfield             string   `protobuf:"bytes,1,opt,name=subfield,proto3" json:"subfield,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMessageRequest_SubMessage) Reset()         { *m = GetMessageRequest_SubMessage{} }
func (m *GetMessageRequest_SubMessage) String() string { return proto.CompactTextString(m) }
func (*GetMessageRequest_SubMessage) ProtoMessage()    {}
func (*GetMessageRequest_SubMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{0, 0}
}

func (m *GetMessageRequest_SubMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMessageRequest_SubMessage.Unmarshal(m, b)
}
func (m *GetMessageRequest_SubMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMessageRequest_SubMessage.Marshal(b, m, deterministic)
}
func (m *GetMessageRequest_SubMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMessageRequest_SubMessage.Merge(m, src)
}
func (m *GetMessageRequest_SubMessage) XXX_Size() int {
	return xxx_messageInfo_GetMessageRequest_SubMessage.Size(m)
}
func (m *GetMessageRequest_SubMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMessageRequest_SubMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMessageRequest_SubMessage proto.InternalMessageInfo

func (m *GetMessageRequest_SubMessage) GetSubfield() string {
	if m != nil {
		return m.Subfield
	}
	return ""
}

type Message struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_72d547f0661199d3, []int{1}
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

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMessageRequest)(nil), "httprule.GetMessageRequest")
	proto.RegisterType((*GetMessageRequest_SubMessage)(nil), "httprule.GetMessageRequest.SubMessage")
	proto.RegisterType((*Message)(nil), "httprule.Message")
}

func init() { proto.RegisterFile("httprule/httprule.proto", fileDescriptor_72d547f0661199d3) }

var fileDescriptor_72d547f0661199d3 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x28, 0x29, 0x29,
	0x28, 0x2a, 0xcd, 0x49, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc,
	0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x3a, 0xa5, 0x69, 0x8c, 0x5c, 0x82,
	0xee, 0xa9, 0x25, 0xbe, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0xb2, 0x5c, 0x5c, 0xb9, 0x10, 0x91, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x4e, 0xa8, 0x88, 0x67, 0x8a, 0x90, 0x05, 0x17, 0x73, 0x71, 0x69, 0x92, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0xb7, 0x91, 0x9a, 0x1e, 0xdc, 0x6a, 0x0c, 0x83, 0xf4, 0x82, 0x4b, 0x93, 0x60,
	0x22, 0x20, 0x2d, 0x52, 0x1a, 0x5c, 0x5c, 0x08, 0x21, 0x21, 0x29, 0x2e, 0x8e, 0xe2, 0xd2, 0xa4,
	0xb4, 0xcc, 0xd4, 0x1c, 0x98, 0x25, 0x70, 0xbe, 0x92, 0x2c, 0x17, 0x3b, 0x4c, 0x99, 0x10, 0x17,
	0x4b, 0x49, 0x6a, 0x45, 0x09, 0x54, 0x09, 0x98, 0x6d, 0x54, 0xcd, 0xc5, 0x09, 0x91, 0xce, 0xcc,
	0x4b, 0x17, 0xca, 0xe3, 0xe2, 0x42, 0x58, 0x2d, 0x24, 0x8d, 0xc7, 0x41, 0x52, 0x82, 0x08, 0x49,
	0xa8, 0x8c, 0x92, 0x41, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xb4, 0x84, 0x34, 0xf4, 0xcb, 0x0c, 0xf5,
	0xa1, 0x9e, 0x2c, 0xd6, 0xaf, 0x46, 0x04, 0x40, 0xad, 0x7e, 0x75, 0x71, 0x69, 0x92, 0x1e, 0xcc,
	0x69, 0xb5, 0x4e, 0x3c, 0x51, 0x5c, 0x30, 0x53, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x21, 0x69, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x99, 0xcf, 0x9b, 0xbf, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessagingClient interface {
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error)
}

type messagingClient struct {
	cc *grpc.ClientConn
}

func NewMessagingClient(cc *grpc.ClientConn) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/httprule.Messaging/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagingServer is the server API for Messaging service.
type MessagingServer interface {
	GetMessage(context.Context, *GetMessageRequest) (*Message, error)
}

// UnimplementedMessagingServer can be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (*UnimplementedMessagingServer) GetMessage(ctx context.Context, req *GetMessageRequest) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}

func RegisterMessagingServer(s *grpc.Server, srv MessagingServer) {
	s.RegisterService(&_Messaging_serviceDesc, srv)
}

func _Messaging_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagingServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httprule.Messaging/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagingServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Messaging_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httprule.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _Messaging_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httprule/httprule.proto",
}
