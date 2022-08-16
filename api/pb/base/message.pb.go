// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: base/message.proto

package base

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

type SampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Data   int32  `protobuf:"varint,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SampleRequest) Reset() {
	*x = SampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequest) ProtoMessage() {}

func (x *SampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequest.ProtoReflect.Descriptor instead.
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return file_base_message_proto_rawDescGZIP(), []int{0}
}

func (x *SampleRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SampleRequest) GetData() int32 {
	if x != nil {
		return x.Data
	}
	return 0
}

type SampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SampleResponse) Reset() {
	*x = SampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_base_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponse) ProtoMessage() {}

func (x *SampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponse.ProtoReflect.Descriptor instead.
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return file_base_message_proto_rawDescGZIP(), []int{1}
}

func (x *SampleResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_base_message_proto protoreflect.FileDescriptor

var file_base_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x68, 0x61, 0x66, 0x38, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x62, 0x61, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x24, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x13, 0x5a, 0x11, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_base_message_proto_rawDescOnce sync.Once
	file_base_message_proto_rawDescData = file_base_message_proto_rawDesc
)

func file_base_message_proto_rawDescGZIP() []byte {
	file_base_message_proto_rawDescOnce.Do(func() {
		file_base_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_base_message_proto_rawDescData)
	})
	return file_base_message_proto_rawDescData
}

var file_base_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_base_message_proto_goTypes = []interface{}{
	(*SampleRequest)(nil),  // 0: haf80.api.base.SampleRequest
	(*SampleResponse)(nil), // 1: haf80.api.base.SampleResponse
}
var file_base_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_message_proto_init() }
func file_base_message_proto_init() {
	if File_base_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_base_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequest); i {
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
		file_base_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResponse); i {
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
			RawDescriptor: file_base_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_base_message_proto_goTypes,
		DependencyIndexes: file_base_message_proto_depIdxs,
		MessageInfos:      file_base_message_proto_msgTypes,
	}.Build()
	File_base_message_proto = out.File
	file_base_message_proto_rawDesc = nil
	file_base_message_proto_goTypes = nil
	file_base_message_proto_depIdxs = nil
}