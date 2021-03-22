// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: bothway-stream.proto

package gen

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

// DemoParma 示例参数
type DemoParma struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *DemoParma) Reset() {
	*x = DemoParma{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoParma) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoParma) ProtoMessage() {}

func (x *DemoParma) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoParma.ProtoReflect.Descriptor instead.
func (*DemoParma) Descriptor() ([]byte, []int) {
	return file_client_stream_proto_rawDescGZIP(), []int{0}
}

func (x *DemoParma) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DemoResponse 示例响应
type DemoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *DemoResponse) Reset() {
	*x = DemoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_stream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DemoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DemoResponse) ProtoMessage() {}

func (x *DemoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_stream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DemoResponse.ProtoReflect.Descriptor instead.
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return file_client_stream_proto_rawDescGZIP(), []int{1}
}

func (x *DemoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_client_stream_proto protoreflect.FileDescriptor

var file_client_stream_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x22, 0x1f, 0x0a, 0x09, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x61, 0x72, 0x6d,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x44, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x5d, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x0a, 0x44, 0x65, 0x6d, 0x6f, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65, 0x6d, 0x6f, 0x50, 0x61, 0x72, 0x6d, 0x61, 0x1a, 0x1b, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x44, 0x65,
	0x6d, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x42, 0x34,
	0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x79,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_stream_proto_rawDescOnce sync.Once
	file_client_stream_proto_rawDescData = file_client_stream_proto_rawDesc
)

func file_client_stream_proto_rawDescGZIP() []byte {
	file_client_stream_proto_rawDescOnce.Do(func() {
		file_client_stream_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_stream_proto_rawDescData)
	})
	return file_client_stream_proto_rawDescData
}

var file_client_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_stream_proto_goTypes = []interface{}{
	(*DemoParma)(nil),    // 0: client_stream.DemoParma
	(*DemoResponse)(nil), // 1: client_stream.DemoResponse
}
var file_client_stream_proto_depIdxs = []int32{
	0, // 0: client_stream.ClientStreamServer.DemoMethod:input_type -> client_stream.DemoParma
	1, // 1: client_stream.ClientStreamServer.DemoMethod:output_type -> client_stream.DemoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_stream_proto_init() }
func file_client_stream_proto_init() {
	if File_client_stream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_stream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoParma); i {
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
		file_client_stream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DemoResponse); i {
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
			RawDescriptor: file_client_stream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_stream_proto_goTypes,
		DependencyIndexes: file_client_stream_proto_depIdxs,
		MessageInfos:      file_client_stream_proto_msgTypes,
	}.Build()
	File_client_stream_proto = out.File
	file_client_stream_proto_rawDesc = nil
	file_client_stream_proto_goTypes = nil
	file_client_stream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientStreamServerClient is the client API for ClientStreamServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientStreamServerClient interface {
	// 请注意这个方法,请求入参有一个 stream 关键字
	// 注意这里我们只标记了请求是流式的,但是响应还是正常形式的
	DemoMethod(ctx context.Context, opts ...grpc.CallOption) (ClientStreamServer_DemoMethodClient, error)
}

type clientStreamServerClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamServerClient(cc grpc.ClientConnInterface) ClientStreamServerClient {
	return &clientStreamServerClient{cc}
}

func (c *clientStreamServerClient) DemoMethod(ctx context.Context, opts ...grpc.CallOption) (ClientStreamServer_DemoMethodClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientStreamServer_serviceDesc.Streams[0], "/client_stream.ClientStreamServer/DemoMethod", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamServerDemoMethodClient{stream}
	return x, nil
}

type ClientStreamServer_DemoMethodClient interface {
	Send(*DemoParma) error
	CloseAndRecv() (*DemoResponse, error)
	grpc.ClientStream
}

type clientStreamServerDemoMethodClient struct {
	grpc.ClientStream
}

func (x *clientStreamServerDemoMethodClient) Send(m *DemoParma) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamServerDemoMethodClient) CloseAndRecv() (*DemoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DemoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamServerServer is the server API for ClientStreamServer service.
type ClientStreamServerServer interface {
	// 请注意这个方法,请求入参有一个 stream 关键字
	// 注意这里我们只标记了请求是流式的,但是响应还是正常形式的
	DemoMethod(ClientStreamServer_DemoMethodServer) error
}

// UnimplementedClientStreamServerServer can be embedded to have forward compatible implementations.
type UnimplementedClientStreamServerServer struct {
}

func (*UnimplementedClientStreamServerServer) DemoMethod(ClientStreamServer_DemoMethodServer) error {
	return status.Errorf(codes.Unimplemented, "method DemoMethod not implemented")
}

func RegisterClientStreamServerServer(s *grpc.Server, srv ClientStreamServerServer) {
	s.RegisterService(&_ClientStreamServer_serviceDesc, srv)
}

func _ClientStreamServer_DemoMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamServerServer).DemoMethod(&clientStreamServerDemoMethodServer{stream})
}

type ClientStreamServer_DemoMethodServer interface {
	SendAndClose(*DemoResponse) error
	Recv() (*DemoParma, error)
	grpc.ServerStream
}

type clientStreamServerDemoMethodServer struct {
	grpc.ServerStream
}

func (x *clientStreamServerDemoMethodServer) SendAndClose(m *DemoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamServerDemoMethodServer) Recv() (*DemoParma, error) {
	m := new(DemoParma)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientStreamServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "client_stream.ClientStreamServer",
	HandlerType: (*ClientStreamServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DemoMethod",
			Handler:       _ClientStreamServer_DemoMethod_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "bothway-stream.proto",
}