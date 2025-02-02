// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/bridge/v1/events.proto

package axelarbridge

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SendBridgeEvent is a Tendermint type for the Msg/SendBridgeEvent
// Bridge is going to listen for this events and enhance the message with
// chain ID and origin Tx ID.
type SendBridgeEvent struct {
	Sender  string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Handler string `protobuf:"bytes,2,opt,name=handler,proto3" json:"handler,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *SendBridgeEvent) Reset()         { *m = SendBridgeEvent{} }
func (m *SendBridgeEvent) String() string { return proto.CompactTextString(m) }
func (*SendBridgeEvent) ProtoMessage()    {}
func (*SendBridgeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_85fbba2e40a9e6c7, []int{0}
}
func (m *SendBridgeEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendBridgeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendBridgeEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendBridgeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendBridgeEvent.Merge(m, src)
}
func (m *SendBridgeEvent) XXX_Size() int {
	return m.Size()
}
func (m *SendBridgeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SendBridgeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SendBridgeEvent proto.InternalMessageInfo

func (m *SendBridgeEvent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *SendBridgeEvent) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *SendBridgeEvent) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*SendBridgeEvent)(nil), "axelar.bridge.v1.SendBridgeEvent")
}

func init() { proto.RegisterFile("axelar/bridge/v1/events.proto", fileDescriptor_85fbba2e40a9e6c7) }

var fileDescriptor_85fbba2e40a9e6c7 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xcf, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0xf1, 0x46, 0x61, 0x62, 0x11, 0x94, 0x22, 0x52, 0x07, 0x86, 0xe1, 0x69, 0x17, 0xfb,
	0x18, 0xfa, 0x09, 0x06, 0x5e, 0x3d, 0xcc, 0x9b, 0x20, 0x92, 0x2e, 0xcf, 0x6c, 0x98, 0xe5, 0x8d,
	0x34, 0xab, 0xdb, 0xb7, 0xf0, 0x63, 0x79, 0xdc, 0xd1, 0xa3, 0xb4, 0x5f, 0x64, 0x34, 0x6f, 0xbd,
	0xe5, 0xcf, 0x2f, 0x24, 0xef, 0xa5, 0x77, 0x6a, 0x8b, 0x56, 0x79, 0x28, 0xfd, 0x52, 0x1b, 0x84,
	0x7a, 0x02, 0x58, 0xa3, 0x0b, 0x55, 0xb1, 0xf6, 0x14, 0x28, 0xbb, 0x62, 0x2e, 0x98, 0x8b, 0x7a,
	0x32, 0xbc, 0x36, 0x64, 0x28, 0x22, 0x74, 0x27, 0xbe, 0x37, 0xbc, 0x9d, 0x53, 0xb5, 0xa2, 0xea,
	0x83, 0x81, 0xa3, 0x27, 0x43, 0x64, 0x2c, 0x42, 0xac, 0x72, 0xf3, 0x09, 0xca, 0xed, 0x98, 0xee,
	0xdf, 0xd3, 0xcb, 0x57, 0x74, 0x7a, 0x1a, 0x1f, 0x7f, 0xee, 0xfe, 0xcd, 0x6e, 0xd2, 0x41, 0x85,
	0x4e, 0xa3, 0xcf, 0xc5, 0x48, 0x8c, 0xcf, 0x67, 0xc7, 0xca, 0xf2, 0xf4, 0x6c, 0xa1, 0x9c, 0xb6,
	0xe8, 0xf3, 0x93, 0x08, 0x7d, 0x76, 0xb2, 0x56, 0x3b, 0x4b, 0x4a, 0xe7, 0xa7, 0x23, 0x31, 0xbe,
	0x98, 0xf5, 0x39, 0x7d, 0xf9, 0x6d, 0xa4, 0xd8, 0x37, 0x52, 0xfc, 0x37, 0x52, 0xfc, 0xb4, 0x32,
	0xd9, 0xb7, 0x32, 0xf9, 0x6b, 0x65, 0xf2, 0xf6, 0x64, 0x96, 0x61, 0xb1, 0x29, 0x8b, 0x39, 0xad,
	0xc0, 0xa3, 0x41, 0xf7, 0xe0, 0x30, 0x7c, 0x93, 0xff, 0x3a, 0x96, 0x45, 0x6d, 0xd0, 0xc3, 0x16,
	0x78, 0x7d, 0xde, 0xbe, 0x1c, 0xc4, 0xa9, 0x1f, 0x0f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xd8,
	0xdf, 0x99, 0x34, 0x01, 0x00, 0x00,
}

func (m *SendBridgeEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendBridgeEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendBridgeEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Handler) > 0 {
		i -= len(m.Handler)
		copy(dAtA[i:], m.Handler)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Handler)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendBridgeEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Handler)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendBridgeEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SendBridgeEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendBridgeEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Handler", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Handler = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
