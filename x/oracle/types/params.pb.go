// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: oracle/params.proto

package types

import (
	fmt "fmt"
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

// Params defines the parameters for the module.
type Params struct {
	MaxOracles             uint64 `protobuf:"varint,1,opt,name=max_oracles,json=maxOracles,proto3" json:"max_oracles,omitempty" yaml:"max_oracles"`
	SlashFractionDowntime  string `protobuf:"bytes,2,opt,name=slash_fraction_downtime,json=slashFractionDowntime,proto3" json:"slash_fraction_downtime,omitempty" yaml:"slash_fraction_downtime"`
	SlashFractionWrongVote string `protobuf:"bytes,3,opt,name=slash_fraction_wrong_vote,json=slashFractionWrongVote,proto3" json:"slash_fraction_wrong_vote,omitempty" yaml:"slash_fraction_wrong_vote"`
	DowntimeDuration       string `protobuf:"bytes,4,opt,name=downtime_duration,json=downtimeDuration,proto3" json:"downtime_duration,omitempty" yaml:"downtime_duration"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db12dafa3fbe2a3, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxOracles() uint64 {
	if m != nil {
		return m.MaxOracles
	}
	return 0
}

func (m *Params) GetSlashFractionDowntime() string {
	if m != nil {
		return m.SlashFractionDowntime
	}
	return ""
}

func (m *Params) GetSlashFractionWrongVote() string {
	if m != nil {
		return m.SlashFractionWrongVote
	}
	return ""
}

func (m *Params) GetDowntimeDuration() string {
	if m != nil {
		return m.DowntimeDuration
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "youngjoonlee.dhub.oracle.Params")
}

func init() { proto.RegisterFile("oracle/params.proto", fileDescriptor_2db12dafa3fbe2a3) }

var fileDescriptor_2db12dafa3fbe2a3 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2f, 0x4a, 0x4c,
	0xce, 0x49, 0xd5, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0xa8, 0xcc, 0x2f, 0xcd, 0x4b, 0xcf, 0xca, 0xcf, 0xcf, 0xcb, 0x49, 0x4d, 0xd5, 0x4b, 0xc9,
	0x28, 0x4d, 0xd2, 0x83, 0x28, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1,
	0x20, 0xea, 0x95, 0xce, 0x33, 0x71, 0xb1, 0x05, 0x80, 0x0d, 0x10, 0x32, 0xe7, 0xe2, 0xce, 0x4d,
	0xac, 0x88, 0x87, 0x28, 0x2f, 0x96, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x71, 0x12, 0xfb, 0x74, 0x4f,
	0x5e, 0xa8, 0x32, 0x31, 0x37, 0xc7, 0x4a, 0x09, 0x49, 0x52, 0x29, 0x88, 0x2b, 0x37, 0xb1, 0xc2,
	0x1f, 0xc2, 0x11, 0x8a, 0xe2, 0x12, 0x2f, 0xce, 0x49, 0x2c, 0xce, 0x88, 0x4f, 0x2b, 0x4a, 0x4c,
	0x2e, 0xc9, 0xcc, 0xcf, 0x8b, 0x4f, 0xc9, 0x2f, 0xcf, 0x2b, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x74, 0x52, 0xfa, 0x74, 0x4f, 0x5e, 0x0e, 0x62, 0x08, 0x0e, 0x85, 0x4a, 0x41,
	0xa2, 0x60, 0x19, 0x37, 0xa8, 0x84, 0x0b, 0x54, 0x5c, 0x28, 0x9e, 0x4b, 0x12, 0x4d, 0x4b, 0x79,
	0x51, 0x7e, 0x5e, 0x7a, 0x7c, 0x59, 0x7e, 0x49, 0xaa, 0x04, 0x33, 0xd8, 0x74, 0x95, 0x4f, 0xf7,
	0xe4, 0x15, 0xb0, 0x9a, 0x8e, 0x50, 0xaa, 0x14, 0x24, 0x86, 0x62, 0x7e, 0x38, 0x48, 0x26, 0x2c,
	0xbf, 0x24, 0x55, 0xc8, 0x93, 0x4b, 0x10, 0xe6, 0x88, 0xf8, 0x94, 0xd2, 0xa2, 0x44, 0x90, 0xac,
	0x04, 0x0b, 0xd8, 0x60, 0x99, 0x4f, 0xf7, 0xe4, 0x25, 0x20, 0x06, 0x63, 0x28, 0x51, 0x0a, 0x12,
	0x80, 0x89, 0xb9, 0x40, 0x85, 0xac, 0x58, 0x66, 0x2c, 0x90, 0x67, 0x70, 0x72, 0x3b, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x9d, 0xf4, 0xcc, 0x12, 0x50, 0x84, 0x24, 0xe7,
	0xe7, 0xea, 0xc3, 0xa3, 0x49, 0x37, 0x27, 0x35, 0x55, 0x1f, 0x14, 0x4f, 0xfa, 0x15, 0xfa, 0xd0,
	0x08, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x47, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x0c, 0xb1, 0x49, 0xe7, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DowntimeDuration) > 0 {
		i -= len(m.DowntimeDuration)
		copy(dAtA[i:], m.DowntimeDuration)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DowntimeDuration)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SlashFractionWrongVote) > 0 {
		i -= len(m.SlashFractionWrongVote)
		copy(dAtA[i:], m.SlashFractionWrongVote)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashFractionWrongVote)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SlashFractionDowntime) > 0 {
		i -= len(m.SlashFractionDowntime)
		copy(dAtA[i:], m.SlashFractionDowntime)
		i = encodeVarintParams(dAtA, i, uint64(len(m.SlashFractionDowntime)))
		i--
		dAtA[i] = 0x12
	}
	if m.MaxOracles != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxOracles))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxOracles != 0 {
		n += 1 + sovParams(uint64(m.MaxOracles))
	}
	l = len(m.SlashFractionDowntime)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.SlashFractionWrongVote)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.DowntimeDuration)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxOracles", wireType)
			}
			m.MaxOracles = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxOracles |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionDowntime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashFractionDowntime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFractionWrongVote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SlashFractionWrongVote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DowntimeDuration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DowntimeDuration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
