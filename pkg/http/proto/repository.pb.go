// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repository.proto

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_10d86afa5a89ec9d, []int{1}
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
	return fileDescriptor_10d86afa5a89ec9d, []int{2}
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
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
}

func init() { proto.RegisterFile("repository.proto", fileDescriptor_10d86afa5a89ec9d) }

var fileDescriptor_10d86afa5a89ec9d = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0x49, 0x93, 0xa6, 0x99, 0xfc, 0x11, 0x19, 0x45, 0x96, 0x9e, 0x42, 0x40, 0x88,
	0x97, 0x52, 0xd4, 0x9b, 0x47, 0x11, 0x2f, 0x82, 0x12, 0xf0, 0x01, 0x62, 0x32, 0x95, 0x85, 0x26,
	0x13, 0x77, 0xb7, 0x42, 0x7c, 0x08, 0xdf, 0xd5, 0x37, 0x90, 0xdd, 0xa4, 0xb5, 0x28, 0x88, 0xa7,
	0xec, 0xef, 0x9b, 0x6f, 0x66, 0x98, 0x2f, 0x70, 0xa8, 0xa8, 0x63, 0x2d, 0x0d, 0xab, 0x7e, 0xd1,
	0x29, 0x36, 0x8c, 0xa1, 0xfb, 0x64, 0xef, 0x3e, 0x04, 0x0f, 0xac, 0x0c, 0x1e, 0x80, 0x2f, 0x6b,
	0xe1, 0xa5, 0x5e, 0x1e, 0x17, 0xbe, 0xac, 0x11, 0x21, 0x68, 0xcb, 0x86, 0x84, 0xef, 0x14, 0xf7,
	0xb6, 0x5a, 0x25, 0x4d, 0x2f, 0x26, 0x83, 0x66, 0xdf, 0x28, 0x20, 0xaa, 0x78, 0xd3, 0x1a, 0xd5,
	0x8b, 0xc0, 0xc9, 0x5b, 0xc4, 0x63, 0x08, 0xcb, 0xb5, 0x2c, 0xb5, 0x08, 0xd3, 0x49, 0x1e, 0x17,
	0x03, 0x58, 0xbf, 0xa2, 0x67, 0xc9, 0xad, 0x16, 0x53, 0xa7, 0x6f, 0x11, 0x53, 0x48, 0x2a, 0x66,
	0x55, 0xcb, 0xb6, 0x34, 0xa4, 0x45, 0x94, 0x4e, 0x72, 0xaf, 0xd8, 0x97, 0x70, 0x0e, 0xb3, 0x4e,
	0xf1, 0xab, 0x6c, 0x2b, 0x12, 0x33, 0xb7, 0x6c, 0xc7, 0xb6, 0x66, 0x64, 0x43, 0x6f, 0xdc, 0x92,
	0x88, 0x87, 0xda, 0x96, 0xf1, 0x04, 0xa6, 0x9b, 0x76, 0xcd, 0x95, 0x16, 0xe0, 0x56, 0x8e, 0xe4,
	0xee, 0xe1, 0x9a, 0x44, 0x32, 0xde, 0xc3, 0x35, 0x65, 0x11, 0x84, 0x37, 0x4d, 0x67, 0xfa, 0xec,
	0x0a, 0x92, 0x3b, 0xa9, 0x4d, 0x41, 0x2f, 0x1b, 0xd2, 0xc6, 0x5e, 0xb3, 0x96, 0x8d, 0x34, 0x2e,
	0xa2, 0xa0, 0x18, 0xc0, 0x4e, 0xe6, 0xd5, 0x4a, 0x93, 0x71, 0x39, 0x05, 0xc5, 0x48, 0xe7, 0x1f,
	0x1e, 0x40, 0xb1, 0x8b, 0x1c, 0x97, 0x10, 0xdb, 0x59, 0x36, 0x68, 0x8d, 0x38, 0xfc, 0x81, 0xc5,
	0xde, 0xf4, 0x79, 0x32, 0x6a, 0xd6, 0x91, 0xfd, 0x5b, 0x7a, 0x78, 0x0a, 0xd1, 0x2d, 0xb9, 0x06,
	0xdc, 0xaf, 0x7d, 0x33, 0xe2, 0x19, 0xc0, 0xb5, 0xa2, 0xd2, 0xd0, 0x4f, 0xe7, 0xff, 0x11, 0x86,
	0x6b, 0x9c, 0xf5, 0xb1, 0xab, 0xff, 0x64, 0xbd, 0x84, 0xa3, 0x61, 0xea, 0xbd, 0xfa, 0x6a, 0xd1,
	0xbf, 0xf6, 0xe4, 0xde, 0xd3, 0xd4, 0x09, 0x17, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x9b,
	0xb9, 0xd9, 0x6c, 0x02, 0x00, 0x00,
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
	ListPorts(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Repository_ListPortsClient, error)
	GetPort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
	CreatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error)
	UpdatePort(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Empty, error)
	CreateOrUpdatePorts(ctx context.Context, opts ...grpc.CallOption) (Repository_CreateOrUpdatePortsClient, error)
}

type repositoryClient struct {
	cc *grpc.ClientConn
}

func NewRepositoryClient(cc *grpc.ClientConn) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) ListPorts(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (Repository_ListPortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Repository_serviceDesc.Streams[0], "/proto.Repository/ListPorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryListPortsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Repository_ListPortsClient interface {
	Recv() (*Port, error)
	grpc.ClientStream
}

type repositoryListPortsClient struct {
	grpc.ClientStream
}

func (x *repositoryListPortsClient) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
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

func (c *repositoryClient) CreateOrUpdatePorts(ctx context.Context, opts ...grpc.CallOption) (Repository_CreateOrUpdatePortsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Repository_serviceDesc.Streams[1], "/proto.Repository/CreateOrUpdatePorts", opts...)
	if err != nil {
		return nil, err
	}
	x := &repositoryCreateOrUpdatePortsClient{stream}
	return x, nil
}

type Repository_CreateOrUpdatePortsClient interface {
	Send(*Port) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type repositoryCreateOrUpdatePortsClient struct {
	grpc.ClientStream
}

func (x *repositoryCreateOrUpdatePortsClient) Send(m *Port) error {
	return x.ClientStream.SendMsg(m)
}

func (x *repositoryCreateOrUpdatePortsClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RepositoryServer is the server API for Repository service.
type RepositoryServer interface {
	ListPorts(*ListRequest, Repository_ListPortsServer) error
	GetPort(context.Context, *Port) (*Port, error)
	CreatePort(context.Context, *Port) (*Empty, error)
	UpdatePort(context.Context, *Port) (*Empty, error)
	CreateOrUpdatePorts(Repository_CreateOrUpdatePortsServer) error
}

func RegisterRepositoryServer(s *grpc.Server, srv RepositoryServer) {
	s.RegisterService(&_Repository_serviceDesc, srv)
}

func _Repository_ListPorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RepositoryServer).ListPorts(m, &repositoryListPortsServer{stream})
}

type Repository_ListPortsServer interface {
	Send(*Port) error
	grpc.ServerStream
}

type repositoryListPortsServer struct {
	grpc.ServerStream
}

func (x *repositoryListPortsServer) Send(m *Port) error {
	return x.ServerStream.SendMsg(m)
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

func _Repository_CreateOrUpdatePorts_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RepositoryServer).CreateOrUpdatePorts(&repositoryCreateOrUpdatePortsServer{stream})
}

type Repository_CreateOrUpdatePortsServer interface {
	SendAndClose(*Empty) error
	Recv() (*Port, error)
	grpc.ServerStream
}

type repositoryCreateOrUpdatePortsServer struct {
	grpc.ServerStream
}

func (x *repositoryCreateOrUpdatePortsServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *repositoryCreateOrUpdatePortsServer) Recv() (*Port, error) {
	m := new(Port)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Repository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
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
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPorts",
			Handler:       _Repository_ListPorts_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "CreateOrUpdatePorts",
			Handler:       _Repository_CreateOrUpdatePorts_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "repository.proto",
}