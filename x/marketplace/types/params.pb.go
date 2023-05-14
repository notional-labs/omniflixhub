// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: marketplace/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Params struct {
	SaleCommission   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=sale_commission,json=saleCommission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sale_commission" yaml:"sale_commission"`
	Distribution     Distribution                           `protobuf:"bytes,2,opt,name=distribution,proto3" json:"distribution"`
	BidCloseDuration time.Duration                          `protobuf:"bytes,3,opt,name=bid_close_duration,json=bidCloseDuration,proto3,stdduration" json:"bid_close_duration" yaml:"bid_close_duration"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_930108a08bf00b01, []int{0}
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

type Distribution struct {
	Staking       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=staking,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking" yaml:"staking"`
	CommunityPool github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=community_pool,json=communityPool,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"community_pool" yaml:"community_pool"`
}

func (m *Distribution) Reset()         { *m = Distribution{} }
func (m *Distribution) String() string { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()    {}
func (*Distribution) Descriptor() ([]byte, []int) {
	return fileDescriptor_930108a08bf00b01, []int{1}
}
func (m *Distribution) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Distribution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Distribution.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Distribution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Distribution.Merge(m, src)
}
func (m *Distribution) XXX_Size() int {
	return m.Size()
}
func (m *Distribution) XXX_DiscardUnknown() {
	xxx_messageInfo_Distribution.DiscardUnknown(m)
}

var xxx_messageInfo_Distribution proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "OmniFlix.marketplace.v1beta1.Params")
	proto.RegisterType((*Distribution)(nil), "OmniFlix.marketplace.v1beta1.Distribution")
}

func init() { proto.RegisterFile("marketplace/v1beta1/params.proto", fileDescriptor_930108a08bf00b01) }

var fileDescriptor_930108a08bf00b01 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0x87, 0x93, 0x2a, 0x95, 0xc6, 0xba, 0x4a, 0x50, 0xd9, 0x16, 0x99, 0x2c, 0x01, 0xa5, 0x08,
	0xce, 0x50, 0x7b, 0xf3, 0x24, 0xe9, 0xa2, 0xde, 0xac, 0xc1, 0x53, 0x2f, 0xcb, 0x4c, 0x32, 0x4d,
	0x87, 0x9d, 0x99, 0x37, 0x66, 0x26, 0xb2, 0xfb, 0x2d, 0x3c, 0xfa, 0x75, 0xbc, 0xed, 0xb1, 0x47,
	0xe9, 0x61, 0xd5, 0xdd, 0x6f, 0xd0, 0x4f, 0x20, 0xf9, 0x57, 0x52, 0x85, 0xc2, 0x9e, 0x92, 0xbc,
	0x79, 0xdf, 0xe7, 0xf7, 0xce, 0x93, 0x78, 0x23, 0x45, 0x8b, 0x29, 0xb7, 0xb9, 0xa4, 0x09, 0x27,
	0x5f, 0x0f, 0x19, 0xb7, 0xf4, 0x90, 0xe4, 0xb4, 0xa0, 0xca, 0xe0, 0xbc, 0x00, 0x0b, 0xfe, 0xb3,
	0x8f, 0x4a, 0x8b, 0x77, 0x52, 0xcc, 0x70, 0xaf, 0x15, 0xb7, 0xad, 0xfb, 0x8f, 0x33, 0xc8, 0xa0,
	0x6e, 0x24, 0xd5, 0x5d, 0x33, 0xb3, 0x8f, 0x32, 0x80, 0x4c, 0x72, 0x52, 0x3f, 0xb1, 0xf2, 0x8c,
	0xa4, 0x65, 0x41, 0xad, 0x00, 0xdd, 0xbc, 0x0f, 0x7f, 0x6c, 0x79, 0xdb, 0x27, 0x75, 0x88, 0xff,
	0xc5, 0x7b, 0x68, 0xa8, 0xe4, 0x93, 0x04, 0x94, 0x12, 0xc6, 0x08, 0xd0, 0x43, 0x77, 0xe4, 0x1e,
	0xec, 0x44, 0x1f, 0x16, 0xcb, 0xc0, 0xb9, 0x5c, 0x06, 0x2f, 0x32, 0x61, 0xcf, 0x4b, 0x86, 0x13,
	0x50, 0x24, 0x01, 0xa3, 0xc0, 0xb4, 0x97, 0x57, 0x26, 0x9d, 0x12, 0x3b, 0xcf, 0xb9, 0xc1, 0x63,
	0x9e, 0x5c, 0x2d, 0x83, 0xa7, 0x73, 0xaa, 0xe4, 0x9b, 0xf0, 0x1f, 0x5c, 0x18, 0x0f, 0xaa, 0xca,
	0xf1, 0x75, 0xc1, 0xff, 0xec, 0xed, 0xa6, 0xc2, 0xd8, 0x42, 0xb0, 0xb2, 0xda, 0x69, 0xb8, 0x35,
	0x72, 0x0f, 0xee, 0xbf, 0x7e, 0x89, 0x6f, 0x3b, 0x28, 0x1e, 0xf7, 0x26, 0xa2, 0xbb, 0xd5, 0x6e,
	0xf1, 0x0d, 0x8a, 0xaf, 0x3d, 0x9f, 0x89, 0x74, 0x92, 0x48, 0x30, 0x7c, 0xd2, 0x9d, 0x77, 0x78,
	0xa7, 0x66, 0xef, 0xe1, 0x46, 0x08, 0xee, 0x84, 0xe0, 0x71, 0xdb, 0x10, 0x3d, 0xaf, 0x50, 0x57,
	0xcb, 0x60, 0xaf, 0x59, 0xfe, 0x7f, 0x44, 0xf8, 0xfd, 0x57, 0xe0, 0xc6, 0x8f, 0x98, 0x48, 0x8f,
	0xab, 0x7a, 0x37, 0x18, 0x5e, 0xba, 0xde, 0x6e, 0x7f, 0x29, 0xff, 0xd4, 0xbb, 0x67, 0x2c, 0x9d,
	0x0a, 0x9d, 0xb5, 0x06, 0xdf, 0x6e, 0x6c, 0x70, 0xd0, 0x1a, 0x6c, 0x30, 0x61, 0xdc, 0x01, 0x7d,
	0xed, 0x0d, 0x2a, 0xa3, 0xa5, 0x16, 0x76, 0x3e, 0xc9, 0x01, 0x64, 0x2d, 0x6d, 0x27, 0x7a, 0xbf,
	0x71, 0xc4, 0x93, 0x26, 0xe2, 0x26, 0x2d, 0x8c, 0x1f, 0x5c, 0x17, 0x4e, 0x00, 0x64, 0xf4, 0x69,
	0xf1, 0x07, 0x39, 0x8b, 0x15, 0x72, 0x2f, 0x56, 0xc8, 0xfd, 0xbd, 0x42, 0xee, 0xb7, 0x35, 0x72,
	0x2e, 0xd6, 0xc8, 0xf9, 0xb9, 0x46, 0xce, 0xe9, 0x51, 0x2f, 0xad, 0xfb, 0x68, 0x04, 0x94, 0x16,
	0x67, 0x52, 0xcc, 0xce, 0x4b, 0x46, 0x66, 0xa4, 0xff, 0x5b, 0xd7, 0xf1, 0x6c, 0xbb, 0x76, 0x7f,
	0xf4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x55, 0x22, 0x04, 0xf2, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BidCloseDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.BidCloseDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Distribution.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.SaleCommission.Size()
		i -= size
		if _, err := m.SaleCommission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Distribution) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Distribution) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Distribution) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CommunityPool.Size()
		i -= size
		if _, err := m.CommunityPool.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.SaleCommission.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Distribution.Size()
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BidCloseDuration)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *Distribution) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Staking.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CommunityPool.Size()
	n += 1 + l + sovParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field SaleCommission", wireType)
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
			if err := m.SaleCommission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Distribution", wireType)
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
			if err := m.Distribution.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidCloseDuration", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BidCloseDuration, dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Distribution) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Distribution: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Distribution: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
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
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommunityPool", wireType)
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
			if err := m.CommunityPool.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
