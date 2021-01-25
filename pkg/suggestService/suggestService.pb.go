// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: suggestService.proto

package suggestService

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SuggestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SuggestRequest) Reset() {
	*x = SuggestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suggestService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestRequest) ProtoMessage() {}

func (x *SuggestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_suggestService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestRequest.ProtoReflect.Descriptor instead.
func (*SuggestRequest) Descriptor() ([]byte, []int) {
	return file_suggestService_proto_rawDescGZIP(), []int{0}
}

func (x *SuggestRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type Suggest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Position uint32 `protobuf:"varint,2,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *Suggest) Reset() {
	*x = Suggest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suggestService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Suggest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Suggest) ProtoMessage() {}

func (x *Suggest) ProtoReflect() protoreflect.Message {
	mi := &file_suggestService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Suggest.ProtoReflect.Descriptor instead.
func (*Suggest) Descriptor() ([]byte, []int) {
	return file_suggestService_proto_rawDescGZIP(), []int{1}
}

func (x *Suggest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Suggest) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type SuggestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suggestions []*Suggest `protobuf:"bytes,1,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
}

func (x *SuggestResponse) Reset() {
	*x = SuggestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_suggestService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuggestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuggestResponse) ProtoMessage() {}

func (x *SuggestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_suggestService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuggestResponse.ProtoReflect.Descriptor instead.
func (*SuggestResponse) Descriptor() ([]byte, []int) {
	return file_suggestService_proto_rawDescGZIP(), []int{2}
}

func (x *SuggestResponse) GetSuggestions() []*Suggest {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

var File_suggestService_proto protoreflect.FileDescriptor

var file_suggestService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x39, 0x0a, 0x07,
	0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4c, 0x0a, 0x0f, 0x53, 0x75, 0x67, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x73, 0x75,
	0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x76, 0x0a, 0x0e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1e, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_suggestService_proto_rawDescOnce sync.Once
	file_suggestService_proto_rawDescData = file_suggestService_proto_rawDesc
)

func file_suggestService_proto_rawDescGZIP() []byte {
	file_suggestService_proto_rawDescOnce.Do(func() {
		file_suggestService_proto_rawDescData = protoimpl.X.CompressGZIP(file_suggestService_proto_rawDescData)
	})
	return file_suggestService_proto_rawDescData
}

var file_suggestService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_suggestService_proto_goTypes = []interface{}{
	(*SuggestRequest)(nil),  // 0: suggestService.SuggestRequest
	(*Suggest)(nil),         // 1: suggestService.Suggest
	(*SuggestResponse)(nil), // 2: suggestService.SuggestResponse
}
var file_suggestService_proto_depIdxs = []int32{
	1, // 0: suggestService.SuggestResponse.suggestions:type_name -> suggestService.Suggest
	0, // 1: suggestService.SuggestService.Query:input_type -> suggestService.SuggestRequest
	2, // 2: suggestService.SuggestService.Query:output_type -> suggestService.SuggestResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_suggestService_proto_init() }
func file_suggestService_proto_init() {
	if File_suggestService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_suggestService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestRequest); i {
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
		file_suggestService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Suggest); i {
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
		file_suggestService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuggestResponse); i {
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
			RawDescriptor: file_suggestService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_suggestService_proto_goTypes,
		DependencyIndexes: file_suggestService_proto_depIdxs,
		MessageInfos:      file_suggestService_proto_msgTypes,
	}.Build()
	File_suggestService_proto = out.File
	file_suggestService_proto_rawDesc = nil
	file_suggestService_proto_goTypes = nil
	file_suggestService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SuggestServiceClient is the client API for SuggestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SuggestServiceClient interface {
	Query(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error)
}

type suggestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuggestServiceClient(cc grpc.ClientConnInterface) SuggestServiceClient {
	return &suggestServiceClient{cc}
}

func (c *suggestServiceClient) Query(ctx context.Context, in *SuggestRequest, opts ...grpc.CallOption) (*SuggestResponse, error) {
	out := new(SuggestResponse)
	err := c.cc.Invoke(ctx, "/suggestService.SuggestService/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuggestServiceServer is the server API for SuggestService service.
type SuggestServiceServer interface {
	Query(context.Context, *SuggestRequest) (*SuggestResponse, error)
}

// UnimplementedSuggestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSuggestServiceServer struct {
}

func (*UnimplementedSuggestServiceServer) Query(context.Context, *SuggestRequest) (*SuggestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

func RegisterSuggestServiceServer(s *grpc.Server, srv SuggestServiceServer) {
	s.RegisterService(&_SuggestService_serviceDesc, srv)
}

func _SuggestService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuggestServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/suggestService.SuggestService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuggestServiceServer).Query(ctx, req.(*SuggestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SuggestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "suggestService.SuggestService",
	HandlerType: (*SuggestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _SuggestService_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "suggestService.proto",
}
