// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: cmdService.proto

package pbs

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{0}
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{1}
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AccessAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adddr string `protobuf:"bytes,1,opt,name=Adddr,proto3" json:"Adddr,omitempty"`
	Op    int32  `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
}

func (x *AccessAddress) Reset() {
	*x = AccessAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessAddress) ProtoMessage() {}

func (x *AccessAddress) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessAddress.ProtoReflect.Descriptor instead.
func (*AccessAddress) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{2}
}

func (x *AccessAddress) GetAdddr() string {
	if x != nil {
		return x.Adddr
	}
	return ""
}

func (x *AccessAddress) GetOp() int32 {
	if x != nil {
		return x.Op
	}
	return 0
}

var File_cmdService_proto protoreflect.FileDescriptor

var file_cmdService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x35, 0x0a, 0x0d, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x41, 0x64, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x41, 0x64, 0x64,
	0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x6f, 0x70, 0x32, 0x80, 0x01, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x77, 0x41, 0x6c, 0x6c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x07, 0x43, 0x68, 0x67, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x13,
	0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x70, 0x62, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdService_proto_rawDescOnce sync.Once
	file_cmdService_proto_rawDescData = file_cmdService_proto_rawDesc
)

func file_cmdService_proto_rawDescGZIP() []byte {
	file_cmdService_proto_rawDescOnce.Do(func() {
		file_cmdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdService_proto_rawDescData)
	})
	return file_cmdService_proto_rawDescData
}

var file_cmdService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmdService_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),   // 0: pbs.EmptyRequest
	(*CommonResponse)(nil), // 1: pbs.CommonResponse
	(*AccessAddress)(nil),  // 2: pbs.AccessAddress
}
var file_cmdService_proto_depIdxs = []int32{
	0, // 0: pbs.CmdService.ShowAllAdminUser:input_type -> pbs.EmptyRequest
	2, // 1: pbs.CmdService.ChgUser:input_type -> pbs.AccessAddress
	1, // 2: pbs.CmdService.ShowAllAdminUser:output_type -> pbs.CommonResponse
	1, // 3: pbs.CmdService.ChgUser:output_type -> pbs.CommonResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmdService_proto_init() }
func file_cmdService_proto_init() {
	if File_cmdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cmdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessAddress); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmdService_proto_goTypes,
		DependencyIndexes: file_cmdService_proto_depIdxs,
		MessageInfos:      file_cmdService_proto_msgTypes,
	}.Build()
	File_cmdService_proto = out.File
	file_cmdService_proto_rawDesc = nil
	file_cmdService_proto_goTypes = nil
	file_cmdService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CmdServiceClient is the client API for CmdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CmdServiceClient interface {
	ShowAllAdminUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	ChgUser(ctx context.Context, in *AccessAddress, opts ...grpc.CallOption) (*CommonResponse, error)
}

type cmdServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCmdServiceClient(cc grpc.ClientConnInterface) CmdServiceClient {
	return &cmdServiceClient{cc}
}

func (c *cmdServiceClient) ShowAllAdminUser(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/ShowAllAdminUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cmdServiceClient) ChgUser(ctx context.Context, in *AccessAddress, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, "/pbs.CmdService/ChgUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CmdServiceServer is the server API for CmdService service.
type CmdServiceServer interface {
	ShowAllAdminUser(context.Context, *EmptyRequest) (*CommonResponse, error)
	ChgUser(context.Context, *AccessAddress) (*CommonResponse, error)
}

// UnimplementedCmdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCmdServiceServer struct {
}

func (*UnimplementedCmdServiceServer) ShowAllAdminUser(context.Context, *EmptyRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowAllAdminUser not implemented")
}
func (*UnimplementedCmdServiceServer) ChgUser(context.Context, *AccessAddress) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChgUser not implemented")
}

func RegisterCmdServiceServer(s *grpc.Server, srv CmdServiceServer) {
	s.RegisterService(&_CmdService_serviceDesc, srv)
}

func _CmdService_ShowAllAdminUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ShowAllAdminUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/ShowAllAdminUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ShowAllAdminUser(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CmdService_ChgUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CmdServiceServer).ChgUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbs.CmdService/ChgUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CmdServiceServer).ChgUser(ctx, req.(*AccessAddress))
	}
	return interceptor(ctx, in, info, handler)
}

var _CmdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pbs.CmdService",
	HandlerType: (*CmdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShowAllAdminUser",
			Handler:    _CmdService_ShowAllAdminUser_Handler,
		},
		{
			MethodName: "ChgUser",
			Handler:    _CmdService_ChgUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmdService.proto",
}
