// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: user/session/v1/session.proto

package v1

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

// define session type and status
type SessionType int32

const (
	SessionType_SingleChat SessionType = 0
	SessionType_GroupChat  SessionType = 1
	// broadcast actually not a standard chat type, but we still use it here
	SessionType_Broadcast SessionType = 2
)

// Enum value maps for SessionType.
var (
	SessionType_name = map[int32]string{
		0: "SingleChat",
		1: "GroupChat",
		2: "Broadcast",
	}
	SessionType_value = map[string]int32{
		"SingleChat": 0,
		"GroupChat":  1,
		"Broadcast":  2,
	}
)

func (x SessionType) Enum() *SessionType {
	p := new(SessionType)
	*p = x
	return p
}

func (x SessionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SessionType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_session_v1_session_proto_enumTypes[0].Descriptor()
}

func (SessionType) Type() protoreflect.EnumType {
	return &file_user_session_v1_session_proto_enumTypes[0]
}

func (x SessionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SessionType.Descriptor instead.
func (SessionType) EnumDescriptor() ([]byte, []int) {
	return file_user_session_v1_session_proto_rawDescGZIP(), []int{0}
}

var File_user_session_v1_session_proto protoreflect.FileDescriptor

var file_user_session_v1_session_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2a, 0x3b, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x43, 0x68, 0x61,
	0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x68, 0x61, 0x74,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x10,
	0x02, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_user_session_v1_session_proto_rawDescOnce sync.Once
	file_user_session_v1_session_proto_rawDescData = file_user_session_v1_session_proto_rawDesc
)

func file_user_session_v1_session_proto_rawDescGZIP() []byte {
	file_user_session_v1_session_proto_rawDescOnce.Do(func() {
		file_user_session_v1_session_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_session_v1_session_proto_rawDescData)
	})
	return file_user_session_v1_session_proto_rawDescData
}

var file_user_session_v1_session_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_session_v1_session_proto_goTypes = []interface{}{
	(SessionType)(0), // 0: api.user.session.v1.SessionType
}
var file_user_session_v1_session_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_session_v1_session_proto_init() }
func file_user_session_v1_session_proto_init() {
	if File_user_session_v1_session_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_session_v1_session_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_session_v1_session_proto_goTypes,
		DependencyIndexes: file_user_session_v1_session_proto_depIdxs,
		EnumInfos:         file_user_session_v1_session_proto_enumTypes,
	}.Build()
	File_user_session_v1_session_proto = out.File
	file_user_session_v1_session_proto_rawDesc = nil
	file_user_session_v1_session_proto_goTypes = nil
	file_user_session_v1_session_proto_depIdxs = nil
}
