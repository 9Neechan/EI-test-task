// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: get_stats.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос на получение статистики
type GetStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *int64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3,oneof" json:"user_id,omitempty"`          // Фильтр по user_id
	ServiceId     *int64                 `protobuf:"varint,2,opt,name=service_id,json=serviceId,proto3,oneof" json:"service_id,omitempty"` // Фильтр по service_id
	Page          int32                  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Limit         int32                  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	mi := &file_get_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_get_stats_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatsRequest) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *GetStatsRequest) GetServiceId() int64 {
	if x != nil && x.ServiceId != nil {
		return *x.ServiceId
	}
	return 0
}

func (x *GetStatsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetStatsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Stats         []*StatRecord          `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
	Total         float64                `protobuf:"fixed64,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	mi := &file_get_stats_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_get_stats_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_get_stats_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatsResponse) GetStats() []*StatRecord {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *GetStatsResponse) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// Структура статистики вызовов
type StatRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        int64                  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ServiceId     int64                  `protobuf:"varint,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Count         int64                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	TotalOneRec   float64                `protobuf:"fixed64,4,opt,name=total_one_rec,json=totalOneRec,proto3" json:"total_one_rec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatRecord) Reset() {
	*x = StatRecord{}
	mi := &file_get_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatRecord) ProtoMessage() {}

func (x *StatRecord) ProtoReflect() protoreflect.Message {
	mi := &file_get_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatRecord.ProtoReflect.Descriptor instead.
func (*StatRecord) Descriptor() ([]byte, []int) {
	return file_get_stats_proto_rawDescGZIP(), []int{2}
}

func (x *StatRecord) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *StatRecord) GetServiceId() int64 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

func (x *StatRecord) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *StatRecord) GetTotalOneRec() float64 {
	if x != nil {
		return x.TotalOneRec
	}
	return 0
}

var File_get_stats_proto protoreflect.FileDescriptor

var file_get_stats_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x7e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x6e, 0x65, 0x5f,
	0x72, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4f, 0x6e, 0x65, 0x52, 0x65, 0x63, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x39, 0x4e, 0x65, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x2f, 0x45, 0x49,
	0x2d, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_get_stats_proto_rawDescOnce sync.Once
	file_get_stats_proto_rawDescData []byte
)

func file_get_stats_proto_rawDescGZIP() []byte {
	file_get_stats_proto_rawDescOnce.Do(func() {
		file_get_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_get_stats_proto_rawDesc), len(file_get_stats_proto_rawDesc)))
	})
	return file_get_stats_proto_rawDescData
}

var file_get_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_get_stats_proto_goTypes = []any{
	(*GetStatsRequest)(nil),  // 0: proto.GetStatsRequest
	(*GetStatsResponse)(nil), // 1: proto.GetStatsResponse
	(*StatRecord)(nil),       // 2: proto.StatRecord
}
var file_get_stats_proto_depIdxs = []int32{
	2, // 0: proto.GetStatsResponse.stats:type_name -> proto.StatRecord
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_get_stats_proto_init() }
func file_get_stats_proto_init() {
	if File_get_stats_proto != nil {
		return
	}
	file_get_stats_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_get_stats_proto_rawDesc), len(file_get_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_stats_proto_goTypes,
		DependencyIndexes: file_get_stats_proto_depIdxs,
		MessageInfos:      file_get_stats_proto_msgTypes,
	}.Build()
	File_get_stats_proto = out.File
	file_get_stats_proto_goTypes = nil
	file_get_stats_proto_depIdxs = nil
}
