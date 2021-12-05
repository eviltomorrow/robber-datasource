// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: datasource.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code            string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Open            float64 `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`
	YesterdayClosed float64 `protobuf:"fixed64,4,opt,name=yesterday_closed,json=yesterdayClosed,proto3" json:"yesterday_closed,omitempty"`
	Latest          float64 `protobuf:"fixed64,5,opt,name=latest,proto3" json:"latest,omitempty"`
	High            float64 `protobuf:"fixed64,6,opt,name=high,proto3" json:"high,omitempty"`
	Low             float64 `protobuf:"fixed64,7,opt,name=low,proto3" json:"low,omitempty"`
	Volume          uint64  `protobuf:"varint,8,opt,name=volume,proto3" json:"volume,omitempty"`
	Account         float64 `protobuf:"fixed64,9,opt,name=account,proto3" json:"account,omitempty"`
	Date            string  `protobuf:"bytes,10,opt,name=date,proto3" json:"date,omitempty"`
	Time            string  `protobuf:"bytes,11,opt,name=time,proto3" json:"time,omitempty"`
	Suspend         string  `protobuf:"bytes,12,opt,name=suspend,proto3" json:"suspend,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_datasource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_datasource_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Metadata) GetYesterdayClosed() float64 {
	if x != nil {
		return x.YesterdayClosed
	}
	return 0
}

func (x *Metadata) GetLatest() float64 {
	if x != nil {
		return x.Latest
	}
	return 0
}

func (x *Metadata) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Metadata) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Metadata) GetVolume() uint64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Metadata) GetAccount() float64 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *Metadata) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Metadata) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Metadata) GetSuspend() string {
	if x != nil {
		return x.Suspend
	}
	return ""
}

var File_datasource_proto protoreflect.FileDescriptor

var file_datasource_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x08,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6f, 0x70, 0x65, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61,
	0x79, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f,
	0x79, 0x65, 0x73, 0x74, 0x65, 0x72, 0x64, 0x61, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x6f, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x73, 0x70, 0x65,
	0x6e, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e,
	0x64, 0x32, 0xd3, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x07, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x50, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x14, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datasource_proto_rawDescOnce sync.Once
	file_datasource_proto_rawDescData = file_datasource_proto_rawDesc
)

func file_datasource_proto_rawDescGZIP() []byte {
	file_datasource_proto_rawDescOnce.Do(func() {
		file_datasource_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasource_proto_rawDescData)
	})
	return file_datasource_proto_rawDescData
}

var file_datasource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datasource_proto_goTypes = []interface{}{
	(*Metadata)(nil),               // 0: datasource.Metadata
	(*emptypb.Empty)(nil),          // 1: google.protobuf.Empty
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_datasource_proto_depIdxs = []int32{
	1, // 0: datasource.Service.Version:input_type -> google.protobuf.Empty
	1, // 1: datasource.Service.Collect:input_type -> google.protobuf.Empty
	2, // 2: datasource.Service.PullData:input_type -> google.protobuf.StringValue
	2, // 3: datasource.Service.Version:output_type -> google.protobuf.StringValue
	2, // 4: datasource.Service.Collect:output_type -> google.protobuf.StringValue
	0, // 5: datasource.Service.PullData:output_type -> datasource.Metadata
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_datasource_proto_init() }
func file_datasource_proto_init() {
	if File_datasource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasource_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
			RawDescriptor: file_datasource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datasource_proto_goTypes,
		DependencyIndexes: file_datasource_proto_depIdxs,
		MessageInfos:      file_datasource_proto_msgTypes,
	}.Build()
	File_datasource_proto = out.File
	file_datasource_proto_rawDesc = nil
	file_datasource_proto_goTypes = nil
	file_datasource_proto_depIdxs = nil
}