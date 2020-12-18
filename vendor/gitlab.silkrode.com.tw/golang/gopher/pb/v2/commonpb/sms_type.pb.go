// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/v2/commonpb/sms_type.proto

package commonpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SMSType int32

const (
	SMSTypeUnknown SMSType = 0
	SMSTypeCloud   SMSType = 1
	SMSTypeLocal   SMSType = 2
)

var SMSType_name = map[int32]string{
	0: "SMS_TYPE_UNKNOWN",
	1: "SMS_TYPE_CLOUD",
	2: "SMS_TYPE_LOCAL",
}

var SMSType_value = map[string]int32{
	"SMS_TYPE_UNKNOWN": 0,
	"SMS_TYPE_CLOUD":   1,
	"SMS_TYPE_LOCAL":   2,
}

func (x SMSType) String() string {
	return proto.EnumName(SMSType_name, int32(x))
}

func (SMSType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e6e27d338b2f2dfa, []int{0}
}

func init() {
	proto.RegisterEnum("commonpbv2.SMSType", SMSType_name, SMSType_value)
	golang_proto.RegisterEnum("commonpbv2.SMSType", SMSType_name, SMSType_value)
}

func init() { proto.RegisterFile("pb/v2/commonpb/sms_type.proto", fileDescriptor_e6e27d338b2f2dfa) }
func init() {
	golang_proto.RegisterFile("pb/v2/commonpb/sms_type.proto", fileDescriptor_e6e27d338b2f2dfa)
}

var fileDescriptor_e6e27d338b2f2dfa = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x48, 0xd2, 0x2f,
	0x33, 0xd2, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x2b, 0x48, 0xd2, 0x2f, 0xce, 0x2d, 0x8e, 0x2f,
	0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0x49, 0x94, 0x19, 0x49,
	0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xa7, 0xe7,
	0xeb, 0x83, 0x95, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0xaa, 0xd5, 0xca,
	0xc8, 0xc5, 0x1e, 0xec, 0x1b, 0x1c, 0x52, 0x59, 0x90, 0x2a, 0xa4, 0xc1, 0x25, 0x10, 0xec, 0x1b,
	0x1c, 0x1f, 0x12, 0x19, 0xe0, 0x1a, 0x1f, 0xea, 0xe7, 0xed, 0xe7, 0x1f, 0xee, 0x27, 0xc0, 0x20,
	0x25, 0xd4, 0x35, 0x57, 0x81, 0x0f, 0xaa, 0x24, 0x34, 0x2f, 0x3b, 0x2f, 0xbf, 0x3c, 0x4f, 0x48,
	0x85, 0x8b, 0x0f, 0xae, 0xd2, 0xd9, 0xc7, 0x3f, 0xd4, 0x45, 0x80, 0x51, 0x4a, 0xa0, 0x6b, 0xae,
	0x02, 0x0f, 0x54, 0x9d, 0x73, 0x4e, 0x7e, 0x69, 0x0a, 0x8a, 0x2a, 0x1f, 0x7f, 0x67, 0x47, 0x1f,
	0x01, 0x26, 0x14, 0x55, 0x3e, 0xf9, 0xc9, 0x89, 0x39, 0x52, 0x2c, 0x1d, 0x8b, 0xe5, 0x18, 0x9c,
	0x2c, 0x0f, 0x3c, 0x96, 0x63, 0x8c, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x49, 0x4c, 0xd2, 0x2b, 0xce,
	0xcc, 0xc9, 0x2e, 0xca, 0x4f, 0x49, 0x05, 0xf9, 0x42, 0xaf, 0xa4, 0x5c, 0x3f, 0x3d, 0x3f, 0x27,
	0x31, 0x2f, 0x5d, 0x3f, 0x3d, 0xbf, 0x20, 0x23, 0xb5, 0x48, 0x1f, 0x35, 0x30, 0x92, 0xd8, 0xc0,
	0x3e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xf2, 0x16, 0x36, 0x25, 0x01, 0x00, 0x00,
}
