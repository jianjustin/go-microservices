// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/mul.proto

package mul

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type MulRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *MulRequest) Reset() {
	*x = MulRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mul_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MulRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulRequest) ProtoMessage() {}

func (x *MulRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mul_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulRequest.ProtoReflect.Descriptor instead.
func (*MulRequest) Descriptor() ([]byte, []int) {
	return file_proto_mul_proto_rawDescGZIP(), []int{0}
}

func (x *MulRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *MulRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type MulReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MulReply) Reset() {
	*x = MulReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mul_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MulReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MulReply) ProtoMessage() {}

func (x *MulReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mul_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MulReply.ProtoReflect.Descriptor instead.
func (*MulReply) Descriptor() ([]byte, []int) {
	return file_proto_mul_proto_rawDescGZIP(), []int{1}
}

func (x *MulReply) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_mul_proto protoreflect.FileDescriptor

var file_proto_mul_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x75, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x6d, 0x75, 0x6c, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x61,
	0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x62, 0x22, 0x22,
	0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x57, 0x0a, 0x0a, 0x4d, 0x75, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x75, 0x6c, 0x12, 0x0f, 0x2e,
	0x6d, 0x75, 0x6c, 0x2e, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x6d, 0x75, 0x6c, 0x2e, 0x4d, 0x75, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x75, 0x6c, 0x2f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x75, 0x6c, 0x3a, 0x01, 0x2a, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x3b, 0x6d, 0x75, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mul_proto_rawDescOnce sync.Once
	file_proto_mul_proto_rawDescData = file_proto_mul_proto_rawDesc
)

func file_proto_mul_proto_rawDescGZIP() []byte {
	file_proto_mul_proto_rawDescOnce.Do(func() {
		file_proto_mul_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mul_proto_rawDescData)
	})
	return file_proto_mul_proto_rawDescData
}

var file_proto_mul_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_mul_proto_goTypes = []interface{}{
	(*MulRequest)(nil), // 0: mul.MulRequest
	(*MulReply)(nil),   // 1: mul.MulReply
}
var file_proto_mul_proto_depIdxs = []int32{
	0, // 0: mul.MulService.HandleMul:input_type -> mul.MulRequest
	1, // 1: mul.MulService.HandleMul:output_type -> mul.MulReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mul_proto_init() }
func file_proto_mul_proto_init() {
	if File_proto_mul_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mul_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MulRequest); i {
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
		file_proto_mul_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MulReply); i {
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
			RawDescriptor: file_proto_mul_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mul_proto_goTypes,
		DependencyIndexes: file_proto_mul_proto_depIdxs,
		MessageInfos:      file_proto_mul_proto_msgTypes,
	}.Build()
	File_proto_mul_proto = out.File
	file_proto_mul_proto_rawDesc = nil
	file_proto_mul_proto_goTypes = nil
	file_proto_mul_proto_depIdxs = nil
}
