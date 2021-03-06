// Code generated by protoc-gen-go. DO NOT EDIT.
// source: httprule/all_pattern.proto

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

type AllPatternRequest struct {
	Double               float64   `protobuf:"fixed64,1,opt,name=double,proto3" json:"double,omitempty"`
	Float                float32   `protobuf:"fixed32,2,opt,name=float,proto3" json:"float,omitempty"`
	Int32                int32     `protobuf:"varint,3,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64                int64     `protobuf:"varint,4,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint32               uint32    `protobuf:"varint,5,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64               uint64    `protobuf:"varint,6,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Fixed32              uint32    `protobuf:"fixed32,7,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64              uint64    `protobuf:"fixed64,8,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32             int32     `protobuf:"fixed32,9,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64             int64     `protobuf:"fixed64,10,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Bool                 bool      `protobuf:"varint,11,opt,name=bool,proto3" json:"bool,omitempty"`
	String_              string    `protobuf:"bytes,12,opt,name=string,proto3" json:"string,omitempty"`
	Bytes                []byte    `protobuf:"bytes,14,opt,name=bytes,proto3" json:"bytes,omitempty"`
	RepeatedDouble       []float64 `protobuf:"fixed64,15,rep,packed,name=repeated_double,json=repeatedDouble,proto3" json:"repeated_double,omitempty"`
	RepeatedFloat        []float32 `protobuf:"fixed32,16,rep,packed,name=repeated_float,json=repeatedFloat,proto3" json:"repeated_float,omitempty"`
	RepeatedInt32        []int32   `protobuf:"varint,17,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedInt64        []int64   `protobuf:"varint,18,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedUint32       []uint32  `protobuf:"varint,19,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedUint64       []uint64  `protobuf:"varint,20,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedFixed32      []uint32  `protobuf:"fixed32,21,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64      []uint64  `protobuf:"fixed64,22,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32     []int32   `protobuf:"fixed32,23,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64     []int64   `protobuf:"fixed64,24,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedBool         []bool    `protobuf:"varint,25,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedString       []string  `protobuf:"bytes,26,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes        [][]byte  `protobuf:"bytes,28,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AllPatternRequest) Reset()         { *m = AllPatternRequest{} }
func (m *AllPatternRequest) String() string { return proto.CompactTextString(m) }
func (*AllPatternRequest) ProtoMessage()    {}
func (*AllPatternRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c72a0d5013ebec1, []int{0}
}

func (m *AllPatternRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllPatternRequest.Unmarshal(m, b)
}
func (m *AllPatternRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllPatternRequest.Marshal(b, m, deterministic)
}
func (m *AllPatternRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllPatternRequest.Merge(m, src)
}
func (m *AllPatternRequest) XXX_Size() int {
	return xxx_messageInfo_AllPatternRequest.Size(m)
}
func (m *AllPatternRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllPatternRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllPatternRequest proto.InternalMessageInfo

func (m *AllPatternRequest) GetDouble() float64 {
	if m != nil {
		return m.Double
	}
	return 0
}

func (m *AllPatternRequest) GetFloat() float32 {
	if m != nil {
		return m.Float
	}
	return 0
}

func (m *AllPatternRequest) GetInt32() int32 {
	if m != nil {
		return m.Int32
	}
	return 0
}

func (m *AllPatternRequest) GetInt64() int64 {
	if m != nil {
		return m.Int64
	}
	return 0
}

func (m *AllPatternRequest) GetUint32() uint32 {
	if m != nil {
		return m.Uint32
	}
	return 0
}

func (m *AllPatternRequest) GetUint64() uint64 {
	if m != nil {
		return m.Uint64
	}
	return 0
}

func (m *AllPatternRequest) GetFixed32() uint32 {
	if m != nil {
		return m.Fixed32
	}
	return 0
}

func (m *AllPatternRequest) GetFixed64() uint64 {
	if m != nil {
		return m.Fixed64
	}
	return 0
}

func (m *AllPatternRequest) GetSfixed32() int32 {
	if m != nil {
		return m.Sfixed32
	}
	return 0
}

func (m *AllPatternRequest) GetSfixed64() int64 {
	if m != nil {
		return m.Sfixed64
	}
	return 0
}

func (m *AllPatternRequest) GetBool() bool {
	if m != nil {
		return m.Bool
	}
	return false
}

func (m *AllPatternRequest) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

func (m *AllPatternRequest) GetBytes() []byte {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedDouble() []float64 {
	if m != nil {
		return m.RepeatedDouble
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedFloat() []float32 {
	if m != nil {
		return m.RepeatedFloat
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedInt32() []int32 {
	if m != nil {
		return m.RepeatedInt32
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedInt64() []int64 {
	if m != nil {
		return m.RepeatedInt64
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedUint32() []uint32 {
	if m != nil {
		return m.RepeatedUint32
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedUint64() []uint64 {
	if m != nil {
		return m.RepeatedUint64
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedFixed32() []uint32 {
	if m != nil {
		return m.RepeatedFixed32
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedFixed64() []uint64 {
	if m != nil {
		return m.RepeatedFixed64
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedSfixed32() []int32 {
	if m != nil {
		return m.RepeatedSfixed32
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedSfixed64() []int64 {
	if m != nil {
		return m.RepeatedSfixed64
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedBool() []bool {
	if m != nil {
		return m.RepeatedBool
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedString() []string {
	if m != nil {
		return m.RepeatedString
	}
	return nil
}

func (m *AllPatternRequest) GetRepeatedBytes() [][]byte {
	if m != nil {
		return m.RepeatedBytes
	}
	return nil
}

type AllPatternResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllPatternResponse) Reset()         { *m = AllPatternResponse{} }
func (m *AllPatternResponse) String() string { return proto.CompactTextString(m) }
func (*AllPatternResponse) ProtoMessage()    {}
func (*AllPatternResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c72a0d5013ebec1, []int{1}
}

func (m *AllPatternResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllPatternResponse.Unmarshal(m, b)
}
func (m *AllPatternResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllPatternResponse.Marshal(b, m, deterministic)
}
func (m *AllPatternResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllPatternResponse.Merge(m, src)
}
func (m *AllPatternResponse) XXX_Size() int {
	return xxx_messageInfo_AllPatternResponse.Size(m)
}
func (m *AllPatternResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllPatternResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllPatternResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AllPatternRequest)(nil), "httprule.AllPatternRequest")
	proto.RegisterType((*AllPatternResponse)(nil), "httprule.AllPatternResponse")
}

func init() { proto.RegisterFile("httprule/all_pattern.proto", fileDescriptor_5c72a0d5013ebec1) }

var fileDescriptor_5c72a0d5013ebec1 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0xab, 0x13, 0x31,
	0x14, 0x25, 0x4d, 0x3b, 0xed, 0x8b, 0xf3, 0x5e, 0xdb, 0x58, 0x9f, 0xb1, 0x76, 0x71, 0xa9, 0x88,
	0x57, 0x84, 0x16, 0xda, 0x90, 0xbd, 0x45, 0x1e, 0xb8, 0x93, 0x3c, 0xdc, 0x08, 0x52, 0xa6, 0x74,
	0x5e, 0x2d, 0x0e, 0x33, 0x63, 0x27, 0x05, 0xdd, 0xfa, 0x17, 0xfc, 0x69, 0xfe, 0x01, 0x17, 0xfe,
	0x10, 0x99, 0x26, 0x99, 0xe9, 0x87, 0xbb, 0x39, 0xe7, 0x9e, 0x5c, 0x72, 0xcf, 0x99, 0x1b, 0x36,
	0xfc, 0x62, 0x4c, 0xbe, 0xdb, 0x27, 0xf1, 0x34, 0x4a, 0x92, 0x65, 0x1e, 0x19, 0x13, 0xef, 0xd2,
	0x49, 0xbe, 0xcb, 0x4c, 0xc6, 0x3b, 0xbe, 0x36, 0x1c, 0x6d, 0xb2, 0x6c, 0x53, 0x6a, 0xf2, 0xed,
	0x34, 0x4a, 0xd3, 0xcc, 0x44, 0x66, 0x9b, 0xa5, 0x85, 0xd5, 0x8d, 0xff, 0x04, 0xac, 0xff, 0x36,
	0x49, 0x3e, 0xd8, 0xc3, 0x3a, 0xfe, 0xb6, 0x8f, 0x0b, 0xc3, 0x6f, 0x59, 0xb0, 0xce, 0xf6, 0xab,
	0x24, 0x16, 0x04, 0x08, 0x12, 0xed, 0x10, 0x1f, 0xb0, 0xd6, 0x43, 0x92, 0x45, 0x46, 0x34, 0x80,
	0x60, 0x43, 0x5b, 0x50, 0xb2, 0xdb, 0xd4, 0xcc, 0x67, 0x82, 0x02, 0xc1, 0x96, 0xb6, 0xc0, 0xb1,
	0x4a, 0x8a, 0x26, 0x10, 0xa4, 0xda, 0x82, 0xb2, 0xf3, 0xde, 0x8a, 0x5b, 0x40, 0xf0, 0x5a, 0x3b,
	0xe4, 0x79, 0x25, 0x45, 0x00, 0x04, 0x9b, 0xda, 0x21, 0x2e, 0x58, 0xfb, 0x61, 0xfb, 0x3d, 0x5e,
	0xcf, 0x67, 0xa2, 0x0d, 0x04, 0xdb, 0xda, 0xc3, 0xaa, 0xa2, 0xa4, 0xe8, 0x00, 0xc1, 0x40, 0x7b,
	0xc8, 0x87, 0xac, 0x53, 0xf8, 0x43, 0x57, 0x40, 0xb0, 0xab, 0x2b, 0x5c, 0xd7, 0x94, 0x14, 0x0c,
	0x08, 0xf6, 0x74, 0x85, 0x39, 0x67, 0xcd, 0x55, 0x96, 0x25, 0xe2, 0x11, 0x10, 0xec, 0xe8, 0xc3,
	0x77, 0x79, 0xaf, 0xc2, 0xec, 0xb6, 0xe9, 0x46, 0x84, 0x40, 0xf0, 0x4a, 0x3b, 0x54, 0x4e, 0xb7,
	0xfa, 0x61, 0xe2, 0x42, 0xdc, 0x00, 0xc1, 0x50, 0x5b, 0xc0, 0x5f, 0xb1, 0xee, 0x2e, 0xce, 0xe3,
	0xc8, 0xc4, 0xeb, 0xa5, 0x33, 0xb0, 0x0b, 0x14, 0x89, 0xbe, 0xf1, 0xf4, 0x3b, 0x6b, 0xe4, 0x4b,
	0x56, 0x31, 0x4b, 0xeb, 0x68, 0x0f, 0x28, 0x36, 0xf4, 0xb5, 0x67, 0xef, 0x0e, 0xce, 0x1e, 0xcb,
	0xac, 0x6b, 0x7d, 0xa0, 0xd8, 0xaa, 0x65, 0xef, 0x0f, 0xe6, 0x9d, 0xc9, 0x94, 0x14, 0x1c, 0x28,
	0xd2, 0x13, 0x99, 0x92, 0x27, 0xb7, 0x73, 0x21, 0x3c, 0x06, 0x8a, 0xd7, 0xf5, 0xed, 0x3e, 0xda,
	0x30, 0xce, 0x85, 0x4a, 0x8a, 0x01, 0x50, 0x6c, 0x9e, 0x0a, 0x95, 0xe4, 0xaf, 0x59, 0xaf, 0x1e,
	0xc3, 0x39, 0xfe, 0x04, 0x28, 0xb6, 0x75, 0xd5, 0xe0, 0xce, 0x19, 0x7f, 0x21, 0x55, 0x52, 0xdc,
	0x02, 0xc5, 0xe0, 0x4c, 0xaa, 0x24, 0x7f, 0xc3, 0xfa, 0x95, 0xb4, 0x0a, 0xf2, 0x29, 0x50, 0xec,
	0xea, 0xaa, 0xc7, 0xbd, 0x0f, 0xf4, 0x52, 0xac, 0xa4, 0x10, 0x40, 0xb1, 0x77, 0x2e, 0x56, 0x92,
	0xbf, 0x60, 0x95, 0x25, 0xcb, 0x43, 0xd4, 0xcf, 0x80, 0x62, 0x47, 0x87, 0x9e, 0x5c, 0x94, 0x91,
	0x1f, 0x4f, 0xef, 0xb2, 0x1f, 0x02, 0xc5, 0xab, 0x7a, 0xfa, 0x7b, 0xfb, 0x0f, 0x1c, 0xdb, 0x6e,
	0x7f, 0x86, 0x11, 0x50, 0x0c, 0x6b, 0xdb, 0x17, 0x25, 0x39, 0x1e, 0x30, 0x7e, 0xbc, 0x61, 0x45,
	0x9e, 0xa5, 0x45, 0x3c, 0xfb, 0xca, 0x58, 0xcd, 0xf2, 0xcf, 0x27, 0xe8, 0xf9, 0xc4, 0x6f, 0xef,
	0xe4, 0x62, 0x37, 0x87, 0xa3, 0xff, 0x17, 0x6d, 0xdb, 0xf1, 0xe0, 0xe7, 0xef, 0xbf, 0xbf, 0x1a,
	0x37, 0x3c, 0x2c, 0xdf, 0x84, 0xa9, 0x7b, 0x13, 0x16, 0xe1, 0x27, 0xe6, 0x0f, 0xe5, 0xab, 0x55,
	0x70, 0x58, 0xfd, 0xf9, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x1f, 0x60, 0xa8, 0x40, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AllPatternClient is the client API for AllPattern service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AllPatternClient interface {
	AllPattern(ctx context.Context, in *AllPatternRequest, opts ...grpc.CallOption) (*AllPatternResponse, error)
}

type allPatternClient struct {
	cc *grpc.ClientConn
}

func NewAllPatternClient(cc *grpc.ClientConn) AllPatternClient {
	return &allPatternClient{cc}
}

func (c *allPatternClient) AllPattern(ctx context.Context, in *AllPatternRequest, opts ...grpc.CallOption) (*AllPatternResponse, error) {
	out := new(AllPatternResponse)
	err := c.cc.Invoke(ctx, "/httprule.AllPattern/AllPattern", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllPatternServer is the server API for AllPattern service.
type AllPatternServer interface {
	AllPattern(context.Context, *AllPatternRequest) (*AllPatternResponse, error)
}

// UnimplementedAllPatternServer can be embedded to have forward compatible implementations.
type UnimplementedAllPatternServer struct {
}

func (*UnimplementedAllPatternServer) AllPattern(ctx context.Context, req *AllPatternRequest) (*AllPatternResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllPattern not implemented")
}

func RegisterAllPatternServer(s *grpc.Server, srv AllPatternServer) {
	s.RegisterService(&_AllPattern_serviceDesc, srv)
}

func _AllPattern_AllPattern_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllPatternRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllPatternServer).AllPattern(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httprule.AllPattern/AllPattern",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllPatternServer).AllPattern(ctx, req.(*AllPatternRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AllPattern_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httprule.AllPattern",
	HandlerType: (*AllPatternServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllPattern",
			Handler:    _AllPattern_AllPattern_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httprule/all_pattern.proto",
}
