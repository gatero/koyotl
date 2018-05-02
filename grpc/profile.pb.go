// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
	Profiles
*/
package grpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc1 "google.golang.org/grpc"
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

type Profile struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FirebaseId string `protobuf:"bytes,2,opt,name=firebaseId" json:"firebaseId,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Alias      string `protobuf:"bytes,4,opt,name=alias" json:"alias,omitempty"`
	Username   string `protobuf:"bytes,5,opt,name=username" json:"username,omitempty"`
	Lastname   string `protobuf:"bytes,6,opt,name=lastname" json:"lastname,omitempty"`
	Email      string `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	Birthday   string `protobuf:"bytes,8,opt,name=birthday" json:"birthday,omitempty"`
	Status     string `protobuf:"bytes,9,opt,name=status" json:"status,omitempty"`
	Role       string `protobuf:"bytes,10,opt,name=role" json:"role,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Profile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Profile) GetFirebaseId() string {
	if m != nil {
		return m.FirebaseId
	}
	return ""
}

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *Profile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Profile) GetLastname() string {
	if m != nil {
		return m.Lastname
	}
	return ""
}

func (m *Profile) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Profile) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *Profile) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Profile) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type Profiles struct {
	Offset   int32      `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit    int32      `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Page     int32      `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	Count    int32      `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	Profiles []*Profile `protobuf:"bytes,5,rep,name=profiles" json:"profiles,omitempty"`
}

func (m *Profiles) Reset()                    { *m = Profiles{} }
func (m *Profiles) String() string            { return proto.CompactTextString(m) }
func (*Profiles) ProtoMessage()               {}
func (*Profiles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Profiles) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Profiles) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *Profiles) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Profiles) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Profiles) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "grpc.Profile")
	proto.RegisterType((*Profiles)(nil), "grpc.Profiles")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc1.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc1.SupportPackageIsVersion4

// Client API for ProfileService service

type ProfileServiceClient interface {
	Create(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error)
	Update(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error)
	Find(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profiles, error)
	FindById(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error)
	DeleteById(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error)
}

type profileServiceClient struct {
	cc *grpc1.ClientConn
}

func NewProfileServiceClient(cc *grpc1.ClientConn) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) Create(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Update(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Find(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profiles, error) {
	out := new(Profiles)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) FindById(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/FindById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) DeleteById(ctx context.Context, in *Profile, opts ...grpc1.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/DeleteById", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProfileService service

type ProfileServiceServer interface {
	Create(context.Context, *Profile) (*Profile, error)
	Update(context.Context, *Profile) (*Profile, error)
	Find(context.Context, *Profile) (*Profiles, error)
	FindById(context.Context, *Profile) (*Profile, error)
	DeleteById(context.Context, *Profile) (*Profile, error)
}

func RegisterProfileServiceServer(s *grpc1.Server, srv ProfileServiceServer) {
	s.RegisterService(&_ProfileService_serviceDesc, srv)
}

func _ProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Create(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Create(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Update(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Update(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).Find(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProfileService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).Find(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_FindById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).FindById(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProfileService/FindById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).FindById(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).DeleteById(ctx, in)
	}
	info := &grpc1.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.ProfileService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).DeleteById(ctx, req.(*Profile))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfileService_serviceDesc = grpc1.ServiceDesc{
	ServiceName: "grpc.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc1.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProfileService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfileService_Update_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _ProfileService_Find_Handler,
		},
		{
			MethodName: "FindById",
			Handler:    _ProfileService_FindById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _ProfileService_DeleteById_Handler,
		},
	},
	Streams:  []grpc1.StreamDesc{},
	Metadata: "profile.proto",
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0x3a, 0x31,
	0x10, 0xc6, 0x59, 0xd8, 0x2e, 0xcb, 0xfc, 0x03, 0x87, 0xe6, 0x1f, 0xd3, 0x70, 0x30, 0x64, 0x2f,
	0xa2, 0x31, 0x1c, 0xf0, 0x0d, 0xd4, 0x98, 0x70, 0x33, 0x6b, 0x7c, 0x80, 0x42, 0x67, 0xb1, 0x49,
	0x61, 0x37, 0x6d, 0x31, 0xe1, 0x21, 0x7c, 0x4e, 0x9f, 0xc1, 0x9b, 0xe9, 0x74, 0x21, 0xe2, 0x45,
	0x6e, 0xfd, 0x7d, 0xdf, 0x37, 0xb3, 0x33, 0x93, 0x85, 0x61, 0x63, 0xeb, 0x4a, 0x1b, 0x9c, 0x35,
	0xb6, 0xf6, 0x35, 0x4f, 0xd7, 0xb6, 0x59, 0x15, 0x5f, 0x09, 0xf4, 0x9f, 0xa3, 0xce, 0x47, 0xd0,
	0xd5, 0x4a, 0x24, 0x93, 0x64, 0x3a, 0x28, 0xbb, 0x5a, 0xf1, 0x4b, 0x80, 0x4a, 0x5b, 0x5c, 0x4a,
	0x87, 0x0b, 0x25, 0xba, 0xa4, 0xff, 0x50, 0x38, 0x87, 0x74, 0x2b, 0x37, 0x28, 0x7a, 0xe4, 0xd0,
	0x9b, 0xff, 0x07, 0x26, 0x8d, 0x96, 0x4e, 0xa4, 0x24, 0x46, 0xe0, 0x63, 0xc8, 0x77, 0x0e, 0x2d,
	0xa5, 0x19, 0x19, 0x47, 0x0e, 0x9e, 0x91, 0xce, 0x93, 0x97, 0x45, 0xef, 0xc0, 0xa1, 0x1b, 0x6e,
	0xa4, 0x36, 0xa2, 0x1f, 0xbb, 0x11, 0x84, 0x8a, 0xa5, 0xb6, 0xfe, 0x4d, 0xc9, 0xbd, 0xc8, 0x63,
	0xc5, 0x81, 0xf9, 0x05, 0x64, 0xce, 0x4b, 0xbf, 0x73, 0x62, 0x40, 0x4e, 0x4b, 0x61, 0x56, 0x5b,
	0x1b, 0x14, 0x10, 0x67, 0x0d, 0xef, 0xe2, 0x23, 0x81, 0xbc, 0xdd, 0xdd, 0x85, 0xc2, 0xba, 0xaa,
	0x1c, 0x7a, 0x3a, 0x00, 0x2b, 0x5b, 0x0a, 0x23, 0x18, 0xbd, 0xd1, 0x9e, 0xf6, 0x67, 0x65, 0x84,
	0xd0, 0xae, 0x91, 0xeb, 0xb8, 0x3a, 0x2b, 0xe9, 0x1d, 0x92, 0xab, 0x7a, 0xb7, 0xf5, 0xb4, 0x3a,
	0x2b, 0x23, 0xf0, 0x6b, 0xc8, 0xdb, 0xbb, 0x3b, 0xc1, 0x26, 0xbd, 0xe9, 0xbf, 0xf9, 0x70, 0x16,
	0x2e, 0x3f, 0x6b, 0xbf, 0x5c, 0x1e, 0xed, 0xf9, 0x67, 0x02, 0xa3, 0x56, 0x7d, 0x41, 0xfb, 0xae,
	0x57, 0xc8, 0xa7, 0x90, 0x3d, 0x58, 0x94, 0x1e, 0xf9, 0x69, 0xd5, 0xf8, 0x14, 0x8b, 0x4e, 0x48,
	0xbe, 0x36, 0xea, 0x9c, 0xe4, 0x15, 0xa4, 0x4f, 0x7a, 0xab, 0x7e, 0xe7, 0x46, 0x27, 0xe8, 0x8a,
	0x0e, 0xbf, 0x81, 0x3c, 0x04, 0xef, 0xf7, 0x0b, 0xf5, 0x67, 0xd3, 0x5b, 0x80, 0x47, 0x34, 0xe8,
	0xf1, 0x9c, 0xf4, 0x32, 0xa3, 0x5f, 0xf0, 0xee, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x48, 0x07, 0x70,
	0xb1, 0x93, 0x02, 0x00, 0x00,
}
