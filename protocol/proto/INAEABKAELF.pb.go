// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: INAEABKAELF.proto

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

type INAEABKAELF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HCHKNOEIFHA uint32 `protobuf:"varint,9,opt,name=HCHKNOEIFHA,proto3" json:"HCHKNOEIFHA,omitempty"`
	OHLEAABAGKF uint32 `protobuf:"varint,14,opt,name=OHLEAABAGKF,proto3" json:"OHLEAABAGKF,omitempty"`
}

func (x *INAEABKAELF) Reset() {
	*x = INAEABKAELF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_INAEABKAELF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *INAEABKAELF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*INAEABKAELF) ProtoMessage() {}

func (x *INAEABKAELF) ProtoReflect() protoreflect.Message {
	mi := &file_INAEABKAELF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use INAEABKAELF.ProtoReflect.Descriptor instead.
func (*INAEABKAELF) Descriptor() ([]byte, []int) {
	return file_INAEABKAELF_proto_rawDescGZIP(), []int{0}
}

func (x *INAEABKAELF) GetHCHKNOEIFHA() uint32 {
	if x != nil {
		return x.HCHKNOEIFHA
	}
	return 0
}

func (x *INAEABKAELF) GetOHLEAABAGKF() uint32 {
	if x != nil {
		return x.OHLEAABAGKF
	}
	return 0
}

var File_INAEABKAELF_proto protoreflect.FileDescriptor

var file_INAEABKAELF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x49, 0x4e, 0x41, 0x45, 0x41, 0x42, 0x4b, 0x41, 0x45, 0x4c, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x49, 0x4e, 0x41, 0x45, 0x41, 0x42, 0x4b, 0x41, 0x45,
	0x4c, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x43, 0x48, 0x4b, 0x4e, 0x4f, 0x45, 0x49, 0x46, 0x48,
	0x41, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x43, 0x48, 0x4b, 0x4e, 0x4f, 0x45,
	0x49, 0x46, 0x48, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x48, 0x4c, 0x45, 0x41, 0x41, 0x42, 0x41,
	0x47, 0x4b, 0x46, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x48, 0x4c, 0x45, 0x41,
	0x41, 0x42, 0x41, 0x47, 0x4b, 0x46, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_INAEABKAELF_proto_rawDescOnce sync.Once
	file_INAEABKAELF_proto_rawDescData = file_INAEABKAELF_proto_rawDesc
)

func file_INAEABKAELF_proto_rawDescGZIP() []byte {
	file_INAEABKAELF_proto_rawDescOnce.Do(func() {
		file_INAEABKAELF_proto_rawDescData = protoimpl.X.CompressGZIP(file_INAEABKAELF_proto_rawDescData)
	})
	return file_INAEABKAELF_proto_rawDescData
}

var file_INAEABKAELF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_INAEABKAELF_proto_goTypes = []interface{}{
	(*INAEABKAELF)(nil), // 0: INAEABKAELF
}
var file_INAEABKAELF_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_INAEABKAELF_proto_init() }
func file_INAEABKAELF_proto_init() {
	if File_INAEABKAELF_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_INAEABKAELF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*INAEABKAELF); i {
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
			RawDescriptor: file_INAEABKAELF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_INAEABKAELF_proto_goTypes,
		DependencyIndexes: file_INAEABKAELF_proto_depIdxs,
		MessageInfos:      file_INAEABKAELF_proto_msgTypes,
	}.Build()
	File_INAEABKAELF_proto = out.File
	file_INAEABKAELF_proto_rawDesc = nil
	file_INAEABKAELF_proto_goTypes = nil
	file_INAEABKAELF_proto_depIdxs = nil
}