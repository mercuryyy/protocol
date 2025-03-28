// Copyright 2023 LiveKit, Inc.
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
// 	protoc-gen-go v1.36.5
// 	protoc        v4.23.4
// source: livekit_webhook.proto

package livekit

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

type WebhookEvent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// one of room_started, room_finished, participant_joined, participant_left,
	// track_published, track_unpublished, egress_started, egress_updated, egress_ended,
	// ingress_started, ingress_ended
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Room  *Room  `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	// set when event is participant_* or track_*
	Participant *ParticipantInfo `protobuf:"bytes,3,opt,name=participant,proto3" json:"participant,omitempty"`
	// set when event is egress_*
	EgressInfo *EgressInfo `protobuf:"bytes,9,opt,name=egress_info,json=egressInfo,proto3" json:"egress_info,omitempty"`
	// set when event is ingress_*
	IngressInfo *IngressInfo `protobuf:"bytes,10,opt,name=ingress_info,json=ingressInfo,proto3" json:"ingress_info,omitempty"`
	// set when event is track_*
	Track *TrackInfo `protobuf:"bytes,8,opt,name=track,proto3" json:"track,omitempty"`
	// unique event uuid
	Id string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// timestamp in seconds
	CreatedAt int64 `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Deprecated: Marked as deprecated in livekit_webhook.proto.
	NumDropped    int32 `protobuf:"varint,11,opt,name=num_dropped,json=numDropped,proto3" json:"num_dropped,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WebhookEvent) Reset() {
	*x = WebhookEvent{}
	mi := &file_livekit_webhook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WebhookEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookEvent) ProtoMessage() {}

func (x *WebhookEvent) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_webhook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookEvent.ProtoReflect.Descriptor instead.
func (*WebhookEvent) Descriptor() ([]byte, []int) {
	return file_livekit_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *WebhookEvent) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *WebhookEvent) GetParticipant() *ParticipantInfo {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *WebhookEvent) GetEgressInfo() *EgressInfo {
	if x != nil {
		return x.EgressInfo
	}
	return nil
}

func (x *WebhookEvent) GetIngressInfo() *IngressInfo {
	if x != nil {
		return x.IngressInfo
	}
	return nil
}

func (x *WebhookEvent) GetTrack() *TrackInfo {
	if x != nil {
		return x.Track
	}
	return nil
}

func (x *WebhookEvent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WebhookEvent) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

// Deprecated: Marked as deprecated in livekit_webhook.proto.
func (x *WebhookEvent) GetNumDropped() int32 {
	if x != nil {
		return x.NumDropped
	}
	return 0
}

var File_livekit_webhook_proto protoreflect.FileDescriptor

var file_livekit_webhook_proto_rawDesc = string([]byte{
	0x0a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x02, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3a, 0x0a,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x65, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x37, 0x0a, 0x0c, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x69, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69,
	0x74, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x23, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x44,
	0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x42, 0x46, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0xaa, 0x02, 0x0d,
	0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xea, 0x02, 0x0e,
	0x4c, 0x69, 0x76, 0x65, 0x4b, 0x69, 0x74, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_livekit_webhook_proto_rawDescOnce sync.Once
	file_livekit_webhook_proto_rawDescData []byte
)

func file_livekit_webhook_proto_rawDescGZIP() []byte {
	file_livekit_webhook_proto_rawDescOnce.Do(func() {
		file_livekit_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_livekit_webhook_proto_rawDesc), len(file_livekit_webhook_proto_rawDesc)))
	})
	return file_livekit_webhook_proto_rawDescData
}

var file_livekit_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_livekit_webhook_proto_goTypes = []any{
	(*WebhookEvent)(nil),    // 0: livekit.WebhookEvent
	(*Room)(nil),            // 1: livekit.Room
	(*ParticipantInfo)(nil), // 2: livekit.ParticipantInfo
	(*EgressInfo)(nil),      // 3: livekit.EgressInfo
	(*IngressInfo)(nil),     // 4: livekit.IngressInfo
	(*TrackInfo)(nil),       // 5: livekit.TrackInfo
}
var file_livekit_webhook_proto_depIdxs = []int32{
	1, // 0: livekit.WebhookEvent.room:type_name -> livekit.Room
	2, // 1: livekit.WebhookEvent.participant:type_name -> livekit.ParticipantInfo
	3, // 2: livekit.WebhookEvent.egress_info:type_name -> livekit.EgressInfo
	4, // 3: livekit.WebhookEvent.ingress_info:type_name -> livekit.IngressInfo
	5, // 4: livekit.WebhookEvent.track:type_name -> livekit.TrackInfo
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_livekit_webhook_proto_init() }
func file_livekit_webhook_proto_init() {
	if File_livekit_webhook_proto != nil {
		return
	}
	file_livekit_models_proto_init()
	file_livekit_egress_proto_init()
	file_livekit_ingress_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_livekit_webhook_proto_rawDesc), len(file_livekit_webhook_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_webhook_proto_goTypes,
		DependencyIndexes: file_livekit_webhook_proto_depIdxs,
		MessageInfos:      file_livekit_webhook_proto_msgTypes,
	}.Build()
	File_livekit_webhook_proto = out.File
	file_livekit_webhook_proto_goTypes = nil
	file_livekit_webhook_proto_depIdxs = nil
}
