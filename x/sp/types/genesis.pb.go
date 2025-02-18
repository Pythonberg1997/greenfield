// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: greenfield/sp/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the sp module's genesis state.
type GenesisState struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// this used by starport scaffolding # genesis/proto/state
	StorageProviders   []StorageProvider `protobuf:"bytes,2,rep,name=storage_providers,json=storageProviders,proto3" json:"storage_providers"`
	SpStoragePriceList []SpStoragePrice  `protobuf:"bytes,3,rep,name=sp_storage_price_list,json=spStoragePriceList,proto3" json:"sp_storage_price_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cf352e27d3a7d62, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStorageProviders() []StorageProvider {
	if m != nil {
		return m.StorageProviders
	}
	return nil
}

func (m *GenesisState) GetSpStoragePriceList() []SpStoragePrice {
	if m != nil {
		return m.SpStoragePriceList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "bnbchain.greenfield.sp.GenesisState")
}

func init() { proto.RegisterFile("greenfield/sp/genesis.proto", fileDescriptor_3cf352e27d3a7d62) }

var fileDescriptor_3cf352e27d3a7d62 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x93, 0x56, 0x3a, 0xa4, 0x0e, 0x1a, 0x54, 0x62, 0x84, 0xb3, 0x38, 0x68, 0x11, 0xbc,
	0x40, 0x5d, 0x9d, 0x8a, 0xe0, 0xe2, 0x50, 0xec, 0xd6, 0x25, 0xe4, 0xd2, 0xe7, 0xf5, 0xa0, 0xcd,
	0x1d, 0xf7, 0x4e, 0xd1, 0x8f, 0xe0, 0xe6, 0xc7, 0xea, 0xd8, 0xd1, 0x49, 0x24, 0xf9, 0x22, 0xe2,
	0xe5, 0x68, 0x0d, 0xb4, 0xdb, 0xf1, 0xde, 0xef, 0xfd, 0xfe, 0xc7, 0x3f, 0x38, 0xe3, 0x1a, 0xa0,
	0x78, 0x16, 0x30, 0x9f, 0x26, 0xa8, 0x12, 0x0e, 0x05, 0xa0, 0x40, 0xaa, 0xb4, 0x34, 0x32, 0x3c,
	0x61, 0x05, 0xcb, 0x67, 0x99, 0x28, 0xe8, 0x86, 0xa2, 0xa8, 0xe2, 0x23, 0x2e, 0xb9, 0xb4, 0x48,
	0xf2, 0xf7, 0xaa, 0xe9, 0x38, 0x6e, 0xaa, 0x54, 0xa6, 0xb3, 0x85, 0x33, 0xc5, 0xa7, 0xcd, 0x9d,
	0x79, 0x57, 0xe0, 0x56, 0x17, 0x1f, 0xad, 0x60, 0xff, 0xa1, 0x8e, 0x1d, 0x9b, 0xcc, 0x40, 0x78,
	0x17, 0x74, 0xea, 0xdb, 0xc8, 0xef, 0xf9, 0xfd, 0xee, 0x80, 0xd0, 0xed, 0xdf, 0xa0, 0x23, 0x4b,
	0x0d, 0xf7, 0x96, 0xdf, 0xe7, 0xde, 0x93, 0xbb, 0x09, 0x27, 0xc1, 0x21, 0x1a, 0xa9, 0x33, 0x0e,
	0xa9, 0xd2, 0xf2, 0x55, 0x4c, 0x41, 0x63, 0xd4, 0xea, 0xb5, 0xfb, 0xdd, 0xc1, 0xd5, 0x2e, 0xd1,
	0xb8, 0x3e, 0x18, 0x39, 0xde, 0x19, 0x0f, 0xb0, 0x39, 0xc6, 0x30, 0x0d, 0x8e, 0x51, 0xa5, 0x1b,
	0xbd, 0xc8, 0x21, 0x9d, 0x0b, 0x34, 0x51, 0xdb, 0xfa, 0x2f, 0x77, 0xfa, 0xd5, 0x3a, 0x41, 0xe4,
	0xe0, 0xf4, 0x21, 0x36, 0xa6, 0x8f, 0x02, 0xcd, 0xf0, 0x7e, 0x59, 0x12, 0x7f, 0x55, 0x12, 0xff,
	0xa7, 0x24, 0xfe, 0x67, 0x45, 0xbc, 0x55, 0x45, 0xbc, 0xaf, 0x8a, 0x78, 0x93, 0x6b, 0x2e, 0xcc,
	0xec, 0x85, 0xd1, 0x5c, 0x2e, 0x12, 0x56, 0xb0, 0x1b, 0x1b, 0x93, 0xfc, 0x6b, 0xf5, 0x6d, 0xdd,
	0x2b, 0xeb, 0xd8, 0x62, 0x6f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x7a, 0x5d, 0x31, 0xdc,
	0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpStoragePriceList) > 0 {
		for iNdEx := len(m.SpStoragePriceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpStoragePriceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StorageProviders) > 0 {
		for iNdEx := len(m.StorageProviders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StorageProviders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StorageProviders) > 0 {
		for _, e := range m.StorageProviders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SpStoragePriceList) > 0 {
		for _, e := range m.SpStoragePriceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StorageProviders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StorageProviders = append(m.StorageProviders, StorageProvider{})
			if err := m.StorageProviders[len(m.StorageProviders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpStoragePriceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpStoragePriceList = append(m.SpStoragePriceList, SpStoragePrice{})
			if err := m.SpStoragePriceList[len(m.SpStoragePriceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
