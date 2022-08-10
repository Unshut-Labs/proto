// Authn protocol

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: authn/v1/authn.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Signature represents a generalized public key signature,
// defined as a union to support cryptographic algorithm agility.
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Union:
	//	*Signature_EcdsaCompact
	Union isSignature_Union `protobuf_oneof:"union"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{0}
}

func (m *Signature) GetUnion() isSignature_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *Signature) GetEcdsaCompact() *Signature_ECDSACompact {
	if x, ok := x.GetUnion().(*Signature_EcdsaCompact); ok {
		return x.EcdsaCompact
	}
	return nil
}

type isSignature_Union interface {
	isSignature_Union()
}

type Signature_EcdsaCompact struct {
	EcdsaCompact *Signature_ECDSACompact `protobuf:"bytes,1,opt,name=ecdsa_compact,json=ecdsaCompact,proto3,oneof"`
}

func (*Signature_EcdsaCompact) isSignature_Union() {}

type Secp256K1Uncompresed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uncompressed point with prefix (0x04) [ P || X || Y ], 65 bytes
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Secp256K1Uncompresed) Reset() {
	*x = Secp256K1Uncompresed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Secp256K1Uncompresed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secp256K1Uncompresed) ProtoMessage() {}

func (x *Secp256K1Uncompresed) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secp256K1Uncompresed.ProtoReflect.Descriptor instead.
func (*Secp256K1Uncompresed) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{1}
}

func (x *Secp256K1Uncompresed) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

// PublicKey represents a generalized public key,
// defined as a union to support cryptographic algorithm agility.
type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64     `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature *Signature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"` // signature is optional in the xmtp-js version
	// Types that are assignable to Union:
	//	*PublicKey_Secp256K1Uncompressed
	Union isPublicKey_Union `protobuf_oneof:"union"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{2}
}

func (x *PublicKey) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PublicKey) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *PublicKey) GetUnion() isPublicKey_Union {
	if m != nil {
		return m.Union
	}
	return nil
}

func (x *PublicKey) GetSecp256K1Uncompressed() *Secp256K1Uncompresed {
	if x, ok := x.GetUnion().(*PublicKey_Secp256K1Uncompressed); ok {
		return x.Secp256K1Uncompressed
	}
	return nil
}

type isPublicKey_Union interface {
	isPublicKey_Union()
}

type PublicKey_Secp256K1Uncompressed struct {
	Secp256K1Uncompressed *Secp256K1Uncompresed `protobuf:"bytes,3,opt,name=secp256k1_uncompressed,json=secp256k1Uncompressed,proto3,oneof"`
}

func (*PublicKey_Secp256K1Uncompressed) isPublicKey_Union() {}

// Token is used by the clients to prove to the node that they are serving a specific wallet.
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdentityKey       *PublicKey `protobuf:"bytes,1,opt,name=identity_key,json=identityKey,proto3" json:"identity_key,omitempty"`                     // identity key signed by a wallet
	AuthDataBytes     []byte     `protobuf:"bytes,2,opt,name=auth_data_bytes,json=authDataBytes,proto3" json:"auth_data_bytes,omitempty"`             // encoded bytes of AuthData
	AuthDataSignature *Signature `protobuf:"bytes,3,opt,name=auth_data_signature,json=authDataSignature,proto3" json:"auth_data_signature,omitempty"` // identity key signature of AuthData bytes
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetIdentityKey() *PublicKey {
	if x != nil {
		return x.IdentityKey
	}
	return nil
}

func (x *Token) GetAuthDataBytes() []byte {
	if x != nil {
		return x.AuthDataBytes
	}
	return nil
}

func (x *Token) GetAuthDataSignature() *Signature {
	if x != nil {
		return x.AuthDataSignature
	}
	return nil
}

// AuthData carries the token parameters that are authenticated by the wallet signature.
// It is embedded in the Token structure as bytes so that the bytes don't need to be
// reconstructed to verify the token signature.
type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddr string `protobuf:"bytes,1,opt,name=wallet_addr,json=walletAddr,proto3" json:"wallet_addr,omitempty"` // address of the wallet
	Timestamp  uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`                    //
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{4}
}

func (x *AuthData) GetWalletAddr() string {
	if x != nil {
		return x.WalletAddr
	}
	return ""
}

func (x *AuthData) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Signature_ECDSACompact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes    []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`        // compact representation [ R || S ], 64 bytes
	Recovery uint32 `protobuf:"varint,2,opt,name=recovery,proto3" json:"recovery,omitempty"` // recovery bit
}

func (x *Signature_ECDSACompact) Reset() {
	*x = Signature_ECDSACompact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authn_v1_authn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature_ECDSACompact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature_ECDSACompact) ProtoMessage() {}

func (x *Signature_ECDSACompact) ProtoReflect() protoreflect.Message {
	mi := &file_authn_v1_authn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature_ECDSACompact.ProtoReflect.Descriptor instead.
func (*Signature_ECDSACompact) Descriptor() ([]byte, []int) {
	return file_authn_v1_authn_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Signature_ECDSACompact) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Signature_ECDSACompact) GetRecovery() uint32 {
	if x != nil {
		return x.Recovery
	}
	return 0
}

var File_authn_v1_authn_proto protoreflect.FileDescriptor

var file_authn_v1_authn_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xa4, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x65, 0x63, 0x64, 0x73, 0x61, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x78, 0x6d, 0x74,
	0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x48, 0x00, 0x52, 0x0c, 0x65, 0x63, 0x64, 0x73, 0x61, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x1a, 0x40, 0x0a, 0x0c, 0x45, 0x43, 0x44, 0x53, 0x41, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x14,
	0x53, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0xc8, 0x01, 0x0a, 0x09, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x5c,
	0x0a, 0x16, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x5f, 0x75, 0x6e, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x65, 0x64, 0x48, 0x00, 0x52, 0x15, 0x73, 0x65, 0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31,
	0x55, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x22, 0xb6, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x3b, 0x0a, 0x0c, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x0b, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x0f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x11, 0x61, 0x75, 0x74,
	0x68, 0x44, 0x61, 0x74, 0x61, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x49,
	0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authn_v1_authn_proto_rawDescOnce sync.Once
	file_authn_v1_authn_proto_rawDescData = file_authn_v1_authn_proto_rawDesc
)

func file_authn_v1_authn_proto_rawDescGZIP() []byte {
	file_authn_v1_authn_proto_rawDescOnce.Do(func() {
		file_authn_v1_authn_proto_rawDescData = protoimpl.X.CompressGZIP(file_authn_v1_authn_proto_rawDescData)
	})
	return file_authn_v1_authn_proto_rawDescData
}

var file_authn_v1_authn_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_authn_v1_authn_proto_goTypes = []interface{}{
	(*Signature)(nil),              // 0: xmtp.authn.v1.Signature
	(*Secp256K1Uncompresed)(nil),   // 1: xmtp.authn.v1.Secp256k1Uncompresed
	(*PublicKey)(nil),              // 2: xmtp.authn.v1.PublicKey
	(*Token)(nil),                  // 3: xmtp.authn.v1.Token
	(*AuthData)(nil),               // 4: xmtp.authn.v1.AuthData
	(*Signature_ECDSACompact)(nil), // 5: xmtp.authn.v1.Signature.ECDSACompact
}
var file_authn_v1_authn_proto_depIdxs = []int32{
	5, // 0: xmtp.authn.v1.Signature.ecdsa_compact:type_name -> xmtp.authn.v1.Signature.ECDSACompact
	0, // 1: xmtp.authn.v1.PublicKey.signature:type_name -> xmtp.authn.v1.Signature
	1, // 2: xmtp.authn.v1.PublicKey.secp256k1_uncompressed:type_name -> xmtp.authn.v1.Secp256k1Uncompresed
	2, // 3: xmtp.authn.v1.Token.identity_key:type_name -> xmtp.authn.v1.PublicKey
	0, // 4: xmtp.authn.v1.Token.auth_data_signature:type_name -> xmtp.authn.v1.Signature
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_authn_v1_authn_proto_init() }
func file_authn_v1_authn_proto_init() {
	if File_authn_v1_authn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authn_v1_authn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
		file_authn_v1_authn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Secp256K1Uncompresed); i {
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
		file_authn_v1_authn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicKey); i {
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
		file_authn_v1_authn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_authn_v1_authn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_authn_v1_authn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature_ECDSACompact); i {
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
	file_authn_v1_authn_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Signature_EcdsaCompact)(nil),
	}
	file_authn_v1_authn_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PublicKey_Secp256K1Uncompressed)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authn_v1_authn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_authn_v1_authn_proto_goTypes,
		DependencyIndexes: file_authn_v1_authn_proto_depIdxs,
		MessageInfos:      file_authn_v1_authn_proto_msgTypes,
	}.Build()
	File_authn_v1_authn_proto = out.File
	file_authn_v1_authn_proto_rawDesc = nil
	file_authn_v1_authn_proto_goTypes = nil
	file_authn_v1_authn_proto_depIdxs = nil
}
