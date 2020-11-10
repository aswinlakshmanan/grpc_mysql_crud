// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: articles/blogpb/blog.proto

package articles

import (
	context "context"
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

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_blogpb_blog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_articles_blogpb_blog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_articles_blogpb_blog_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_blogpb_blog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_articles_blogpb_blog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleRequest.ProtoReflect.Descriptor instead.
func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_articles_blogpb_blog_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleRequest) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type CreateArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *CreateArticleResponse) Reset() {
	*x = CreateArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_blogpb_blog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleResponse) ProtoMessage() {}

func (x *CreateArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_articles_blogpb_blog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArticleResponse.ProtoReflect.Descriptor instead.
func (*CreateArticleResponse) Descriptor() ([]byte, []int) {
	return file_articles_blogpb_blog_proto_rawDescGZIP(), []int{2}
}

func (x *CreateArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type ListArticleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListArticleRequest) Reset() {
	*x = ListArticleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_blogpb_blog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleRequest) ProtoMessage() {}

func (x *ListArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_articles_blogpb_blog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleRequest.ProtoReflect.Descriptor instead.
func (*ListArticleRequest) Descriptor() ([]byte, []int) {
	return file_articles_blogpb_blog_proto_rawDescGZIP(), []int{3}
}

func (x *ListArticleRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type ListArticleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *Article `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *ListArticleResponse) Reset() {
	*x = ListArticleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_articles_blogpb_blog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArticleResponse) ProtoMessage() {}

func (x *ListArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_articles_blogpb_blog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArticleResponse.ProtoReflect.Descriptor instead.
func (*ListArticleResponse) Descriptor() ([]byte, []int) {
	return file_articles_blogpb_blog_proto_rawDescGZIP(), []int{4}
}

func (x *ListArticleResponse) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

var File_articles_blogpb_blog_proto protoreflect.FileDescriptor

var file_articles_blogpb_blog_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70,
	0x62, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x43, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x22, 0x44, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x22, 0x42, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x32, 0xad, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_articles_blogpb_blog_proto_rawDescOnce sync.Once
	file_articles_blogpb_blog_proto_rawDescData = file_articles_blogpb_blog_proto_rawDesc
)

func file_articles_blogpb_blog_proto_rawDescGZIP() []byte {
	file_articles_blogpb_blog_proto_rawDescOnce.Do(func() {
		file_articles_blogpb_blog_proto_rawDescData = protoimpl.X.CompressGZIP(file_articles_blogpb_blog_proto_rawDescData)
	})
	return file_articles_blogpb_blog_proto_rawDescData
}

var file_articles_blogpb_blog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_articles_blogpb_blog_proto_goTypes = []interface{}{
	(*Article)(nil),               // 0: articles.Article
	(*CreateArticleRequest)(nil),  // 1: articles.CreateArticleRequest
	(*CreateArticleResponse)(nil), // 2: articles.CreateArticleResponse
	(*ListArticleRequest)(nil),    // 3: articles.ListArticleRequest
	(*ListArticleResponse)(nil),   // 4: articles.ListArticleResponse
}
var file_articles_blogpb_blog_proto_depIdxs = []int32{
	0, // 0: articles.CreateArticleRequest.article:type_name -> articles.Article
	0, // 1: articles.CreateArticleResponse.article:type_name -> articles.Article
	0, // 2: articles.ListArticleResponse.article:type_name -> articles.Article
	1, // 3: articles.ServiceName.CreateArticle:input_type -> articles.CreateArticleRequest
	3, // 4: articles.ServiceName.ListArticle:input_type -> articles.ListArticleRequest
	2, // 5: articles.ServiceName.CreateArticle:output_type -> articles.CreateArticleResponse
	4, // 6: articles.ServiceName.ListArticle:output_type -> articles.ListArticleResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_articles_blogpb_blog_proto_init() }
func file_articles_blogpb_blog_proto_init() {
	if File_articles_blogpb_blog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_articles_blogpb_blog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_articles_blogpb_blog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleRequest); i {
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
		file_articles_blogpb_blog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArticleResponse); i {
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
		file_articles_blogpb_blog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleRequest); i {
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
		file_articles_blogpb_blog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArticleResponse); i {
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
			RawDescriptor: file_articles_blogpb_blog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_articles_blogpb_blog_proto_goTypes,
		DependencyIndexes: file_articles_blogpb_blog_proto_depIdxs,
		MessageInfos:      file_articles_blogpb_blog_proto_msgTypes,
	}.Build()
	File_articles_blogpb_blog_proto = out.File
	file_articles_blogpb_blog_proto_rawDesc = nil
	file_articles_blogpb_blog_proto_goTypes = nil
	file_articles_blogpb_blog_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceNameClient is the client API for ServiceName service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceNameClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error)
	ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (ServiceName_ListArticleClient, error)
}

type serviceNameClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceNameClient(cc grpc.ClientConnInterface) ServiceNameClient {
	return &serviceNameClient{cc}
}

func (c *serviceNameClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*CreateArticleResponse, error) {
	out := new(CreateArticleResponse)
	err := c.cc.Invoke(ctx, "/articles.ServiceName/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceNameClient) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...grpc.CallOption) (ServiceName_ListArticleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServiceName_serviceDesc.Streams[0], "/articles.ServiceName/ListArticle", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceNameListArticleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceName_ListArticleClient interface {
	Recv() (*ListArticleResponse, error)
	grpc.ClientStream
}

type serviceNameListArticleClient struct {
	grpc.ClientStream
}

func (x *serviceNameListArticleClient) Recv() (*ListArticleResponse, error) {
	m := new(ListArticleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceNameServer is the server API for ServiceName service.
type ServiceNameServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error)
	ListArticle(*ListArticleRequest, ServiceName_ListArticleServer) error
}

// UnimplementedServiceNameServer can be embedded to have forward compatible implementations.
type UnimplementedServiceNameServer struct {
}

func (*UnimplementedServiceNameServer) CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (*UnimplementedServiceNameServer) ListArticle(*ListArticleRequest, ServiceName_ListArticleServer) error {
	return status.Errorf(codes.Unimplemented, "method ListArticle not implemented")
}

func RegisterServiceNameServer(s *grpc.Server, srv ServiceNameServer) {
	s.RegisterService(&_ServiceName_serviceDesc, srv)
}

func _ServiceName_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/articles.ServiceName/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceName_ListArticle_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListArticleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceNameServer).ListArticle(m, &serviceNameListArticleServer{stream})
}

type ServiceName_ListArticleServer interface {
	Send(*ListArticleResponse) error
	grpc.ServerStream
}

type serviceNameListArticleServer struct {
	grpc.ServerStream
}

func (x *serviceNameListArticleServer) Send(m *ListArticleResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ServiceName_serviceDesc = grpc.ServiceDesc{
	ServiceName: "articles.ServiceName",
	HandlerType: (*ServiceNameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ServiceName_CreateArticle_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListArticle",
			Handler:       _ServiceName_ListArticle_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "articles/blogpb/blog.proto",
}