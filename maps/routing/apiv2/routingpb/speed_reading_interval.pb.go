// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.2
// source: google/maps/routing/v2/speed_reading_interval.proto

package routingpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The classification of polyline speed based on traffic data.
type SpeedReadingInterval_Speed int32

const (
	// Default value. This value is unused.
	SpeedReadingInterval_SPEED_UNSPECIFIED SpeedReadingInterval_Speed = 0
	// Normal speed, no slowdown is detected.
	SpeedReadingInterval_NORMAL SpeedReadingInterval_Speed = 1
	// Slowdown detected, but no traffic jam formed.
	SpeedReadingInterval_SLOW SpeedReadingInterval_Speed = 2
	// Traffic jam detected.
	SpeedReadingInterval_TRAFFIC_JAM SpeedReadingInterval_Speed = 3
)

// Enum value maps for SpeedReadingInterval_Speed.
var (
	SpeedReadingInterval_Speed_name = map[int32]string{
		0: "SPEED_UNSPECIFIED",
		1: "NORMAL",
		2: "SLOW",
		3: "TRAFFIC_JAM",
	}
	SpeedReadingInterval_Speed_value = map[string]int32{
		"SPEED_UNSPECIFIED": 0,
		"NORMAL":            1,
		"SLOW":              2,
		"TRAFFIC_JAM":       3,
	}
)

func (x SpeedReadingInterval_Speed) Enum() *SpeedReadingInterval_Speed {
	p := new(SpeedReadingInterval_Speed)
	*p = x
	return p
}

func (x SpeedReadingInterval_Speed) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SpeedReadingInterval_Speed) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routing_v2_speed_reading_interval_proto_enumTypes[0].Descriptor()
}

func (SpeedReadingInterval_Speed) Type() protoreflect.EnumType {
	return &file_google_maps_routing_v2_speed_reading_interval_proto_enumTypes[0]
}

func (x SpeedReadingInterval_Speed) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SpeedReadingInterval_Speed.Descriptor instead.
func (SpeedReadingInterval_Speed) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_speed_reading_interval_proto_rawDescGZIP(), []int{0, 0}
}

// Traffic density indicator on a contiguous segment of a polyline or path.
// Given a path with points P_0, P_1, ... , P_N (zero-based index), the
// SpeedReadingInterval defines an interval and describes its traffic using the
// following categories.
type SpeedReadingInterval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The starting index of this interval in the polyline.
	StartPolylinePointIndex *int32 `protobuf:"varint,1,opt,name=start_polyline_point_index,json=startPolylinePointIndex,proto3,oneof" json:"start_polyline_point_index,omitempty"`
	// The ending index of this interval in the polyline.
	EndPolylinePointIndex *int32 `protobuf:"varint,2,opt,name=end_polyline_point_index,json=endPolylinePointIndex,proto3,oneof" json:"end_polyline_point_index,omitempty"`
	// Types that are assignable to SpeedType:
	//
	//	*SpeedReadingInterval_Speed_
	SpeedType isSpeedReadingInterval_SpeedType `protobuf_oneof:"speed_type"`
}

func (x *SpeedReadingInterval) Reset() {
	*x = SpeedReadingInterval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeedReadingInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeedReadingInterval) ProtoMessage() {}

func (x *SpeedReadingInterval) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeedReadingInterval.ProtoReflect.Descriptor instead.
func (*SpeedReadingInterval) Descriptor() ([]byte, []int) {
	return file_google_maps_routing_v2_speed_reading_interval_proto_rawDescGZIP(), []int{0}
}

func (x *SpeedReadingInterval) GetStartPolylinePointIndex() int32 {
	if x != nil && x.StartPolylinePointIndex != nil {
		return *x.StartPolylinePointIndex
	}
	return 0
}

func (x *SpeedReadingInterval) GetEndPolylinePointIndex() int32 {
	if x != nil && x.EndPolylinePointIndex != nil {
		return *x.EndPolylinePointIndex
	}
	return 0
}

func (m *SpeedReadingInterval) GetSpeedType() isSpeedReadingInterval_SpeedType {
	if m != nil {
		return m.SpeedType
	}
	return nil
}

func (x *SpeedReadingInterval) GetSpeed() SpeedReadingInterval_Speed {
	if x, ok := x.GetSpeedType().(*SpeedReadingInterval_Speed_); ok {
		return x.Speed
	}
	return SpeedReadingInterval_SPEED_UNSPECIFIED
}

type isSpeedReadingInterval_SpeedType interface {
	isSpeedReadingInterval_SpeedType()
}

type SpeedReadingInterval_Speed_ struct {
	// Traffic speed in this interval.
	Speed SpeedReadingInterval_Speed `protobuf:"varint,3,opt,name=speed,proto3,enum=google.maps.routing.v2.SpeedReadingInterval_Speed,oneof"`
}

func (*SpeedReadingInterval_Speed_) isSpeedReadingInterval_SpeedType() {}

var File_google_maps_routing_v2_speed_reading_interval_proto protoreflect.FileDescriptor

var file_google_maps_routing_v2_speed_reading_interval_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61,
	0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x22, 0xf3, 0x02,
	0x0a, 0x14, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x1a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x17, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x3c, 0x0a, 0x18, 0x65, 0x6e, 0x64, 0x5f,
	0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x15, 0x65, 0x6e,
	0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x12, 0x4a, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x64, 0x48, 0x00, 0x52, 0x05, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x22, 0x45, 0x0a, 0x05, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x50, 0x45, 0x45, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08,
	0x0a, 0x04, 0x53, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x46,
	0x46, 0x49, 0x43, 0x5f, 0x4a, 0x41, 0x4d, 0x10, 0x03, 0x42, 0x0c, 0x0a, 0x0a, 0x73, 0x70, 0x65,
	0x65, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x70,
	0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x42, 0xce, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x42, 0x19, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70,
	0x62, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xa2, 0x02,
	0x05, 0x47, 0x4d, 0x52, 0x56, 0x32, 0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x32, 0xca,
	0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f,
	0x75, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x3a, 0x3a, 0x4d, 0x61, 0x70, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67,
	0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routing_v2_speed_reading_interval_proto_rawDescOnce sync.Once
	file_google_maps_routing_v2_speed_reading_interval_proto_rawDescData = file_google_maps_routing_v2_speed_reading_interval_proto_rawDesc
)

func file_google_maps_routing_v2_speed_reading_interval_proto_rawDescGZIP() []byte {
	file_google_maps_routing_v2_speed_reading_interval_proto_rawDescOnce.Do(func() {
		file_google_maps_routing_v2_speed_reading_interval_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routing_v2_speed_reading_interval_proto_rawDescData)
	})
	return file_google_maps_routing_v2_speed_reading_interval_proto_rawDescData
}

var file_google_maps_routing_v2_speed_reading_interval_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_routing_v2_speed_reading_interval_proto_goTypes = []interface{}{
	(SpeedReadingInterval_Speed)(0), // 0: google.maps.routing.v2.SpeedReadingInterval.Speed
	(*SpeedReadingInterval)(nil),    // 1: google.maps.routing.v2.SpeedReadingInterval
}
var file_google_maps_routing_v2_speed_reading_interval_proto_depIdxs = []int32{
	0, // 0: google.maps.routing.v2.SpeedReadingInterval.speed:type_name -> google.maps.routing.v2.SpeedReadingInterval.Speed
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_maps_routing_v2_speed_reading_interval_proto_init() }
func file_google_maps_routing_v2_speed_reading_interval_proto_init() {
	if File_google_maps_routing_v2_speed_reading_interval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeedReadingInterval); i {
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
	file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SpeedReadingInterval_Speed_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routing_v2_speed_reading_interval_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routing_v2_speed_reading_interval_proto_goTypes,
		DependencyIndexes: file_google_maps_routing_v2_speed_reading_interval_proto_depIdxs,
		EnumInfos:         file_google_maps_routing_v2_speed_reading_interval_proto_enumTypes,
		MessageInfos:      file_google_maps_routing_v2_speed_reading_interval_proto_msgTypes,
	}.Build()
	File_google_maps_routing_v2_speed_reading_interval_proto = out.File
	file_google_maps_routing_v2_speed_reading_interval_proto_rawDesc = nil
	file_google_maps_routing_v2_speed_reading_interval_proto_goTypes = nil
	file_google_maps_routing_v2_speed_reading_interval_proto_depIdxs = nil
}
