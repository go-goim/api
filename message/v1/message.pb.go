// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: message/v1/message.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	response "github.com/go-goim/api/transport/response"
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

type MessageContentType int32

const (
	MessageContentType_UnknownContentType MessageContentType = 0
	MessageContentType_Text               MessageContentType = 1
	MessageContentType_Image              MessageContentType = 2
	MessageContentType_Voice              MessageContentType = 3
)

// Enum value maps for MessageContentType.
var (
	MessageContentType_name = map[int32]string{
		0: "UnknownContentType",
		1: "Text",
		2: "Image",
		3: "Voice",
	}
	MessageContentType_value = map[string]int32{
		"UnknownContentType": 0,
		"Text":               1,
		"Image":              2,
		"Voice":              3,
	}
)

func (x MessageContentType) Enum() *MessageContentType {
	p := new(MessageContentType)
	*p = x
	return p
}

func (x MessageContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_v1_message_proto_enumTypes[0].Descriptor()
}

func (MessageContentType) Type() protoreflect.EnumType {
	return &file_message_v1_message_proto_enumTypes[0]
}

func (x MessageContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageContentType.Descriptor instead.
func (MessageContentType) EnumDescriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{0}
}

type PushMessageType int32

const (
	PushMessageType_UnknownUserType PushMessageType = 0
	// user to user
	PushMessageType_User PushMessageType = 1
	// user to group
	PushMessageType_Group PushMessageType = 2
	// global broadcast
	PushMessageType_Broadcast PushMessageType = 3
)

// Enum value maps for PushMessageType.
var (
	PushMessageType_name = map[int32]string{
		0: "UnknownUserType",
		1: "User",
		2: "Group",
		3: "Broadcast",
	}
	PushMessageType_value = map[string]int32{
		"UnknownUserType": 0,
		"User":            1,
		"Group":           2,
		"Broadcast":       3,
	}
)

func (x PushMessageType) Enum() *PushMessageType {
	p := new(PushMessageType)
	*p = x
	return p
}

func (x PushMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_v1_message_proto_enumTypes[1].Descriptor()
}

func (PushMessageType) Type() protoreflect.EnumType {
	return &file_message_v1_message_proto_enumTypes[1]
}

func (x PushMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushMessageType.Descriptor instead.
func (PushMessageType) EnumDescriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{1}
}

// SendMessageReq receive data from gateway
type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser    string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser      string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	ContentType MessageContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content     string             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageReq) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *SendMessageReq) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *SendMessageReq) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *SendMessageReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// SendMessageResp is response body for sendMessage
type SendMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *response.BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	// MsgSeq is unique seq of a message
	MsgSeq string `protobuf:"bytes,3,opt,name=msg_seq,json=msgSeq,proto3" json:"msg_seq,omitempty"`
}

func (x *SendMessageResp) Reset() {
	*x = SendMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResp) ProtoMessage() {}

func (x *SendMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResp.ProtoReflect.Descriptor instead.
func (*SendMessageResp) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResp) GetResponse() *response.BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *SendMessageResp) GetMsgSeq() string {
	if x != nil {
		return x.MsgSeq
	}
	return ""
}

// MqMessage is message protocol when pub/sub msg to mq
type MqMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser        string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser          string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	PushMessageType PushMessageType    `protobuf:"varint,3,opt,name=push_message_type,json=pushMessageType,proto3,enum=api.message.v1.PushMessageType" json:"push_message_type,omitempty"`
	ContentType     MessageContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content         string             `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *MqMessage) Reset() {
	*x = MqMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MqMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MqMessage) ProtoMessage() {}

func (x *MqMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MqMessage.ProtoReflect.Descriptor instead.
func (*MqMessage) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *MqMessage) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *MqMessage) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *MqMessage) GetPushMessageType() PushMessageType {
	if x != nil {
		return x.PushMessageType
	}
	return PushMessageType_UnknownUserType
}

func (x *MqMessage) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *MqMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// PushMessage use for push a message to persistence connection server
type PushMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser        string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser          string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	PushMessageType PushMessageType    `protobuf:"varint,3,opt,name=push_message_type,json=pushMessageType,proto3,enum=api.message.v1.PushMessageType" json:"push_message_type,omitempty"`
	ContentType     MessageContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content         string             `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	MsgSeq          string             `protobuf:"bytes,6,opt,name=msg_seq,json=msgSeq,proto3" json:"msg_seq,omitempty"`
}

func (x *PushMessageReq) Reset() {
	*x = PushMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageReq) ProtoMessage() {}

func (x *PushMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageReq.ProtoReflect.Descriptor instead.
func (*PushMessageReq) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *PushMessageReq) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *PushMessageReq) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *PushMessageReq) GetPushMessageType() PushMessageType {
	if x != nil {
		return x.PushMessageType
	}
	return PushMessageType_UnknownUserType
}

func (x *PushMessageReq) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *PushMessageReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PushMessageReq) GetMsgSeq() string {
	if x != nil {
		return x.MsgSeq
	}
	return ""
}

// BriefMessage is message with basic information
type BriefMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser    string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser      string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	ContentType MessageContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content     string             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	MsgSeq      string             `protobuf:"bytes,5,opt,name=msg_seq,json=msgSeq,proto3" json:"msg_seq,omitempty"`
}

func (x *BriefMessage) Reset() {
	*x = BriefMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BriefMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BriefMessage) ProtoMessage() {}

func (x *BriefMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BriefMessage.ProtoReflect.Descriptor instead.
func (*BriefMessage) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *BriefMessage) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *BriefMessage) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *BriefMessage) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *BriefMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *BriefMessage) GetMsgSeq() string {
	if x != nil {
		return x.MsgSeq
	}
	return ""
}

type QueryOfflineMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LastMsgSeq string `protobuf:"bytes,2,opt,name=last_msg_seq,json=lastMsgSeq,proto3" json:"last_msg_seq,omitempty"`
	OnlyCount  bool   `protobuf:"varint,3,opt,name=onlyCount,proto3" json:"onlyCount,omitempty"`
	Page       int32  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageSize   int32  `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *QueryOfflineMessageReq) Reset() {
	*x = QueryOfflineMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOfflineMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOfflineMessageReq) ProtoMessage() {}

func (x *QueryOfflineMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOfflineMessageReq.ProtoReflect.Descriptor instead.
func (*QueryOfflineMessageReq) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *QueryOfflineMessageReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryOfflineMessageReq) GetLastMsgSeq() string {
	if x != nil {
		return x.LastMsgSeq
	}
	return ""
}

func (x *QueryOfflineMessageReq) GetOnlyCount() bool {
	if x != nil {
		return x.OnlyCount
	}
	return false
}

func (x *QueryOfflineMessageReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryOfflineMessageReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type QueryOfflineMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *response.BaseResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Total    int32                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Messages []*BriefMessage        `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *QueryOfflineMessageResp) Reset() {
	*x = QueryOfflineMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_v1_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOfflineMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOfflineMessageResp) ProtoMessage() {}

func (x *QueryOfflineMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_message_v1_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOfflineMessageResp.ProtoReflect.Descriptor instead.
func (*QueryOfflineMessageResp) Descriptor() ([]byte, []int) {
	return file_message_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *QueryOfflineMessageResp) GetResponse() *response.BaseResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *QueryOfflineMessageResp) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryOfflineMessageResp) GetMessages() []*BriefMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_message_v1_message_proto protoreflect.FileDescriptor

var file_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x09, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x14, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x14, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x4f, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0x80, 0x08, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x6c, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x22, 0xef, 0x01, 0x0a, 0x09, 0x4d, 0x71, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x11, 0x70, 0x75,
	0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8d, 0x02, 0x0a, 0x0e, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x4b, 0x0a, 0x11, 0x70, 0x75, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x70,
	0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x22, 0xbe, 0x01, 0x0a, 0x0c, 0x42, 0x72, 0x69,
	0x65, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72,
	0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x71, 0x22, 0xbf, 0x01, 0x0a, 0x16, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x14, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d,
	0x73, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x0a, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x53, 0x65,
	0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x6e, 0x6c, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x64, 0x28,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x17,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x40, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x38, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x72, 0x69, 0x65, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2a, 0x4c, 0x0a, 0x12, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10,
	0x01, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x56, 0x6f, 0x69, 0x63, 0x65, 0x10, 0x03, 0x2a, 0x4a, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x10, 0x03, 0x32, 0xac, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x32, 0x63, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x53, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x78, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x6c, 0x69,
	0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x66, 0x0a, 0x13, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f,
	0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_v1_message_proto_rawDescOnce sync.Once
	file_message_v1_message_proto_rawDescData = file_message_v1_message_proto_rawDesc
)

func file_message_v1_message_proto_rawDescGZIP() []byte {
	file_message_v1_message_proto_rawDescOnce.Do(func() {
		file_message_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_v1_message_proto_rawDescData)
	})
	return file_message_v1_message_proto_rawDescData
}

var file_message_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_v1_message_proto_goTypes = []interface{}{
	(MessageContentType)(0),         // 0: api.message.v1.MessageContentType
	(PushMessageType)(0),            // 1: api.message.v1.PushMessageType
	(*SendMessageReq)(nil),          // 2: api.message.v1.SendMessageReq
	(*SendMessageResp)(nil),         // 3: api.message.v1.SendMessageResp
	(*MqMessage)(nil),               // 4: api.message.v1.MqMessage
	(*PushMessageReq)(nil),          // 5: api.message.v1.PushMessageReq
	(*BriefMessage)(nil),            // 6: api.message.v1.BriefMessage
	(*QueryOfflineMessageReq)(nil),  // 7: api.message.v1.QueryOfflineMessageReq
	(*QueryOfflineMessageResp)(nil), // 8: api.message.v1.QueryOfflineMessageResp
	(*response.BaseResponse)(nil),   // 9: api.transport.response.BaseResponse
}
var file_message_v1_message_proto_depIdxs = []int32{
	0,  // 0: api.message.v1.SendMessageReq.content_type:type_name -> api.message.v1.MessageContentType
	9,  // 1: api.message.v1.SendMessageResp.response:type_name -> api.transport.response.BaseResponse
	1,  // 2: api.message.v1.MqMessage.push_message_type:type_name -> api.message.v1.PushMessageType
	0,  // 3: api.message.v1.MqMessage.content_type:type_name -> api.message.v1.MessageContentType
	1,  // 4: api.message.v1.PushMessageReq.push_message_type:type_name -> api.message.v1.PushMessageType
	0,  // 5: api.message.v1.PushMessageReq.content_type:type_name -> api.message.v1.MessageContentType
	0,  // 6: api.message.v1.BriefMessage.content_type:type_name -> api.message.v1.MessageContentType
	9,  // 7: api.message.v1.QueryOfflineMessageResp.response:type_name -> api.transport.response.BaseResponse
	6,  // 8: api.message.v1.QueryOfflineMessageResp.messages:type_name -> api.message.v1.BriefMessage
	2,  // 9: api.message.v1.SendMessager.SendMessage:input_type -> api.message.v1.SendMessageReq
	2,  // 10: api.message.v1.SendMessager.Broadcast:input_type -> api.message.v1.SendMessageReq
	5,  // 11: api.message.v1.PushMessager.PushMessage:input_type -> api.message.v1.PushMessageReq
	7,  // 12: api.message.v1.OfflineMessage.QueryOfflineMessage:input_type -> api.message.v1.QueryOfflineMessageReq
	3,  // 13: api.message.v1.SendMessager.SendMessage:output_type -> api.message.v1.SendMessageResp
	3,  // 14: api.message.v1.SendMessager.Broadcast:output_type -> api.message.v1.SendMessageResp
	9,  // 15: api.message.v1.PushMessager.PushMessage:output_type -> api.transport.response.BaseResponse
	8,  // 16: api.message.v1.OfflineMessage.QueryOfflineMessage:output_type -> api.message.v1.QueryOfflineMessageResp
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_message_v1_message_proto_init() }
func file_message_v1_message_proto_init() {
	if File_message_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReq); i {
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
		file_message_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResp); i {
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
		file_message_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MqMessage); i {
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
		file_message_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageReq); i {
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
		file_message_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BriefMessage); i {
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
		file_message_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOfflineMessageReq); i {
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
		file_message_v1_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOfflineMessageResp); i {
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
			RawDescriptor: file_message_v1_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_message_v1_message_proto_goTypes,
		DependencyIndexes: file_message_v1_message_proto_depIdxs,
		EnumInfos:         file_message_v1_message_proto_enumTypes,
		MessageInfos:      file_message_v1_message_proto_msgTypes,
	}.Build()
	File_message_v1_message_proto = out.File
	file_message_v1_message_proto_rawDesc = nil
	file_message_v1_message_proto_goTypes = nil
	file_message_v1_message_proto_depIdxs = nil
}
