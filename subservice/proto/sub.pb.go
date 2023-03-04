// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/sub.proto

package sub

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

type SubRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SubRequest) Reset() {
	*x = SubRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubRequest) ProtoMessage() {}

func (x *SubRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubRequest.ProtoReflect.Descriptor instead.
func (*SubRequest) Descriptor() ([]byte, []int) {
	return file_proto_sub_proto_rawDescGZIP(), []int{0}
}

func (x *SubRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SubRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type SubReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SubReply) Reset() {
	*x = SubReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sub_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubReply) ProtoMessage() {}

func (x *SubReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sub_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubReply.ProtoReflect.Descriptor instead.
func (*SubReply) Descriptor() ([]byte, []int) {
	return file_proto_sub_proto_rawDescGZIP(), []int{1}
}

func (x *SubReply) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_sub_proto protoreflect.FileDescriptor

var file_proto_sub_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x73, 0x75, 0x62, 0x22, 0x28, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x62,
	0x22, 0x22, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x3b, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x75, 0x62, 0x12,
	0x0f, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x75, 0x62, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_sub_proto_rawDescOnce sync.Once
	file_proto_sub_proto_rawDescData = file_proto_sub_proto_rawDesc
)

func file_proto_sub_proto_rawDescGZIP() []byte {
	file_proto_sub_proto_rawDescOnce.Do(func() {
		file_proto_sub_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sub_proto_rawDescData)
	})
	return file_proto_sub_proto_rawDescData
}

var file_proto_sub_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_sub_proto_goTypes = []interface{}{
	(*SubRequest)(nil), // 0: sub.SubRequest
	(*SubReply)(nil),   // 1: sub.SubReply
}
var file_proto_sub_proto_depIdxs = []int32{
	0, // 0: sub.SubService.HandleSub:input_type -> sub.SubRequest
	1, // 1: sub.SubService.HandleSub:output_type -> sub.SubReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sub_proto_init() }
func file_proto_sub_proto_init() {
	if File_proto_sub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubRequest); i {
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
		file_proto_sub_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubReply); i {
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
			RawDescriptor: file_proto_sub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sub_proto_goTypes,
		DependencyIndexes: file_proto_sub_proto_depIdxs,
		MessageInfos:      file_proto_sub_proto_msgTypes,
	}.Build()
	File_proto_sub_proto = out.File
	file_proto_sub_proto_rawDesc = nil
	file_proto_sub_proto_goTypes = nil
	file_proto_sub_proto_depIdxs = nil
}
