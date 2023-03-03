// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/v1/bls_key.proto

package types

import (
	fmt "fmt"
	github_com_babylonchain_babylon_crypto_bls12381 "github.com/babylonchain/babylon/crypto/bls12381"
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

// BlsKey wraps BLS public key with PoP
type BlsKey struct {
	// pubkey is the BLS public key of a validator
	Pubkey *github_com_babylonchain_babylon_crypto_bls12381.PublicKey `protobuf:"bytes,2,opt,name=pubkey,proto3,customtype=github.com/babylonchain/babylon/crypto/bls12381.PublicKey" json:"pubkey,omitempty"`
	// pop is the proof-of-possession of the BLS key
	Pop *ProofOfPossession `protobuf:"bytes,3,opt,name=pop,proto3" json:"pop,omitempty"`
}

func (m *BlsKey) Reset()         { *m = BlsKey{} }
func (m *BlsKey) String() string { return proto.CompactTextString(m) }
func (*BlsKey) ProtoMessage()    {}
func (*BlsKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a8c0d37ce63f038, []int{0}
}
func (m *BlsKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlsKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlsKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlsKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlsKey.Merge(m, src)
}
func (m *BlsKey) XXX_Size() int {
	return m.Size()
}
func (m *BlsKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BlsKey.DiscardUnknown(m)
}

var xxx_messageInfo_BlsKey proto.InternalMessageInfo

func (m *BlsKey) GetPop() *ProofOfPossession {
	if m != nil {
		return m.Pop
	}
	return nil
}

// ProofOfPossession defines proof for the ownership of Ed25519 and BLS private keys
type ProofOfPossession struct {
	// ed25519_sig is used for verification, ed25519_sig = sign(key = Ed25519_sk, data = BLS_pk)
	Ed25519Sig []byte `protobuf:"bytes,2,opt,name=ed25519_sig,json=ed25519Sig,proto3" json:"ed25519_sig,omitempty"`
	// bls_sig is the result of PoP, bls_sig = sign(key = BLS_sk, data = ed25519_sig)
	BlsSig *github_com_babylonchain_babylon_crypto_bls12381.Signature `protobuf:"bytes,3,opt,name=bls_sig,json=blsSig,proto3,customtype=github.com/babylonchain/babylon/crypto/bls12381.Signature" json:"bls_sig,omitempty"`
}

func (m *ProofOfPossession) Reset()         { *m = ProofOfPossession{} }
func (m *ProofOfPossession) String() string { return proto.CompactTextString(m) }
func (*ProofOfPossession) ProtoMessage()    {}
func (*ProofOfPossession) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a8c0d37ce63f038, []int{1}
}
func (m *ProofOfPossession) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProofOfPossession) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProofOfPossession.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProofOfPossession) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProofOfPossession.Merge(m, src)
}
func (m *ProofOfPossession) XXX_Size() int {
	return m.Size()
}
func (m *ProofOfPossession) XXX_DiscardUnknown() {
	xxx_messageInfo_ProofOfPossession.DiscardUnknown(m)
}

var xxx_messageInfo_ProofOfPossession proto.InternalMessageInfo

func (m *ProofOfPossession) GetEd25519Sig() []byte {
	if m != nil {
		return m.Ed25519Sig
	}
	return nil
}

// ValidatorWithBLSSet defines a set of validators with their BLS public keys
type ValidatorWithBlsKeySet struct {
	ValSet []*ValidatorWithBlsKey `protobuf:"bytes,1,rep,name=val_set,json=valSet,proto3" json:"val_set,omitempty"`
}

func (m *ValidatorWithBlsKeySet) Reset()         { *m = ValidatorWithBlsKeySet{} }
func (m *ValidatorWithBlsKeySet) String() string { return proto.CompactTextString(m) }
func (*ValidatorWithBlsKeySet) ProtoMessage()    {}
func (*ValidatorWithBlsKeySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a8c0d37ce63f038, []int{2}
}
func (m *ValidatorWithBlsKeySet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorWithBlsKeySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorWithBlsKeySet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorWithBlsKeySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorWithBlsKeySet.Merge(m, src)
}
func (m *ValidatorWithBlsKeySet) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorWithBlsKeySet) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorWithBlsKeySet.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorWithBlsKeySet proto.InternalMessageInfo

func (m *ValidatorWithBlsKeySet) GetValSet() []*ValidatorWithBlsKey {
	if m != nil {
		return m.ValSet
	}
	return nil
}

// ValidatorWithBlsKey couples validator address, voting power, and its bls public key
type ValidatorWithBlsKey struct {
	ValidatorAddress string `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	BlsPubKey        []byte `protobuf:"bytes,2,opt,name=bls_pub_key,json=blsPubKey,proto3" json:"bls_pub_key,omitempty"`
	VotingPower      uint64 `protobuf:"varint,3,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *ValidatorWithBlsKey) Reset()         { *m = ValidatorWithBlsKey{} }
func (m *ValidatorWithBlsKey) String() string { return proto.CompactTextString(m) }
func (*ValidatorWithBlsKey) ProtoMessage()    {}
func (*ValidatorWithBlsKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a8c0d37ce63f038, []int{3}
}
func (m *ValidatorWithBlsKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorWithBlsKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorWithBlsKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorWithBlsKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorWithBlsKey.Merge(m, src)
}
func (m *ValidatorWithBlsKey) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorWithBlsKey) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorWithBlsKey.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorWithBlsKey proto.InternalMessageInfo

func (m *ValidatorWithBlsKey) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *ValidatorWithBlsKey) GetBlsPubKey() []byte {
	if m != nil {
		return m.BlsPubKey
	}
	return nil
}

func (m *ValidatorWithBlsKey) GetVotingPower() uint64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*BlsKey)(nil), "babylon.checkpointing.v1.BlsKey")
	proto.RegisterType((*ProofOfPossession)(nil), "babylon.checkpointing.v1.ProofOfPossession")
	proto.RegisterType((*ValidatorWithBlsKeySet)(nil), "babylon.checkpointing.v1.ValidatorWithBlsKeySet")
	proto.RegisterType((*ValidatorWithBlsKey)(nil), "babylon.checkpointing.v1.ValidatorWithBlsKey")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/v1/bls_key.proto", fileDescriptor_3a8c0d37ce63f038)
}

var fileDescriptor_3a8c0d37ce63f038 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0xdd, 0x71, 0x25, 0xa5, 0x93, 0x1e, 0x6c, 0x14, 0x09, 0x1e, 0xd2, 0x75, 0x0f, 0xb2, 0x50,
	0x4c, 0xc8, 0x96, 0x05, 0x7b, 0xe8, 0xc1, 0x3d, 0x78, 0xe9, 0xa1, 0x21, 0xc1, 0x0a, 0x5e, 0xe2,
	0x4c, 0x76, 0x3a, 0x3b, 0xec, 0x98, 0x6f, 0xc8, 0x4c, 0xa2, 0xf9, 0x01, 0xde, 0x3c, 0xf8, 0x0b,
	0xfc, 0x3d, 0x1e, 0x7b, 0x14, 0x0f, 0x22, 0xbb, 0x7f, 0x44, 0x66, 0x13, 0x17, 0xd4, 0x16, 0xc1,
	0x5b, 0x78, 0xef, 0xe5, 0xf1, 0xde, 0x9b, 0x0f, 0x3f, 0xa1, 0x84, 0xb6, 0x12, 0xca, 0xa8, 0x58,
	0xb2, 0x62, 0xa5, 0x40, 0x94, 0x46, 0x94, 0x3c, 0x6a, 0xe2, 0x88, 0x4a, 0x9d, 0xaf, 0x58, 0x1b,
	0xaa, 0x0a, 0x0c, 0x78, 0x7e, 0xaf, 0x0b, 0x7f, 0xd3, 0x85, 0x4d, 0xfc, 0xe8, 0x01, 0x07, 0x0e,
	0x5b, 0x51, 0x64, 0xbf, 0x3a, 0xfd, 0xf8, 0x33, 0xc2, 0xce, 0x5c, 0xea, 0x73, 0xd6, 0x7a, 0x2f,
	0xb1, 0xa3, 0x6a, 0xba, 0x62, 0xad, 0x7f, 0x67, 0x84, 0x26, 0x07, 0xf3, 0xb3, 0x6f, 0xdf, 0x8f,
	0x4e, 0xb9, 0x30, 0xcb, 0x9a, 0x86, 0x05, 0xbc, 0x8d, 0x7a, 0xe7, 0x62, 0x49, 0x44, 0x19, 0xed,
	0xe2, 0x54, 0xad, 0x32, 0x60, 0x43, 0xc4, 0xd3, 0x93, 0x67, 0x71, 0x98, 0xd4, 0x54, 0x8a, 0xe2,
	0x9c, 0xb5, 0x69, 0x6f, 0xe6, 0x9d, 0xe1, 0xa1, 0x02, 0xe5, 0x0f, 0x47, 0x68, 0xe2, 0x4e, 0x8f,
	0xc3, 0xdb, 0xf2, 0x85, 0x49, 0x05, 0x70, 0x75, 0x71, 0x95, 0x80, 0xd6, 0x4c, 0x6b, 0x01, 0x65,
	0x6a, 0xff, 0x1b, 0x7f, 0x44, 0xf8, 0xf0, 0x2f, 0xca, 0x3b, 0xc2, 0x2e, 0x5b, 0x4c, 0x67, 0xb3,
	0xf8, 0x34, 0xd7, 0x82, 0x77, 0x81, 0x53, 0xdc, 0x43, 0x99, 0xe0, 0xde, 0x25, 0xde, 0xb3, 0xc3,
	0x58, 0x72, 0xf8, 0xff, 0x6d, 0x32, 0xc1, 0x4b, 0x62, 0xea, 0x8a, 0xa5, 0x0e, 0x95, 0x3a, 0x13,
	0x7c, 0xfc, 0x06, 0x3f, 0xbc, 0x24, 0x52, 0x2c, 0x88, 0x81, 0xea, 0x95, 0x30, 0xcb, 0x6e, 0xbb,
	0x8c, 0x19, 0xef, 0x05, 0xde, 0x6b, 0x88, 0xcc, 0x35, 0x33, 0x3e, 0x1a, 0x0d, 0x27, 0xee, 0xf4,
	0xe9, 0xed, 0x5d, 0x6f, 0xb0, 0x48, 0x9d, 0x86, 0xc8, 0x8c, 0x99, 0xf1, 0x07, 0x84, 0xef, 0xdf,
	0xc0, 0x7b, 0xc7, 0xf8, 0xb0, 0xf9, 0x05, 0xe7, 0x64, 0xb1, 0xa8, 0x98, 0xd6, 0x3e, 0x1a, 0xa1,
	0xc9, 0x7e, 0x7a, 0x6f, 0x47, 0x3c, 0xef, 0x70, 0x2f, 0xc0, 0xae, 0xad, 0xaf, 0x6a, 0x9a, 0xef,
	0x1e, 0x34, 0xdd, 0xa7, 0x52, 0x27, 0x35, 0xb5, 0x66, 0x8f, 0xf1, 0x41, 0x03, 0x36, 0x4d, 0xae,
	0xe0, 0x1d, 0xab, 0xb6, 0x1b, 0xdd, 0x4d, 0xdd, 0x0e, 0x4b, 0x2c, 0x34, 0xbf, 0xf8, 0xb2, 0x0e,
	0xd0, 0xf5, 0x3a, 0x40, 0x3f, 0xd6, 0x01, 0xfa, 0xb4, 0x09, 0x06, 0xd7, 0x9b, 0x60, 0xf0, 0x75,
	0x13, 0x0c, 0x5e, 0xcf, 0xfe, 0x35, 0xe3, 0xfb, 0x3f, 0xae, 0xd4, 0xb4, 0x8a, 0x69, 0xea, 0x6c,
	0x2f, 0xee, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x90, 0xb1, 0x5b, 0xcb, 0x02, 0x00,
	0x00,
}

func (m *BlsKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlsKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlsKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pop != nil {
		{
			size, err := m.Pop.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Pubkey != nil {
		{
			size := m.Pubkey.Size()
			i -= size
			if _, err := m.Pubkey.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ProofOfPossession) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProofOfPossession) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProofOfPossession) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlsSig != nil {
		{
			size := m.BlsSig.Size()
			i -= size
			if _, err := m.BlsSig.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBlsKey(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Ed25519Sig) > 0 {
		i -= len(m.Ed25519Sig)
		copy(dAtA[i:], m.Ed25519Sig)
		i = encodeVarintBlsKey(dAtA, i, uint64(len(m.Ed25519Sig)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorWithBlsKeySet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorWithBlsKeySet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorWithBlsKeySet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValSet) > 0 {
		for iNdEx := len(m.ValSet) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValSet[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlsKey(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorWithBlsKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorWithBlsKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorWithBlsKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintBlsKey(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x18
	}
	if len(m.BlsPubKey) > 0 {
		i -= len(m.BlsPubKey)
		copy(dAtA[i:], m.BlsPubKey)
		i = encodeVarintBlsKey(dAtA, i, uint64(len(m.BlsPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintBlsKey(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlsKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlsKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlsKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pubkey != nil {
		l = m.Pubkey.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	if m.Pop != nil {
		l = m.Pop.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	return n
}

func (m *ProofOfPossession) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ed25519Sig)
	if l > 0 {
		n += 1 + l + sovBlsKey(uint64(l))
	}
	if m.BlsSig != nil {
		l = m.BlsSig.Size()
		n += 1 + l + sovBlsKey(uint64(l))
	}
	return n
}

func (m *ValidatorWithBlsKeySet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValSet) > 0 {
		for _, e := range m.ValSet {
			l = e.Size()
			n += 1 + l + sovBlsKey(uint64(l))
		}
	}
	return n
}

func (m *ValidatorWithBlsKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovBlsKey(uint64(l))
	}
	l = len(m.BlsPubKey)
	if l > 0 {
		n += 1 + l + sovBlsKey(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovBlsKey(uint64(m.VotingPower))
	}
	return n
}

func sovBlsKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlsKey(x uint64) (n int) {
	return sovBlsKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlsKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: BlsKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlsKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_crypto_bls12381.PublicKey
			m.Pubkey = &v
			if err := m.Pubkey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pop", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pop == nil {
				m.Pop = &ProofOfPossession{}
			}
			if err := m.Pop.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func (m *ProofOfPossession) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: ProofOfPossession: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProofOfPossession: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ed25519Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ed25519Sig = append(m.Ed25519Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Ed25519Sig == nil {
				m.Ed25519Sig = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_babylonchain_babylon_crypto_bls12381.Signature
			m.BlsSig = &v
			if err := m.BlsSig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func (m *ValidatorWithBlsKeySet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: ValidatorWithBlsKeySet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorWithBlsKeySet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValSet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValSet = append(m.ValSet, &ValidatorWithBlsKey{})
			if err := m.ValSet[len(m.ValSet)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func (m *ValidatorWithBlsKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlsKey
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
			return fmt.Errorf("proto: ValidatorWithBlsKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorWithBlsKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsPubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
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
				return ErrInvalidLengthBlsKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlsKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlsPubKey = append(m.BlsPubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.BlsPubKey == nil {
				m.BlsPubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlsKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBlsKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlsKey
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
func skipBlsKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlsKey
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
					return 0, ErrIntOverflowBlsKey
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
					return 0, ErrIntOverflowBlsKey
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
				return 0, ErrInvalidLengthBlsKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlsKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlsKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlsKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlsKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlsKey = fmt.Errorf("proto: unexpected end of group")
)
