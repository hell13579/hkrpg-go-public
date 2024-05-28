// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SectionStatus.proto

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

type SectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId     uint32               `protobuf:"varint,13,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	SectionStatus MessageSectionStatus `protobuf:"varint,3,opt,name=section_status,json=sectionStatus,proto3,enum=MessageSectionStatus" json:"section_status,omitempty"`
}

func (x *SectionStatus) Reset() {
	*x = SectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SectionStatus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SectionStatus) ProtoMessage() {}

func (x *SectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_SectionStatus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SectionStatus.ProtoReflect.Descriptor instead.
func (*SectionStatus) Descriptor() ([]byte, []int) {
	return file_SectionStatus_proto_rawDescGZIP(), []int{0}
}

func (x *SectionStatus) GetSectionId() uint32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

func (x *SectionStatus) GetSectionStatus() MessageSectionStatus {
	if x != nil {
		return x.SectionStatus
	}
	return MessageSectionStatus_MESSAGE_SECTION_NONE
}

var File_SectionStatus_proto protoreflect.FileDescriptor

var file_SectionStatus_proto_rawDesc = []byte{
	0x0a, 0x13, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6c, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x3c, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x0d, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_SectionStatus_proto_rawDescOnce sync.Once
	file_SectionStatus_proto_rawDescData = file_SectionStatus_proto_rawDesc
)

func file_SectionStatus_proto_rawDescGZIP() []byte {
	file_SectionStatus_proto_rawDescOnce.Do(func() {
		file_SectionStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_SectionStatus_proto_rawDescData)
	})
	return file_SectionStatus_proto_rawDescData
}

var file_SectionStatus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SectionStatus_proto_goTypes = []interface{}{
	(*SectionStatus)(nil),     // 0: SectionStatus
	(MessageSectionStatus)(0), // 1: MessageSectionStatus
}
var file_SectionStatus_proto_depIdxs = []int32{
	1, // 0: SectionStatus.section_status:type_name -> MessageSectionStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SectionStatus_proto_init() }
func file_SectionStatus_proto_init() {
	if File_SectionStatus_proto != nil {
		return
	}
	file_MessageSectionStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SectionStatus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SectionStatus); i {
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
			RawDescriptor: file_SectionStatus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SectionStatus_proto_goTypes,
		DependencyIndexes: file_SectionStatus_proto_depIdxs,
		MessageInfos:      file_SectionStatus_proto_msgTypes,
	}.Build()
	File_SectionStatus_proto = out.File
	file_SectionStatus_proto_rawDesc = nil
	file_SectionStatus_proto_goTypes = nil
	file_SectionStatus_proto_depIdxs = nil
}