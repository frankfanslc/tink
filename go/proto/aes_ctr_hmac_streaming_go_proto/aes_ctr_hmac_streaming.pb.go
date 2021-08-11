// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: third_party/tink/proto/aes_ctr_hmac_streaming.proto

package aes_ctr_hmac_streaming_go_proto

import (
	common_go_proto "github.com/google/tink/go/proto/common_go_proto"
	hmac_go_proto "github.com/google/tink/go/proto/hmac_go_proto"
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

type AesCtrHmacStreamingParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CiphertextSegmentSize uint32                    `protobuf:"varint,1,opt,name=ciphertext_segment_size,json=ciphertextSegmentSize,proto3" json:"ciphertext_segment_size,omitempty"`
	DerivedKeySize        uint32                    `protobuf:"varint,2,opt,name=derived_key_size,json=derivedKeySize,proto3" json:"derived_key_size,omitempty"`                            // size of AES-CTR keys derived for each segment
	HkdfHashType          common_go_proto.HashType  `protobuf:"varint,3,opt,name=hkdf_hash_type,json=hkdfHashType,proto3,enum=google.crypto.tink.HashType" json:"hkdf_hash_type,omitempty"` // hash function for key derivation via HKDF
	HmacParams            *hmac_go_proto.HmacParams `protobuf:"bytes,4,opt,name=hmac_params,json=hmacParams,proto3" json:"hmac_params,omitempty"`                                           // params for authentication tags
}

func (x *AesCtrHmacStreamingParams) Reset() {
	*x = AesCtrHmacStreamingParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrHmacStreamingParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrHmacStreamingParams) ProtoMessage() {}

func (x *AesCtrHmacStreamingParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrHmacStreamingParams.ProtoReflect.Descriptor instead.
func (*AesCtrHmacStreamingParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *AesCtrHmacStreamingParams) GetCiphertextSegmentSize() uint32 {
	if x != nil {
		return x.CiphertextSegmentSize
	}
	return 0
}

func (x *AesCtrHmacStreamingParams) GetDerivedKeySize() uint32 {
	if x != nil {
		return x.DerivedKeySize
	}
	return 0
}

func (x *AesCtrHmacStreamingParams) GetHkdfHashType() common_go_proto.HashType {
	if x != nil {
		return x.HkdfHashType
	}
	return common_go_proto.HashType(0)
}

func (x *AesCtrHmacStreamingParams) GetHmacParams() *hmac_go_proto.HmacParams {
	if x != nil {
		return x.HmacParams
	}
	return nil
}

type AesCtrHmacStreamingKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32                     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Params  *AesCtrHmacStreamingParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize uint32                     `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"` // size of the main key (aka. "ikm", input key material)
}

func (x *AesCtrHmacStreamingKeyFormat) Reset() {
	*x = AesCtrHmacStreamingKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrHmacStreamingKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrHmacStreamingKeyFormat) ProtoMessage() {}

func (x *AesCtrHmacStreamingKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrHmacStreamingKeyFormat.ProtoReflect.Descriptor instead.
func (*AesCtrHmacStreamingKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *AesCtrHmacStreamingKeyFormat) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AesCtrHmacStreamingKeyFormat) GetParams() *AesCtrHmacStreamingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AesCtrHmacStreamingKeyFormat) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.AesCtrHmacStreamingKey
type AesCtrHmacStreamingKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint32                     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params   *AesCtrHmacStreamingParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue []byte                     `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"` // the main key, aka. "ikm", input key material
}

func (x *AesCtrHmacStreamingKey) Reset() {
	*x = AesCtrHmacStreamingKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrHmacStreamingKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrHmacStreamingKey) ProtoMessage() {}

func (x *AesCtrHmacStreamingKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrHmacStreamingKey.ProtoReflect.Descriptor instead.
func (*AesCtrHmacStreamingKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescGZIP(), []int{2}
}

func (x *AesCtrHmacStreamingKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AesCtrHmacStreamingKey) GetParams() *AesCtrHmacStreamingParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AesCtrHmacStreamingKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

var File_third_party_tink_proto_aes_ctr_hmac_streaming_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDesc = []byte{
	0x0a, 0x33, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73, 0x5f, 0x63, 0x74, 0x72,
	0x5f, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x23, 0x74, 0x68, 0x69, 0x72, 0x64,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6d, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x02, 0x0a, 0x19, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x48, 0x6d, 0x61, 0x63,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x36, 0x0a, 0x17, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x15, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x72, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x42, 0x0a, 0x0e, 0x68, 0x6b, 0x64, 0x66, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x68, 0x6b, 0x64, 0x66, 0x48, 0x61, 0x73,
	0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e,
	0x48, 0x6d, 0x61, 0x63, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x68, 0x6d, 0x61, 0x63,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x9a, 0x01, 0x0a, 0x1c, 0x41, 0x65, 0x73, 0x43, 0x74,
	0x72, 0x48, 0x6d, 0x61, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4b, 0x65,
	0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x45, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x48, 0x6d, 0x61,
	0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x16, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x48, 0x6d,
	0x61, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x65,
	0x73, 0x43, 0x74, 0x72, 0x48, 0x6d, 0x61, 0x63, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e,
	0x67, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x5e, 0x0a, 0x1c,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73,
	0x5f, 0x63, 0x74, 0x72, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x69, 0x6e, 0x67, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescData = file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDesc
)

func file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescData)
	})
	return file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDescData
}

var file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_goTypes = []interface{}{
	(*AesCtrHmacStreamingParams)(nil),    // 0: google.crypto.tink.AesCtrHmacStreamingParams
	(*AesCtrHmacStreamingKeyFormat)(nil), // 1: google.crypto.tink.AesCtrHmacStreamingKeyFormat
	(*AesCtrHmacStreamingKey)(nil),       // 2: google.crypto.tink.AesCtrHmacStreamingKey
	(common_go_proto.HashType)(0),        // 3: google.crypto.tink.HashType
	(*hmac_go_proto.HmacParams)(nil),     // 4: google.crypto.tink.HmacParams
}
var file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_depIdxs = []int32{
	3, // 0: google.crypto.tink.AesCtrHmacStreamingParams.hkdf_hash_type:type_name -> google.crypto.tink.HashType
	4, // 1: google.crypto.tink.AesCtrHmacStreamingParams.hmac_params:type_name -> google.crypto.tink.HmacParams
	0, // 2: google.crypto.tink.AesCtrHmacStreamingKeyFormat.params:type_name -> google.crypto.tink.AesCtrHmacStreamingParams
	0, // 3: google.crypto.tink.AesCtrHmacStreamingKey.params:type_name -> google.crypto.tink.AesCtrHmacStreamingParams
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_init() }
func file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_init() {
	if File_third_party_tink_proto_aes_ctr_hmac_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrHmacStreamingParams); i {
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
		file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrHmacStreamingKeyFormat); i {
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
		file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrHmacStreamingKey); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_aes_ctr_hmac_streaming_proto = out.File
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_rawDesc = nil
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_goTypes = nil
	file_third_party_tink_proto_aes_ctr_hmac_streaming_proto_depIdxs = nil
}
