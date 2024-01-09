// Messages used for transport and storage of user conversations.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: message_contents/message.proto

package message_contents

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

// Message header is encoded separately as the bytes are also used
// as associated data for authenticated encryption
type MessageHeaderV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender    *PublicKeyBundle `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient *PublicKeyBundle `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Timestamp uint64           `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *MessageHeaderV1) Reset() {
	*x = MessageHeaderV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageHeaderV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageHeaderV1) ProtoMessage() {}

func (x *MessageHeaderV1) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageHeaderV1.ProtoReflect.Descriptor instead.
func (*MessageHeaderV1) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{0}
}

func (x *MessageHeaderV1) GetSender() *PublicKeyBundle {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *MessageHeaderV1) GetRecipient() *PublicKeyBundle {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *MessageHeaderV1) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// Message is the top level protocol element
type MessageV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderBytes []byte `protobuf:"bytes,1,opt,name=header_bytes,json=headerBytes,proto3" json:"header_bytes,omitempty"` // encapsulates encoded MessageHeaderV1
	// Ciphertext.payload MUST contain encrypted EncodedContent
	Ciphertext *Ciphertext `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *MessageV1) Reset() {
	*x = MessageV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageV1) ProtoMessage() {}

func (x *MessageV1) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageV1.ProtoReflect.Descriptor instead.
func (*MessageV1) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageV1) GetHeaderBytes() []byte {
	if x != nil {
		return x.HeaderBytes
	}
	return nil
}

func (x *MessageV1) GetCiphertext() *Ciphertext {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

// Message header carries information that is not encrypted, and is therefore
// observable by the network. It is however authenticated as associated data
// of the AEAD encryption used to protect the message,
// thus providing tamper evidence.
type MessageHeaderV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender specified message creation time
	CreatedNs uint64 `protobuf:"varint,1,opt,name=created_ns,json=createdNs,proto3" json:"created_ns,omitempty"`
	// the topic the message belongs to
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *MessageHeaderV2) Reset() {
	*x = MessageHeaderV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageHeaderV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageHeaderV2) ProtoMessage() {}

func (x *MessageHeaderV2) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageHeaderV2.ProtoReflect.Descriptor instead.
func (*MessageHeaderV2) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{2}
}

func (x *MessageHeaderV2) GetCreatedNs() uint64 {
	if x != nil {
		return x.CreatedNs
	}
	return 0
}

func (x *MessageHeaderV2) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

// Message combines the encoded header with the encrypted payload.
type MessageV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderBytes []byte `protobuf:"bytes,1,opt,name=header_bytes,json=headerBytes,proto3" json:"header_bytes,omitempty"` // encapsulates encoded MessageHeaderV2
	// Ciphertext.payload MUST contain encrypted SignedContent
	Ciphertext *Ciphertext `protobuf:"bytes,2,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
}

func (x *MessageV2) Reset() {
	*x = MessageV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageV2) ProtoMessage() {}

func (x *MessageV2) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageV2.ProtoReflect.Descriptor instead.
func (*MessageV2) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{3}
}

func (x *MessageV2) GetHeaderBytes() []byte {
	if x != nil {
		return x.HeaderBytes
	}
	return nil
}

func (x *MessageV2) GetCiphertext() *Ciphertext {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

// Versioned Message
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Version:
	//
	//	*Message_V1
	//	*Message_V2
	Version isMessage_Version `protobuf_oneof:"version"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{4}
}

func (m *Message) GetVersion() isMessage_Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (x *Message) GetV1() *MessageV1 {
	if x, ok := x.GetVersion().(*Message_V1); ok {
		return x.V1
	}
	return nil
}

func (x *Message) GetV2() *MessageV2 {
	if x, ok := x.GetVersion().(*Message_V2); ok {
		return x.V2
	}
	return nil
}

type isMessage_Version interface {
	isMessage_Version()
}

type Message_V1 struct {
	V1 *MessageV1 `protobuf:"bytes,1,opt,name=v1,proto3,oneof"`
}

type Message_V2 struct {
	V2 *MessageV2 `protobuf:"bytes,2,opt,name=v2,proto3,oneof"`
}

func (*Message_V1) isMessage_Version() {}

func (*Message_V2) isMessage_Version() {}

// DecodedMessage represents the decrypted message contents.
// DecodedMessage instances are not stored on the network, but
// may be serialized and stored by clients
type DecodedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MessageVersion   string                 `protobuf:"bytes,2,opt,name=message_version,json=messageVersion,proto3" json:"message_version,omitempty"`
	SenderAddress    string                 `protobuf:"bytes,3,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	RecipientAddress *string                `protobuf:"bytes,4,opt,name=recipient_address,json=recipientAddress,proto3,oneof" json:"recipient_address,omitempty"`
	SentNs           uint64                 `protobuf:"varint,5,opt,name=sent_ns,json=sentNs,proto3" json:"sent_ns,omitempty"`
	ContentTopic     string                 `protobuf:"bytes,6,opt,name=content_topic,json=contentTopic,proto3" json:"content_topic,omitempty"`
	Conversation     *ConversationReference `protobuf:"bytes,7,opt,name=conversation,proto3" json:"conversation,omitempty"`
	ContentBytes     []byte                 `protobuf:"bytes,8,opt,name=content_bytes,json=contentBytes,proto3" json:"content_bytes,omitempty"` // encapsulates EncodedContent
}

func (x *DecodedMessage) Reset() {
	*x = DecodedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_contents_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedMessage) ProtoMessage() {}

func (x *DecodedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_contents_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedMessage.ProtoReflect.Descriptor instead.
func (*DecodedMessage) Descriptor() ([]byte, []int) {
	return file_message_contents_message_proto_rawDescGZIP(), []int{5}
}

func (x *DecodedMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DecodedMessage) GetMessageVersion() string {
	if x != nil {
		return x.MessageVersion
	}
	return ""
}

func (x *DecodedMessage) GetSenderAddress() string {
	if x != nil {
		return x.SenderAddress
	}
	return ""
}

func (x *DecodedMessage) GetRecipientAddress() string {
	if x != nil && x.RecipientAddress != nil {
		return *x.RecipientAddress
	}
	return ""
}

func (x *DecodedMessage) GetSentNs() uint64 {
	if x != nil {
		return x.SentNs
	}
	return 0
}

func (x *DecodedMessage) GetContentTopic() string {
	if x != nil {
		return x.ContentTopic
	}
	return ""
}

func (x *DecodedMessage) GetConversation() *ConversationReference {
	if x != nil {
		return x.Conversation
	}
	return nil
}

func (x *DecodedMessage) GetContentBytes() []byte {
	if x != nil {
		return x.ContentBytes
	}
	return nil
}

var File_message_contents_message_proto protoreflect.FileDescriptor

var file_message_contents_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a,
	0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x31,
	0x12, 0x3e, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x44, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x22, 0x71, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56,
	0x31, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x22, 0x46, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x32, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22,
	0x71, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x32, 0x12, 0x21, 0x0a, 0x0c,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x41, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65,
	0x78, 0x74, 0x22, 0x7c, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x02, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x31, 0x48, 0x00, 0x52, 0x02, 0x76,
	0x31, 0x12, 0x32, 0x0a, 0x02, 0x76, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x78, 0x6d, 0x74, 0x70, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x32, 0x48,
	0x00, 0x52, 0x02, 0x76, 0x32, 0x42, 0x09, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xed, 0x02, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x10, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x4e, 0x73, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x50, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x78, 0x6d, 0x74, 0x70,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x4f, 0x0a, 0x1f, 0x6f, 0x72, 0x67, 0x2e, 0x78, 0x6d, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x78, 0x6d, 0x74, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x67, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_contents_message_proto_rawDescOnce sync.Once
	file_message_contents_message_proto_rawDescData = file_message_contents_message_proto_rawDesc
)

func file_message_contents_message_proto_rawDescGZIP() []byte {
	file_message_contents_message_proto_rawDescOnce.Do(func() {
		file_message_contents_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_contents_message_proto_rawDescData)
	})
	return file_message_contents_message_proto_rawDescData
}

var file_message_contents_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_contents_message_proto_goTypes = []interface{}{
	(*MessageHeaderV1)(nil),       // 0: xmtp.message_contents.MessageHeaderV1
	(*MessageV1)(nil),             // 1: xmtp.message_contents.MessageV1
	(*MessageHeaderV2)(nil),       // 2: xmtp.message_contents.MessageHeaderV2
	(*MessageV2)(nil),             // 3: xmtp.message_contents.MessageV2
	(*Message)(nil),               // 4: xmtp.message_contents.Message
	(*DecodedMessage)(nil),        // 5: xmtp.message_contents.DecodedMessage
	(*PublicKeyBundle)(nil),       // 6: xmtp.message_contents.PublicKeyBundle
	(*Ciphertext)(nil),            // 7: xmtp.message_contents.Ciphertext
	(*ConversationReference)(nil), // 8: xmtp.message_contents.ConversationReference
}
var file_message_contents_message_proto_depIdxs = []int32{
	6, // 0: xmtp.message_contents.MessageHeaderV1.sender:type_name -> xmtp.message_contents.PublicKeyBundle
	6, // 1: xmtp.message_contents.MessageHeaderV1.recipient:type_name -> xmtp.message_contents.PublicKeyBundle
	7, // 2: xmtp.message_contents.MessageV1.ciphertext:type_name -> xmtp.message_contents.Ciphertext
	7, // 3: xmtp.message_contents.MessageV2.ciphertext:type_name -> xmtp.message_contents.Ciphertext
	1, // 4: xmtp.message_contents.Message.v1:type_name -> xmtp.message_contents.MessageV1
	3, // 5: xmtp.message_contents.Message.v2:type_name -> xmtp.message_contents.MessageV2
	8, // 6: xmtp.message_contents.DecodedMessage.conversation:type_name -> xmtp.message_contents.ConversationReference
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_message_contents_message_proto_init() }
func file_message_contents_message_proto_init() {
	if File_message_contents_message_proto != nil {
		return
	}
	file_message_contents_ciphertext_proto_init()
	file_message_contents_conversation_reference_proto_init()
	file_message_contents_public_key_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_contents_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageHeaderV1); i {
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
		file_message_contents_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageV1); i {
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
		file_message_contents_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageHeaderV2); i {
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
		file_message_contents_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageV2); i {
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
		file_message_contents_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_message_contents_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedMessage); i {
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
	file_message_contents_message_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Message_V1)(nil),
		(*Message_V2)(nil),
	}
	file_message_contents_message_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_contents_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_contents_message_proto_goTypes,
		DependencyIndexes: file_message_contents_message_proto_depIdxs,
		MessageInfos:      file_message_contents_message_proto_msgTypes,
	}.Build()
	File_message_contents_message_proto = out.File
	file_message_contents_message_proto_rawDesc = nil
	file_message_contents_message_proto_goTypes = nil
	file_message_contents_message_proto_depIdxs = nil
}
