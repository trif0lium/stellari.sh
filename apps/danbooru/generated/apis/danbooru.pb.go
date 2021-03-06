// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: apis/danbooru.proto

package danbooru_v1

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

type Rating int32

const (
	Rating_ALL          Rating = 0
	Rating_SAFE         Rating = 1
	Rating_QUESTIONABLE Rating = 2
	Rating_EXPLICIT     Rating = 3
)

// Enum value maps for Rating.
var (
	Rating_name = map[int32]string{
		0: "ALL",
		1: "SAFE",
		2: "QUESTIONABLE",
		3: "EXPLICIT",
	}
	Rating_value = map[string]int32{
		"ALL":          0,
		"SAFE":         1,
		"QUESTIONABLE": 2,
		"EXPLICIT":     3,
	}
)

func (x Rating) Enum() *Rating {
	p := new(Rating)
	*p = x
	return p
}

func (x Rating) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rating) Descriptor() protoreflect.EnumDescriptor {
	return file_apis_danbooru_proto_enumTypes[0].Descriptor()
}

func (Rating) Type() protoreflect.EnumType {
	return &file_apis_danbooru_proto_enumTypes[0]
}

func (x Rating) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rating.Descriptor instead.
func (Rating) EnumDescriptor() ([]byte, []int) {
	return file_apis_danbooru_proto_rawDescGZIP(), []int{0}
}

type PostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *PostsRequest) Reset() {
	*x = PostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_danbooru_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsRequest) ProtoMessage() {}

func (x *PostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apis_danbooru_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsRequest.ProtoReflect.Descriptor instead.
func (*PostsRequest) Descriptor() ([]byte, []int) {
	return file_apis_danbooru_proto_rawDescGZIP(), []int{0}
}

func (x *PostsRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type PostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*PostsResponse_Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *PostsResponse) Reset() {
	*x = PostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_danbooru_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse) ProtoMessage() {}

func (x *PostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apis_danbooru_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsResponse.ProtoReflect.Descriptor instead.
func (*PostsResponse) Descriptor() ([]byte, []int) {
	return file_apis_danbooru_proto_rawDescGZIP(), []int{1}
}

func (x *PostsResponse) GetPosts() []*PostsResponse_Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type PostsResponse_Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags      []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	FileUrl   string   `protobuf:"bytes,3,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	CreatedAt int32    `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PostsResponse_Post) Reset() {
	*x = PostsResponse_Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apis_danbooru_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostsResponse_Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostsResponse_Post) ProtoMessage() {}

func (x *PostsResponse_Post) ProtoReflect() protoreflect.Message {
	mi := &file_apis_danbooru_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostsResponse_Post.ProtoReflect.Descriptor instead.
func (*PostsResponse_Post) Descriptor() ([]byte, []int) {
	return file_apis_danbooru_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PostsResponse_Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostsResponse_Post) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *PostsResponse_Post) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *PostsResponse_Post) GetCreatedAt() int32 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_apis_danbooru_proto protoreflect.FileDescriptor

var file_apis_danbooru_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f, 0x72, 0x75, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f, 0x72, 0x75, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0xac, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f,
	0x72, 0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x1a,
	0x64, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x3b, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x41, 0x46, 0x45,
	0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x51, 0x55, 0x45, 0x53, 0x54, 0x49, 0x4f, 0x4e, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x50, 0x4c, 0x49, 0x43, 0x49, 0x54,
	0x10, 0x03, 0x32, 0x4b, 0x0a, 0x07, 0x47, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x79, 0x12, 0x40, 0x0a,
	0x05, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f, 0x72,
	0x75, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f, 0x72, 0x75, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x3b, 0x64, 0x61, 0x6e, 0x62, 0x6f, 0x6f, 0x72, 0x75, 0x5f, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apis_danbooru_proto_rawDescOnce sync.Once
	file_apis_danbooru_proto_rawDescData = file_apis_danbooru_proto_rawDesc
)

func file_apis_danbooru_proto_rawDescGZIP() []byte {
	file_apis_danbooru_proto_rawDescOnce.Do(func() {
		file_apis_danbooru_proto_rawDescData = protoimpl.X.CompressGZIP(file_apis_danbooru_proto_rawDescData)
	})
	return file_apis_danbooru_proto_rawDescData
}

var file_apis_danbooru_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apis_danbooru_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apis_danbooru_proto_goTypes = []interface{}{
	(Rating)(0),                // 0: danbooru.v1.Rating
	(*PostsRequest)(nil),       // 1: danbooru.v1.PostsRequest
	(*PostsResponse)(nil),      // 2: danbooru.v1.PostsResponse
	(*PostsResponse_Post)(nil), // 3: danbooru.v1.PostsResponse.Post
}
var file_apis_danbooru_proto_depIdxs = []int32{
	3, // 0: danbooru.v1.PostsResponse.posts:type_name -> danbooru.v1.PostsResponse.Post
	1, // 1: danbooru.v1.Gallery.Posts:input_type -> danbooru.v1.PostsRequest
	2, // 2: danbooru.v1.Gallery.Posts:output_type -> danbooru.v1.PostsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_apis_danbooru_proto_init() }
func file_apis_danbooru_proto_init() {
	if File_apis_danbooru_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apis_danbooru_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsRequest); i {
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
		file_apis_danbooru_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsResponse); i {
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
		file_apis_danbooru_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostsResponse_Post); i {
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
			RawDescriptor: file_apis_danbooru_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_danbooru_proto_goTypes,
		DependencyIndexes: file_apis_danbooru_proto_depIdxs,
		EnumInfos:         file_apis_danbooru_proto_enumTypes,
		MessageInfos:      file_apis_danbooru_proto_msgTypes,
	}.Build()
	File_apis_danbooru_proto = out.File
	file_apis_danbooru_proto_rawDesc = nil
	file_apis_danbooru_proto_goTypes = nil
	file_apis_danbooru_proto_depIdxs = nil
}
