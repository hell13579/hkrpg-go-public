// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TrainVisitorInfo.proto

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

type TrainVisitorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      TrainVisitorStatus `protobuf:"varint,7,opt,name=status,proto3,enum=TrainVisitorStatus" json:"status,omitempty"`
	ADNDLJGBNGA uint32             `protobuf:"varint,1,opt,name=ADNDLJGBNGA,proto3" json:"ADNDLJGBNGA,omitempty"`
	KIMGCNFAANM []uint32           `protobuf:"varint,14,rep,packed,name=KIMGCNFAANM,proto3" json:"KIMGCNFAANM,omitempty"`
	PPGJNEMBPIG uint32             `protobuf:"varint,10,opt,name=PPGJNEMBPIG,proto3" json:"PPGJNEMBPIG,omitempty"`
	NCAGKIGLMNO bool               `protobuf:"varint,9,opt,name=NCAGKIGLMNO,proto3" json:"NCAGKIGLMNO,omitempty"`
}

func (x *TrainVisitorInfo) Reset() {
	*x = TrainVisitorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TrainVisitorInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainVisitorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainVisitorInfo) ProtoMessage() {}

func (x *TrainVisitorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_TrainVisitorInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainVisitorInfo.ProtoReflect.Descriptor instead.
func (*TrainVisitorInfo) Descriptor() ([]byte, []int) {
	return file_TrainVisitorInfo_proto_rawDescGZIP(), []int{0}
}

func (x *TrainVisitorInfo) GetStatus() TrainVisitorStatus {
	if x != nil {
		return x.Status
	}
	return TrainVisitorStatus_TRAIN_VISITOR_STATUS_NONE
}

func (x *TrainVisitorInfo) GetADNDLJGBNGA() uint32 {
	if x != nil {
		return x.ADNDLJGBNGA
	}
	return 0
}

func (x *TrainVisitorInfo) GetKIMGCNFAANM() []uint32 {
	if x != nil {
		return x.KIMGCNFAANM
	}
	return nil
}

func (x *TrainVisitorInfo) GetPPGJNEMBPIG() uint32 {
	if x != nil {
		return x.PPGJNEMBPIG
	}
	return 0
}

func (x *TrainVisitorInfo) GetNCAGKIGLMNO() bool {
	if x != nil {
		return x.NCAGKIGLMNO
	}
	return false
}

var File_TrainVisitorInfo_proto protoreflect.FileDescriptor

var file_TrainVisitorInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x56,
	0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x56, 0x69, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x56,
	0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x44, 0x4e, 0x44, 0x4c, 0x4a, 0x47, 0x42,
	0x4e, 0x47, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x44, 0x4e, 0x44, 0x4c,
	0x4a, 0x47, 0x42, 0x4e, 0x47, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x49, 0x4d, 0x47, 0x43, 0x4e,
	0x46, 0x41, 0x41, 0x4e, 0x4d, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4b, 0x49, 0x4d,
	0x47, 0x43, 0x4e, 0x46, 0x41, 0x41, 0x4e, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x50, 0x47, 0x4a,
	0x4e, 0x45, 0x4d, 0x42, 0x50, 0x49, 0x47, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50,
	0x50, 0x47, 0x4a, 0x4e, 0x45, 0x4d, 0x42, 0x50, 0x49, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x43,
	0x41, 0x47, 0x4b, 0x49, 0x47, 0x4c, 0x4d, 0x4e, 0x4f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x4e, 0x43, 0x41, 0x47, 0x4b, 0x49, 0x47, 0x4c, 0x4d, 0x4e, 0x4f, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrainVisitorInfo_proto_rawDescOnce sync.Once
	file_TrainVisitorInfo_proto_rawDescData = file_TrainVisitorInfo_proto_rawDesc
)

func file_TrainVisitorInfo_proto_rawDescGZIP() []byte {
	file_TrainVisitorInfo_proto_rawDescOnce.Do(func() {
		file_TrainVisitorInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrainVisitorInfo_proto_rawDescData)
	})
	return file_TrainVisitorInfo_proto_rawDescData
}

var file_TrainVisitorInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TrainVisitorInfo_proto_goTypes = []interface{}{
	(*TrainVisitorInfo)(nil), // 0: TrainVisitorInfo
	(TrainVisitorStatus)(0),  // 1: TrainVisitorStatus
}
var file_TrainVisitorInfo_proto_depIdxs = []int32{
	1, // 0: TrainVisitorInfo.status:type_name -> TrainVisitorStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TrainVisitorInfo_proto_init() }
func file_TrainVisitorInfo_proto_init() {
	if File_TrainVisitorInfo_proto != nil {
		return
	}
	file_TrainVisitorStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TrainVisitorInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainVisitorInfo); i {
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
			RawDescriptor: file_TrainVisitorInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrainVisitorInfo_proto_goTypes,
		DependencyIndexes: file_TrainVisitorInfo_proto_depIdxs,
		MessageInfos:      file_TrainVisitorInfo_proto_msgTypes,
	}.Build()
	File_TrainVisitorInfo_proto = out.File
	file_TrainVisitorInfo_proto_rawDesc = nil
	file_TrainVisitorInfo_proto_goTypes = nil
	file_TrainVisitorInfo_proto_depIdxs = nil
}