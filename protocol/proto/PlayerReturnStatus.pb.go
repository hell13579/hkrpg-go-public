// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerReturnStatus.proto

package proto

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

type PlayerReturnStatus int32

const (
	PlayerReturnStatus_PLAYER_RETURN_NONE       PlayerReturnStatus = 0
	PlayerReturnStatus_PLAYER_RETURN_PROCESSING PlayerReturnStatus = 1
	PlayerReturnStatus_PLAYER_RETURN_FINISH     PlayerReturnStatus = 2
)

// Enum value maps for PlayerReturnStatus.
var (
	PlayerReturnStatus_name = map[int32]string{
		0: "PLAYER_RETURN_NONE",
		1: "PLAYER_RETURN_PROCESSING",
		2: "PLAYER_RETURN_FINISH",
	}
	PlayerReturnStatus_value = map[string]int32{
		"PLAYER_RETURN_NONE":       0,
		"PLAYER_RETURN_PROCESSING": 1,
		"PLAYER_RETURN_FINISH":     2,
	}
)

func (x PlayerReturnStatus) Enum() *PlayerReturnStatus {
	p := new(PlayerReturnStatus)
	*p = x
	return p
}

func (x PlayerReturnStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayerReturnStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_PlayerReturnStatus_proto_enumTypes[0].Descriptor()
}

func (PlayerReturnStatus) Type() protoreflect.EnumType {
	return &file_PlayerReturnStatus_proto_enumTypes[0]
}

func (x PlayerReturnStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayerReturnStatus.Descriptor instead.
func (PlayerReturnStatus) EnumDescriptor() ([]byte, []int) {
	return file_PlayerReturnStatus_proto_rawDescGZIP(), []int{0}
}

var File_PlayerReturnStatus_proto protoreflect.FileDescriptor

var file_PlayerReturnStatus_proto_rawDesc = []byte{
	0x0a, 0x18, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x64, 0x0a, 0x12, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x12, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52,
	0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x4c, 0x41, 0x59,
	0x45, 0x52, 0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53,
	0x53, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x4c, 0x41, 0x59, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x02,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerReturnStatus_proto_rawDescOnce sync.Once
	file_PlayerReturnStatus_proto_rawDescData = file_PlayerReturnStatus_proto_rawDesc
)

func file_PlayerReturnStatus_proto_rawDescGZIP() []byte {
	file_PlayerReturnStatus_proto_rawDescOnce.Do(func() {
		file_PlayerReturnStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerReturnStatus_proto_rawDescData)
	})
	return file_PlayerReturnStatus_proto_rawDescData
}

var file_PlayerReturnStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_PlayerReturnStatus_proto_goTypes = []interface{}{
	(PlayerReturnStatus)(0), // 0: PlayerReturnStatus
}
var file_PlayerReturnStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerReturnStatus_proto_init() }
func file_PlayerReturnStatus_proto_init() {
	if File_PlayerReturnStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_PlayerReturnStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerReturnStatus_proto_goTypes,
		DependencyIndexes: file_PlayerReturnStatus_proto_depIdxs,
		EnumInfos:         file_PlayerReturnStatus_proto_enumTypes,
	}.Build()
	File_PlayerReturnStatus_proto = out.File
	file_PlayerReturnStatus_proto_rawDesc = nil
	file_PlayerReturnStatus_proto_goTypes = nil
	file_PlayerReturnStatus_proto_depIdxs = nil
}