// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: restaurant_like_service.proto

package proto

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

type RestaurantLikeStatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResIds []int32 `protobuf:"varint,1,rep,packed,name=resIds,proto3" json:"resIds,omitempty"`
}

func (x *RestaurantLikeStatRequest) Reset() {
	*x = RestaurantLikeStatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_like_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantLikeStatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantLikeStatRequest) ProtoMessage() {}

func (x *RestaurantLikeStatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_like_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantLikeStatRequest.ProtoReflect.Descriptor instead.
func (*RestaurantLikeStatRequest) Descriptor() ([]byte, []int) {
	return file_restaurant_like_service_proto_rawDescGZIP(), []int{0}
}

func (x *RestaurantLikeStatRequest) GetResIds() []int32 {
	if x != nil {
		return x.ResIds
	}
	return nil
}

type RestaurantLikeStatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result map[int32]int32 `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *RestaurantLikeStatResponse) Reset() {
	*x = RestaurantLikeStatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_like_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantLikeStatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantLikeStatResponse) ProtoMessage() {}

func (x *RestaurantLikeStatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_like_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantLikeStatResponse.ProtoReflect.Descriptor instead.
func (*RestaurantLikeStatResponse) Descriptor() ([]byte, []int) {
	return file_restaurant_like_service_proto_rawDescGZIP(), []int{1}
}

func (x *RestaurantLikeStatResponse) GetResult() map[int32]int32 {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_restaurant_like_service_proto protoreflect.FileDescriptor

var file_restaurant_like_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6b,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x64, 0x65, 0x6d, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e,
	0x74, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x49, 0x64, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x1a, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x39, 0x0a,
	0x0b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x94, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x7b, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1f, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x6b,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64,
	0x65, 0x6d, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x4c, 0x69,
	0x6b, 0x65, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_restaurant_like_service_proto_rawDescOnce sync.Once
	file_restaurant_like_service_proto_rawDescData = file_restaurant_like_service_proto_rawDesc
)

func file_restaurant_like_service_proto_rawDescGZIP() []byte {
	file_restaurant_like_service_proto_rawDescOnce.Do(func() {
		file_restaurant_like_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_like_service_proto_rawDescData)
	})
	return file_restaurant_like_service_proto_rawDescData
}

var file_restaurant_like_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_restaurant_like_service_proto_goTypes = []interface{}{
	(*RestaurantLikeStatRequest)(nil),  // 0: demo.RestaurantLikeStatRequest
	(*RestaurantLikeStatResponse)(nil), // 1: demo.RestaurantLikeStatResponse
	nil,                                // 2: demo.RestaurantLikeStatResponse.ResultEntry
}
var file_restaurant_like_service_proto_depIdxs = []int32{
	2, // 0: demo.RestaurantLikeStatResponse.result:type_name -> demo.RestaurantLikeStatResponse.ResultEntry
	0, // 1: demo.RestaurantLikeService.GetRestaurantLikeStat:input_type -> demo.RestaurantLikeStatRequest
	1, // 2: demo.RestaurantLikeService.GetRestaurantLikeStat:output_type -> demo.RestaurantLikeStatResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_restaurant_like_service_proto_init() }
func file_restaurant_like_service_proto_init() {
	if File_restaurant_like_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_like_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantLikeStatRequest); i {
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
		file_restaurant_like_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantLikeStatResponse); i {
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
			RawDescriptor: file_restaurant_like_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_like_service_proto_goTypes,
		DependencyIndexes: file_restaurant_like_service_proto_depIdxs,
		MessageInfos:      file_restaurant_like_service_proto_msgTypes,
	}.Build()
	File_restaurant_like_service_proto = out.File
	file_restaurant_like_service_proto_rawDesc = nil
	file_restaurant_like_service_proto_goTypes = nil
	file_restaurant_like_service_proto_depIdxs = nil
}