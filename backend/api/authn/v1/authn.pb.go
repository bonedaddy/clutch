// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authn/v1/authn.proto

package authnv1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/lyft/clutch/backend/api/api/v1"
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

type LoginRequest struct {
	RedirectUrl          string   `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f51d149111bdf3, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

type LoginResponse struct {
	AuthUrl              string   `protobuf:"bytes,1,opt,name=auth_url,json=authUrl,proto3" json:"auth_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f51d149111bdf3, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAuthUrl() string {
	if m != nil {
		return m.AuthUrl
	}
	return ""
}

// See https://www.oauth.com/oauth2-servers/authorization/the-authorization-response/ for description of the parameters.
type CallbackRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	Error                string   `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	ErrorDescription     string   `protobuf:"bytes,4,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackRequest) Reset()         { *m = CallbackRequest{} }
func (m *CallbackRequest) String() string { return proto.CompactTextString(m) }
func (*CallbackRequest) ProtoMessage()    {}
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f51d149111bdf3, []int{2}
}

func (m *CallbackRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackRequest.Unmarshal(m, b)
}
func (m *CallbackRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackRequest.Marshal(b, m, deterministic)
}
func (m *CallbackRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackRequest.Merge(m, src)
}
func (m *CallbackRequest) XXX_Size() int {
	return xxx_messageInfo_CallbackRequest.Size(m)
}
func (m *CallbackRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackRequest proto.InternalMessageInfo

func (m *CallbackRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CallbackRequest) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *CallbackRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CallbackRequest) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

type CallbackResponse struct {
	// This is the token that the user should present. Note: this response is only valid in a gRPC context. In an HTTP
	// context the user will be redirected.
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackResponse) Reset()         { *m = CallbackResponse{} }
func (m *CallbackResponse) String() string { return proto.CompactTextString(m) }
func (*CallbackResponse) ProtoMessage()    {}
func (*CallbackResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15f51d149111bdf3, []int{3}
}

func (m *CallbackResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackResponse.Unmarshal(m, b)
}
func (m *CallbackResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackResponse.Marshal(b, m, deterministic)
}
func (m *CallbackResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackResponse.Merge(m, src)
}
func (m *CallbackResponse) XXX_Size() int {
	return xxx_messageInfo_CallbackResponse.Size(m)
}
func (m *CallbackResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackResponse proto.InternalMessageInfo

func (m *CallbackResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "clutch.authn.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "clutch.authn.v1.LoginResponse")
	proto.RegisterType((*CallbackRequest)(nil), "clutch.authn.v1.CallbackRequest")
	proto.RegisterType((*CallbackResponse)(nil), "clutch.authn.v1.CallbackResponse")
}

func init() {
	proto.RegisterFile("authn/v1/authn.proto", fileDescriptor_15f51d149111bdf3)
}

var fileDescriptor_15f51d149111bdf3 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x4d, 0x2b, 0x68, 0x19, 0x31, 0xc0, 0xa6, 0x87, 0xda, 0x80, 0x81, 0x9e, 0x88, 0x26, 0x6d,
	0xaa, 0x5f, 0x80, 0x7a, 0x31, 0xf1, 0x60, 0x48, 0xbc, 0x78, 0x21, 0x65, 0xd9, 0x40, 0xc3, 0x66,
	0xb7, 0x6c, 0xb7, 0x1c, 0x3d, 0xf8, 0x0b, 0xfe, 0x85, 0xbf, 0xe3, 0x17, 0x98, 0xf0, 0x21, 0x66,
	0xa7, 0x45, 0x08, 0x46, 0x6f, 0x33, 0x6f, 0x5e, 0xe7, 0xbd, 0x7d, 0x53, 0x70, 0x93, 0x42, 0x2f,
	0x44, 0xb4, 0x8e, 0x23, 0x2c, 0xc2, 0x4c, 0x49, 0x2d, 0x49, 0x8b, 0xf2, 0x42, 0xd3, 0x45, 0x58,
	0x62, 0xeb, 0xd8, 0xef, 0xce, 0xa5, 0x9c, 0x73, 0x16, 0x25, 0x59, 0x1a, 0x25, 0x42, 0x48, 0x9d,
	0xe8, 0x54, 0x8a, 0xbc, 0xa4, 0xfb, 0x9e, 0x81, 0xcd, 0x8a, 0xc3, 0x49, 0x10, 0x43, 0xf3, 0x51,
	0xce, 0x53, 0x31, 0x66, 0xab, 0x82, 0xe5, 0x9a, 0x0c, 0xa0, 0xa9, 0xd8, 0x2c, 0x55, 0x8c, 0xea,
	0x49, 0xa1, 0xb8, 0x67, 0xf5, 0xad, 0x61, 0x63, 0x7c, 0xba, 0xc5, 0x9e, 0x15, 0x0f, 0x2e, 0xe1,
	0xac, 0xfa, 0x24, 0xcf, 0xa4, 0xc8, 0x19, 0x39, 0x07, 0xc7, 0xf8, 0xd8, 0xe3, 0x9f, 0x98, 0xde,
	0x70, 0x5f, 0xa1, 0x75, 0x97, 0x70, 0x3e, 0x4d, 0xe8, 0x72, 0xab, 0x40, 0xa0, 0x46, 0xe5, 0x8c,
	0x55, 0x4c, 0xac, 0x89, 0x0b, 0xf5, 0x5c, 0x27, 0x9a, 0x79, 0x36, 0x82, 0x65, 0x63, 0x50, 0xa6,
	0x94, 0x54, 0xde, 0x51, 0x89, 0x62, 0x43, 0xae, 0xa0, 0x83, 0xc5, 0x64, 0xc6, 0x72, 0xaa, 0xd2,
	0xcc, 0xbc, 0xc6, 0xab, 0x21, 0xa3, 0x8d, 0x83, 0xfb, 0x1d, 0x1e, 0x0c, 0xa1, 0xbd, 0xd3, 0xaf,
	0xec, 0xba, 0x50, 0xd7, 0x72, 0xc9, 0x44, 0xe5, 0xa0, 0x6c, 0xae, 0x37, 0x16, 0x38, 0x23, 0x93,
	0xe6, 0xe8, 0xe9, 0x81, 0x30, 0xa8, 0xe3, 0x13, 0x49, 0x2f, 0x3c, 0x08, 0x3a, 0xdc, 0x4f, 0xcb,
	0xbf, 0xf8, 0x6b, 0x5c, 0x4a, 0x05, 0xbd, 0x8f, 0xaf, 0xae, 0xed, 0xd8, 0x6f, 0x9f, 0x9b, 0x77,
	0xbb, 0x43, 0x5a, 0x3f, 0x57, 0x8c, 0x38, 0x6e, 0x5f, 0x81, 0xb3, 0x75, 0x47, 0xfa, 0xbf, 0x56,
	0x1d, 0x04, 0xe7, 0x0f, 0xfe, 0x61, 0x54, 0x7a, 0x7d, 0xd4, 0xb3, 0x50, 0xcf, 0x25, 0x64, 0xa7,
	0x47, 0x2b, 0xe6, 0x6d, 0xe3, 0x05, 0x6f, 0x23, 0xd6, 0xf1, 0xf4, 0x18, 0xff, 0x80, 0x9b, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x53, 0x76, 0x73, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthnAPIClient is the client API for AuthnAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthnAPIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
}

type authnAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnAPIClient(cc grpc.ClientConnInterface) AuthnAPIClient {
	return &authnAPIClient{cc}
}

func (c *authnAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/clutch.authn.v1.AuthnAPI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnAPIClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, "/clutch.authn.v1.AuthnAPI/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnAPIServer is the server API for AuthnAPI service.
type AuthnAPIServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
}

// UnimplementedAuthnAPIServer can be embedded to have forward compatible implementations.
type UnimplementedAuthnAPIServer struct {
}

func (*UnimplementedAuthnAPIServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAuthnAPIServer) Callback(ctx context.Context, req *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}

func RegisterAuthnAPIServer(s *grpc.Server, srv AuthnAPIServer) {
	s.RegisterService(&_AuthnAPI_serviceDesc, srv)
}

func _AuthnAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.authn.v1.AuthnAPI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnAPI_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.authn.v1.AuthnAPI/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthnAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.authn.v1.AuthnAPI",
	HandlerType: (*AuthnAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthnAPI_Login_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _AuthnAPI_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authn/v1/authn.proto",
}
