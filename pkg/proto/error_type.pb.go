// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.1
// source: error_type.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_INTERNAL_ERROR ErrorReason = 0
	ErrorReason_VALIDATE_FAIL  ErrorReason = 1
	ErrorReason_NOT_FOUND      ErrorReason = 2
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "INTERNAL_ERROR",
		1: "VALIDATE_FAIL",
		2: "NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"INTERNAL_ERROR": 0,
		"VALIDATE_FAIL":  1,
		"NOT_FOUND":      2,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_error_type_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_error_type_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_error_type_proto_rawDescGZIP(), []int{0}
}

var File_error_type_proto protoreflect.FileDescriptor

var file_error_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x63, 0x6b, 0x1a,
	0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x5b, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x17, 0x0a,
	0x0d, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x13, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4,
	0x03, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_error_type_proto_rawDescOnce sync.Once
	file_error_type_proto_rawDescData = file_error_type_proto_rawDesc
)

func file_error_type_proto_rawDescGZIP() []byte {
	file_error_type_proto_rawDescOnce.Do(func() {
		file_error_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_type_proto_rawDescData)
	})
	return file_error_type_proto_rawDescData
}

var file_error_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_type_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: candlestick.ErrorReason
}
var file_error_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_type_proto_init() }
func file_error_type_proto_init() {
	if File_error_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_type_proto_goTypes,
		DependencyIndexes: file_error_type_proto_depIdxs,
		EnumInfos:         file_error_type_proto_enumTypes,
	}.Build()
	File_error_type_proto = out.File
	file_error_type_proto_rawDesc = nil
	file_error_type_proto_goTypes = nil
	file_error_type_proto_depIdxs = nil
}
