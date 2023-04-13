// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: github.com/dipdup-io/starknet-indexer/pkg/grpc/proto/indexer.proto

package pb

import (
	pb "github.com/dipdup-net/indexer-sdk/pkg/modules/grpc/pb"
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

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head           bool                  `protobuf:"varint,1,opt,name=head,proto3" json:"head,omitempty"`
	Invokes        *InvokeFilters        `protobuf:"bytes,2,opt,name=invokes,proto3" json:"invokes,omitempty"`
	Declares       *DeclareFilters       `protobuf:"bytes,3,opt,name=declares,proto3" json:"declares,omitempty"`
	Deploys        *DeployFilters        `protobuf:"bytes,4,opt,name=deploys,proto3" json:"deploys,omitempty"`
	DeployAccounts *DeployAccountFilters `protobuf:"bytes,5,opt,name=deploy_accounts,json=deployAccounts,proto3" json:"deploy_accounts,omitempty"`
	L1Handlers     *L1HandlerFilter      `protobuf:"bytes,6,opt,name=l1_handlers,json=l1Handlers,proto3" json:"l1_handlers,omitempty"`
	Internals      *InternalFilter       `protobuf:"bytes,7,opt,name=internals,proto3" json:"internals,omitempty"`
	Fees           *FeeFilter            `protobuf:"bytes,8,opt,name=fees,proto3" json:"fees,omitempty"`
	Events         *EventFilter          `protobuf:"bytes,9,opt,name=events,proto3" json:"events,omitempty"`
	Msgs           *MessageFilter        `protobuf:"bytes,10,opt,name=msgs,proto3" json:"msgs,omitempty"`
	Transfers      *TransferFilter       `protobuf:"bytes,11,opt,name=transfers,proto3" json:"transfers,omitempty"`
	StorageDiffs   *StorageDiffFilter    `protobuf:"bytes,12,opt,name=storage_diffs,json=storageDiffs,proto3" json:"storage_diffs,omitempty"`
	TokenBalances  *TokenBalanceFilter   `protobuf:"bytes,13,opt,name=token_balances,json=tokenBalances,proto3" json:"token_balances,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetHead() bool {
	if x != nil {
		return x.Head
	}
	return false
}

func (x *SubscribeRequest) GetInvokes() *InvokeFilters {
	if x != nil {
		return x.Invokes
	}
	return nil
}

func (x *SubscribeRequest) GetDeclares() *DeclareFilters {
	if x != nil {
		return x.Declares
	}
	return nil
}

func (x *SubscribeRequest) GetDeploys() *DeployFilters {
	if x != nil {
		return x.Deploys
	}
	return nil
}

func (x *SubscribeRequest) GetDeployAccounts() *DeployAccountFilters {
	if x != nil {
		return x.DeployAccounts
	}
	return nil
}

func (x *SubscribeRequest) GetL1Handlers() *L1HandlerFilter {
	if x != nil {
		return x.L1Handlers
	}
	return nil
}

func (x *SubscribeRequest) GetInternals() *InternalFilter {
	if x != nil {
		return x.Internals
	}
	return nil
}

func (x *SubscribeRequest) GetFees() *FeeFilter {
	if x != nil {
		return x.Fees
	}
	return nil
}

func (x *SubscribeRequest) GetEvents() *EventFilter {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *SubscribeRequest) GetMsgs() *MessageFilter {
	if x != nil {
		return x.Msgs
	}
	return nil
}

func (x *SubscribeRequest) GetTransfers() *TransferFilter {
	if x != nil {
		return x.Transfers
	}
	return nil
}

func (x *SubscribeRequest) GetStorageDiffs() *StorageDiffFilter {
	if x != nil {
		return x.StorageDiffs
	}
	return nil
}

func (x *SubscribeRequest) GetTokenBalances() *TokenBalanceFilter {
	if x != nil {
		return x.TokenBalances
	}
	return nil
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response      *pb.SubscribeResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	Block         *Block                `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Declare       *Declare              `protobuf:"bytes,3,opt,name=declare,proto3" json:"declare,omitempty"`
	Deploy        *Deploy               `protobuf:"bytes,4,opt,name=deploy,proto3" json:"deploy,omitempty"`
	DeployAccount *DeployAccount        `protobuf:"bytes,5,opt,name=deploy_account,json=deployAccount,proto3" json:"deploy_account,omitempty"`
	Event         *Event                `protobuf:"bytes,6,opt,name=event,proto3" json:"event,omitempty"`
	Fee           *Fee                  `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
	Internal      *Internal             `protobuf:"bytes,8,opt,name=internal,proto3" json:"internal,omitempty"`
	Invoke        *Invoke               `protobuf:"bytes,9,opt,name=invoke,proto3" json:"invoke,omitempty"`
	L1Handler     *L1Handler            `protobuf:"bytes,10,opt,name=l1_handler,json=l1Handler,proto3" json:"l1_handler,omitempty"`
	Message       *StarknetMessage      `protobuf:"bytes,11,opt,name=message,proto3" json:"message,omitempty"`
	StorageDiff   *StorageDiff          `protobuf:"bytes,12,opt,name=storage_diff,json=storageDiff,proto3" json:"storage_diff,omitempty"`
	TokenBalance  *TokenBalance         `protobuf:"bytes,13,opt,name=token_balance,json=tokenBalance,proto3" json:"token_balance,omitempty"`
	Transfer      *Transfer             `protobuf:"bytes,14,opt,name=transfer,proto3" json:"transfer,omitempty"`
	EndOfBlock    bool                  `protobuf:"varint,15,opt,name=end_of_block,json=endOfBlock,proto3" json:"end_of_block,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *Subscription) GetResponse() *pb.SubscribeResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Subscription) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Subscription) GetDeclare() *Declare {
	if x != nil {
		return x.Declare
	}
	return nil
}

func (x *Subscription) GetDeploy() *Deploy {
	if x != nil {
		return x.Deploy
	}
	return nil
}

func (x *Subscription) GetDeployAccount() *DeployAccount {
	if x != nil {
		return x.DeployAccount
	}
	return nil
}

func (x *Subscription) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *Subscription) GetFee() *Fee {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *Subscription) GetInternal() *Internal {
	if x != nil {
		return x.Internal
	}
	return nil
}

func (x *Subscription) GetInvoke() *Invoke {
	if x != nil {
		return x.Invoke
	}
	return nil
}

func (x *Subscription) GetL1Handler() *L1Handler {
	if x != nil {
		return x.L1Handler
	}
	return nil
}

func (x *Subscription) GetMessage() *StarknetMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *Subscription) GetStorageDiff() *StorageDiff {
	if x != nil {
		return x.StorageDiff
	}
	return nil
}

func (x *Subscription) GetTokenBalance() *TokenBalance {
	if x != nil {
		return x.TokenBalance
	}
	return nil
}

func (x *Subscription) GetTransfer() *Transfer {
	if x != nil {
		return x.Transfer
	}
	return nil
}

func (x *Subscription) GetEndOfBlock() bool {
	if x != nil {
		return x.EndOfBlock
	}
	return false
}

type Bytes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bytes.ProtoReflect.Descriptor instead.
func (*Bytes) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *Bytes) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type JsonSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions    []*JsonSchemaItem `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty"`
	L1Handlers   []*JsonSchemaItem `protobuf:"bytes,2,rep,name=l1_handlers,json=l1Handlers,proto3" json:"l1_handlers,omitempty"`
	Constructors []*JsonSchemaItem `protobuf:"bytes,3,rep,name=constructors,proto3" json:"constructors,omitempty"`
	Events       []*JsonSchemaItem `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
	Structs      []*JsonSchemaItem `protobuf:"bytes,5,rep,name=structs,proto3" json:"structs,omitempty"`
}

func (x *JsonSchema) Reset() {
	*x = JsonSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonSchema) ProtoMessage() {}

func (x *JsonSchema) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonSchema.ProtoReflect.Descriptor instead.
func (*JsonSchema) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{3}
}

func (x *JsonSchema) GetFunctions() []*JsonSchemaItem {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *JsonSchema) GetL1Handlers() []*JsonSchemaItem {
	if x != nil {
		return x.L1Handlers
	}
	return nil
}

func (x *JsonSchema) GetConstructors() []*JsonSchemaItem {
	if x != nil {
		return x.Constructors
	}
	return nil
}

func (x *JsonSchema) GetEvents() []*JsonSchemaItem {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *JsonSchema) GetStructs() []*JsonSchemaItem {
	if x != nil {
		return x.Structs
	}
	return nil
}

type JsonSchemaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schema []byte `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *JsonSchemaItem) Reset() {
	*x = JsonSchemaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonSchemaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonSchemaItem) ProtoMessage() {}

func (x *JsonSchemaItem) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonSchemaItem.ProtoReflect.Descriptor instead.
func (*JsonSchemaItem) Descriptor() ([]byte, []int) {
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP(), []int{4}
}

func (x *JsonSchemaItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *JsonSchemaItem) GetSchema() []byte {
	if x != nil {
		return x.Schema
	}
	return nil
}

var File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto protoreflect.FileDescriptor

var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc = []byte{
	0x0a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70,
	0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x46, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x6e,
	0x65, 0x74, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e,
	0x65, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75,
	0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x05, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x68, 0x65, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x07,
	0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x08,
	0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x08, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x07, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x73, 0x12,
	0x44, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x6c, 0x31, 0x5f, 0x68, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x31, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x0a, 0x6c, 0x31, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x12, 0x33,
	0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x04, 0x66, 0x65, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x52, 0x04, 0x66, 0x65, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6d, 0x73, 0x67, 0x73, 0x12,
	0x33, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x64, 0x69, 0x66, 0x66, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69, 0x66, 0x66, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69,
	0x66, 0x66, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xaf, 0x05, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x28, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72,
	0x65, 0x52, 0x07, 0x64, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x52, 0x06, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x12, 0x3b, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x52, 0x03, 0x66, 0x65, 0x65,
	0x12, 0x2b, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x25, 0x0a,
	0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x06, 0x69, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x31, 0x5f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x31, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x09, 0x6c, 0x31, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69, 0x66,
	0x66, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x69, 0x66, 0x66, 0x12, 0x38,
	0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x08, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x65, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x64,
	0x4f, 0x66, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x1b, 0x0a, 0x05, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x94, 0x02, 0x0a, 0x0a, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a,
	0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x6c, 0x31, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x6c, 0x31, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x73,
	0x12, 0x39, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a,
	0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0c, 0x63,
	0x6f, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x07, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x73, 0x22, 0x3c, 0x0a, 0x0e, 0x4a,
	0x73, 0x6f, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x32, 0xfa, 0x01, 0x0a, 0x0e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x0b, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x30, 0x0a, 0x12, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x6f, 0x72,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x15, 0x4a, 0x53, 0x4f, 0x4e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46,
	0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x70, 0x64, 0x75, 0x70, 0x2d, 0x69, 0x6f, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescOnce sync.Once
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData = file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc
)

func file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescGZIP() []byte {
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescOnce.Do(func() {
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData)
	})
	return file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDescData
}

var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes = []interface{}{
	(*SubscribeRequest)(nil),       // 0: proto.SubscribeRequest
	(*Subscription)(nil),           // 1: proto.Subscription
	(*Bytes)(nil),                  // 2: proto.Bytes
	(*JsonSchema)(nil),             // 3: proto.JsonSchema
	(*JsonSchemaItem)(nil),         // 4: proto.JsonSchemaItem
	(*InvokeFilters)(nil),          // 5: proto.InvokeFilters
	(*DeclareFilters)(nil),         // 6: proto.DeclareFilters
	(*DeployFilters)(nil),          // 7: proto.DeployFilters
	(*DeployAccountFilters)(nil),   // 8: proto.DeployAccountFilters
	(*L1HandlerFilter)(nil),        // 9: proto.L1HandlerFilter
	(*InternalFilter)(nil),         // 10: proto.InternalFilter
	(*FeeFilter)(nil),              // 11: proto.FeeFilter
	(*EventFilter)(nil),            // 12: proto.EventFilter
	(*MessageFilter)(nil),          // 13: proto.MessageFilter
	(*TransferFilter)(nil),         // 14: proto.TransferFilter
	(*StorageDiffFilter)(nil),      // 15: proto.StorageDiffFilter
	(*TokenBalanceFilter)(nil),     // 16: proto.TokenBalanceFilter
	(*pb.SubscribeResponse)(nil),   // 17: proto.SubscribeResponse
	(*Block)(nil),                  // 18: proto.Block
	(*Declare)(nil),                // 19: proto.Declare
	(*Deploy)(nil),                 // 20: proto.Deploy
	(*DeployAccount)(nil),          // 21: proto.DeployAccount
	(*Event)(nil),                  // 22: proto.Event
	(*Fee)(nil),                    // 23: proto.Fee
	(*Internal)(nil),               // 24: proto.Internal
	(*Invoke)(nil),                 // 25: proto.Invoke
	(*L1Handler)(nil),              // 26: proto.L1Handler
	(*StarknetMessage)(nil),        // 27: proto.StarknetMessage
	(*StorageDiff)(nil),            // 28: proto.StorageDiff
	(*TokenBalance)(nil),           // 29: proto.TokenBalance
	(*Transfer)(nil),               // 30: proto.Transfer
	(*pb.UnsubscribeRequest)(nil),  // 31: proto.UnsubscribeRequest
	(*pb.UnsubscribeResponse)(nil), // 32: proto.UnsubscribeResponse
}
var file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs = []int32{
	5,  // 0: proto.SubscribeRequest.invokes:type_name -> proto.InvokeFilters
	6,  // 1: proto.SubscribeRequest.declares:type_name -> proto.DeclareFilters
	7,  // 2: proto.SubscribeRequest.deploys:type_name -> proto.DeployFilters
	8,  // 3: proto.SubscribeRequest.deploy_accounts:type_name -> proto.DeployAccountFilters
	9,  // 4: proto.SubscribeRequest.l1_handlers:type_name -> proto.L1HandlerFilter
	10, // 5: proto.SubscribeRequest.internals:type_name -> proto.InternalFilter
	11, // 6: proto.SubscribeRequest.fees:type_name -> proto.FeeFilter
	12, // 7: proto.SubscribeRequest.events:type_name -> proto.EventFilter
	13, // 8: proto.SubscribeRequest.msgs:type_name -> proto.MessageFilter
	14, // 9: proto.SubscribeRequest.transfers:type_name -> proto.TransferFilter
	15, // 10: proto.SubscribeRequest.storage_diffs:type_name -> proto.StorageDiffFilter
	16, // 11: proto.SubscribeRequest.token_balances:type_name -> proto.TokenBalanceFilter
	17, // 12: proto.Subscription.response:type_name -> proto.SubscribeResponse
	18, // 13: proto.Subscription.block:type_name -> proto.Block
	19, // 14: proto.Subscription.declare:type_name -> proto.Declare
	20, // 15: proto.Subscription.deploy:type_name -> proto.Deploy
	21, // 16: proto.Subscription.deploy_account:type_name -> proto.DeployAccount
	22, // 17: proto.Subscription.event:type_name -> proto.Event
	23, // 18: proto.Subscription.fee:type_name -> proto.Fee
	24, // 19: proto.Subscription.internal:type_name -> proto.Internal
	25, // 20: proto.Subscription.invoke:type_name -> proto.Invoke
	26, // 21: proto.Subscription.l1_handler:type_name -> proto.L1Handler
	27, // 22: proto.Subscription.message:type_name -> proto.StarknetMessage
	28, // 23: proto.Subscription.storage_diff:type_name -> proto.StorageDiff
	29, // 24: proto.Subscription.token_balance:type_name -> proto.TokenBalance
	30, // 25: proto.Subscription.transfer:type_name -> proto.Transfer
	4,  // 26: proto.JsonSchema.functions:type_name -> proto.JsonSchemaItem
	4,  // 27: proto.JsonSchema.l1_handlers:type_name -> proto.JsonSchemaItem
	4,  // 28: proto.JsonSchema.constructors:type_name -> proto.JsonSchemaItem
	4,  // 29: proto.JsonSchema.events:type_name -> proto.JsonSchemaItem
	4,  // 30: proto.JsonSchema.structs:type_name -> proto.JsonSchemaItem
	0,  // 31: proto.IndexerService.Subscribe:input_type -> proto.SubscribeRequest
	31, // 32: proto.IndexerService.Unsubscribe:input_type -> proto.UnsubscribeRequest
	2,  // 33: proto.IndexerService.JSONSchemaForClass:input_type -> proto.Bytes
	2,  // 34: proto.IndexerService.JSONSchemaForContract:input_type -> proto.Bytes
	1,  // 35: proto.IndexerService.Subscribe:output_type -> proto.Subscription
	32, // 36: proto.IndexerService.Unsubscribe:output_type -> proto.UnsubscribeResponse
	2,  // 37: proto.IndexerService.JSONSchemaForClass:output_type -> proto.Bytes
	2,  // 38: proto.IndexerService.JSONSchemaForContract:output_type -> proto.Bytes
	35, // [35:39] is the sub-list for method output_type
	31, // [31:35] is the sub-list for method input_type
	31, // [31:31] is the sub-list for extension type_name
	31, // [31:31] is the sub-list for extension extendee
	0,  // [0:31] is the sub-list for field type_name
}

func init() { file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_init() }
func file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_init() {
	if File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto != nil {
		return
	}
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_entity_filters_proto_init()
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bytes); i {
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
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonSchema); i {
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
		file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonSchemaItem); i {
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
			RawDescriptor: file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes,
		DependencyIndexes: file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs,
		MessageInfos:      file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_msgTypes,
	}.Build()
	File_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto = out.File
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_rawDesc = nil
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_goTypes = nil
	file_github_com_dipdup_io_starknet_indexer_pkg_grpc_proto_indexer_proto_depIdxs = nil
}
