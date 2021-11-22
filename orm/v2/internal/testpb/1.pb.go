// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: 1.proto

package testpb

import (
	_ "github.com/regen-network/regen-ledger/orm/v2/types/ormpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Enumeration int32

const (
	Enumeration_One Enumeration = 0
	Enumeration_Two Enumeration = 1
)

// Enum value maps for Enumeration.
var (
	Enumeration_name = map[int32]string{
		0: "One",
		1: "Two",
	}
	Enumeration_value = map[string]int32{
		"One": 0,
		"Two": 1,
	}
)

func (x Enumeration) Enum() *Enumeration {
	p := new(Enumeration)
	*p = x
	return p
}

func (x Enumeration) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enumeration) Descriptor() protoreflect.EnumDescriptor {
	return file__1_proto_enumTypes[0].Descriptor()
}

func (Enumeration) Type() protoreflect.EnumType {
	return &file__1_proto_enumTypes[0]
}

func (x Enumeration) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enumeration.Descriptor instead.
func (Enumeration) EnumDescriptor() ([]byte, []int) {
	return file__1_proto_rawDescGZIP(), []int{0}
}

type A struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ENUM     Enumeration   `protobuf:"varint,1,opt,name=ENUM,proto3,enum=Enumeration" json:"ENUM,omitempty"`
	BOOL     bool          `protobuf:"varint,2,opt,name=BOOL,proto3" json:"BOOL,omitempty"`
	INT32    int32         `protobuf:"varint,3,opt,name=INT32,proto3" json:"INT32,omitempty"`
	SINT32   int32         `protobuf:"zigzag32,4,opt,name=SINT32,proto3" json:"SINT32,omitempty"`
	UINT32   uint32        `protobuf:"varint,5,opt,name=UINT32,proto3" json:"UINT32,omitempty"`
	INT64    int64         `protobuf:"varint,6,opt,name=INT64,proto3" json:"INT64,omitempty"`
	SINT64   int64         `protobuf:"zigzag64,7,opt,name=SINT64,proto3" json:"SINT64,omitempty"`
	UINT64   uint64        `protobuf:"varint,8,opt,name=UINT64,proto3" json:"UINT64,omitempty"`
	SFIXED32 int32         `protobuf:"fixed32,9,opt,name=SFIXED32,proto3" json:"SFIXED32,omitempty"`
	FIXED32  uint32        `protobuf:"fixed32,10,opt,name=FIXED32,proto3" json:"FIXED32,omitempty"`
	FLOAT    float32       `protobuf:"fixed32,11,opt,name=FLOAT,proto3" json:"FLOAT,omitempty"`
	SFIXED64 int64         `protobuf:"fixed64,12,opt,name=SFIXED64,proto3" json:"SFIXED64,omitempty"`
	FIXED64  uint64        `protobuf:"fixed64,13,opt,name=FIXED64,proto3" json:"FIXED64,omitempty"`
	DOUBLE   float64       `protobuf:"fixed64,14,opt,name=DOUBLE,proto3" json:"DOUBLE,omitempty"`
	STRING   string        `protobuf:"bytes,15,opt,name=STRING,proto3" json:"STRING,omitempty"`
	BYTES    []byte        `protobuf:"bytes,16,opt,name=BYTES,proto3" json:"BYTES,omitempty"`
	MESSAGE  *B            `protobuf:"bytes,17,opt,name=MESSAGE,proto3" json:"MESSAGE,omitempty"`
	MAP      map[string]*B `protobuf:"bytes,18,rep,name=MAP,proto3" json:"MAP,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LIST     []*B          `protobuf:"bytes,19,rep,name=LIST,proto3" json:"LIST,omitempty"`
	// Types that are assignable to ONEOF:
	//	*A_ONEOF_B
	//	*A_ONEOF_STRING
	ONEOF     isA_ONEOF              `protobuf_oneof:"ONEOF"`
	LIST_ENUM []Enumeration          `protobuf:"varint,22,rep,packed,name=LIST_ENUM,json=LISTENUM,proto3,enum=Enumeration" json:"LIST_ENUM,omitempty"`
	TIMESTAMP *timestamppb.Timestamp `protobuf:"bytes,23,opt,name=TIMESTAMP,proto3" json:"TIMESTAMP,omitempty"`
	DURATION  *durationpb.Duration   `protobuf:"bytes,24,opt,name=DURATION,proto3" json:"DURATION,omitempty"`
}

func (x *A) Reset() {
	*x = A{}
	if protoimpl.UnsafeEnabled {
		mi := &file__1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *A) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*A) ProtoMessage() {}

func (x *A) ProtoReflect() protoreflect.Message {
	mi := &file__1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use A.ProtoReflect.Descriptor instead.
func (*A) Descriptor() ([]byte, []int) {
	return file__1_proto_rawDescGZIP(), []int{0}
}

func (x *A) GetENUM() Enumeration {
	if x != nil {
		return x.ENUM
	}
	return Enumeration_One
}

func (x *A) GetBOOL() bool {
	if x != nil {
		return x.BOOL
	}
	return false
}

func (x *A) GetINT32() int32 {
	if x != nil {
		return x.INT32
	}
	return 0
}

func (x *A) GetSINT32() int32 {
	if x != nil {
		return x.SINT32
	}
	return 0
}

func (x *A) GetUINT32() uint32 {
	if x != nil {
		return x.UINT32
	}
	return 0
}

func (x *A) GetINT64() int64 {
	if x != nil {
		return x.INT64
	}
	return 0
}

func (x *A) GetSINT64() int64 {
	if x != nil {
		return x.SINT64
	}
	return 0
}

func (x *A) GetUINT64() uint64 {
	if x != nil {
		return x.UINT64
	}
	return 0
}

func (x *A) GetSFIXED32() int32 {
	if x != nil {
		return x.SFIXED32
	}
	return 0
}

func (x *A) GetFIXED32() uint32 {
	if x != nil {
		return x.FIXED32
	}
	return 0
}

func (x *A) GetFLOAT() float32 {
	if x != nil {
		return x.FLOAT
	}
	return 0
}

func (x *A) GetSFIXED64() int64 {
	if x != nil {
		return x.SFIXED64
	}
	return 0
}

func (x *A) GetFIXED64() uint64 {
	if x != nil {
		return x.FIXED64
	}
	return 0
}

func (x *A) GetDOUBLE() float64 {
	if x != nil {
		return x.DOUBLE
	}
	return 0
}

func (x *A) GetSTRING() string {
	if x != nil {
		return x.STRING
	}
	return ""
}

func (x *A) GetBYTES() []byte {
	if x != nil {
		return x.BYTES
	}
	return nil
}

func (x *A) GetMESSAGE() *B {
	if x != nil {
		return x.MESSAGE
	}
	return nil
}

func (x *A) GetMAP() map[string]*B {
	if x != nil {
		return x.MAP
	}
	return nil
}

func (x *A) GetLIST() []*B {
	if x != nil {
		return x.LIST
	}
	return nil
}

func (m *A) GetONEOF() isA_ONEOF {
	if m != nil {
		return m.ONEOF
	}
	return nil
}

func (x *A) GetONEOF_B() *B {
	if x, ok := x.GetONEOF().(*A_ONEOF_B); ok {
		return x.ONEOF_B
	}
	return nil
}

func (x *A) GetONEOF_STRING() string {
	if x, ok := x.GetONEOF().(*A_ONEOF_STRING); ok {
		return x.ONEOF_STRING
	}
	return ""
}

func (x *A) GetLIST_ENUM() []Enumeration {
	if x != nil {
		return x.LIST_ENUM
	}
	return nil
}

func (x *A) GetTIMESTAMP() *timestamppb.Timestamp {
	if x != nil {
		return x.TIMESTAMP
	}
	return nil
}

func (x *A) GetDURATION() *durationpb.Duration {
	if x != nil {
		return x.DURATION
	}
	return nil
}

type isA_ONEOF interface {
	isA_ONEOF()
}

type A_ONEOF_B struct {
	ONEOF_B *B `protobuf:"bytes,20,opt,name=ONEOF_B,json=ONEOFB,proto3,oneof"`
}

type A_ONEOF_STRING struct {
	ONEOF_STRING string `protobuf:"bytes,21,opt,name=ONEOF_STRING,json=ONEOFSTRING,proto3,oneof"`
}

func (*A_ONEOF_B) isA_ONEOF() {}

func (*A_ONEOF_STRING) isA_ONEOF() {}

type B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X string `protobuf:"bytes,1,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *B) Reset() {
	*x = B{}
	if protoimpl.UnsafeEnabled {
		mi := &file__1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*B) ProtoMessage() {}

func (x *B) ProtoReflect() protoreflect.Message {
	mi := &file__1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use B.ProtoReflect.Descriptor instead.
func (*B) Descriptor() ([]byte, []int) {
	return file__1_proto_rawDescGZIP(), []int{1}
}

func (x *B) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

type C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	X  string `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
}

func (x *C) Reset() {
	*x = C{}
	if protoimpl.UnsafeEnabled {
		mi := &file__1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C) ProtoMessage() {}

func (x *C) ProtoReflect() protoreflect.Message {
	mi := &file__1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C.ProtoReflect.Descriptor instead.
func (*C) Descriptor() ([]byte, []int) {
	return file__1_proto_rawDescGZIP(), []int{2}
}

func (x *C) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *C) GetX() string {
	if x != nil {
		return x.X
	}
	return ""
}

var File__1_proto protoreflect.FileDescriptor

var file__1_proto_rawDesc = []byte{
	0x0a, 0x07, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6f,
	0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x06, 0x0a, 0x01, 0x41, 0x12, 0x20,
	0x0a, 0x04, 0x45, 0x4e, 0x55, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x45, 0x4e, 0x55, 0x4d,
	0x12, 0x12, 0x0a, 0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x42, 0x4f, 0x4f, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x49,
	0x4e, 0x54, 0x33, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x53, 0x49, 0x4e, 0x54,
	0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x4e,
	0x54, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x49, 0x4e, 0x54, 0x36, 0x34,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x18, 0x07, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x06, 0x53, 0x49, 0x4e, 0x54, 0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54,
	0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36, 0x34,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0f, 0x52, 0x08, 0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07,
	0x46, 0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x07, 0x52, 0x07, 0x46,
	0x49, 0x58, 0x45, 0x44, 0x33, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x12, 0x1a, 0x0a, 0x08,
	0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x10, 0x52, 0x08,
	0x53, 0x46, 0x49, 0x58, 0x45, 0x44, 0x36, 0x34, 0x12, 0x18, 0x0a, 0x07, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x36, 0x34, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x46, 0x49, 0x58, 0x45, 0x44,
	0x36, 0x34, 0x12, 0x16, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x12, 0x1c, 0x0a, 0x07, 0x4d, 0x45, 0x53, 0x53,
	0x41, 0x47, 0x45, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x07, 0x4d,
	0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x12, 0x1d, 0x0a, 0x03, 0x4d, 0x41, 0x50, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x2e, 0x4d, 0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x03, 0x4d, 0x41, 0x50, 0x12, 0x16, 0x0a, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x18, 0x13, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x04, 0x4c, 0x49, 0x53, 0x54, 0x12, 0x1d, 0x0a,
	0x07, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x42, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02,
	0x2e, 0x42, 0x48, 0x00, 0x52, 0x06, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x42, 0x12, 0x23, 0x0a, 0x0c,
	0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x12, 0x29, 0x0a, 0x09, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x18, 0x16,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x55, 0x4d, 0x12, 0x38, 0x0a, 0x09,
	0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x49, 0x4d,
	0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x12, 0x35, 0x0a, 0x08, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x1a, 0x3a, 0x0a,
	0x08, 0x4d, 0x41, 0x50, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x02, 0x2e, 0x42, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x58, 0x82, 0x90, 0xbc, 0xfd, 0x02,
	0x52, 0x0a, 0x16, 0x0a, 0x14, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x2c, 0x55, 0x49, 0x4e, 0x54,
	0x36, 0x34, 0x2c, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x49, 0x4e,
	0x54, 0x36, 0x34, 0x2c, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x2c, 0x55, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x02, 0x12,
	0x10, 0x0a, 0x0c, 0x42, 0x59, 0x54, 0x45, 0x53, 0x2c, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x18, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x4f, 0x4e, 0x45, 0x4f, 0x46, 0x22, 0x1b, 0x0a, 0x01,
	0x42, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x3a,
	0x08, 0x8a, 0x90, 0xbc, 0xfd, 0x02, 0x02, 0x08, 0x02, 0x22, 0x33, 0x0a, 0x01, 0x43, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x78, 0x3a, 0x10, 0x82, 0x90,
	0xbc, 0xfd, 0x02, 0x0a, 0x0a, 0x06, 0x0a, 0x02, 0x69, 0x64, 0x10, 0x01, 0x18, 0x03, 0x2a, 0x1f,
	0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x07, 0x0a,
	0x03, 0x4f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x01, 0x42,
	0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file__1_proto_rawDescOnce sync.Once
	file__1_proto_rawDescData = file__1_proto_rawDesc
)

func file__1_proto_rawDescGZIP() []byte {
	file__1_proto_rawDescOnce.Do(func() {
		file__1_proto_rawDescData = protoimpl.X.CompressGZIP(file__1_proto_rawDescData)
	})
	return file__1_proto_rawDescData
}

var file__1_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file__1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file__1_proto_goTypes = []interface{}{
	(Enumeration)(0),              // 0: Enumeration
	(*A)(nil),                     // 1: A
	(*B)(nil),                     // 2: B
	(*C)(nil),                     // 3: C
	nil,                           // 4: A.MAPEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file__1_proto_depIdxs = []int32{
	0, // 0: A.ENUM:type_name -> Enumeration
	2, // 1: A.MESSAGE:type_name -> B
	4, // 2: A.MAP:type_name -> A.MAPEntry
	2, // 3: A.LIST:type_name -> B
	2, // 4: A.ONEOF_B:type_name -> B
	0, // 5: A.LIST_ENUM:type_name -> Enumeration
	5, // 6: A.TIMESTAMP:type_name -> google.protobuf.Timestamp
	6, // 7: A.DURATION:type_name -> google.protobuf.Duration
	2, // 8: A.MAPEntry.value:type_name -> B
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file__1_proto_init() }
func file__1_proto_init() {
	if File__1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file__1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*A); i {
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
		file__1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*B); i {
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
		file__1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C); i {
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
	file__1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*A_ONEOF_B)(nil),
		(*A_ONEOF_STRING)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file__1_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file__1_proto_goTypes,
		DependencyIndexes: file__1_proto_depIdxs,
		EnumInfos:         file__1_proto_enumTypes,
		MessageInfos:      file__1_proto_msgTypes,
	}.Build()
	File__1_proto = out.File
	file__1_proto_rawDesc = nil
	file__1_proto_goTypes = nil
	file__1_proto_depIdxs = nil
}
