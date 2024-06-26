// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: error_details.proto

package ngrpc

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

type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string               `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Source  *ErrorDetails_Source `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorDetails) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorDetails) GetSource() *ErrorDetails_Source {
	if x != nil {
		return x.Source
	}
	return nil
}

type ErrorDetails_Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message  string            `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Traces   []string          `protobuf:"bytes,2,rep,name=traces,proto3" json:"traces,omitempty"`
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ErrorDetails_Source) Reset() {
	*x = ErrorDetails_Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails_Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails_Source) ProtoMessage() {}

func (x *ErrorDetails_Source) ProtoReflect() protoreflect.Message {
	mi := &file_error_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails_Source.ProtoReflect.Descriptor instead.
func (*ErrorDetails_Source) Descriptor() ([]byte, []int) {
	return file_error_details_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ErrorDetails_Source) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorDetails_Source) GetTraces() []string {
	if x != nil {
		return x.Traces
	}
	return nil
}

func (x *ErrorDetails_Source) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_error_details_proto protoreflect.FileDescriptor

var file_error_details_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x22, 0xb0, 0x02, 0x0a,
	0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a,
	0xbd, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x19, 0x5a, 0x17, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x62,
	0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x6e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_error_details_proto_rawDescOnce sync.Once
	file_error_details_proto_rawDescData = file_error_details_proto_rawDesc
)

func file_error_details_proto_rawDescGZIP() []byte {
	file_error_details_proto_rawDescOnce.Do(func() {
		file_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_details_proto_rawDescData)
	})
	return file_error_details_proto_rawDescData
}

var file_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_error_details_proto_goTypes = []interface{}{
	(*ErrorDetails)(nil),        // 0: ngrpc.ErrorDetails
	(*ErrorDetails_Source)(nil), // 1: ngrpc.ErrorDetails.Source
	nil,                         // 2: ngrpc.ErrorDetails.Source.MetadataEntry
}
var file_error_details_proto_depIdxs = []int32{
	1, // 0: ngrpc.ErrorDetails.source:type_name -> ngrpc.ErrorDetails.Source
	2, // 1: ngrpc.ErrorDetails.Source.metadata:type_name -> ngrpc.ErrorDetails.Source.MetadataEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_error_details_proto_init() }
func file_error_details_proto_init() {
	if File_error_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
		file_error_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails_Source); i {
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
			RawDescriptor: file_error_details_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_details_proto_goTypes,
		DependencyIndexes: file_error_details_proto_depIdxs,
		MessageInfos:      file_error_details_proto_msgTypes,
	}.Build()
	File_error_details_proto = out.File
	file_error_details_proto_rawDesc = nil
	file_error_details_proto_goTypes = nil
	file_error_details_proto_depIdxs = nil
}
