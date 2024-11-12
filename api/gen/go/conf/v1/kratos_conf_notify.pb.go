// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: conf/v1/kratos_conf_notify.proto

package confv1

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

// 通知消息
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   bool                   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	DingTalk *Notification_DingTalk `protobuf:"bytes,2,opt,name=ding_talk,json=dingTalk,proto3" json:"ding_talk,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_conf_v1_kratos_conf_notify_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_notify_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_notify_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Notification) GetDingTalk() *Notification_DingTalk {
	if x != nil {
		return x.DingTalk
	}
	return nil
}

// 钉钉
type Notification_DingTalk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens []string `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *Notification_DingTalk) Reset() {
	*x = Notification_DingTalk{}
	mi := &file_conf_v1_kratos_conf_notify_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification_DingTalk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification_DingTalk) ProtoMessage() {}

func (x *Notification_DingTalk) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_notify_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification_DingTalk.ProtoReflect.Descriptor instead.
func (*Notification_DingTalk) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_notify_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Notification_DingTalk) GetTokens() []string {
	if x != nil {
		return x.Tokens
	}
	return nil
}

var File_conf_v1_kratos_conf_notify_proto protoreflect.FileDescriptor

var file_conf_v1_kratos_conf_notify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76, 0x31, 0x22, 0x87, 0x01, 0x0a, 0x0c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x6c,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x52, 0x08, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c,
	0x6b, 0x1a, 0x22, 0x0a, 0x08, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x9c, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x63, 0x34,
	0x30, 0x34, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa,
	0x02, 0x07, 0x43, 0x6f, 0x6e, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x43, 0x6f, 0x6e, 0x66,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x43, 0x6f, 0x6e, 0x66,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_kratos_conf_notify_proto_rawDescOnce sync.Once
	file_conf_v1_kratos_conf_notify_proto_rawDescData = file_conf_v1_kratos_conf_notify_proto_rawDesc
)

func file_conf_v1_kratos_conf_notify_proto_rawDescGZIP() []byte {
	file_conf_v1_kratos_conf_notify_proto_rawDescOnce.Do(func() {
		file_conf_v1_kratos_conf_notify_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_kratos_conf_notify_proto_rawDescData)
	})
	return file_conf_v1_kratos_conf_notify_proto_rawDescData
}

var file_conf_v1_kratos_conf_notify_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_conf_v1_kratos_conf_notify_proto_goTypes = []any{
	(*Notification)(nil),          // 0: conf.v1.Notification
	(*Notification_DingTalk)(nil), // 1: conf.v1.Notification.DingTalk
}
var file_conf_v1_kratos_conf_notify_proto_depIdxs = []int32{
	1, // 0: conf.v1.Notification.ding_talk:type_name -> conf.v1.Notification.DingTalk
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conf_v1_kratos_conf_notify_proto_init() }
func file_conf_v1_kratos_conf_notify_proto_init() {
	if File_conf_v1_kratos_conf_notify_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conf_v1_kratos_conf_notify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_kratos_conf_notify_proto_goTypes,
		DependencyIndexes: file_conf_v1_kratos_conf_notify_proto_depIdxs,
		MessageInfos:      file_conf_v1_kratos_conf_notify_proto_msgTypes,
	}.Build()
	File_conf_v1_kratos_conf_notify_proto = out.File
	file_conf_v1_kratos_conf_notify_proto_rawDesc = nil
	file_conf_v1_kratos_conf_notify_proto_goTypes = nil
	file_conf_v1_kratos_conf_notify_proto_depIdxs = nil
}