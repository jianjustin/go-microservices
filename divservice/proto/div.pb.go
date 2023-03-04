// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: proto/div.proto

package div

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

type DivRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A float64 `protobuf:"fixed64,1,opt,name=a,proto3" json:"a,omitempty"`
	B float64 `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *DivRequest) Reset() {
	*x = DivRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_div_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivRequest) ProtoMessage() {}

func (x *DivRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_div_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivRequest.ProtoReflect.Descriptor instead.
func (*DivRequest) Descriptor() ([]byte, []int) {
	return file_proto_div_proto_rawDescGZIP(), []int{0}
}

func (x *DivRequest) GetA() float64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *DivRequest) GetB() float64 {
	if x != nil {
		return x.B
	}
	return 0
}

type DivReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float64 `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DivReply) Reset() {
	*x = DivReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_div_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivReply) ProtoMessage() {}

func (x *DivReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_div_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivReply.ProtoReflect.Descriptor instead.
func (*DivReply) Descriptor() ([]byte, []int) {
	return file_proto_div_proto_rawDescGZIP(), []int{1}
}

func (x *DivReply) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_div_proto protoreflect.FileDescriptor

var file_proto_div_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x64, 0x69, 0x76, 0x22, 0x28, 0x0a, 0x0a, 0x44, 0x69, 0x76, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x62,
	0x22, 0x22, 0x0a, 0x08, 0x44, 0x69, 0x76, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x32, 0x3b, 0x0a, 0x0a, 0x44, 0x69, 0x76, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x44, 0x69, 0x76, 0x12,
	0x0f, 0x2e, 0x64, 0x69, 0x76, 0x2e, 0x44, 0x69, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x64, 0x69, 0x76, 0x2e, 0x44, 0x69, 0x76, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x64, 0x69, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_div_proto_rawDescOnce sync.Once
	file_proto_div_proto_rawDescData = file_proto_div_proto_rawDesc
)

func file_proto_div_proto_rawDescGZIP() []byte {
	file_proto_div_proto_rawDescOnce.Do(func() {
		file_proto_div_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_div_proto_rawDescData)
	})
	return file_proto_div_proto_rawDescData
}

var file_proto_div_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_div_proto_goTypes = []interface{}{
	(*DivRequest)(nil), // 0: div.DivRequest
	(*DivReply)(nil),   // 1: div.DivReply
}
var file_proto_div_proto_depIdxs = []int32{
	0, // 0: div.DivService.HandleDiv:input_type -> div.DivRequest
	1, // 1: div.DivService.HandleDiv:output_type -> div.DivReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_div_proto_init() }
func file_proto_div_proto_init() {
	if File_proto_div_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_div_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivRequest); i {
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
		file_proto_div_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivReply); i {
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
			RawDescriptor: file_proto_div_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_div_proto_goTypes,
		DependencyIndexes: file_proto_div_proto_depIdxs,
		MessageInfos:      file_proto_div_proto_msgTypes,
	}.Build()
	File_proto_div_proto = out.File
	file_proto_div_proto_rawDesc = nil
	file_proto_div_proto_goTypes = nil
	file_proto_div_proto_depIdxs = nil
}