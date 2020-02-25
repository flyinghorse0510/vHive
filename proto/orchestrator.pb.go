// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orchestrator.proto

package proto

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

type StartVMReq struct {
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartVMReq) Reset()         { *m = StartVMReq{} }
func (m *StartVMReq) String() string { return proto.CompactTextString(m) }
func (*StartVMReq) ProtoMessage()    {}
func (*StartVMReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b6e6782baaa298, []int{0}
}

func (m *StartVMReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartVMReq.Unmarshal(m, b)
}
func (m *StartVMReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartVMReq.Marshal(b, m, deterministic)
}
func (m *StartVMReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartVMReq.Merge(m, src)
}
func (m *StartVMReq) XXX_Size() int {
	return xxx_messageInfo_StartVMReq.Size(m)
}
func (m *StartVMReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartVMReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartVMReq proto.InternalMessageInfo

func (m *StartVMReq) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *StartVMReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StopVMsReq struct {
	AllVms               bool     `protobuf:"varint,1,opt,name=all_vms,json=allVms,proto3" json:"all_vms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopVMsReq) Reset()         { *m = StopVMsReq{} }
func (m *StopVMsReq) String() string { return proto.CompactTextString(m) }
func (*StopVMsReq) ProtoMessage()    {}
func (*StopVMsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b6e6782baaa298, []int{1}
}

func (m *StopVMsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopVMsReq.Unmarshal(m, b)
}
func (m *StopVMsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopVMsReq.Marshal(b, m, deterministic)
}
func (m *StopVMsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopVMsReq.Merge(m, src)
}
func (m *StopVMsReq) XXX_Size() int {
	return xxx_messageInfo_StopVMsReq.Size(m)
}
func (m *StopVMsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopVMsReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopVMsReq proto.InternalMessageInfo

func (m *StopVMsReq) GetAllVms() bool {
	if m != nil {
		return m.AllVms
	}
	return false
}

type StopSingleVMReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopSingleVMReq) Reset()         { *m = StopSingleVMReq{} }
func (m *StopSingleVMReq) String() string { return proto.CompactTextString(m) }
func (*StopSingleVMReq) ProtoMessage()    {}
func (*StopSingleVMReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b6e6782baaa298, []int{2}
}

func (m *StopSingleVMReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopSingleVMReq.Unmarshal(m, b)
}
func (m *StopSingleVMReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopSingleVMReq.Marshal(b, m, deterministic)
}
func (m *StopSingleVMReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopSingleVMReq.Merge(m, src)
}
func (m *StopSingleVMReq) XXX_Size() int {
	return xxx_messageInfo_StopSingleVMReq.Size(m)
}
func (m *StopSingleVMReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopSingleVMReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopSingleVMReq proto.InternalMessageInfo

func (m *StopSingleVMReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Status struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b6e6782baaa298, []int{3}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StartVMResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Profile              string   `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartVMResp) Reset()         { *m = StartVMResp{} }
func (m *StartVMResp) String() string { return proto.CompactTextString(m) }
func (*StartVMResp) ProtoMessage()    {}
func (*StartVMResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b6e6782baaa298, []int{4}
}

func (m *StartVMResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartVMResp.Unmarshal(m, b)
}
func (m *StartVMResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartVMResp.Marshal(b, m, deterministic)
}
func (m *StartVMResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartVMResp.Merge(m, src)
}
func (m *StartVMResp) XXX_Size() int {
	return xxx_messageInfo_StartVMResp.Size(m)
}
func (m *StartVMResp) XXX_DiscardUnknown() {
	xxx_messageInfo_StartVMResp.DiscardUnknown(m)
}

var xxx_messageInfo_StartVMResp proto.InternalMessageInfo

func (m *StartVMResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StartVMResp) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func init() {
	proto.RegisterType((*StartVMReq)(nil), "proto.StartVMReq")
	proto.RegisterType((*StopVMsReq)(nil), "proto.StopVMsReq")
	proto.RegisterType((*StopSingleVMReq)(nil), "proto.StopSingleVMReq")
	proto.RegisterType((*Status)(nil), "proto.Status")
	proto.RegisterType((*StartVMResp)(nil), "proto.StartVMResp")
}

func init() { proto.RegisterFile("orchestrator.proto", fileDescriptor_96b6e6782baaa298) }

var fileDescriptor_96b6e6782baaa298 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x6d, 0x61, 0xad, 0x3e, 0xa7, 0xb2, 0x20, 0x5a, 0x04, 0x41, 0x03, 0x82, 0x17, 0x7b,
	0x98, 0x82, 0x67, 0x77, 0x1f, 0x8e, 0x16, 0x7a, 0x95, 0xb8, 0xc5, 0x1a, 0x48, 0x4d, 0xcc, 0xcb,
	0xc4, 0x7f, 0xc8, 0xff, 0x53, 0x92, 0xac, 0x6b, 0x60, 0xec, 0x14, 0xde, 0xcb, 0xf7, 0xfb, 0xf2,
	0xde, 0x17, 0x20, 0xca, 0x2c, 0x3f, 0x39, 0x5a, 0xc3, 0xac, 0x32, 0xa5, 0x36, 0xca, 0x2a, 0x32,
	0xf2, 0x07, 0x9d, 0x02, 0xd4, 0x96, 0x19, 0xdb, 0xcc, 0x2b, 0xfe, 0x4d, 0xce, 0x61, 0x24, 0x3a,
	0xd6, 0xf2, 0x22, 0xb9, 0x49, 0xee, 0x8f, 0xaa, 0x50, 0x90, 0x53, 0x48, 0xc5, 0xaa, 0x48, 0x7d,
	0x2b, 0x15, 0x2b, 0x7a, 0xe7, 0x18, 0xa5, 0x9b, 0x39, 0x3a, 0xe6, 0x12, 0x72, 0x26, 0xe5, 0xdb,
	0x4f, 0x87, 0x9e, 0x3a, 0xac, 0x32, 0x26, 0x65, 0xd3, 0x21, 0xbd, 0x85, 0x33, 0x27, 0xab, 0xc5,
	0x57, 0x2b, 0x79, 0xf0, 0x0f, 0x4e, 0xc9, 0xd6, 0x89, 0x42, 0x56, 0x5b, 0x66, 0xd7, 0x48, 0x0a,
	0xc8, 0x3b, 0x8e, 0x38, 0xbc, 0xdd, 0x97, 0xf4, 0x05, 0x8e, 0xb7, 0x13, 0xa2, 0xde, 0x2f, 0x74,
	0x37, 0xda, 0xa8, 0x0f, 0x21, 0xf9, 0x66, 0xd6, 0xbe, 0x9c, 0xfe, 0x25, 0x30, 0x7e, 0x8d, 0x22,
	0x20, 0x0f, 0x90, 0x6f, 0x3c, 0xc9, 0x24, 0xe4, 0x51, 0x0e, 0x29, 0x5c, 0x9d, 0x0c, 0x2d, 0xbb,
	0x46, 0x7a, 0x10, 0xe4, 0x7e, 0xe1, 0x48, 0xde, 0x07, 0xb0, 0x2b, 0x7f, 0x86, 0x71, 0xbc, 0x38,
	0xb9, 0x88, 0x98, 0x28, 0x8d, 0x1d, 0x70, 0xf6, 0x04, 0xd7, 0x42, 0x95, 0xad, 0xd1, 0xcb, 0x92,
	0xff, 0xb2, 0x4e, 0x4b, 0x8e, 0x65, 0xfc, 0x75, 0xb3, 0x49, 0xbc, 0xc5, 0xc2, 0xc1, 0x8b, 0xe4,
	0x3d, 0xf3, 0x2e, 0x8f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x71, 0x2a, 0x19, 0xe6, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrchestratorClient interface {
	StartVM(ctx context.Context, in *StartVMReq, opts ...grpc.CallOption) (*Status, error)
	StopVMs(ctx context.Context, in *StopVMsReq, opts ...grpc.CallOption) (*Status, error)
	StopSingleVM(ctx context.Context, in *StopSingleVMReq, opts ...grpc.CallOption) (*Status, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) StartVM(ctx context.Context, in *StartVMReq, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.Orchestrator/StartVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) StopVMs(ctx context.Context, in *StopVMsReq, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.Orchestrator/StopVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) StopSingleVM(ctx context.Context, in *StopSingleVMReq, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/proto.Orchestrator/StopSingleVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
type OrchestratorServer interface {
	StartVM(context.Context, *StartVMReq) (*Status, error)
	StopVMs(context.Context, *StopVMsReq) (*Status, error)
	StopSingleVM(context.Context, *StopSingleVMReq) (*Status, error)
}

// UnimplementedOrchestratorServer can be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServer struct {
}

func (*UnimplementedOrchestratorServer) StartVM(ctx context.Context, req *StartVMReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVM not implemented")
}
func (*UnimplementedOrchestratorServer) StopVMs(ctx context.Context, req *StopVMsReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVMs not implemented")
}
func (*UnimplementedOrchestratorServer) StopSingleVM(ctx context.Context, req *StopSingleVMReq) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSingleVM not implemented")
}

func RegisterOrchestratorServer(s *grpc.Server, srv OrchestratorServer) {
	s.RegisterService(&_Orchestrator_serviceDesc, srv)
}

func _Orchestrator_StartVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartVMReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).StartVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Orchestrator/StartVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).StartVM(ctx, req.(*StartVMReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_StopVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVMsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).StopVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Orchestrator/StopVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).StopVMs(ctx, req.(*StopVMsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_StopSingleVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopSingleVMReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).StopSingleVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Orchestrator/StopSingleVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).StopSingleVM(ctx, req.(*StopSingleVMReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Orchestrator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartVM",
			Handler:    _Orchestrator_StartVM_Handler,
		},
		{
			MethodName: "StopVMs",
			Handler:    _Orchestrator_StopVMs_Handler,
		},
		{
			MethodName: "StopSingleVM",
			Handler:    _Orchestrator_StopSingleVM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orchestrator.proto",
}
