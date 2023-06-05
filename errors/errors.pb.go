// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: errors/errors.proto

package errors

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

type ErrorCode int32

const (
	ErrorCode_OK ErrorCode = 0
	// common error codes
	ErrorCode_InternalError ErrorCode = 10001
	ErrorCode_InvalidParams ErrorCode = 10002
	ErrorCode_DBError       ErrorCode = 10003
	ErrorCode_CacheError    ErrorCode = 10004
	// user error codes
	ErrorCode_UserNotExist              ErrorCode = 20001
	ErrorCode_UserExist                 ErrorCode = 20002
	ErrorCode_InvalidUsernameOrPassword ErrorCode = 20003
	// relation error codes
	ErrorCode_RelationNotExist            ErrorCode = 30001
	ErrorCode_RelationExist               ErrorCode = 30002
	ErrorCode_InvalidUpdateRelationAction ErrorCode = 30003
	// friend request error codes
	ErrorCode_FriendRequestNotExist    ErrorCode = 31001
	ErrorCode_FriendRequestStatusError ErrorCode = 31002
	// push server error codes
	ErrorCode_UserNotOnline ErrorCode = 40001
	// group error codes
	ErrorCode_GroupNotExist    ErrorCode = 50001
	ErrorCode_GroupExist       ErrorCode = 50002
	ErrorCode_NotGroupMember   ErrorCode = 50003
	ErrorCode_NotGroupOwner    ErrorCode = 50004
	ErrorCode_GroupLimitExceed ErrorCode = 50005
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:     "OK",
		10001: "InternalError",
		10002: "InvalidParams",
		10003: "DBError",
		10004: "CacheError",
		20001: "UserNotExist",
		20002: "UserExist",
		20003: "InvalidUsernameOrPassword",
		30001: "RelationNotExist",
		30002: "RelationExist",
		30003: "InvalidUpdateRelationAction",
		31001: "FriendRequestNotExist",
		31002: "FriendRequestStatusError",
		40001: "UserNotOnline",
		50001: "GroupNotExist",
		50002: "GroupExist",
		50003: "NotGroupMember",
		50004: "NotGroupOwner",
		50005: "GroupLimitExceed",
	}
	ErrorCode_value = map[string]int32{
		"OK":                          0,
		"InternalError":               10001,
		"InvalidParams":               10002,
		"DBError":                     10003,
		"CacheError":                  10004,
		"UserNotExist":                20001,
		"UserExist":                   20002,
		"InvalidUsernameOrPassword":   20003,
		"RelationNotExist":            30001,
		"RelationExist":               30002,
		"InvalidUpdateRelationAction": 30003,
		"FriendRequestNotExist":       31001,
		"FriendRequestStatusError":    31002,
		"UserNotOnline":               40001,
		"GroupNotExist":               50001,
		"GroupExist":                  50002,
		"NotGroupMember":              50003,
		"NotGroupOwner":               50004,
		"GroupLimitExceed":            50005,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_errors_errors_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_errors_errors_proto_rawDescGZIP(), []int{0}
}

// Error use as define error code and message.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=api.errors.ErrorCode" json:"error_code,omitempty"`
	Reason    string    `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message   string    `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errors_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errors_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errors_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetErrorCode() ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return ErrorCode_OK
}

func (x *Error) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_errors_errors_proto protoreflect.FileDescriptor

var file_errors_errors_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x22, 0x6f, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0xac, 0x03, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x91, 0x4e, 0x12, 0x12, 0x0a, 0x0d,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x10, 0x92, 0x4e,
	0x12, 0x0c, 0x0a, 0x07, 0x44, 0x42, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x93, 0x4e, 0x12, 0x0f,
	0x0a, 0x0a, 0x43, 0x61, 0x63, 0x68, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x94, 0x4e, 0x12,
	0x12, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10,
	0xa1, 0x9c, 0x01, 0x12, 0x0f, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74,
	0x10, 0xa2, 0x9c, 0x01, 0x12, 0x1f, 0x0a, 0x19, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x4f, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x10, 0xa3, 0x9c, 0x01, 0x12, 0x16, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xb1, 0xea, 0x01, 0x12, 0x13, 0x0a,
	0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xb2,
	0xea, 0x01, 0x12, 0x21, 0x0a, 0x1b, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0xb3, 0xea, 0x01, 0x12, 0x1b, 0x0a, 0x15, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0x99,
	0xf2, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x9a,
	0xf2, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x10, 0xc1, 0xb8, 0x02, 0x12, 0x13, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x10, 0x0a, 0x0a,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x78, 0x69, 0x73, 0x74, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x14,
	0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x10, 0xd3, 0x86, 0x03, 0x12, 0x13, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x10, 0xd4, 0x86, 0x03, 0x12, 0x16, 0x0a, 0x10, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x45, 0x78, 0x63, 0x65, 0x65, 0x64, 0x10, 0xd5, 0x86,
	0x03, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_errors_proto_rawDescOnce sync.Once
	file_errors_errors_proto_rawDescData = file_errors_errors_proto_rawDesc
)

func file_errors_errors_proto_rawDescGZIP() []byte {
	file_errors_errors_proto_rawDescOnce.Do(func() {
		file_errors_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_errors_proto_rawDescData)
	})
	return file_errors_errors_proto_rawDescData
}

var file_errors_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_errors_errors_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: api.errors.ErrorCode
	(*Error)(nil),  // 1: api.errors.Error
}
var file_errors_errors_proto_depIdxs = []int32{
	0, // 0: api.errors.Error.error_code:type_name -> api.errors.ErrorCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_errors_errors_proto_init() }
func file_errors_errors_proto_init() {
	if File_errors_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errors_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_errors_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_errors_proto_goTypes,
		DependencyIndexes: file_errors_errors_proto_depIdxs,
		EnumInfos:         file_errors_errors_proto_enumTypes,
		MessageInfos:      file_errors_errors_proto_msgTypes,
	}.Build()
	File_errors_errors_proto = out.File
	file_errors_errors_proto_rawDesc = nil
	file_errors_errors_proto_goTypes = nil
	file_errors_errors_proto_depIdxs = nil
}