// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

/*
Package backend is a generated protocol buffer package.

It is generated from these files:
	backend.proto

It has these top-level messages:
	Empty
	IDResponse
	Device
	CreateUserRequest
	AddDeviceRequest
	LoginRequest
	ConfirmRequest
	UserPersonalInfoRequest
	UserCountryRequest
*/
package backend

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type IDResponse struct {
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
}

func (m *IDResponse) Reset()                    { *m = IDResponse{} }
func (m *IDResponse) String() string            { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()               {}
func (*IDResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IDResponse) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Device struct {
	ID     string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Locale string `protobuf:"bytes,3,opt,name=locale" json:"locale,omitempty"`
	Lang   string `protobuf:"bytes,4,opt,name=lang" json:"lang,omitempty"`
}

func (m *Device) Reset()                    { *m = Device{} }
func (m *Device) String() string            { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()               {}
func (*Device) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Device) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

func (m *Device) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type CreateUserRequest struct {
	Username     string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	PasswordHash string  `protobuf:"bytes,2,opt,name=passwordHash" json:"passwordHash,omitempty"`
	Device       *Device `protobuf:"bytes,3,opt,name=device" json:"device,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *CreateUserRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type AddDeviceRequest struct {
	UserID string  `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Device *Device `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
}

func (m *AddDeviceRequest) Reset()                    { *m = AddDeviceRequest{} }
func (m *AddDeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*AddDeviceRequest) ProtoMessage()               {}
func (*AddDeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *AddDeviceRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type LoginRequest struct {
	SessionID   string  `protobuf:"bytes,1,opt,name=sessionID" json:"sessionID,omitempty"`
	UserID      string  `protobuf:"bytes,2,opt,name=userID" json:"userID,omitempty"`
	Device      *Device `protobuf:"bytes,3,opt,name=device" json:"device,omitempty"`
	UserAgent   string  `protobuf:"bytes,4,opt,name=userAgent" json:"userAgent,omitempty"`
	LoginIP     string  `protobuf:"bytes,5,opt,name=loginIP" json:"loginIP,omitempty"`
	LoginRegion string  `protobuf:"bytes,6,opt,name=loginRegion" json:"loginRegion,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginRequest) GetSessionID() string {
	if m != nil {
		return m.SessionID
	}
	return ""
}

func (m *LoginRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *LoginRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *LoginRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *LoginRequest) GetLoginIP() string {
	if m != nil {
		return m.LoginIP
	}
	return ""
}

func (m *LoginRequest) GetLoginRegion() string {
	if m != nil {
		return m.LoginRegion
	}
	return ""
}

type ConfirmRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	SessionId string `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=userId" json:"userId,omitempty"`
	DeviceId  string `protobuf:"bytes,4,opt,name=deviceId" json:"deviceId,omitempty"`
	Code      string `protobuf:"bytes,5,opt,name=code" json:"code,omitempty"`
}

func (m *ConfirmRequest) Reset()                    { *m = ConfirmRequest{} }
func (m *ConfirmRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfirmRequest) ProtoMessage()               {}
func (*ConfirmRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ConfirmRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ConfirmRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *ConfirmRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ConfirmRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *ConfirmRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type UserPersonalInfoRequest struct {
	UserID    string                     `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	FirstName string                     `protobuf:"bytes,2,opt,name=firstName" json:"firstName,omitempty"`
	LastName  string                     `protobuf:"bytes,3,opt,name=lastName" json:"lastName,omitempty"`
	BirthDate *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=birthDate" json:"birthDate,omitempty"`
}

func (m *UserPersonalInfoRequest) Reset()                    { *m = UserPersonalInfoRequest{} }
func (m *UserPersonalInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserPersonalInfoRequest) ProtoMessage()               {}
func (*UserPersonalInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserPersonalInfoRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserPersonalInfoRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserPersonalInfoRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserPersonalInfoRequest) GetBirthDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.BirthDate
	}
	return nil
}

type UserCountryRequest struct {
	UserID  string `protobuf:"bytes,1,opt,name=userID" json:"userID,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country" json:"country,omitempty"`
}

func (m *UserCountryRequest) Reset()                    { *m = UserCountryRequest{} }
func (m *UserCountryRequest) String() string            { return proto.CompactTextString(m) }
func (*UserCountryRequest) ProtoMessage()               {}
func (*UserCountryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UserCountryRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserCountryRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "backend.Empty")
	proto.RegisterType((*IDResponse)(nil), "backend.IDResponse")
	proto.RegisterType((*Device)(nil), "backend.Device")
	proto.RegisterType((*CreateUserRequest)(nil), "backend.CreateUserRequest")
	proto.RegisterType((*AddDeviceRequest)(nil), "backend.AddDeviceRequest")
	proto.RegisterType((*LoginRequest)(nil), "backend.LoginRequest")
	proto.RegisterType((*ConfirmRequest)(nil), "backend.ConfirmRequest")
	proto.RegisterType((*UserPersonalInfoRequest)(nil), "backend.UserPersonalInfoRequest")
	proto.RegisterType((*UserCountryRequest)(nil), "backend.UserCountryRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Spawn service

type SpawnClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*IDResponse, error)
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*Empty, error)
	HandleLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Empty, error)
	DoConfirm(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateUserCountry(ctx context.Context, in *UserCountryRequest, opts ...grpc.CallOption) (*Empty, error)
	UpdateUserPersonalInfo(ctx context.Context, in *UserPersonalInfoRequest, opts ...grpc.CallOption) (*Empty, error)
}

type spawnClient struct {
	cc *grpc.ClientConn
}

func NewSpawnClient(cc *grpc.ClientConn) SpawnClient {
	return &spawnClient{cc}
}

func (c *spawnClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := grpc.Invoke(ctx, "/backend.Spawn/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/backend.Spawn/AddDevice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnClient) HandleLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/backend.Spawn/HandleLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnClient) DoConfirm(ctx context.Context, in *ConfirmRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/backend.Spawn/DoConfirm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnClient) UpdateUserCountry(ctx context.Context, in *UserCountryRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/backend.Spawn/UpdateUserCountry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spawnClient) UpdateUserPersonalInfo(ctx context.Context, in *UserPersonalInfoRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/backend.Spawn/UpdateUserPersonalInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Spawn service

type SpawnServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*IDResponse, error)
	AddDevice(context.Context, *AddDeviceRequest) (*Empty, error)
	HandleLogin(context.Context, *LoginRequest) (*Empty, error)
	DoConfirm(context.Context, *ConfirmRequest) (*Empty, error)
	UpdateUserCountry(context.Context, *UserCountryRequest) (*Empty, error)
	UpdateUserPersonalInfo(context.Context, *UserPersonalInfoRequest) (*Empty, error)
}

func RegisterSpawnServer(s *grpc.Server, srv SpawnServer) {
	s.RegisterService(&_Spawn_serviceDesc, srv)
}

func _Spawn_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spawn_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spawn_HandleLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).HandleLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/HandleLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).HandleLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spawn_DoConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).DoConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/DoConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).DoConfirm(ctx, req.(*ConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spawn_UpdateUserCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCountryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).UpdateUserCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/UpdateUserCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).UpdateUserCountry(ctx, req.(*UserCountryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Spawn_UpdateUserPersonalInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPersonalInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpawnServer).UpdateUserPersonalInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backend.Spawn/UpdateUserPersonalInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpawnServer).UpdateUserPersonalInfo(ctx, req.(*UserPersonalInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Spawn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backend.Spawn",
	HandlerType: (*SpawnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Spawn_CreateUser_Handler,
		},
		{
			MethodName: "AddDevice",
			Handler:    _Spawn_AddDevice_Handler,
		},
		{
			MethodName: "HandleLogin",
			Handler:    _Spawn_HandleLogin_Handler,
		},
		{
			MethodName: "DoConfirm",
			Handler:    _Spawn_DoConfirm_Handler,
		},
		{
			MethodName: "UpdateUserCountry",
			Handler:    _Spawn_UpdateUserCountry_Handler,
		},
		{
			MethodName: "UpdateUserPersonalInfo",
			Handler:    _Spawn_UpdateUserPersonalInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xbb, 0x36, 0xa5, 0xb7, 0xa3, 0x30, 0x23, 0xba, 0x10, 0x2a, 0x51, 0xe5, 0x05, 0x9e,
	0x3a, 0xa9, 0x48, 0x88, 0x37, 0x34, 0x1a, 0xd0, 0x22, 0x4d, 0x68, 0xea, 0x98, 0xc4, 0xab, 0x5b,
	0xbb, 0x59, 0xb4, 0xd4, 0x0e, 0xb1, 0xcb, 0x34, 0x89, 0x47, 0x9e, 0xf8, 0x1f, 0xfc, 0x16, 0xfe,
	0x16, 0xf2, 0x47, 0x9c, 0x94, 0x76, 0xdb, 0x9b, 0xef, 0x47, 0xce, 0x3d, 0xc7, 0x3e, 0x37, 0xf0,
	0x78, 0x8e, 0x17, 0xd7, 0x94, 0x91, 0x71, 0x5e, 0x70, 0xc9, 0x51, 0xc7, 0x86, 0xc1, 0xab, 0x84,
	0xf3, 0x24, 0xa3, 0xc7, 0x3a, 0x3d, 0x5f, 0x2f, 0x8f, 0x65, 0xba, 0xa2, 0x42, 0xe2, 0x55, 0x6e,
	0x3a, 0xc3, 0x0e, 0xb4, 0x3f, 0xad, 0x72, 0x79, 0x1b, 0x0e, 0x01, 0xe2, 0x68, 0x46, 0x45, 0xce,
	0x99, 0xa0, 0xa8, 0x0f, 0xcd, 0x38, 0xf2, 0x1b, 0xa3, 0xc6, 0x9b, 0xee, 0xac, 0x19, 0x47, 0xe1,
	0x37, 0xf0, 0x22, 0xfa, 0x23, 0x5d, 0x6c, 0x55, 0x10, 0x82, 0x16, 0xc3, 0x2b, 0xea, 0x37, 0x75,
	0x46, 0x9f, 0xd1, 0x00, 0xbc, 0x8c, 0x2f, 0x70, 0x46, 0xfd, 0x7d, 0x9d, 0xb5, 0x91, 0xea, 0xcd,
	0x30, 0x4b, 0xfc, 0x96, 0xe9, 0x55, 0xe7, 0xf0, 0x27, 0x1c, 0x4e, 0x0b, 0x8a, 0x25, 0xbd, 0x14,
	0xb4, 0x98, 0xd1, 0xef, 0x6b, 0x2a, 0x24, 0x0a, 0xe0, 0xd1, 0x5a, 0xd0, 0x42, 0x03, 0x9b, 0x51,
	0x2e, 0x46, 0x21, 0x1c, 0xe4, 0x58, 0x88, 0x1b, 0x5e, 0x90, 0x53, 0x2c, 0xae, 0xec, 0xe0, 0x8d,
	0x1c, 0x7a, 0x0d, 0x1e, 0xd1, 0x74, 0x35, 0x81, 0xde, 0xe4, 0xc9, 0xb8, 0xbc, 0x1f, 0xa3, 0x62,
	0x66, 0xcb, 0xe1, 0x05, 0x3c, 0x3d, 0x21, 0xc4, 0x26, 0xed, 0xf0, 0x01, 0x78, 0x6a, 0x98, 0x53,
	0x69, 0xa3, 0x1a, 0x68, 0xf3, 0x7e, 0xd0, 0xbf, 0x0d, 0x38, 0x38, 0xe3, 0x49, 0xca, 0x4a, 0xc4,
	0x21, 0x74, 0x05, 0x15, 0x22, 0xe5, 0xcc, 0x81, 0x56, 0x89, 0xda, 0xbc, 0xe6, 0x1d, 0xf3, 0xee,
	0x17, 0xa1, 0xe0, 0xd5, 0x27, 0x27, 0x09, 0x65, 0xd2, 0xde, 0x6d, 0x95, 0x40, 0x3e, 0x74, 0x32,
	0x45, 0x26, 0x3e, 0xf7, 0xdb, 0xba, 0x56, 0x86, 0x68, 0x04, 0xbd, 0xcc, 0xd0, 0x4c, 0x52, 0xce,
	0x7c, 0x4f, 0x57, 0xeb, 0xa9, 0xf0, 0x77, 0x03, 0xfa, 0x53, 0xce, 0x96, 0x69, 0xb1, 0x2a, 0xb5,
	0x20, 0x68, 0x5d, 0xa7, 0x8c, 0x58, 0x19, 0xfa, 0x5c, 0xd7, 0x47, 0xac, 0x88, 0x2a, 0xe1, 0xf4,
	0x91, 0xd2, 0x0d, 0x26, 0x52, 0x8f, 0x6c, 0x04, 0xc4, 0xc4, 0xb2, 0x76, 0xb1, 0x9a, 0xb2, 0xe0,
	0x84, 0x5a, 0xc6, 0xfa, 0x1c, 0xfe, 0x69, 0xc0, 0x91, 0x32, 0xc9, 0x39, 0x2d, 0x04, 0x67, 0x38,
	0x8b, 0xd9, 0x92, 0x3f, 0xf4, 0x66, 0x43, 0xe8, 0x2e, 0xd3, 0x42, 0xc8, 0x2f, 0x95, 0x45, 0xab,
	0x84, 0x62, 0x90, 0x61, 0x5b, 0x34, 0xdc, 0x5c, 0x8c, 0xde, 0x43, 0x77, 0x9e, 0x16, 0xf2, 0x2a,
	0xc2, 0x92, 0x6a, 0x7a, 0xbd, 0x49, 0x30, 0x36, 0xdb, 0x34, 0x2e, 0xb7, 0x69, 0xfc, 0xb5, 0xdc,
	0xa6, 0x59, 0xd5, 0x1c, 0x7e, 0x06, 0xa4, 0x68, 0x4e, 0xf9, 0x9a, 0xc9, 0xe2, 0xf6, 0x21, 0x86,
	0x3e, 0x74, 0x16, 0xa6, 0xd3, 0xf2, 0x2b, 0xc3, 0xc9, 0xaf, 0x7d, 0x68, 0x5f, 0xe4, 0xf8, 0x86,
	0xa1, 0x0f, 0x00, 0xd5, 0x8e, 0xa0, 0xc0, 0xf9, 0x60, 0x6b, 0x71, 0x82, 0x67, 0xae, 0x56, 0x2d,
	0x73, 0xb8, 0xa7, 0xc4, 0x38, 0x9b, 0xa3, 0x17, 0xae, 0xe7, 0x7f, 0xeb, 0x07, 0x7d, 0x57, 0x32,
	0x3f, 0x85, 0x3d, 0xf4, 0x0e, 0x7a, 0xa7, 0x98, 0x91, 0x8c, 0x6a, 0x43, 0xa3, 0xe7, 0xae, 0xa1,
	0x6e, 0xf0, 0x9d, 0xdf, 0x75, 0x23, 0x6e, 0xad, 0x83, 0x8e, 0x2a, 0xc6, 0x1b, 0x66, 0xda, 0xf1,
	0xdd, 0x47, 0x38, 0xbc, 0xcc, 0x89, 0x55, 0x65, 0xaf, 0x10, 0xbd, 0x74, 0x6d, 0xdb, 0x17, 0xbb,
	0x03, 0xe3, 0x0c, 0x06, 0x15, 0x46, 0xdd, 0x2d, 0x68, 0xb4, 0x01, 0xb4, 0xc3, 0x48, 0xdb, 0x68,
	0x73, 0x4f, 0xbf, 0xf6, 0xdb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x05, 0x88, 0x80, 0x63,
	0x05, 0x00, 0x00,
}
