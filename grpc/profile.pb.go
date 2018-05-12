// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

/*
Package grpc is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
	UpdateProfile
	DeleteProfile
	FindProfile
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

type UpdateProfile struct {
	Id      string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Profile *Profile `protobuf:"bytes,2,opt,name=profile" json:"profile,omitempty"`
}

func (m *UpdateProfile) Reset()                    { *m = UpdateProfile{} }
func (m *UpdateProfile) String() string            { return proto.CompactTextString(m) }
func (*UpdateProfile) ProtoMessage()               {}
func (*UpdateProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpdateProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateProfile) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

type DeleteProfile struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteProfile) Reset()                    { *m = DeleteProfile{} }
func (m *DeleteProfile) String() string            { return proto.CompactTextString(m) }
func (*DeleteProfile) ProtoMessage()               {}
func (*DeleteProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type FindProfile struct {
	Offset  int32    `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit   int32    `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Page    int32    `protobuf:"varint,3,opt,name=page" json:"page,omitempty"`
	OrderBy string   `protobuf:"bytes,4,opt,name=orderBy" json:"orderBy,omitempty"`
	Order   string   `protobuf:"bytes,5,opt,name=order" json:"order,omitempty"`
	Query   *Profile `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *FindProfile) Reset()                    { *m = FindProfile{} }
func (m *FindProfile) String() string            { return proto.CompactTextString(m) }
func (*FindProfile) ProtoMessage()               {}
func (*FindProfile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FindProfile) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *FindProfile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FindProfile) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *FindProfile) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

func (m *FindProfile) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *FindProfile) GetQuery() *Profile {
	if m != nil {
		return m.Query
	}
	return nil
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
func (*Profiles) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	proto.RegisterType((*UpdateProfile)(nil), "grpc.UpdateProfile")
	proto.RegisterType((*DeleteProfile)(nil), "grpc.DeleteProfile")
	proto.RegisterType((*FindProfile)(nil), "grpc.FindProfile")
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
	Find(ctx context.Context, in *FindProfile, opts ...grpc1.CallOption) (*Profiles, error)
	Update(ctx context.Context, in *UpdateProfile, opts ...grpc1.CallOption) (*Profile, error)
	DeleteById(ctx context.Context, in *DeleteProfile, opts ...grpc1.CallOption) (*Profile, error)
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

func (c *profileServiceClient) Find(ctx context.Context, in *FindProfile, opts ...grpc1.CallOption) (*Profiles, error) {
	out := new(Profiles)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/Find", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) Update(ctx context.Context, in *UpdateProfile, opts ...grpc1.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc1.Invoke(ctx, "/grpc.ProfileService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) DeleteById(ctx context.Context, in *DeleteProfile, opts ...grpc1.CallOption) (*Profile, error) {
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
	Find(context.Context, *FindProfile) (*Profiles, error)
	Update(context.Context, *UpdateProfile) (*Profile, error)
	DeleteById(context.Context, *DeleteProfile) (*Profile, error)
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

func _ProfileService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProfile)
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
		return srv.(ProfileServiceServer).Find(ctx, req.(*FindProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfile)
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
		return srv.(ProfileServiceServer).Update(ctx, req.(*UpdateProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc1.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfile)
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
		return srv.(ProfileServiceServer).DeleteById(ctx, req.(*DeleteProfile))
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
			MethodName: "Find",
			Handler:    _ProfileService_Find_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProfileService_Update_Handler,
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
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x51, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0x34, 0xeb, 0xa4, 0x53, 0x25, 0x12, 0x43, 0x55, 0xad, 0xf2, 0x00, 0x95, 0x79,
	0xa0, 0x08, 0x29, 0x0f, 0xe1, 0x06, 0x05, 0x21, 0xfa, 0x86, 0x8c, 0x38, 0xc0, 0x26, 0x3b, 0x2e,
	0x2b, 0x39, 0xb1, 0xd9, 0x5d, 0x23, 0xf9, 0x10, 0xdc, 0x82, 0xdb, 0x70, 0x12, 0x6e, 0x81, 0x76,
	0x76, 0x5d, 0xd5, 0xb2, 0x78, 0xe3, 0x6d, 0xfe, 0xf9, 0x66, 0x26, 0x33, 0x7f, 0xd6, 0xb0, 0x6a,
	0x6d, 0x53, 0x99, 0x9a, 0xb6, 0xad, 0x6d, 0x7c, 0x83, 0xf3, 0x07, 0xdb, 0x1e, 0x8a, 0x3f, 0x19,
	0x2c, 0x3e, 0xc7, 0x3c, 0xbe, 0x00, 0xa8, 0x8c, 0xa5, 0xbd, 0x72, 0x74, 0xaf, 0xe5, 0xec, 0x26,
	0xbb, 0xbd, 0x28, 0x9f, 0x64, 0x10, 0x61, 0x7e, 0x52, 0x47, 0x92, 0xe7, 0x4c, 0x38, 0xc6, 0x2b,
	0x10, 0xaa, 0x36, 0xca, 0xc9, 0x39, 0x27, 0xa3, 0xc0, 0x0d, 0x2c, 0x3b, 0x47, 0x96, 0xab, 0x05,
	0x83, 0x47, 0x1d, 0x58, 0xad, 0x9c, 0x67, 0x96, 0x47, 0x36, 0xe8, 0x30, 0x8d, 0x8e, 0xca, 0xd4,
	0x72, 0x11, 0xa7, 0xb1, 0x08, 0x1d, 0x7b, 0x63, 0xfd, 0x37, 0xad, 0x7a, 0xb9, 0x8c, 0x1d, 0x83,
	0xc6, 0x6b, 0xc8, 0x9d, 0x57, 0xbe, 0x73, 0xf2, 0x82, 0x49, 0x52, 0x61, 0x57, 0xdb, 0xd4, 0x24,
	0x21, 0xee, 0x1a, 0xe2, 0xe2, 0x13, 0xac, 0xbe, 0xb6, 0x5a, 0x79, 0x1a, 0x0e, 0x5e, 0xc3, 0xcc,
	0x68, 0x99, 0x71, 0xc9, 0xcc, 0x68, 0x7c, 0x0d, 0x8b, 0xe4, 0x11, 0x5f, 0x7f, 0xb9, 0x5b, 0x6d,
	0x83, 0x49, 0xdb, 0x54, 0x5f, 0x0e, 0xb4, 0x78, 0x09, 0xab, 0x0f, 0x54, 0xd3, 0x3f, 0x27, 0x15,
	0xbf, 0x32, 0xb8, 0xfc, 0x68, 0x4e, 0x7a, 0xe0, 0xd7, 0x90, 0x37, 0x55, 0xe5, 0xc8, 0x73, 0x8d,
	0x28, 0x93, 0x0a, 0x07, 0xd7, 0xe6, 0x68, 0x3c, 0xff, 0x9e, 0x28, 0xa3, 0x08, 0xcb, 0xb7, 0xea,
	0x21, 0x1a, 0x2d, 0x4a, 0x8e, 0x51, 0xc2, 0xa2, 0xb1, 0x9a, 0xec, 0x5d, 0x9f, 0xac, 0x1e, 0x64,
	0x98, 0xc1, 0x61, 0x72, 0x3a, 0x0a, 0x7c, 0x05, 0xe2, 0x7b, 0x47, 0xb6, 0x67, 0x8f, 0x27, 0x97,
	0x44, 0x56, 0xfc, 0xcc, 0x60, 0x99, 0x52, 0xee, 0x3f, 0xec, 0x78, 0x05, 0xe2, 0xd0, 0x74, 0x27,
	0xcf, 0x1b, 0x8a, 0x32, 0x0a, 0x7c, 0x03, 0xcb, 0xe4, 0x9b, 0x93, 0xe2, 0xe6, 0x7c, 0xba, 0xcc,
	0x23, 0xde, 0xfd, 0xce, 0x60, 0x9d, 0xb2, 0x5f, 0xc8, 0xfe, 0x30, 0x07, 0xc2, 0x5b, 0xc8, 0xdf,
	0x5b, 0x52, 0x9e, 0x70, 0xdc, 0xb5, 0x19, 0xcb, 0xe2, 0x0c, 0xdf, 0xc2, 0x3c, 0x58, 0x8e, 0xcf,
	0x22, 0x78, 0x62, 0xff, 0x66, 0x3d, 0xaa, 0x75, 0xc5, 0x19, 0x6e, 0x21, 0x8f, 0x6f, 0x01, 0x9f,
	0x47, 0x36, 0x7a, 0x19, 0xd3, 0xe1, 0x3b, 0x80, 0xf8, 0x8f, 0xdf, 0xf5, 0xf7, 0x7a, 0xe8, 0x19,
	0xbd, 0x81, 0x49, 0xcf, 0x3e, 0xe7, 0x0f, 0xed, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9,
	0xf0, 0x79, 0xc7, 0x79, 0x03, 0x00, 0x00,
}
