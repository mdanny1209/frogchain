// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: frogchain/amm/pool.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Pool struct {
	Id          uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PoolParam   *PoolParam `protobuf:"bytes,2,opt,name=poolParam,proto3" json:"poolParam,omitempty"`
	PoolAssets  *PoolAsset `protobuf:"bytes,3,opt,name=poolAssets,proto3" json:"poolAssets,omitempty"`
	ShareToken  types.Coin `protobuf:"bytes,4,opt,name=shareToken,proto3" json:"shareToken"`
	IsActivated bool       `protobuf:"varint,5,opt,name=isActivated,proto3" json:"isActivated,omitempty"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59f92a00a9fe694, []int{0}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pool) GetPoolParam() *PoolParam {
	if m != nil {
		return m.PoolParam
	}
	return nil
}

func (m *Pool) GetPoolAssets() *PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func (m *Pool) GetShareToken() types.Coin {
	if m != nil {
		return m.ShareToken
	}
	return types.Coin{}
}

func (m *Pool) GetIsActivated() bool {
	if m != nil {
		return m.IsActivated
	}
	return false
}

func init() {
	proto.RegisterType((*Pool)(nil), "frogchain.amm.Pool")
}

func init() { proto.RegisterFile("frogchain/amm/pool.proto", fileDescriptor_b59f92a00a9fe694) }

var fileDescriptor_b59f92a00a9fe694 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x18, 0x85, 0xe3, 0xde, 0x5c, 0x04, 0xae, 0x60, 0xb0, 0x40, 0x32, 0x0c, 0x26, 0x62, 0xea, 0x64,
	0xab, 0x20, 0x21, 0x36, 0xd4, 0xf2, 0x02, 0x55, 0xc4, 0xc4, 0x82, 0xdc, 0xc4, 0xa4, 0x16, 0x75,
	0xfe, 0x28, 0xb6, 0x2a, 0x78, 0x0b, 0x1e, 0xab, 0x63, 0x47, 0x26, 0x84, 0x92, 0x91, 0x97, 0x40,
	0x4e, 0x5a, 0x1a, 0x04, 0x6c, 0x7f, 0x72, 0xbe, 0x23, 0x7f, 0x3a, 0x98, 0x3e, 0x94, 0x90, 0x25,
	0x33, 0xa9, 0x73, 0x21, 0x8d, 0x11, 0x05, 0xc0, 0x9c, 0x17, 0x25, 0x38, 0x20, 0xfb, 0x5f, 0x09,
	0x97, 0xc6, 0x9c, 0xb0, 0x9f, 0xe0, 0x7d, 0x21, 0x4b, 0x69, 0x5a, 0xfc, 0xd7, 0x5c, 0x5a, 0xab,
	0xdc, 0x3a, 0x3f, 0xcc, 0x20, 0x83, 0xe6, 0x14, 0xfe, 0xda, 0xb4, 0x12, 0xb0, 0x06, 0xac, 0x98,
	0x4a, 0xab, 0xc4, 0x62, 0x38, 0x55, 0x4e, 0x0e, 0x45, 0x02, 0x3a, 0x6f, 0xf3, 0xb3, 0x0f, 0x84,
	0xc3, 0x09, 0xc0, 0x9c, 0x1c, 0xe0, 0x9e, 0x4e, 0x29, 0x8a, 0xd0, 0x20, 0x8c, 0x7b, 0x3a, 0x25,
	0x97, 0x78, 0xcf, 0x3f, 0x31, 0xf1, 0x06, 0xb4, 0x17, 0xa1, 0x41, 0xff, 0x9c, 0xf2, 0x6f, 0xc6,
	0x7c, 0xb2, 0xc9, 0xe3, 0x2d, 0x4a, 0xae, 0x30, 0xf6, 0x1f, 0x23, 0x6f, 0x66, 0xe9, 0xbf, 0x3f,
	0x8b, 0x0d, 0x10, 0x77, 0x58, 0x72, 0x8d, 0xb1, 0x9d, 0xc9, 0x52, 0xdd, 0xc2, 0xa3, 0xca, 0x69,
	0xd8, 0x34, 0x8f, 0x79, 0xeb, 0xcf, 0xbd, 0x3f, 0x5f, 0xfb, 0xf3, 0x1b, 0xd0, 0xf9, 0x38, 0x5c,
	0xbe, 0x9d, 0x06, 0x71, 0xa7, 0x42, 0x22, 0xdc, 0xd7, 0x76, 0x94, 0x38, 0xbd, 0x90, 0x4e, 0xa5,
	0xf4, 0x7f, 0x84, 0x06, 0xbb, 0x71, 0xf7, 0xd7, 0x58, 0x2c, 0x2b, 0x86, 0x56, 0x15, 0x43, 0xef,
	0x15, 0x43, 0x2f, 0x35, 0x0b, 0x56, 0x35, 0x0b, 0x5e, 0x6b, 0x16, 0xdc, 0x1d, 0x6d, 0xd7, 0x7d,
	0x6a, 0xf6, 0x75, 0xcf, 0x85, 0xb2, 0xd3, 0x9d, 0x66, 0xa5, 0x8b, 0xcf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x96, 0x1c, 0xb4, 0x11, 0xc6, 0x01, 0x00, 0x00,
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsActivated {
		i--
		if m.IsActivated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.ShareToken.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PoolAssets != nil {
		{
			size, err := m.PoolAssets.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.PoolParam != nil {
		{
			size, err := m.PoolParam.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	if m.PoolParam != nil {
		l = m.PoolParam.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.PoolAssets != nil {
		l = m.PoolAssets.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.ShareToken.Size()
	n += 1 + l + sovPool(uint64(l))
	if m.IsActivated {
		n += 2
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParam", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParam == nil {
				m.PoolParam = &PoolParam{}
			}
			if err := m.PoolParam.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolAssets == nil {
				m.PoolAssets = &PoolAsset{}
			}
			if err := m.PoolAssets.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShareToken", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ShareToken.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActivated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActivated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)