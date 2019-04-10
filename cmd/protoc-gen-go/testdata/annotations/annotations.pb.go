// Code generated by protoc-gen-go. DO NOT EDIT.
// source: annotations/annotations.proto

package annotations

import (
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoregistry "github.com/golang/protobuf/v2/reflect/protoregistry"
	protoiface "github.com/golang/protobuf/v2/runtime/protoiface"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type AnnotationsTestEnum int32

const (
	AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE AnnotationsTestEnum = 0
)

// Deprecated: Use AnnotationsTestEnum.Type.Values instead.
var AnnotationsTestEnum_name = map[int32]string{
	0: "ANNOTATIONS_TEST_ENUM_VALUE",
}

// Deprecated: Use AnnotationsTestEnum.Type.Values instead.
var AnnotationsTestEnum_value = map[string]int32{
	"ANNOTATIONS_TEST_ENUM_VALUE": 0,
}

func (x AnnotationsTestEnum) Enum() *AnnotationsTestEnum {
	p := new(AnnotationsTestEnum)
	*p = x
	return p
}

func (x AnnotationsTestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

func (AnnotationsTestEnum) Type() protoreflect.EnumType {
	return xxx_File_annotations_annotations_proto_enumTypes[0]
}

func (x AnnotationsTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AnnotationsTestEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Type(), b)
	if err != nil {
		return err
	}
	*x = AnnotationsTestEnum(num)
	return nil
}

// Deprecated: Use AnnotationsTestEnum.Type instead.
func (AnnotationsTestEnum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

type AnnotationsTestMessage struct {
	AnnotationsTestField *string  `protobuf:"bytes,1,opt,name=AnnotationsTestField" json:"AnnotationsTestField,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (x *AnnotationsTestMessage) Reset() {
	*x = AnnotationsTestMessage{}
}

func (x *AnnotationsTestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationsTestMessage) ProtoMessage() {}

func (x *AnnotationsTestMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_annotations_annotations_proto_messageTypes[0].MessageOf(x)
}

func (m *AnnotationsTestMessage) XXX_Methods() *protoiface.Methods {
	return xxx_File_annotations_annotations_proto_messageTypes[0].Methods()
}

// Deprecated: Use AnnotationsTestMessage.ProtoReflect.Type instead.
func (*AnnotationsTestMessage) Descriptor() ([]byte, []int) {
	return xxx_File_annotations_annotations_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotationsTestMessage) GetAnnotationsTestField() string {
	if x != nil && x.AnnotationsTestField != nil {
		return *x.AnnotationsTestField
	}
	return ""
}

var File_annotations_annotations_proto protoreflect.FileDescriptor

var xxx_File_annotations_annotations_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4c, 0x0a, 0x16, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x36, 0x0a, 0x13, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d,
	0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10,
	0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
}

var (
	xxx_File_annotations_annotations_proto_rawDesc_once sync.Once
	xxx_File_annotations_annotations_proto_rawDesc_data = xxx_File_annotations_annotations_proto_rawDesc
)

func xxx_File_annotations_annotations_proto_rawDescGZIP() []byte {
	xxx_File_annotations_annotations_proto_rawDesc_once.Do(func() {
		xxx_File_annotations_annotations_proto_rawDesc_data = protoimpl.X.CompressGZIP(xxx_File_annotations_annotations_proto_rawDesc_data)
	})
	return xxx_File_annotations_annotations_proto_rawDesc_data
}

var xxx_File_annotations_annotations_proto_enumTypes = make([]protoreflect.EnumType, 1)
var xxx_File_annotations_annotations_proto_messageTypes = make([]protoimpl.MessageType, 1)
var xxx_File_annotations_annotations_proto_goTypes = []interface{}{
	(AnnotationsTestEnum)(0),       // 0: goproto.protoc.annotations.AnnotationsTestEnum
	(*AnnotationsTestMessage)(nil), // 1: goproto.protoc.annotations.AnnotationsTestMessage
}
var xxx_File_annotations_annotations_proto_depIdxs = []int32{}

func init() { xxx_File_annotations_annotations_proto_init() }
func xxx_File_annotations_annotations_proto_init() {
	if File_annotations_annotations_proto != nil {
		return
	}
	File_annotations_annotations_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_annotations_annotations_proto_rawDesc,
		GoTypes:            xxx_File_annotations_annotations_proto_goTypes,
		DependencyIndexes:  xxx_File_annotations_annotations_proto_depIdxs,
		EnumOutputTypes:    xxx_File_annotations_annotations_proto_enumTypes,
		MessageOutputTypes: xxx_File_annotations_annotations_proto_messageTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	xxx_File_annotations_annotations_proto_rawDesc = nil
	xxx_File_annotations_annotations_proto_goTypes = nil
	xxx_File_annotations_annotations_proto_depIdxs = nil
}
