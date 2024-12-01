// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc3
// source: nihao.proto

package fuwu

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

type NihaoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestName string `protobuf:"bytes,1,opt,name=requestName,proto3" json:"requestName,omitempty"`
	Age         int64  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *NihaoRequest) Reset() {
	*x = NihaoRequest{}
	mi := &file_nihao_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NihaoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NihaoRequest) ProtoMessage() {}

func (x *NihaoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nihao_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NihaoRequest.ProtoReflect.Descriptor instead.
func (*NihaoRequest) Descriptor() ([]byte, []int) {
	return file_nihao_proto_rawDescGZIP(), []int{0}
}

func (x *NihaoRequest) GetRequestName() string {
	if x != nil {
		return x.RequestName
	}
	return ""
}

func (x *NihaoRequest) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type NihaoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseMsg string   `protobuf:"bytes,1,opt,name=responseMsg,proto3" json:"responseMsg,omitempty"`
	Id          int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Names       []string `protobuf:"bytes,3,rep,name=Names,proto3" json:"Names,omitempty"`
}

func (x *NihaoResponse) Reset() {
	*x = NihaoResponse{}
	mi := &file_nihao_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NihaoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NihaoResponse) ProtoMessage() {}

func (x *NihaoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nihao_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NihaoResponse.ProtoReflect.Descriptor instead.
func (*NihaoResponse) Descriptor() ([]byte, []int) {
	return file_nihao_proto_rawDescGZIP(), []int{1}
}

func (x *NihaoResponse) GetResponseMsg() string {
	if x != nil {
		return x.ResponseMsg
	}
	return ""
}

func (x *NihaoResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NihaoResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_nihao_proto protoreflect.FileDescriptor

var file_nihao_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6e, 0x69, 0x68, 0x61, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x0c, 0x4e, 0x69, 0x68, 0x61, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x22, 0x57, 0x0a, 0x0d, 0x4e, 0x69, 0x68, 0x61, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x73,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x32, 0x31, 0x0a, 0x05, 0x4e, 0x69,
	0x68, 0x61, 0x6f, 0x12, 0x28, 0x0a, 0x05, 0x4e, 0x69, 0x68, 0x61, 0x6f, 0x12, 0x0d, 0x2e, 0x4e,
	0x69, 0x68, 0x61, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4e, 0x69,
	0x68, 0x61, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x66, 0x75, 0x77, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nihao_proto_rawDescOnce sync.Once
	file_nihao_proto_rawDescData = file_nihao_proto_rawDesc
)

func file_nihao_proto_rawDescGZIP() []byte {
	file_nihao_proto_rawDescOnce.Do(func() {
		file_nihao_proto_rawDescData = protoimpl.X.CompressGZIP(file_nihao_proto_rawDescData)
	})
	return file_nihao_proto_rawDescData
}

var file_nihao_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nihao_proto_goTypes = []any{
	(*NihaoRequest)(nil),  // 0: NihaoRequest
	(*NihaoResponse)(nil), // 1: NihaoResponse
}
var file_nihao_proto_depIdxs = []int32{
	0, // 0: Nihao.Nihao:input_type -> NihaoRequest
	1, // 1: Nihao.Nihao:output_type -> NihaoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nihao_proto_init() }
func file_nihao_proto_init() {
	if File_nihao_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_nihao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nihao_proto_goTypes,
		DependencyIndexes: file_nihao_proto_depIdxs,
		MessageInfos:      file_nihao_proto_msgTypes,
	}.Build()
	File_nihao_proto = out.File
	file_nihao_proto_rawDesc = nil
	file_nihao_proto_goTypes = nil
	file_nihao_proto_depIdxs = nil
}