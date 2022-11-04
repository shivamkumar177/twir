// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: eventsub.proto

package eventsub

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

type SubscribeToEvents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelId string `protobuf:"bytes,1,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *SubscribeToEvents) Reset() {
	*x = SubscribeToEvents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToEvents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToEvents) ProtoMessage() {}

func (x *SubscribeToEvents) ProtoReflect() protoreflect.Message {
	mi := &file_eventsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToEvents.ProtoReflect.Descriptor instead.
func (*SubscribeToEvents) Descriptor() ([]byte, []int) {
	return file_eventsub_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeToEvents) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type SubscribeToEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SubscribeToEventsResponse) Reset() {
	*x = SubscribeToEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eventsub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeToEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeToEventsResponse) ProtoMessage() {}

func (x *SubscribeToEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eventsub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeToEventsResponse.ProtoReflect.Descriptor instead.
func (*SubscribeToEventsResponse) Descriptor() ([]byte, []int) {
	return file_eventsub_proto_rawDescGZIP(), []int{1}
}

var File_eventsub_proto protoreflect.FileDescriptor

var file_eventsub_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x31, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x6f, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x54, 0x6f, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x17, 0x5a, 0x15, 0x74, 0x73, 0x75, 0x77, 0x61, 0x72, 0x69, 0x2f, 0x6e, 0x61, 0x74, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eventsub_proto_rawDescOnce sync.Once
	file_eventsub_proto_rawDescData = file_eventsub_proto_rawDesc
)

func file_eventsub_proto_rawDescGZIP() []byte {
	file_eventsub_proto_rawDescOnce.Do(func() {
		file_eventsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_eventsub_proto_rawDescData)
	})
	return file_eventsub_proto_rawDescData
}

var file_eventsub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eventsub_proto_goTypes = []interface{}{
	(*SubscribeToEvents)(nil),         // 0: SubscribeToEvents
	(*SubscribeToEventsResponse)(nil), // 1: SubscribeToEventsResponse
}
var file_eventsub_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eventsub_proto_init() }
func file_eventsub_proto_init() {
	if File_eventsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eventsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToEvents); i {
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
		file_eventsub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeToEventsResponse); i {
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
			RawDescriptor: file_eventsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eventsub_proto_goTypes,
		DependencyIndexes: file_eventsub_proto_depIdxs,
		MessageInfos:      file_eventsub_proto_msgTypes,
	}.Build()
	File_eventsub_proto = out.File
	file_eventsub_proto_rawDesc = nil
	file_eventsub_proto_goTypes = nil
	file_eventsub_proto_depIdxs = nil
}

  const (
    SUBJECTS_SUBSCTUBE_TO_EVENTS_BY_CHANNEL_ID = "subscribeToEventsByChannelId"
    )
    