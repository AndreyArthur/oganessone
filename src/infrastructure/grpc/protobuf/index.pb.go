// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: src/infrastructure/grpc/proto/index.proto

package protobuf

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Error) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Account) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateAccountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateAccountData) Reset() {
	*x = CreateAccountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountData) ProtoMessage() {}

func (x *CreateAccountData) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountData.ProtoReflect.Descriptor instead.
func (*CreateAccountData) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccountData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CreateAccountData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAccountRequest) GetData() *CreateAccountData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *Account `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAccountResponse) GetData() *Account {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateAccountResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type CreateSessionRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateSessionRequestData) Reset() {
	*x = CreateSessionRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequestData) ProtoMessage() {}

func (x *CreateSessionRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequestData.ProtoReflect.Descriptor instead.
func (*CreateSessionRequestData) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSessionRequestData) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CreateSessionRequestData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *CreateSessionRequestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{6}
}

func (x *CreateSessionRequest) GetData() *CreateSessionRequestData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CreateSessionResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account    *Account `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	SessionKey string   `protobuf:"bytes,2,opt,name=sessionKey,proto3" json:"sessionKey,omitempty"`
}

func (x *CreateSessionResponseData) Reset() {
	*x = CreateSessionResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponseData) ProtoMessage() {}

func (x *CreateSessionResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponseData.ProtoReflect.Descriptor instead.
func (*CreateSessionResponseData) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{7}
}

func (x *CreateSessionResponseData) GetAccount() *Account {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *CreateSessionResponseData) GetSessionKey() string {
	if x != nil {
		return x.SessionKey
	}
	return ""
}

type CreateSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *CreateSessionResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error *Error                     `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateSessionResponse) Reset() {
	*x = CreateSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionResponse) ProtoMessage() {}

func (x *CreateSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_infrastructure_grpc_proto_index_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateSessionResponse) Descriptor() ([]byte, []int) {
	return file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP(), []int{8}
}

func (x *CreateSessionResponse) GetData() *CreateSessionResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateSessionResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_src_infrastructure_grpc_proto_index_proto protoreflect.FileDescriptor

var file_src_infrastructure_grpc_proto_index_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x49, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x87, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x61, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x47, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4c, 0x0a,
	0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x4e, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x68, 0x0a, 0x19, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x22, 0x77, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xb9,
	0x01, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x64, 0x72, 0x65, 0x79, 0x41,
	0x72, 0x74, 0x68, 0x75, 0x72, 0x2f, 0x6f, 0x67, 0x61, 0x6e, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x65,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_infrastructure_grpc_proto_index_proto_rawDescOnce sync.Once
	file_src_infrastructure_grpc_proto_index_proto_rawDescData = file_src_infrastructure_grpc_proto_index_proto_rawDesc
)

func file_src_infrastructure_grpc_proto_index_proto_rawDescGZIP() []byte {
	file_src_infrastructure_grpc_proto_index_proto_rawDescOnce.Do(func() {
		file_src_infrastructure_grpc_proto_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_infrastructure_grpc_proto_index_proto_rawDescData)
	})
	return file_src_infrastructure_grpc_proto_index_proto_rawDescData
}

var file_src_infrastructure_grpc_proto_index_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_src_infrastructure_grpc_proto_index_proto_goTypes = []interface{}{
	(*Error)(nil),                     // 0: protobuf.Error
	(*Account)(nil),                   // 1: protobuf.Account
	(*CreateAccountData)(nil),         // 2: protobuf.CreateAccountData
	(*CreateAccountRequest)(nil),      // 3: protobuf.CreateAccountRequest
	(*CreateAccountResponse)(nil),     // 4: protobuf.CreateAccountResponse
	(*CreateSessionRequestData)(nil),  // 5: protobuf.CreateSessionRequestData
	(*CreateSessionRequest)(nil),      // 6: protobuf.CreateSessionRequest
	(*CreateSessionResponseData)(nil), // 7: protobuf.CreateSessionResponseData
	(*CreateSessionResponse)(nil),     // 8: protobuf.CreateSessionResponse
}
var file_src_infrastructure_grpc_proto_index_proto_depIdxs = []int32{
	2, // 0: protobuf.CreateAccountRequest.data:type_name -> protobuf.CreateAccountData
	1, // 1: protobuf.CreateAccountResponse.data:type_name -> protobuf.Account
	0, // 2: protobuf.CreateAccountResponse.error:type_name -> protobuf.Error
	5, // 3: protobuf.CreateSessionRequest.data:type_name -> protobuf.CreateSessionRequestData
	1, // 4: protobuf.CreateSessionResponseData.account:type_name -> protobuf.Account
	7, // 5: protobuf.CreateSessionResponse.data:type_name -> protobuf.CreateSessionResponseData
	0, // 6: protobuf.CreateSessionResponse.error:type_name -> protobuf.Error
	3, // 7: protobuf.AccountsService.CreateAccount:input_type -> protobuf.CreateAccountRequest
	6, // 8: protobuf.AccountsService.CreateSession:input_type -> protobuf.CreateSessionRequest
	4, // 9: protobuf.AccountsService.CreateAccount:output_type -> protobuf.CreateAccountResponse
	8, // 10: protobuf.AccountsService.CreateSession:output_type -> protobuf.CreateSessionResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_src_infrastructure_grpc_proto_index_proto_init() }
func file_src_infrastructure_grpc_proto_index_proto_init() {
	if File_src_infrastructure_grpc_proto_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountData); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequestData); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponseData); i {
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
		file_src_infrastructure_grpc_proto_index_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionResponse); i {
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
			RawDescriptor: file_src_infrastructure_grpc_proto_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_infrastructure_grpc_proto_index_proto_goTypes,
		DependencyIndexes: file_src_infrastructure_grpc_proto_index_proto_depIdxs,
		MessageInfos:      file_src_infrastructure_grpc_proto_index_proto_msgTypes,
	}.Build()
	File_src_infrastructure_grpc_proto_index_proto = out.File
	file_src_infrastructure_grpc_proto_index_proto_rawDesc = nil
	file_src_infrastructure_grpc_proto_index_proto_goTypes = nil
	file_src_infrastructure_grpc_proto_index_proto_depIdxs = nil
}
