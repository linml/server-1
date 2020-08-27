// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: protobuf/server/server_cmd.proto

package inner

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CMD int32

const (
	CMD_CMDUNKNOW          CMD = 0
	CMD_CMD_HB_INFO_DETAIL CMD = 1 // 查看红包详情
)

// Enum value maps for CMD.
var (
	CMD_name = map[int32]string{
		0: "CMDUNKNOW",
		1: "CMD_HB_INFO_DETAIL",
	}
	CMD_value = map[string]int32{
		"CMDUNKNOW":          0,
		"CMD_HB_INFO_DETAIL": 1,
	}
)

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}

func (x CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_server_server_cmd_proto_enumTypes[0].Descriptor()
}

func (CMD) Type() protoreflect.EnumType {
	return &file_protobuf_server_server_cmd_proto_enumTypes[0]
}

func (x CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CMD.Descriptor instead.
func (CMD) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_server_server_cmd_proto_rawDescGZIP(), []int{0}
}

// 查看红包详情
type HB_INFO_DETAIL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HbID uint32 `protobuf:"varint,1,opt,name=hbID,proto3" json:"hbID,omitempty"` // 红包ID
}

func (x *HB_INFO_DETAIL) Reset() {
	*x = HB_INFO_DETAIL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_server_server_cmd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HB_INFO_DETAIL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HB_INFO_DETAIL) ProtoMessage() {}

func (x *HB_INFO_DETAIL) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_server_server_cmd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HB_INFO_DETAIL.ProtoReflect.Descriptor instead.
func (*HB_INFO_DETAIL) Descriptor() ([]byte, []int) {
	return file_protobuf_server_server_cmd_proto_rawDescGZIP(), []int{0}
}

func (x *HB_INFO_DETAIL) GetHbID() uint32 {
	if x != nil {
		return x.HbID
	}
	return 0
}

var File_protobuf_server_server_cmd_proto protoreflect.FileDescriptor

var file_protobuf_server_server_cmd_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73, 0x67, 0x22, 0x24, 0x0a, 0x0e,
	0x48, 0x42, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x62, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x68, 0x62,
	0x49, 0x44, 0x2a, 0x2c, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4d, 0x44,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4d, 0x44, 0x5f,
	0x48, 0x42, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x10, 0x01,
	0x42, 0x15, 0x5a, 0x13, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6d, 0x73,
	0x67, 0x2f, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_server_server_cmd_proto_rawDescOnce sync.Once
	file_protobuf_server_server_cmd_proto_rawDescData = file_protobuf_server_server_cmd_proto_rawDesc
)

func file_protobuf_server_server_cmd_proto_rawDescGZIP() []byte {
	file_protobuf_server_server_cmd_proto_rawDescOnce.Do(func() {
		file_protobuf_server_server_cmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_server_server_cmd_proto_rawDescData)
	})
	return file_protobuf_server_server_cmd_proto_rawDescData
}

var file_protobuf_server_server_cmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobuf_server_server_cmd_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_server_server_cmd_proto_goTypes = []interface{}{
	(CMD)(0),               // 0: protomsg.CMD
	(*HB_INFO_DETAIL)(nil), // 1: protomsg.HB_INFO_DETAIL
}
var file_protobuf_server_server_cmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_server_server_cmd_proto_init() }
func file_protobuf_server_server_cmd_proto_init() {
	if File_protobuf_server_server_cmd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_server_server_cmd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HB_INFO_DETAIL); i {
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
			RawDescriptor: file_protobuf_server_server_cmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_server_server_cmd_proto_goTypes,
		DependencyIndexes: file_protobuf_server_server_cmd_proto_depIdxs,
		EnumInfos:         file_protobuf_server_server_cmd_proto_enumTypes,
		MessageInfos:      file_protobuf_server_server_cmd_proto_msgTypes,
	}.Build()
	File_protobuf_server_server_cmd_proto = out.File
	file_protobuf_server_server_cmd_proto_rawDesc = nil
	file_protobuf_server_server_cmd_proto_goTypes = nil
	file_protobuf_server_server_cmd_proto_depIdxs = nil
}
