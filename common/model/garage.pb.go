// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.14.0
// source: garage.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GarageCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GarageCoordinate) Reset() {
	*x = GarageCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageCoordinate) ProtoMessage() {}

func (x *GarageCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageCoordinate.ProtoReflect.Descriptor instead.
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{0}
}

func (x *GarageCoordinate) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GarageCoordinate) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Garage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate *GarageCoordinate `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *Garage) Reset() {
	*x = Garage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Garage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Garage) ProtoMessage() {}

func (x *Garage) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Garage.ProtoReflect.Descriptor instead.
func (*Garage) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{1}
}

func (x *Garage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Garage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Garage) GetCoordinate() *GarageCoordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type GarageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Garage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GarageList) Reset() {
	*x = GarageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageList) ProtoMessage() {}

func (x *GarageList) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageList.ProtoReflect.Descriptor instead.
func (*GarageList) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{2}
}

func (x *GarageList) GetList() []*Garage {
	if x != nil {
		return x.List
	}
	return nil
}

type GarageListByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List map[string]*GarageList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GarageListByUser) Reset() {
	*x = GarageListByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageListByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageListByUser) ProtoMessage() {}

func (x *GarageListByUser) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageListByUser.ProtoReflect.Descriptor instead.
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{3}
}

func (x *GarageListByUser) GetList() map[string]*GarageList {
	if x != nil {
		return x.List
	}
	return nil
}

type GarageUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GarageUserId) Reset() {
	*x = GarageUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageUserId) ProtoMessage() {}

func (x *GarageUserId) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageUserId.ProtoReflect.Descriptor instead.
func (*GarageUserId) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{4}
}

func (x *GarageUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GarageAndUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Garage *Garage `protobuf:"bytes,2,opt,name=garage,proto3" json:"garage,omitempty"`
}

func (x *GarageAndUserId) Reset() {
	*x = GarageAndUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_garage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageAndUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageAndUserId) ProtoMessage() {}

func (x *GarageAndUserId) ProtoReflect() protoreflect.Message {
	mi := &file_garage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageAndUserId.ProtoReflect.Descriptor instead.
func (*GarageAndUserId) Descriptor() ([]byte, []int) {
	return file_garage_proto_rawDescGZIP(), []int{5}
}

func (x *GarageAndUserId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GarageAndUserId) GetGarage() *Garage {
	if x != nil {
		return x.Garage
	}
	return nil
}

var File_garage_proto protoreflect.FileDescriptor

var file_garage_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x10, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x65, 0x0a, 0x06, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37,
	0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x22, 0x2f, 0x0a, 0x0a, 0x47, 0x61, 0x72, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61,
	0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x1a, 0x4a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x27, 0x0a, 0x0c, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x0f, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x06, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x32, 0x74, 0x0a, 0x07,
	0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x13, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_garage_proto_rawDescOnce sync.Once
	file_garage_proto_rawDescData = file_garage_proto_rawDesc
)

func file_garage_proto_rawDescGZIP() []byte {
	file_garage_proto_rawDescOnce.Do(func() {
		file_garage_proto_rawDescData = protoimpl.X.CompressGZIP(file_garage_proto_rawDescData)
	})
	return file_garage_proto_rawDescData
}

var file_garage_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_garage_proto_goTypes = []interface{}{
	(*GarageCoordinate)(nil), // 0: model.GarageCoordinate
	(*Garage)(nil),           // 1: model.Garage
	(*GarageList)(nil),       // 2: model.GarageList
	(*GarageListByUser)(nil), // 3: model.GarageListByUser
	(*GarageUserId)(nil),     // 4: model.GarageUserId
	(*GarageAndUserId)(nil),  // 5: model.GarageAndUserId
	nil,                      // 6: model.GarageListByUser.ListEntry
	(*emptypb.Empty)(nil),    // 7: google.protobuf.Empty
}
var file_garage_proto_depIdxs = []int32{
	0, // 0: model.Garage.coordinate:type_name -> model.GarageCoordinate
	1, // 1: model.GarageList.list:type_name -> model.Garage
	6, // 2: model.GarageListByUser.list:type_name -> model.GarageListByUser.ListEntry
	1, // 3: model.GarageAndUserId.garage:type_name -> model.Garage
	2, // 4: model.GarageListByUser.ListEntry.value:type_name -> model.GarageList
	4, // 5: model.Garages.List:input_type -> model.GarageUserId
	5, // 6: model.Garages.Add:input_type -> model.GarageAndUserId
	2, // 7: model.Garages.List:output_type -> model.GarageList
	7, // 8: model.Garages.Add:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_garage_proto_init() }
func file_garage_proto_init() {
	if File_garage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_garage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageCoordinate); i {
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
		file_garage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Garage); i {
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
		file_garage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageList); i {
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
		file_garage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageListByUser); i {
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
		file_garage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageUserId); i {
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
		file_garage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageAndUserId); i {
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
			RawDescriptor: file_garage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_garage_proto_goTypes,
		DependencyIndexes: file_garage_proto_depIdxs,
		MessageInfos:      file_garage_proto_msgTypes,
	}.Build()
	File_garage_proto = out.File
	file_garage_proto_rawDesc = nil
	file_garage_proto_goTypes = nil
	file_garage_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GaragesClient is the client API for Garages service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GaragesClient interface {
	List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error)
	Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type garagesClient struct {
	cc grpc.ClientConnInterface
}

func NewGaragesClient(cc grpc.ClientConnInterface) GaragesClient {
	return &garagesClient{cc}
}

func (c *garagesClient) List(ctx context.Context, in *GarageUserId, opts ...grpc.CallOption) (*GarageList, error) {
	out := new(GarageList)
	err := c.cc.Invoke(ctx, "/model.Garages/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *garagesClient) Add(ctx context.Context, in *GarageAndUserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/model.Garages/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GaragesServer is the server API for Garages service.
type GaragesServer interface {
	List(context.Context, *GarageUserId) (*GarageList, error)
	Add(context.Context, *GarageAndUserId) (*emptypb.Empty, error)
}

// UnimplementedGaragesServer can be embedded to have forward compatible implementations.
type UnimplementedGaragesServer struct {
}

func (*UnimplementedGaragesServer) List(context.Context, *GarageUserId) (*GarageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedGaragesServer) Add(context.Context, *GarageAndUserId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterGaragesServer(s *grpc.Server, srv GaragesServer) {
	s.RegisterService(&_Garages_serviceDesc, srv)
}

func _Garages_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).List(ctx, req.(*GarageUserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Garages_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GarageAndUserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GaragesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Garages/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GaragesServer).Add(ctx, req.(*GarageAndUserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Garages_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Garages",
	HandlerType: (*GaragesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Garages_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Garages_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "garage.proto",
}
