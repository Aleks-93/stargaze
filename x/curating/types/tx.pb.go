// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stakebird/curating/v1beta1/tx.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type MsgPost struct {
	VendorID      uint32                                        `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id" yaml:"vendor_id"`
	PostID        string                                        `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id" yaml:"post_id"`
	Creator       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty" yaml:"creator"`
	RewardAccount github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=reward_account,json=rewardAccount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"reward_account,omitempty" yaml:"reward_account"`
	Body          string                                        `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty" yaml:"body"`
}

func (m *MsgPost) Reset()         { *m = MsgPost{} }
func (m *MsgPost) String() string { return proto.CompactTextString(m) }
func (*MsgPost) ProtoMessage()    {}
func (*MsgPost) Descriptor() ([]byte, []int) {
	return fileDescriptor_6419e195b675a5f8, []int{0}
}
func (m *MsgPost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgPost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgPost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgPost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPost.Merge(m, src)
}
func (m *MsgPost) XXX_Size() int {
	return m.Size()
}
func (m *MsgPost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPost proto.InternalMessageInfo

func (m *MsgPost) GetVendorID() uint32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgPost) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgPost) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgPost) GetRewardAccount() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.RewardAccount
	}
	return nil
}

func (m *MsgPost) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type MsgUpvote struct {
	VendorID      uint32                                        `protobuf:"varint,1,opt,name=vendor_id,json=vendorId,proto3" json:"vendor_id" yaml:"vendor_id"`
	PostID        string                                        `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id" yaml:"post_id"`
	Curator       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=curator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"curator,omitempty" yaml:"curator"`
	RewardAccount github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,4,opt,name=reward_account,json=rewardAccount,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"reward_account,omitempty" yaml:"reward_account"`
	VoteNum       int32                                         `protobuf:"varint,5,opt,name=vote_num,json=voteNum,proto3" json:"vote_num,omitempty"`
}

func (m *MsgUpvote) Reset()         { *m = MsgUpvote{} }
func (m *MsgUpvote) String() string { return proto.CompactTextString(m) }
func (*MsgUpvote) ProtoMessage()    {}
func (*MsgUpvote) Descriptor() ([]byte, []int) {
	return fileDescriptor_6419e195b675a5f8, []int{1}
}
func (m *MsgUpvote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpvote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpvote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpvote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpvote.Merge(m, src)
}
func (m *MsgUpvote) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpvote) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpvote.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpvote proto.InternalMessageInfo

func (m *MsgUpvote) GetVendorID() uint32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgUpvote) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *MsgUpvote) GetCurator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Curator
	}
	return nil
}

func (m *MsgUpvote) GetRewardAccount() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.RewardAccount
	}
	return nil
}

func (m *MsgUpvote) GetVoteNum() int32 {
	if m != nil {
		return m.VoteNum
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgPost)(nil), "stakebird.user.v1beta1.MsgPost")
	proto.RegisterType((*MsgUpvote)(nil), "stakebird.user.v1beta1.MsgUpvote")
}

func init() {
	proto.RegisterFile("stakebird/curating/v1beta1/tx.proto", fileDescriptor_6419e195b675a5f8)
}

var fileDescriptor_6419e195b675a5f8 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x1c, 0xcd, 0x91, 0x34, 0x7f, 0x0e, 0x5a, 0x90, 0x05, 0x28, 0x30, 0xf8, 0x22, 0x77, 0x20, 0x4b,
	0x6c, 0x55, 0x30, 0x75, 0x22, 0x81, 0x25, 0x95, 0x5a, 0x21, 0x4b, 0x30, 0x20, 0xa1, 0xe8, 0x7c,
	0x77, 0x32, 0x56, 0xeb, 0x9c, 0xb9, 0x3b, 0xa7, 0xcd, 0x77, 0x60, 0x60, 0xe0, 0x03, 0xf0, 0x71,
	0x18, 0x3b, 0x32, 0x9d, 0x90, 0xb3, 0x20, 0xc6, 0x8c, 0x4c, 0xe8, 0xee, 0xd2, 0x1a, 0x24, 0xa6,
	0x2c, 0x74, 0xf2, 0xcf, 0xef, 0x3d, 0xbf, 0x67, 0xdd, 0xbb, 0x1f, 0xdc, 0x97, 0x0a, 0x9f, 0xb2,
	0x24, 0x13, 0x34, 0x22, 0xa5, 0xc0, 0x2a, 0x9b, 0xa7, 0xd1, 0xe2, 0x20, 0x61, 0x0a, 0x1f, 0x44,
	0xea, 0x22, 0x2c, 0x04, 0x57, 0xdc, 0x7b, 0x78, 0x2d, 0x0a, 0x4b, 0xc9, 0x44, 0xb8, 0x11, 0x3c,
	0xbe, 0x9f, 0xf2, 0x94, 0x5b, 0x49, 0x64, 0x26, 0xa7, 0x0e, 0x3e, 0x37, 0x61, 0xe7, 0x58, 0xa6,
	0xaf, 0xb8, 0x54, 0xde, 0x11, 0xec, 0x2d, 0xd8, 0x9c, 0x72, 0x31, 0xcb, 0x68, 0x1f, 0x0c, 0xc0,
	0x70, 0x77, 0x32, 0xaa, 0x34, 0xea, 0xbe, 0xb1, 0xe0, 0xf4, 0xe5, 0x4f, 0x8d, 0x6a, 0xc1, 0x5a,
	0xa3, 0x7b, 0x4b, 0x9c, 0x9f, 0x1d, 0x06, 0xd7, 0x50, 0x10, 0x77, 0xdd, 0x3c, 0xa5, 0xde, 0x73,
	0xd8, 0x29, 0xb8, 0x54, 0xc6, 0xe9, 0xd6, 0x00, 0x0c, 0x7b, 0x93, 0x27, 0x95, 0x46, 0x6d, 0x13,
	0x63, 0x7d, 0xae, 0xc8, 0xb5, 0x46, 0x7b, 0xce, 0x65, 0x03, 0x04, 0x71, 0xdb, 0x4c, 0x53, 0xea,
	0xbd, 0x83, 0x1d, 0x22, 0x18, 0x56, 0x5c, 0xf4, 0x9b, 0x03, 0x30, 0xbc, 0x33, 0x79, 0x51, 0x8b,
	0x37, 0x44, 0xf0, 0x4b, 0xa3, 0x51, 0x9a, 0xa9, 0xf7, 0x65, 0x12, 0x12, 0x9e, 0x47, 0x84, 0xcb,
	0x9c, 0xcb, 0xcd, 0x63, 0x24, 0xe9, 0x69, 0xa4, 0x96, 0x05, 0x93, 0xe1, 0x98, 0x90, 0x31, 0xa5,
	0x82, 0x49, 0x19, 0x5f, 0x79, 0x7a, 0x1f, 0xe0, 0x9e, 0x60, 0xe7, 0x58, 0xd0, 0x19, 0x26, 0x84,
	0x97, 0x73, 0xd5, 0x6f, 0xd9, 0x94, 0xa3, 0xb5, 0x46, 0x0f, 0x5c, 0xca, 0xdf, 0xfc, 0x16, 0x61,
	0xbb, 0xce, 0x61, 0xec, 0x0c, 0xbc, 0x7d, 0xd8, 0x4a, 0x38, 0x5d, 0xf6, 0x77, 0xec, 0x81, 0xdc,
	0x5d, 0x6b, 0x74, 0xdb, 0x05, 0x19, 0x34, 0x88, 0x2d, 0x79, 0xd8, 0xfa, 0xf1, 0x05, 0x81, 0xe0,
	0x63, 0x13, 0xf6, 0x8e, 0x65, 0xfa, 0xba, 0x58, 0x70, 0xc5, 0x6e, 0x60, 0x31, 0xe6, 0xf6, 0xfd,
	0xb3, 0x18, 0x47, 0x6c, 0x55, 0x8c, 0xfb, 0xf4, 0x7f, 0x14, 0xf3, 0x08, 0x76, 0xcd, 0x39, 0xcf,
	0xe6, 0x65, 0x6e, 0xcb, 0xd9, 0x89, 0x3b, 0xe6, 0xfd, 0xa4, 0xcc, 0x5d, 0x1d, 0x93, 0x93, 0xaf,
	0x95, 0x0f, 0x2e, 0x2b, 0x1f, 0x7c, 0xaf, 0x7c, 0xf0, 0x69, 0xe5, 0x37, 0x2e, 0x57, 0x7e, 0xe3,
	0xdb, 0xca, 0x6f, 0xbc, 0x7d, 0xf6, 0x47, 0x70, 0x51, 0x26, 0x67, 0x19, 0x19, 0xe1, 0x73, 0x26,
	0x79, 0xce, 0xa2, 0x7a, 0x59, 0x2f, 0xea, 0x75, 0xb5, 0xbf, 0x92, 0xb4, 0xed, 0xf2, 0x3d, 0xfd,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x40, 0x17, 0x6b, 0xd1, 0x03, 0x00, 0x00,
}

func (this *MsgPost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgPost)
	if !ok {
		that2, ok := that.(MsgPost)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VendorID != that1.VendorID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if !bytes.Equal(this.RewardAccount, that1.RewardAccount) {
		return false
	}
	if this.Body != that1.Body {
		return false
	}
	return true
}
func (this *MsgUpvote) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpvote)
	if !ok {
		that2, ok := that.(MsgUpvote)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.VendorID != that1.VendorID {
		return false
	}
	if this.PostID != that1.PostID {
		return false
	}
	if !bytes.Equal(this.Curator, that1.Curator) {
		return false
	}
	if !bytes.Equal(this.RewardAccount, that1.RewardAccount) {
		return false
	}
	if this.VoteNum != that1.VoteNum {
		return false
	}
	return true
}
func (m *MsgPost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgPost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgPost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RewardAccount) > 0 {
		i -= len(m.RewardAccount)
		copy(dAtA[i:], m.RewardAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RewardAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpvote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpvote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpvote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoteNum != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VoteNum))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RewardAccount) > 0 {
		i -= len(m.RewardAccount)
		copy(dAtA[i:], m.RewardAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.RewardAccount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Curator) > 0 {
		i -= len(m.Curator)
		copy(dAtA[i:], m.Curator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Curator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgPost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RewardAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpvote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Curator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.RewardAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VoteNum != 0 {
		n += 1 + sovTx(uint64(m.VoteNum))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgPost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgPost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgPost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAccount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAccount = append(m.RewardAccount[:0], dAtA[iNdEx:postIndex]...)
			if m.RewardAccount == nil {
				m.RewardAccount = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpvote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpvote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpvote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Curator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Curator = append(m.Curator[:0], dAtA[iNdEx:postIndex]...)
			if m.Curator == nil {
				m.Curator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAccount", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardAccount = append(m.RewardAccount[:0], dAtA[iNdEx:postIndex]...)
			if m.RewardAccount == nil {
				m.RewardAccount = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteNum", wireType)
			}
			m.VoteNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoteNum |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)