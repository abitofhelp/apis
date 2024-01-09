// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.7
// source: proto/hello_world/v1/hello_world.proto

package v1

import (
	v5 "github.com/abitofhelp/apis/proto/enum/access_tier/v5"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tier v5.AccessTier          `protobuf:"varint,2,opt,name=tier,proto3,enum=v5.AccessTier" json:"tier,omitempty"`
	Utc  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=utc,proto3" json:"utc,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetTier() v5.AccessTier {
	if x != nil {
		return x.Tier
	}
	return v5.AccessTier(0)
}

func (x *Person) GetUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.Utc
	}
	return nil
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SentUtc *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=sent_utc,json=sentUtc,proto3" json:"sent_utc,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HelloRequest) GetSentUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.SentUtc
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message           string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MessageArrivedUtc *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=message_arrived_utc,json=messageArrivedUtc,proto3" json:"message_arrived_utc,omitempty"`
	RepliedUtc        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=replied_utc,json=repliedUtc,proto3" json:"replied_utc,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hello_world_v1_hello_world_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_hello_world_v1_hello_world_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *HelloResponse) GetMessageArrivedUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.MessageArrivedUtc
	}
	return nil
}

func (x *HelloResponse) GetRepliedUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.RepliedUtc
	}
	return nil
}

var File_proto_hello_world_v1_hello_world_proto protoreflect.FileDescriptor

var file_proto_hello_world_v1_hello_world_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x69, 0x65, 0x72, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x69,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x74, 0x69, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x76, 0x35, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x69, 0x65, 0x72, 0x52, 0x04, 0x74, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x03, 0x75, 0x74,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x03, 0x75, 0x74, 0x63, 0x22, 0x59, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x74,
	0x55, 0x74, 0x63, 0x22, 0xb2, 0x01, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x4a, 0x0a, 0x13, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x41, 0x72, 0x72, 0x69, 0x76, 0x65, 0x64, 0x55, 0x74, 0x63, 0x12, 0x3b, 0x0a, 0x0b, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x75, 0x74, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6c, 0x69, 0x65, 0x64, 0x55, 0x74, 0x63, 0x32, 0x50, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x10, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22,
	0x09, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x69, 0x74, 0x6f, 0x66, 0x68,
	0x65, 0x6c, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hello_world_v1_hello_world_proto_rawDescOnce sync.Once
	file_proto_hello_world_v1_hello_world_proto_rawDescData = file_proto_hello_world_v1_hello_world_proto_rawDesc
)

func file_proto_hello_world_v1_hello_world_proto_rawDescGZIP() []byte {
	file_proto_hello_world_v1_hello_world_proto_rawDescOnce.Do(func() {
		file_proto_hello_world_v1_hello_world_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hello_world_v1_hello_world_proto_rawDescData)
	})
	return file_proto_hello_world_v1_hello_world_proto_rawDescData
}

var file_proto_hello_world_v1_hello_world_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hello_world_v1_hello_world_proto_goTypes = []interface{}{
	(*Person)(nil),                // 0: v1.Person
	(*HelloRequest)(nil),          // 1: v1.HelloRequest
	(*HelloResponse)(nil),         // 2: v1.HelloResponse
	(v5.AccessTier)(0),            // 3: v5.AccessTier
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_proto_hello_world_v1_hello_world_proto_depIdxs = []int32{
	3, // 0: v1.Person.tier:type_name -> v5.AccessTier
	4, // 1: v1.Person.utc:type_name -> google.protobuf.Timestamp
	4, // 2: v1.HelloRequest.sent_utc:type_name -> google.protobuf.Timestamp
	4, // 3: v1.HelloResponse.message_arrived_utc:type_name -> google.protobuf.Timestamp
	4, // 4: v1.HelloResponse.replied_utc:type_name -> google.protobuf.Timestamp
	1, // 5: v1.Greeter.SayHello:input_type -> v1.HelloRequest
	2, // 6: v1.Greeter.SayHello:output_type -> v1.HelloResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_hello_world_v1_hello_world_proto_init() }
func file_proto_hello_world_v1_hello_world_proto_init() {
	if File_proto_hello_world_v1_hello_world_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hello_world_v1_hello_world_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_proto_hello_world_v1_hello_world_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_proto_hello_world_v1_hello_world_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_proto_hello_world_v1_hello_world_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hello_world_v1_hello_world_proto_goTypes,
		DependencyIndexes: file_proto_hello_world_v1_hello_world_proto_depIdxs,
		MessageInfos:      file_proto_hello_world_v1_hello_world_proto_msgTypes,
	}.Build()
	File_proto_hello_world_v1_hello_world_proto = out.File
	file_proto_hello_world_v1_hello_world_proto_rawDesc = nil
	file_proto_hello_world_v1_hello_world_proto_goTypes = nil
	file_proto_hello_world_v1_hello_world_proto_depIdxs = nil
}
