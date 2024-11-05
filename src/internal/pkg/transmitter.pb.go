// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: transmitter.proto

package pkg

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

type FrequencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumValues int32 `protobuf:"varint,1,opt,name=num_values,json=numValues,proto3" json:"num_values,omitempty"` // Количество значений для генерации
}

func (x *FrequencyRequest) Reset() {
	*x = FrequencyRequest{}
	mi := &file_transmitter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FrequencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrequencyRequest) ProtoMessage() {}

func (x *FrequencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrequencyRequest.ProtoReflect.Descriptor instead.
func (*FrequencyRequest) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{0}
}

func (x *FrequencyRequest) GetNumValues() int32 {
	if x != nil {
		return x.NumValues
	}
	return 0
}

type Frequency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string  `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Frequency float64 `protobuf:"fixed64,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	Timestamp int64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Frequency) Reset() {
	*x = Frequency{}
	mi := &file_transmitter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Frequency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Frequency) ProtoMessage() {}

func (x *Frequency) ProtoReflect() protoreflect.Message {
	mi := &file_transmitter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Frequency.ProtoReflect.Descriptor instead.
func (*Frequency) Descriptor() ([]byte, []int) {
	return file_transmitter_proto_rawDescGZIP(), []int{1}
}

func (x *Frequency) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Frequency) GetFrequency() float64 {
	if x != nil {
		return x.Frequency
	}
	return 0
}

func (x *Frequency) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_transmitter_proto protoreflect.FileDescriptor

var file_transmitter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x10, 0x46, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x66, 0x0a,
	0x09, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x66, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x4f, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x11, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x79, 0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x30, 0x01, 0x42, 0x25, 0x5a, 0x23, 0x61, 0x6c, 0x69, 0x65, 0x6e, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2d, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transmitter_proto_rawDescOnce sync.Once
	file_transmitter_proto_rawDescData = file_transmitter_proto_rawDesc
)

func file_transmitter_proto_rawDescGZIP() []byte {
	file_transmitter_proto_rawDescOnce.Do(func() {
		file_transmitter_proto_rawDescData = protoimpl.X.CompressGZIP(file_transmitter_proto_rawDescData)
	})
	return file_transmitter_proto_rawDescData
}

var file_transmitter_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transmitter_proto_goTypes = []any{
	(*FrequencyRequest)(nil), // 0: proto.FrequencyRequest
	(*Frequency)(nil),        // 1: proto.Frequency
}
var file_transmitter_proto_depIdxs = []int32{
	1, // 0: proto.TransmitterService.GenerateFrequency:input_type -> proto.Frequency
	1, // 1: proto.TransmitterService.GenerateFrequency:output_type -> proto.Frequency
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transmitter_proto_init() }
func file_transmitter_proto_init() {
	if File_transmitter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transmitter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transmitter_proto_goTypes,
		DependencyIndexes: file_transmitter_proto_depIdxs,
		MessageInfos:      file_transmitter_proto_msgTypes,
	}.Build()
	File_transmitter_proto = out.File
	file_transmitter_proto_rawDesc = nil
	file_transmitter_proto_goTypes = nil
	file_transmitter_proto_depIdxs = nil
}
