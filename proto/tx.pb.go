// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: proto/tx.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SolaceTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace  string   `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WalletAddr string   `protobuf:"bytes,2,opt,name=walletAddr,proto3" json:"walletAddr,omitempty"`
	SenderAddr string   `protobuf:"bytes,3,opt,name=senderAddr,proto3" json:"senderAddr,omitempty"`
	ToAddr     string   `protobuf:"bytes,4,opt,name=toAddr,proto3" json:"toAddr,omitempty"`
	TokenAddr  string   `protobuf:"bytes,5,opt,name=tokenAddr,proto3" json:"tokenAddr,omitempty"`
	Value      int32    `protobuf:"varint,6,opt,name=value,proto3" json:"value,omitempty"`
	Signatures []string `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty"`
}

func (x *SolaceTx) Reset() {
	*x = SolaceTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolaceTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolaceTx) ProtoMessage() {}

func (x *SolaceTx) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolaceTx.ProtoReflect.Descriptor instead.
func (*SolaceTx) Descriptor() ([]byte, []int) {
	return file_proto_tx_proto_rawDescGZIP(), []int{0}
}

func (x *SolaceTx) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *SolaceTx) GetWalletAddr() string {
	if x != nil {
		return x.WalletAddr
	}
	return ""
}

func (x *SolaceTx) GetSenderAddr() string {
	if x != nil {
		return x.SenderAddr
	}
	return ""
}

func (x *SolaceTx) GetToAddr() string {
	if x != nil {
		return x.ToAddr
	}
	return ""
}

func (x *SolaceTx) GetTokenAddr() string {
	if x != nil {
		return x.TokenAddr
	}
	return ""
}

func (x *SolaceTx) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *SolaceTx) GetSignatures() []string {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature string                 `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Tx        *SolaceTx              `protobuf:"bytes,4,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_tx_proto_msgTypes[1]
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
	return file_proto_tx_proto_rawDescGZIP(), []int{1}
}

func (x *Signature) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Signature) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Signature) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *Signature) GetTx() *SolaceTx {
	if x != nil {
		return x.Tx
	}
	return nil
}

var File_proto_tx_proto protoreflect.FileDescriptor

var file_proto_tx_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x08, 0x53, 0x6f, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x78, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19,
	0x0a, 0x02, 0x74, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x6f, 0x6c,
	0x61, 0x63, 0x65, 0x54, 0x78, 0x52, 0x02, 0x74, 0x78, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_tx_proto_rawDescOnce sync.Once
	file_proto_tx_proto_rawDescData = file_proto_tx_proto_rawDesc
)

func file_proto_tx_proto_rawDescGZIP() []byte {
	file_proto_tx_proto_rawDescOnce.Do(func() {
		file_proto_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_tx_proto_rawDescData)
	})
	return file_proto_tx_proto_rawDescData
}

var file_proto_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_tx_proto_goTypes = []interface{}{
	(*SolaceTx)(nil),              // 0: SolaceTx
	(*Signature)(nil),             // 1: Signature
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_tx_proto_depIdxs = []int32{
	2, // 0: Signature.timestamp:type_name -> google.protobuf.Timestamp
	0, // 1: Signature.tx:type_name -> SolaceTx
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_tx_proto_init() }
func file_proto_tx_proto_init() {
	if File_proto_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolaceTx); i {
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
		file_proto_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_tx_proto_goTypes,
		DependencyIndexes: file_proto_tx_proto_depIdxs,
		MessageInfos:      file_proto_tx_proto_msgTypes,
	}.Build()
	File_proto_tx_proto = out.File
	file_proto_tx_proto_rawDesc = nil
	file_proto_tx_proto_goTypes = nil
	file_proto_tx_proto_depIdxs = nil
}
