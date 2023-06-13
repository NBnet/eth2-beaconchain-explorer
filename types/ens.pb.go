// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: types/ens.proto

package types

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

// Eth1Block is stored in the blocks table under <chainID>:<reversePaddedNumber>
type EnsNameRegistered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash       []byte                 `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockNumber      uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	RegisterContract []byte                 `protobuf:"bytes,3,opt,name=register_contract,json=registerContract,proto3" json:"register_contract,omitempty"`
	ResolverContract []byte                 `protobuf:"bytes,4,opt,name=resolver_contract,json=resolverContract,proto3" json:"resolver_contract,omitempty"`
	Time             *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=time,proto3" json:"time,omitempty"`
	Label            []byte                 `protobuf:"bytes,6,opt,name=label,proto3" json:"label,omitempty"`
	Owner            []byte                 `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Resolver         []byte                 `protobuf:"bytes,8,opt,name=resolver,proto3" json:"resolver,omitempty"`
	Node             []byte                 `protobuf:"bytes,9,opt,name=node,proto3" json:"node,omitempty"`
	Name             []byte                 `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Expires          *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *EnsNameRegistered) Reset() {
	*x = EnsNameRegistered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_ens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsNameRegistered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsNameRegistered) ProtoMessage() {}

func (x *EnsNameRegistered) ProtoReflect() protoreflect.Message {
	mi := &file_types_ens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsNameRegistered.ProtoReflect.Descriptor instead.
func (*EnsNameRegistered) Descriptor() ([]byte, []int) {
	return file_types_ens_proto_rawDescGZIP(), []int{0}
}

func (x *EnsNameRegistered) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *EnsNameRegistered) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EnsNameRegistered) GetRegisterContract() []byte {
	if x != nil {
		return x.RegisterContract
	}
	return nil
}

func (x *EnsNameRegistered) GetResolverContract() []byte {
	if x != nil {
		return x.ResolverContract
	}
	return nil
}

func (x *EnsNameRegistered) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EnsNameRegistered) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *EnsNameRegistered) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *EnsNameRegistered) GetResolver() []byte {
	if x != nil {
		return x.Resolver
	}
	return nil
}

func (x *EnsNameRegistered) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *EnsNameRegistered) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *EnsNameRegistered) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type EnsNameRenewed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash  []byte                 `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockNumber uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	Time        *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Label       []byte                 `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Name        []byte                 `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Expires     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expires,proto3" json:"expires,omitempty"`
}

func (x *EnsNameRenewed) Reset() {
	*x = EnsNameRenewed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_ens_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsNameRenewed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsNameRenewed) ProtoMessage() {}

func (x *EnsNameRenewed) ProtoReflect() protoreflect.Message {
	mi := &file_types_ens_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsNameRenewed.ProtoReflect.Descriptor instead.
func (*EnsNameRenewed) Descriptor() ([]byte, []int) {
	return file_types_ens_proto_rawDescGZIP(), []int{1}
}

func (x *EnsNameRenewed) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *EnsNameRenewed) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EnsNameRenewed) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EnsNameRenewed) GetLabel() []byte {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *EnsNameRenewed) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *EnsNameRenewed) GetExpires() *timestamppb.Timestamp {
	if x != nil {
		return x.Expires
	}
	return nil
}

type EnsAddressChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash      []byte                 `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	ResolveContract []byte                 `protobuf:"bytes,3,opt,name=resolve_contract,json=resolveContract,proto3" json:"resolve_contract,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Node            []byte                 `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	CoinType        uint64                 `protobuf:"varint,6,opt,name=coin_type,json=coinType,proto3" json:"coin_type,omitempty"`
	NewAddress      []byte                 `protobuf:"bytes,7,opt,name=new_address,json=newAddress,proto3" json:"new_address,omitempty"`
}

func (x *EnsAddressChanged) Reset() {
	*x = EnsAddressChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_ens_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsAddressChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsAddressChanged) ProtoMessage() {}

func (x *EnsAddressChanged) ProtoReflect() protoreflect.Message {
	mi := &file_types_ens_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsAddressChanged.ProtoReflect.Descriptor instead.
func (*EnsAddressChanged) Descriptor() ([]byte, []int) {
	return file_types_ens_proto_rawDescGZIP(), []int{2}
}

func (x *EnsAddressChanged) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *EnsAddressChanged) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EnsAddressChanged) GetResolveContract() []byte {
	if x != nil {
		return x.ResolveContract
	}
	return nil
}

func (x *EnsAddressChanged) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EnsAddressChanged) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *EnsAddressChanged) GetCoinType() uint64 {
	if x != nil {
		return x.CoinType
	}
	return 0
}

func (x *EnsAddressChanged) GetNewAddress() []byte {
	if x != nil {
		return x.NewAddress
	}
	return nil
}

type EnsNameChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash      []byte                 `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	BlockNumber     uint64                 `protobuf:"varint,2,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
	ResolveContract []byte                 `protobuf:"bytes,3,opt,name=resolve_contract,json=resolveContract,proto3" json:"resolve_contract,omitempty"`
	Time            *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	Node            []byte                 `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	NewName         []byte                 `protobuf:"bytes,6,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *EnsNameChanged) Reset() {
	*x = EnsNameChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_ens_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnsNameChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsNameChanged) ProtoMessage() {}

func (x *EnsNameChanged) ProtoReflect() protoreflect.Message {
	mi := &file_types_ens_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsNameChanged.ProtoReflect.Descriptor instead.
func (*EnsNameChanged) Descriptor() ([]byte, []int) {
	return file_types_ens_proto_rawDescGZIP(), []int{3}
}

func (x *EnsNameChanged) GetParentHash() []byte {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *EnsNameChanged) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EnsNameChanged) GetResolveContract() []byte {
	if x != nil {
		return x.ResolveContract
	}
	return nil
}

func (x *EnsNameChanged) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EnsNameChanged) GetNode() []byte {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *EnsNameChanged) GetNewName() []byte {
	if x != nil {
		return x.NewName
	}
	return nil
}

var File_types_ens_proto protoreflect.FileDescriptor

var file_types_ens_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x11, 0x45, 0x6e,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x6e, 0x65, 0x77, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x11, 0x45,
	0x6e, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x6e,
	0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63, 0x6f, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0xde, 0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_ens_proto_rawDescOnce sync.Once
	file_types_ens_proto_rawDescData = file_types_ens_proto_rawDesc
)

func file_types_ens_proto_rawDescGZIP() []byte {
	file_types_ens_proto_rawDescOnce.Do(func() {
		file_types_ens_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_ens_proto_rawDescData)
	})
	return file_types_ens_proto_rawDescData
}

var file_types_ens_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_ens_proto_goTypes = []interface{}{
	(*EnsNameRegistered)(nil),     // 0: types.EnsNameRegistered
	(*EnsNameRenewed)(nil),        // 1: types.EnsNameRenewed
	(*EnsAddressChanged)(nil),     // 2: types.EnsAddressChanged
	(*EnsNameChanged)(nil),        // 3: types.EnsNameChanged
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_types_ens_proto_depIdxs = []int32{
	4, // 0: types.EnsNameRegistered.time:type_name -> google.protobuf.Timestamp
	4, // 1: types.EnsNameRegistered.expires:type_name -> google.protobuf.Timestamp
	4, // 2: types.EnsNameRenewed.time:type_name -> google.protobuf.Timestamp
	4, // 3: types.EnsNameRenewed.expires:type_name -> google.protobuf.Timestamp
	4, // 4: types.EnsAddressChanged.time:type_name -> google.protobuf.Timestamp
	4, // 5: types.EnsNameChanged.time:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_types_ens_proto_init() }
func file_types_ens_proto_init() {
	if File_types_ens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_ens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsNameRegistered); i {
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
		file_types_ens_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsNameRenewed); i {
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
		file_types_ens_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsAddressChanged); i {
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
		file_types_ens_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnsNameChanged); i {
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
			RawDescriptor: file_types_ens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_ens_proto_goTypes,
		DependencyIndexes: file_types_ens_proto_depIdxs,
		MessageInfos:      file_types_ens_proto_msgTypes,
	}.Build()
	File_types_ens_proto = out.File
	file_types_ens_proto_rawDesc = nil
	file_types_ens_proto_goTypes = nil
	file_types_ens_proto_depIdxs = nil
}
