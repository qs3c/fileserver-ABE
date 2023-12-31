// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.0
// source: abe.proto

package go_micro_service_encupload

import (
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

type ReqAccessPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessPolicy string `protobuf:"bytes,1,opt,name=access_policy,json=accessPolicy,proto3" json:"access_policy,omitempty"`
}

func (x *ReqAccessPolicy) Reset() {
	*x = ReqAccessPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encupload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqAccessPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqAccessPolicy) ProtoMessage() {}

func (x *ReqAccessPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_encupload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqAccessPolicy.ProtoReflect.Descriptor instead.
func (*ReqAccessPolicy) Descriptor() ([]byte, []int) {
	return file_encupload_proto_rawDescGZIP(), []int{0}
}

func (x *ReqAccessPolicy) GetAccessPolicy() string {
	if x != nil {
		return x.AccessPolicy
	}
	return ""
}

type RespKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RespKey) Reset() {
	*x = RespKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encupload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespKey) ProtoMessage() {}

func (x *RespKey) ProtoReflect() protoreflect.Message {
	mi := &file_encupload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespKey.ProtoReflect.Descriptor instead.
func (*RespKey) Descriptor() ([]byte, []int) {
	return file_encupload_proto_rawDescGZIP(), []int{1}
}

func (x *RespKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_encupload_proto protoreflect.FileDescriptor

var file_encupload_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x63, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x36, 0x0a,
	0x0f, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x70, 0x4b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x32, 0x76, 0x0a, 0x10, 0x45, 0x6e, 0x63, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a, 0x0c, 0x41, 0x62, 0x45, 0x6e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x63, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x2e, 0x52, 0x65, 0x71, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x1a, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x65, 0x6e, 0x63, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x2e, 0x2f,
	0x3b, 0x67, 0x6f, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x65, 0x6e, 0x63, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_encupload_proto_rawDescOnce sync.Once
	file_encupload_proto_rawDescData = file_encupload_proto_rawDesc
)

func file_encupload_proto_rawDescGZIP() []byte {
	file_encupload_proto_rawDescOnce.Do(func() {
		file_encupload_proto_rawDescData = protoimpl.X.CompressGZIP(file_encupload_proto_rawDescData)
	})
	return file_encupload_proto_rawDescData
}

var file_encupload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_encupload_proto_goTypes = []interface{}{
	(*ReqAccessPolicy)(nil), // 0: go.micro.service.abe.ReqAccessPolicy
	(*RespKey)(nil),         // 1: go.micro.service.abe.RespKey
}
var file_encupload_proto_depIdxs = []int32{
	0, // 0: go.micro.service.abe.EncUploadService.AbEncryption:input_type -> go.micro.service.abe.ReqAccessPolicy
	1, // 1: go.micro.service.abe.EncUploadService.AbEncryption:output_type -> go.micro.service.abe.RespKey
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_encupload_proto_init() }
func file_encupload_proto_init() {
	if File_encupload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encupload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqAccessPolicy); i {
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
		file_encupload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespKey); i {
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
			RawDescriptor: file_encupload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_encupload_proto_goTypes,
		DependencyIndexes: file_encupload_proto_depIdxs,
		MessageInfos:      file_encupload_proto_msgTypes,
	}.Build()
	File_encupload_proto = out.File
	file_encupload_proto_rawDesc = nil
	file_encupload_proto_goTypes = nil
	file_encupload_proto_depIdxs = nil
}
