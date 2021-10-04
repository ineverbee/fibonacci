// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: fibonacci.proto

package fibonacci

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

type IntRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *IntRequest) Reset() {
	*x = IntRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fibonacci_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntRequest) ProtoMessage() {}

func (x *IntRequest) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntRequest.ProtoReflect.Descriptor instead.
func (*IntRequest) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{0}
}

func (x *IntRequest) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *IntRequest) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type IntSliceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []uint32 `protobuf:"varint,1,rep,packed,name=result,proto3" json:"result,omitempty"`
}

func (x *IntSliceResponse) Reset() {
	*x = IntSliceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fibonacci_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntSliceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntSliceResponse) ProtoMessage() {}

func (x *IntSliceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fibonacci_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntSliceResponse.ProtoReflect.Descriptor instead.
func (*IntSliceResponse) Descriptor() ([]byte, []int) {
	return file_fibonacci_proto_rawDescGZIP(), []int{1}
}

func (x *IntSliceResponse) GetResult() []uint32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_fibonacci_proto protoreflect.FileDescriptor

var file_fibonacci_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a,
	0x0a, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79, 0x22, 0x2a, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x53, 0x6c,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x5f, 0x0a, 0x09, 0x46, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69,
	0x12, 0x52, 0x0a, 0x04, 0x46, 0x69, 0x62, 0x6f, 0x12, 0x18, 0x2e, 0x66, 0x69, 0x62, 0x6f, 0x6e,
	0x61, 0x63, 0x63, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x66, 0x69, 0x62, 0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x74, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05, 0x2f, 0x66, 0x69, 0x62,
	0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x65, 0x76, 0x65, 0x72, 0x62, 0x65, 0x65, 0x2f, 0x66, 0x69, 0x62,
	0x6f, 0x6e, 0x61, 0x63, 0x63, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fibonacci_proto_rawDescOnce sync.Once
	file_fibonacci_proto_rawDescData = file_fibonacci_proto_rawDesc
)

func file_fibonacci_proto_rawDescGZIP() []byte {
	file_fibonacci_proto_rawDescOnce.Do(func() {
		file_fibonacci_proto_rawDescData = protoimpl.X.CompressGZIP(file_fibonacci_proto_rawDescData)
	})
	return file_fibonacci_proto_rawDescData
}

var file_fibonacci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fibonacci_proto_goTypes = []interface{}{
	(*IntRequest)(nil),       // 0: fibonacci.v1.IntRequest
	(*IntSliceResponse)(nil), // 1: fibonacci.v1.IntSliceResponse
}
var file_fibonacci_proto_depIdxs = []int32{
	0, // 0: fibonacci.v1.Fibonacci.Fibo:input_type -> fibonacci.v1.IntRequest
	1, // 1: fibonacci.v1.Fibonacci.Fibo:output_type -> fibonacci.v1.IntSliceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_fibonacci_proto_init() }
func file_fibonacci_proto_init() {
	if File_fibonacci_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fibonacci_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntRequest); i {
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
		file_fibonacci_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntSliceResponse); i {
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
			RawDescriptor: file_fibonacci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fibonacci_proto_goTypes,
		DependencyIndexes: file_fibonacci_proto_depIdxs,
		MessageInfos:      file_fibonacci_proto_msgTypes,
	}.Build()
	File_fibonacci_proto = out.File
	file_fibonacci_proto_rawDesc = nil
	file_fibonacci_proto_goTypes = nil
	file_fibonacci_proto_depIdxs = nil
}
