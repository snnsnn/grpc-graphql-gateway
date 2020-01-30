// Code generated by protoc-gen-go. DO NOT EDIT.
// source: graphql.proto

package graphql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GraphqlQuery struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Field                string   `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Autoresolve          bool     `protobuf:"varint,3,opt,name=autoresolve,proto3" json:"autoresolve,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlQuery) Reset()         { *m = GraphqlQuery{} }
func (m *GraphqlQuery) String() string { return proto.CompactTextString(m) }
func (*GraphqlQuery) ProtoMessage()    {}
func (*GraphqlQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{0}
}

func (m *GraphqlQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlQuery.Unmarshal(m, b)
}
func (m *GraphqlQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlQuery.Marshal(b, m, deterministic)
}
func (m *GraphqlQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlQuery.Merge(m, src)
}
func (m *GraphqlQuery) XXX_Size() int {
	return xxx_messageInfo_GraphqlQuery.Size(m)
}
func (m *GraphqlQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlQuery.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlQuery proto.InternalMessageInfo

func (m *GraphqlQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphqlQuery) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *GraphqlQuery) GetAutoresolve() bool {
	if m != nil {
		return m.Autoresolve
	}
	return false
}

type GraphqlMutation struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Field                string   `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`
	Autoresolve          bool     `protobuf:"varint,3,opt,name=autoresolve,proto3" json:"autoresolve,omitempty"`
	ExtractField         bool     `protobuf:"varint,4,opt,name=extract_field,json=extractField,proto3" json:"extract_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlMutation) Reset()         { *m = GraphqlMutation{} }
func (m *GraphqlMutation) String() string { return proto.CompactTextString(m) }
func (*GraphqlMutation) ProtoMessage()    {}
func (*GraphqlMutation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{1}
}

func (m *GraphqlMutation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlMutation.Unmarshal(m, b)
}
func (m *GraphqlMutation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlMutation.Marshal(b, m, deterministic)
}
func (m *GraphqlMutation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlMutation.Merge(m, src)
}
func (m *GraphqlMutation) XXX_Size() int {
	return xxx_messageInfo_GraphqlMutation.Size(m)
}
func (m *GraphqlMutation) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlMutation.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlMutation proto.InternalMessageInfo

func (m *GraphqlMutation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphqlMutation) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *GraphqlMutation) GetAutoresolve() bool {
	if m != nil {
		return m.Autoresolve
	}
	return false
}

func (m *GraphqlMutation) GetExtractField() bool {
	if m != nil {
		return m.ExtractField
	}
	return false
}

type GraphqlField struct {
	Optional             bool     `protobuf:"varint,1,opt,name=optional,proto3" json:"optional,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GraphqlField) Reset()         { *m = GraphqlField{} }
func (m *GraphqlField) String() string { return proto.CompactTextString(m) }
func (*GraphqlField) ProtoMessage()    {}
func (*GraphqlField) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce0fc368bdc1a51, []int{2}
}

func (m *GraphqlField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GraphqlField.Unmarshal(m, b)
}
func (m *GraphqlField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GraphqlField.Marshal(b, m, deterministic)
}
func (m *GraphqlField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GraphqlField.Merge(m, src)
}
func (m *GraphqlField) XXX_Size() int {
	return xxx_messageInfo_GraphqlField.Size(m)
}
func (m *GraphqlField) XXX_DiscardUnknown() {
	xxx_messageInfo_GraphqlField.DiscardUnknown(m)
}

var xxx_messageInfo_GraphqlField proto.InternalMessageInfo

func (m *GraphqlField) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

func (m *GraphqlField) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

var E_Query = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*GraphqlQuery)(nil),
	Field:         50001,
	Name:          "graphql.query",
	Tag:           "bytes,50001,opt,name=query",
	Filename:      "graphql.proto",
}

var E_Mutation = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*GraphqlMutation)(nil),
	Field:         50002,
	Name:          "graphql.mutation",
	Tag:           "bytes,50002,opt,name=mutation",
	Filename:      "graphql.proto",
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*GraphqlField)(nil),
	Field:         50003,
	Name:          "graphql.field",
	Tag:           "bytes,50003,opt,name=field",
	Filename:      "graphql.proto",
}

func init() {
	proto.RegisterType((*GraphqlQuery)(nil), "graphql.GraphqlQuery")
	proto.RegisterType((*GraphqlMutation)(nil), "graphql.GraphqlMutation")
	proto.RegisterType((*GraphqlField)(nil), "graphql.GraphqlField")
	proto.RegisterExtension(E_Query)
	proto.RegisterExtension(E_Mutation)
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("graphql.proto", fileDescriptor_3ce0fc368bdc1a51) }

var fileDescriptor_3ce0fc368bdc1a51 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x49, 0x6d, 0x35, 0xbe, 0xb6, 0x08, 0x87, 0x42, 0x28, 0x28, 0xa1, 0x2e, 0x9d, 0x52,
	0xd0, 0x2d, 0x83, 0xa3, 0x2e, 0x06, 0x31, 0x83, 0x83, 0x8b, 0x5c, 0x9b, 0x6b, 0x1a, 0xb8, 0xf6,
	0xa5, 0x97, 0x8b, 0xe8, 0xea, 0xe0, 0xf7, 0x53, 0xbf, 0x90, 0xe4, 0xdd, 0x5d, 0x2c, 0x8a, 0xb8,
	0xb8, 0xe5, 0xfd, 0x79, 0xef, 0xc7, 0xff, 0x7e, 0x81, 0x61, 0xae, 0x78, 0xb9, 0xdc, 0xc8, 0xa8,
	0x54, 0xa8, 0x91, 0xed, 0xd9, 0x71, 0x14, 0xe6, 0x88, 0xb9, 0x14, 0x53, 0x8a, 0x67, 0xf5, 0x62,
	0x9a, 0x89, 0x6a, 0xae, 0x8a, 0x52, 0xa3, 0x32, 0xab, 0xe3, 0x7b, 0x18, 0x5c, 0x99, 0xe5, 0xdb,
	0x5a, 0xa8, 0x67, 0xc6, 0xa0, 0xbb, 0xe6, 0x2b, 0x11, 0x78, 0xa1, 0x37, 0xd9, 0x4f, 0xe9, 0x9b,
	0x1d, 0x42, 0x6f, 0x51, 0x08, 0x99, 0x05, 0x1d, 0x0a, 0xcd, 0xc0, 0x42, 0xe8, 0xf3, 0x5a, 0xa3,
	0x12, 0x15, 0xca, 0x47, 0x11, 0xec, 0x84, 0xde, 0xc4, 0x4f, 0xb7, 0xa3, 0xf1, 0x8b, 0x07, 0x07,
	0x16, 0x9e, 0xd4, 0x9a, 0xeb, 0x02, 0xd7, 0xff, 0xc9, 0x67, 0xa7, 0x30, 0x14, 0x4f, 0x5a, 0xf1,
	0xb9, 0x7e, 0x30, 0xf7, 0x5d, 0xda, 0x19, 0xd8, 0xf0, 0xb2, 0xc9, 0xc6, 0x17, 0xed, 0x03, 0x69,
	0x66, 0x23, 0xf0, 0xb1, 0x6c, 0xaa, 0x70, 0x49, 0x25, 0xfc, 0xb4, 0x9d, 0xdb, 0x72, 0x9d, 0xaf,
	0x72, 0x71, 0x02, 0xbd, 0x0d, 0x99, 0x39, 0x89, 0x8c, 0xcc, 0xc8, 0xc9, 0x8c, 0x12, 0xa1, 0x97,
	0x98, 0xdd, 0xd0, 0x75, 0x15, 0xbc, 0xbd, 0x36, 0x15, 0xfb, 0x67, 0x47, 0x91, 0xfb, 0x19, 0xdb,
	0x62, 0x53, 0x43, 0x89, 0xef, 0xc0, 0x5f, 0x39, 0x17, 0x7f, 0x11, 0xdf, 0x2d, 0x31, 0xf8, 0x4e,
	0x74, 0x36, 0xd3, 0x96, 0x15, 0x5f, 0x5b, 0x87, 0xec, 0xf8, 0x07, 0x94, 0xde, 0xed, 0x98, 0x1f,
	0xbf, 0xb5, 0xa4, 0x2d, 0xeb, 0x7e, 0xb6, 0x4b, 0xc7, 0xe7, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd4, 0x09, 0xef, 0xde, 0x58, 0x02, 0x00, 0x00,
}
