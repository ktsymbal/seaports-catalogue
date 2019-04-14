// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Port struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float64 `protobuf:"fixed64,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{0}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []float64 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Ports struct {
	Ports                []*Port  `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ports) Reset()         { *m = Ports{} }
func (m *Ports) String() string { return proto.CompactTextString(m) }
func (*Ports) ProtoMessage()    {}
func (*Ports) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{1}
}

func (m *Ports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ports.Unmarshal(m, b)
}
func (m *Ports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ports.Marshal(b, m, deterministic)
}
func (m *Ports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ports.Merge(m, src)
}
func (m *Ports) XXX_Size() int {
	return xxx_messageInfo_Ports.Size(m)
}
func (m *Ports) XXX_DiscardUnknown() {
	xxx_messageInfo_Ports.DiscardUnknown(m)
}

var xxx_messageInfo_Ports proto.InternalMessageInfo

func (m *Ports) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ListRequest struct {
	Limit                uint64   `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{3}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Port)(nil), "proto.Port")
	proto.RegisterType((*Ports)(nil), "proto.Ports")
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor_10d86afa5a89ec9d) }

var fileDescriptor_10d86afa5a89ec9d = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xdf, 0x4a, 0xec, 0x30,
	0x10, 0xc6, 0xb7, 0xff, 0xb6, 0xdb, 0xe9, 0xe1, 0x70, 0x98, 0x23, 0x12, 0xf6, 0xaa, 0x16, 0x84,
	0xea, 0xc5, 0x0a, 0xbb, 0x97, 0x5e, 0x8a, 0x78, 0x23, 0x28, 0x01, 0x1f, 0xa0, 0xb6, 0x59, 0x09,
	0x6c, 0x9b, 0x9a, 0x64, 0x85, 0xfa, 0x10, 0xbe, 0xaa, 0xaf, 0x20, 0x49, 0xda, 0xb5, 0xea, 0x8d,
	0x57, 0x9d, 0xdf, 0x37, 0x93, 0x99, 0x6f, 0xa6, 0xf0, 0x4f, 0xb2, 0x4e, 0x28, 0xae, 0x85, 0xec,
	0x57, 0x9d, 0x14, 0x5a, 0x60, 0x64, 0x3f, 0xf9, 0x9b, 0x0f, 0xe1, 0xbd, 0x90, 0x1a, 0xff, 0x82,
	0xcf, 0x6b, 0xe2, 0x65, 0x5e, 0x91, 0x50, 0x9f, 0xd7, 0x88, 0x10, 0xb6, 0x65, 0xc3, 0x88, 0x6f,
	0x15, 0x1b, 0x1b, 0xad, 0xe2, 0xba, 0x27, 0x81, 0xd3, 0x4c, 0x8c, 0x04, 0xe2, 0x4a, 0xec, 0x5b,
	0x2d, 0x7b, 0x12, 0x5a, 0x79, 0x44, 0x3c, 0x82, 0xa8, 0xdc, 0xf1, 0x52, 0x91, 0x28, 0x0b, 0x8a,
	0x84, 0x3a, 0x30, 0xf5, 0x92, 0x3d, 0x71, 0xd1, 0x2a, 0x32, 0xb7, 0xfa, 0x88, 0x98, 0x41, 0x5a,
	0x09, 0x21, 0x6b, 0xde, 0x96, 0x9a, 0x29, 0x12, 0x67, 0x41, 0xe1, 0xd1, 0xa9, 0x84, 0x4b, 0x58,
	0x74, 0x52, 0xbc, 0xf0, 0xb6, 0x62, 0x64, 0x61, 0x87, 0x1d, 0xd8, 0xe4, 0x34, 0x6f, 0xd8, 0xab,
	0x68, 0x19, 0x49, 0x5c, 0x6e, 0x64, 0x3c, 0x86, 0xf9, 0xbe, 0xdd, 0x89, 0x4a, 0x11, 0xb0, 0x23,
	0x07, 0xb2, 0xfb, 0x88, 0x9a, 0x91, 0x74, 0xd8, 0x47, 0xd4, 0x2c, 0x3f, 0x87, 0xc8, 0xdc, 0x43,
	0xe1, 0x09, 0x44, 0x9d, 0x09, 0x88, 0x97, 0x05, 0x45, 0xba, 0x4e, 0xdd, 0xdd, 0x56, 0x26, 0x49,
	0x5d, 0x26, 0x8f, 0x21, 0xba, 0x6e, 0x3a, 0xdd, 0xe7, 0x97, 0x90, 0xde, 0x72, 0xa5, 0x29, 0x7b,
	0xde, 0x33, 0xa5, 0xcd, 0xe6, 0x3b, 0xde, 0x70, 0x6d, 0xcf, 0x19, 0x52, 0x07, 0xc6, 0x85, 0xd8,
	0x6e, 0x15, 0xd3, 0xf6, 0xa6, 0x21, 0x1d, 0x68, 0xfd, 0xee, 0x01, 0xd0, 0xc3, 0xef, 0xc1, 0x0b,
	0x48, 0x4c, 0x2f, 0x67, 0x02, 0x87, 0xa9, 0x93, 0xee, 0xcb, 0x3f, 0x13, 0x27, 0x2a, 0x9f, 0xe1,
	0x29, 0xc4, 0x37, 0xcc, 0xd6, 0xe3, 0xd4, 0xe4, 0x72, 0x0a, 0xf9, 0x0c, 0xcf, 0x00, 0xae, 0x24,
	0x2b, 0x35, 0xfb, 0x59, 0x39, 0x76, 0x74, 0xcb, 0xd8, 0xd2, 0x87, 0xae, 0xfe, 0x55, 0xe9, 0x06,
	0xfe, 0xbb, 0xae, 0x77, 0xf2, 0xf3, 0x89, 0xc2, 0x2f, 0x1e, 0xbf, 0x3f, 0x7a, 0x9c, 0x5b, 0xdc,
	0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x55, 0xc7, 0xaf, 0x07, 0x96, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RepositoryClient interface {
	ListPorts(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Ports, error)
	GetPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
	CreatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error)
	UpdatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error)
	CreateOrUpdatePorts(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Empty, error)
}

type repositoryClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryClient(cc *grpc.ClientConn) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) ListPorts(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*Ports, error) {
	out := new(Ports)
	err := c.cc.Invoke(ctx, "/proto.Repository/ListPorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) GetPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/proto.Repository/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) CreatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Repository/CreatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) UpdatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Repository/UpdatePort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) CreateOrUpdatePorts(ctx context.Context, in *Ports, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Repository/CreateOrUpdatePorts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
type RepositoryServer interface {
	ListPorts(context.Context, *ListRequest) (*Ports, error)
	GetPort(context.Context, *Port) (*Port, error)
	CreatePort(context.Context, *Port) (*Empty, error)
	UpdatePort(context.Context, *Port) (*Empty, error)
	CreateOrUpdatePorts(context.Context, *Ports) (*Empty, error)
}

func RegisterRepositoryServer(s *grpc.Server, srv RepositoryServer) {
	s.RegisterService(&_Repository_serviceDesc, srv)
}

func _Repository_ListPorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).ListPorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repository/ListPorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).ListPorts(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repository/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).GetPort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_CreatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).CreatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repository/CreatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).CreatePort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_UpdatePort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdatePort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repository/UpdatePort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdatePort(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_CreateOrUpdatePorts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ports)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).CreateOrUpdatePorts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Repository/CreateOrUpdatePorts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).CreateOrUpdatePorts(ctx, req.(*Ports))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPorts",
			Handler:    _Repository_ListPorts_Handler,
		},
		{
			MethodName: "GetPort",
			Handler:    _Repository_GetPort_Handler,
		},
		{
			MethodName: "CreatePort",
			Handler:    _Repository_CreatePort_Handler,
		},
		{
			MethodName: "UpdatePort",
			Handler:    _Repository_UpdatePort_Handler,
		},
		{
			MethodName: "CreateOrUpdatePorts",
			Handler:    _Repository_CreateOrUpdatePorts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repository.proto",
}
