// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params holds parameters for the tokenfactory module
type Params struct {
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc8299d306f3ff47, []int{0}
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

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.tokenfactory.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/params.proto", fileDescriptor_cc8299d306f3ff47)
}

var fileDescriptor_cc8299d306f3ff47 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x63, 0x7d, 0x52, 0x87, 0x7e, 0x0b, 0x8a, 0x18, 0x48, 0x85, 0x1c, 0x94, 0xa9, 0x0c,
	0xb5, 0x55, 0xe8, 0xc4, 0xd8, 0x4a, 0x6c, 0x45, 0x88, 0x91, 0x25, 0xba, 0x49, 0xdc, 0xd4, 0x6a,
	0x93, 0x1b, 0xc5, 0x2e, 0x22, 0x6f, 0xc1, 0xc4, 0xce, 0xca, 0x93, 0x74, 0xec, 0xc8, 0x54, 0x50,
	0xf2, 0x06, 0x3c, 0x01, 0xaa, 0xe3, 0xa2, 0x22, 0x24, 0x26, 0xfb, 0xe8, 0x9e, 0xf3, 0xbb, 0x7f,
	0xba, 0xe7, 0xa8, 0x32, 0x54, 0x52, 0x71, 0x8d, 0x0b, 0x91, 0xcf, 0x20, 0xd6, 0x58, 0x56, 0xfc,
	0x61, 0x18, 0x09, 0x0d, 0x43, 0x5e, 0x40, 0x09, 0x99, 0x62, 0x45, 0x89, 0x1a, 0xdd, 0x53, 0x6b,
	0x65, 0x87, 0x56, 0x66, 0xad, 0xbd, 0xe3, 0x14, 0x53, 0x34, 0x46, 0xbe, 0xfb, 0xb5, 0x99, 0xde,
	0xe8, 0x4f, 0x3c, 0xac, 0xf4, 0x1c, 0x4b, 0xa9, 0xab, 0xa9, 0xd0, 0x90, 0x80, 0x06, 0x9b, 0xf2,
	0x62, 0x13, 0x0b, 0x5b, 0x5c, 0x2b, 0x6c, 0x89, 0xb6, 0x8a, 0x47, 0xa0, 0xc4, 0x37, 0x27, 0x46,
	0x99, 0xb7, 0xf5, 0xe0, 0x85, 0x74, 0x3b, 0xb7, 0x66, 0x6a, 0xf7, 0x99, 0x74, 0xdd, 0x44, 0xe4,
	0x98, 0x85, 0x71, 0x29, 0x40, 0x4b, 0xcc, 0xc3, 0x99, 0x10, 0x27, 0xe4, 0xec, 0x5f, 0xff, 0xff,
	0x85, 0xc7, 0x2c, 0x76, 0x07, 0xda, 0x2f, 0xc1, 0x26, 0x28, 0xf3, 0xf1, 0x74, 0xbd, 0xf5, 0x9d,
	0xcf, 0xad, 0xef, 0x55, 0x90, 0x2d, 0xaf, 0x82, 0xdf, 0x88, 0xe0, 0xf5, 0xdd, 0xef, 0xa7, 0x52,
	0xcf, 0x57, 0x11, 0x8b, 0x31, 0xb3, 0x03, 0xda, 0x67, 0xa0, 0x92, 0x05, 0xd7, 0x55, 0x21, 0x94,
	0xa1, 0xa9, 0xbb, 0x23, 0x03, 0x98, 0xd8, 0xfc, 0xb5, 0x10, 0xe3, 0x9b, 0x75, 0x4d, 0xc9, 0xa6,
	0xa6, 0xe4, 0xa3, 0xa6, 0xe4, 0xa9, 0xa1, 0xce, 0xa6, 0xa1, 0xce, 0x5b, 0x43, 0x9d, 0xfb, 0xd1,
	0x01, 0xd5, 0x5e, 0x6e, 0xb0, 0x84, 0x48, 0xed, 0x05, 0x7f, 0xfc, 0x79, 0x48, 0xd3, 0x27, 0xea,
	0x98, 0xd5, 0x2f, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x5e, 0xae, 0x50, 0xcc, 0x01, 0x00,
	0x00,
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
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
