// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: tracking_shipment.proto

package service_v1

import (
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

type ShipmentStatus int32

const (
	ShipmentStatus_SHIPMENT_STATUS_UNSPECIFIED ShipmentStatus = 0
	ShipmentStatus_SHIPMENT_STATUS_REGISTERED  ShipmentStatus = 1
	ShipmentStatus_SHIPMENT_STATUS_ONROUTE     ShipmentStatus = 2
	ShipmentStatus_SHIPMENT_STATUS_SHIPPED     ShipmentStatus = 3
	ShipmentStatus_SHIPMENT_STATUS_NOT_SHIPPED ShipmentStatus = 4
)

// Enum value maps for ShipmentStatus.
var (
	ShipmentStatus_name = map[int32]string{
		0: "SHIPMENT_STATUS_UNSPECIFIED",
		1: "SHIPMENT_STATUS_REGISTERED",
		2: "SHIPMENT_STATUS_ONROUTE",
		3: "SHIPMENT_STATUS_SHIPPED",
		4: "SHIPMENT_STATUS_NOT_SHIPPED",
	}
	ShipmentStatus_value = map[string]int32{
		"SHIPMENT_STATUS_UNSPECIFIED": 0,
		"SHIPMENT_STATUS_REGISTERED":  1,
		"SHIPMENT_STATUS_ONROUTE":     2,
		"SHIPMENT_STATUS_SHIPPED":     3,
		"SHIPMENT_STATUS_NOT_SHIPPED": 4,
	}
)

func (x ShipmentStatus) Enum() *ShipmentStatus {
	p := new(ShipmentStatus)
	*p = x
	return p
}

func (x ShipmentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShipmentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_tracking_shipment_proto_enumTypes[0].Descriptor()
}

func (ShipmentStatus) Type() protoreflect.EnumType {
	return &file_tracking_shipment_proto_enumTypes[0]
}

func (x ShipmentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShipmentStatus.Descriptor instead.
func (ShipmentStatus) EnumDescriptor() ([]byte, []int) {
	return file_tracking_shipment_proto_rawDescGZIP(), []int{0}
}

type Shipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipmentId  string                 `protobuf:"bytes,1,opt,name=shipmentId,proto3" json:"shipmentId,omitempty"`
	Status      ShipmentStatus         `protobuf:"varint,2,opt,name=status,proto3,enum=liaison.v1.ShipmentStatus" json:"status,omitempty"`
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	Location    string                 `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Shipment) Reset() {
	*x = Shipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_shipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipment) ProtoMessage() {}

func (x *Shipment) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_shipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipment.ProtoReflect.Descriptor instead.
func (*Shipment) Descriptor() ([]byte, []int) {
	return file_tracking_shipment_proto_rawDescGZIP(), []int{0}
}

func (x *Shipment) GetShipmentId() string {
	if x != nil {
		return x.ShipmentId
	}
	return ""
}

func (x *Shipment) GetStatus() ShipmentStatus {
	if x != nil {
		return x.Status
	}
	return ShipmentStatus_SHIPMENT_STATUS_UNSPECIFIED
}

func (x *Shipment) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *Shipment) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_tracking_shipment_proto protoreflect.FileDescriptor

var file_tracking_shipment_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c, 0x69, 0x61, 0x69, 0x73,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6c, 0x69, 0x61, 0x69, 0x73, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0xac, 0x01, 0x0a, 0x0e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45,
	0x52, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4e, 0x52, 0x4f, 0x55, 0x54, 0x45,
	0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x1f, 0x0a, 0x1b, 0x53, 0x48, 0x49, 0x50, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x04,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x65, 0x61, 0x64, 0x66, 0x61, 0x73, 0x74, 0x69, 0x65, 0x2f, 0x6c, 0x69, 0x61, 0x69, 0x73,
	0x6f, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_shipment_proto_rawDescOnce sync.Once
	file_tracking_shipment_proto_rawDescData = file_tracking_shipment_proto_rawDesc
)

func file_tracking_shipment_proto_rawDescGZIP() []byte {
	file_tracking_shipment_proto_rawDescOnce.Do(func() {
		file_tracking_shipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_shipment_proto_rawDescData)
	})
	return file_tracking_shipment_proto_rawDescData
}

var file_tracking_shipment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tracking_shipment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tracking_shipment_proto_goTypes = []interface{}{
	(ShipmentStatus)(0),           // 0: liaison.v1.ShipmentStatus
	(*Shipment)(nil),              // 1: liaison.v1.Shipment
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_tracking_shipment_proto_depIdxs = []int32{
	0, // 0: liaison.v1.Shipment.status:type_name -> liaison.v1.ShipmentStatus
	2, // 1: liaison.v1.Shipment.lastUpdated:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tracking_shipment_proto_init() }
func file_tracking_shipment_proto_init() {
	if File_tracking_shipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_shipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipment); i {
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
			RawDescriptor: file_tracking_shipment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracking_shipment_proto_goTypes,
		DependencyIndexes: file_tracking_shipment_proto_depIdxs,
		EnumInfos:         file_tracking_shipment_proto_enumTypes,
		MessageInfos:      file_tracking_shipment_proto_msgTypes,
	}.Build()
	File_tracking_shipment_proto = out.File
	file_tracking_shipment_proto_rawDesc = nil
	file_tracking_shipment_proto_goTypes = nil
	file_tracking_shipment_proto_depIdxs = nil
}
