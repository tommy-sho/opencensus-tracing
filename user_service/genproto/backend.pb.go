// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/backend.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Mail                 string   `protobuf:"bytes,4,opt,name=mail,proto3" json:"mail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *User) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

type SaveUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserRequest) Reset()         { *m = SaveUserRequest{} }
func (m *SaveUserRequest) String() string { return proto.CompactTextString(m) }
func (*SaveUserRequest) ProtoMessage()    {}
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{1}
}

func (m *SaveUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserRequest.Unmarshal(m, b)
}
func (m *SaveUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserRequest.Marshal(b, m, deterministic)
}
func (m *SaveUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserRequest.Merge(m, src)
}
func (m *SaveUserRequest) XXX_Size() int {
	return xxx_messageInfo_SaveUserRequest.Size(m)
}
func (m *SaveUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserRequest proto.InternalMessageInfo

func (m *SaveUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SaveUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveUserResponse) Reset()         { *m = SaveUserResponse{} }
func (m *SaveUserResponse) String() string { return proto.CompactTextString(m) }
func (*SaveUserResponse) ProtoMessage()    {}
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{2}
}

func (m *SaveUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveUserResponse.Unmarshal(m, b)
}
func (m *SaveUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveUserResponse.Marshal(b, m, deterministic)
}
func (m *SaveUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveUserResponse.Merge(m, src)
}
func (m *SaveUserResponse) XXX_Size() int {
	return xxx_messageInfo_SaveUserResponse.Size(m)
}
func (m *SaveUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveUserResponse proto.InternalMessageInfo

func (m *SaveUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{3}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{4}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUsersRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{5}
}

func (m *GetUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersRequest.Unmarshal(m, b)
}
func (m *GetUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersRequest.Marshal(b, m, deterministic)
}
func (m *GetUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersRequest.Merge(m, src)
}
func (m *GetUsersRequest) XXX_Size() int {
	return xxx_messageInfo_GetUsersRequest.Size(m)
}
func (m *GetUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersRequest proto.InternalMessageInfo

type GetUsersResponse struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f614a5ff5361bf, []int{6}
}

func (m *GetUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUsersResponse.Unmarshal(m, b)
}
func (m *GetUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUsersResponse.Marshal(b, m, deterministic)
}
func (m *GetUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUsersResponse.Merge(m, src)
}
func (m *GetUsersResponse) XXX_Size() int {
	return xxx_messageInfo_GetUsersResponse.Size(m)
}
func (m *GetUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUsersResponse proto.InternalMessageInfo

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*SaveUserRequest)(nil), "proto.SaveUserRequest")
	proto.RegisterType((*SaveUserResponse)(nil), "proto.SaveUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "proto.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "proto.GetUserResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "proto.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "proto.GetUsersResponse")
}

func init() { proto.RegisterFile("proto/backend.proto", fileDescriptor_e0f614a5ff5361bf) }

var fileDescriptor_e0f614a5ff5361bf = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0xb3, 0xfe, 0x81, 0x30, 0x96, 0x82, 0x19, 0x44, 0x58, 0xa5, 0xc1, 0x6c, 0x95, 0x2a,
	0x48, 0x8e, 0x68, 0x90, 0x52, 0xd3, 0x22, 0x47, 0x1c, 0x60, 0x13, 0x8f, 0x90, 0x05, 0xb1, 0x83,
	0xd7, 0xc9, 0x15, 0xb9, 0x16, 0xda, 0x1f, 0x6c, 0x58, 0x0a, 0xa8, 0x3c, 0xfb, 0xe6, 0x9b, 0x37,
	0xf3, 0x0c, 0x97, 0xfb, 0xb6, 0xe9, 0x9a, 0xbb, 0x8d, 0xdc, 0xbe, 0x52, 0x5d, 0x2e, 0xcc, 0x0b,
	0x63, 0xf3, 0x11, 0x4f, 0x10, 0x3d, 0x2b, 0x6a, 0x71, 0x02, 0x41, 0x55, 0x72, 0x96, 0xb1, 0xf9,
	0x59, 0x11, 0x54, 0x25, 0x22, 0x44, 0xb5, 0xdc, 0x11, 0x0f, 0x8c, 0x62, 0x6a, 0x4c, 0x21, 0x94,
	0x2f, 0xc4, 0xc3, 0x8c, 0xcd, 0xe3, 0x42, 0x97, 0x9a, 0xda, 0xc9, 0xea, 0x8d, 0x47, 0x96, 0xd2,
	0xb5, 0xc8, 0xe1, 0x7c, 0x2d, 0x8f, 0xa4, 0x5d, 0x0b, 0x7a, 0x3f, 0x90, 0xea, 0xf0, 0x06, 0xa2,
	0x83, 0xa2, 0xd6, 0xd8, 0x27, 0x79, 0x62, 0x2f, 0x58, 0x18, 0xc2, 0x34, 0xc4, 0x12, 0xd2, 0x61,
	0x46, 0xed, 0x9b, 0x5a, 0xd1, 0xdf, 0x43, 0x19, 0x4c, 0x1e, 0xa9, 0xfb, 0xbe, 0xc7, 0x0b, 0xa1,
	0x4f, 0xe9, 0x89, 0xff, 0xba, 0x5e, 0xf4, 0x33, 0xca, 0xd9, 0x8a, 0x7b, 0x48, 0x07, 0xc9, 0xf9,
	0xdc, 0x42, 0xac, 0x71, 0xc5, 0x59, 0x16, 0xfa, 0x46, 0xb6, 0x93, 0x7f, 0x30, 0x48, 0xf4, 0x7b,
	0x4d, 0xed, 0xb1, 0xda, 0x12, 0xae, 0x60, 0xfc, 0x15, 0x12, 0xa7, 0x8e, 0xf7, 0xfe, 0xd4, 0xec,
	0xfa, 0x97, 0x6e, 0xf7, 0x89, 0x11, 0x3e, 0xc0, 0xa9, 0xbb, 0x02, 0xaf, 0x1c, 0xf5, 0x33, 0xfe,
	0x6c, 0xea, 0xcb, 0xfd, 0xec, 0x0a, 0xc6, 0x4e, 0xcc, 0xd1, 0xa3, 0x94, 0xbf, 0xda, 0x8f, 0x2a,
	0x46, 0x9b, 0x13, 0xd3, 0x59, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x35, 0x7f, 0xb8, 0xe9, 0x49,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUser2(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) SaveUser(ctx context.Context, in *SaveUserRequest, opts ...grpc.CallOption) (*SaveUserResponse, error) {
	out := new(SaveUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUser2(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/proto.UserService/GetUser2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	SaveUser(context.Context, *SaveUserRequest) (*SaveUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUser2(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SaveUser(ctx, req.(*SaveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUser2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/GetUser2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser2(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _UserService_SaveUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUser2",
			Handler:    _UserService_GetUser2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/backend.proto",
}