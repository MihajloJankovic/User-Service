// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: app.proto

package main

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

type AuthGet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *AuthGet) Reset() {
	*x = AuthGet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthGet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthGet) ProtoMessage() {}

func (x *AuthGet) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthGet.ProtoReflect.Descriptor instead.
func (*AuthGet) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{0}
}

func (x *AuthGet) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AuthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *AuthRequest) Reset() {
	*x = AuthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthRequest) ProtoMessage() {}

func (x *AuthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthRequest.ProtoReflect.Descriptor instead.
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{1}
}

func (x *AuthRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type AuthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ticket      string `protobuf:"bytes,3,opt,name=ticket,proto3" json:"ticket,omitempty"`
	TicketReset string `protobuf:"bytes,4,opt,name=ticketReset,proto3" json:"ticketReset,omitempty"`
	Activated   bool   `protobuf:"varint,5,opt,name=activated,proto3" json:"activated,omitempty"`
}

func (x *AuthResponse) Reset() {
	*x = AuthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResponse) ProtoMessage() {}

func (x *AuthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResponse.ProtoReflect.Descriptor instead.
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AuthResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AuthResponse) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

func (x *AuthResponse) GetTicketReset() string {
	if x != nil {
		return x.TicketReset
	}
	return ""
}

func (x *AuthResponse) GetActivated() bool {
	if x != nil {
		return x.Activated
	}
	return false
}

type AuthEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthEmpty) Reset() {
	*x = AuthEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthEmpty) ProtoMessage() {}

func (x *AuthEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthEmpty.ProtoReflect.Descriptor instead.
func (*AuthEmpty) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{3}
}

type AuthTicket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ticket string `protobuf:"bytes,1,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *AuthTicket) Reset() {
	*x = AuthTicket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthTicket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthTicket) ProtoMessage() {}

func (x *AuthTicket) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthTicket.ProtoReflect.Descriptor instead.
func (*AuthTicket) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{4}
}

func (x *AuthTicket) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

type ActivateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email  string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (x *ActivateRequest) Reset() {
	*x = ActivateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateRequest) ProtoMessage() {}

func (x *ActivateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateRequest.ProtoReflect.Descriptor instead.
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{5}
}

func (x *ActivateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ActivateRequest) GetTicket() string {
	if x != nil {
		return x.Ticket
	}
	return ""
}

type ChangePasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	CurrentPassword string `protobuf:"bytes,2,opt,name=currentPassword,proto3" json:"currentPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{6}
}

func (x *ChangePasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ChangePasswordRequest) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	TicketReset string `protobuf:"bytes,2,opt,name=ticketReset,proto3" json:"ticketReset,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=newPassword,proto3" json:"newPassword,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_app_proto_rawDescGZIP(), []int{7}
}

func (x *ResetRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ResetRequest) GetTicketReset() string {
	if x != nil {
		return x.TicketReset
	}
	return ""
}

func (x *ResetRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_app_proto protoreflect.FileDescriptor

var file_app_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x07, 0x41,
	0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3f, 0x0a, 0x0b,
	0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x98, 0x01,
	0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x22, 0x0b, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x24, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x3f, 0x0a, 0x0f, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x79, 0x0a, 0x15,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x68, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x32, 0xac, 0x02, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x24, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x1f, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0c, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65,
	0x74, 0x12, 0x22, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x08,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x1a, 0x0b, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x41,
	0x75, 0x74, 0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x12, 0x08, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74, 0x1a, 0x0a, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x47, 0x65, 0x74,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_app_proto_rawDescOnce sync.Once
	file_app_proto_rawDescData = file_app_proto_rawDesc
)

func file_app_proto_rawDescGZIP() []byte {
	file_app_proto_rawDescOnce.Do(func() {
		file_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proto_rawDescData)
	})
	return file_app_proto_rawDescData
}

var file_app_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_proto_goTypes = []interface{}{
	(*AuthGet)(nil),               // 0: AuthGet
	(*AuthRequest)(nil),           // 1: AuthRequest
	(*AuthResponse)(nil),          // 2: AuthResponse
	(*AuthEmpty)(nil),             // 3: AuthEmpty
	(*AuthTicket)(nil),            // 4: AuthTicket
	(*ActivateRequest)(nil),       // 5: ActivateRequest
	(*ChangePasswordRequest)(nil), // 6: ChangePasswordRequest
	(*ResetRequest)(nil),          // 7: ResetRequest
}
var file_app_proto_depIdxs = []int32{
	1, // 0: auth.Register:input_type -> AuthRequest
	1, // 1: auth.Login:input_type -> AuthRequest
	0, // 2: auth.GetTicket:input_type -> AuthGet
	5, // 3: auth.Activate:input_type -> ActivateRequest
	6, // 4: auth.ChangePassword:input_type -> ChangePasswordRequest
	0, // 5: auth.RequestPasswordReset:input_type -> AuthGet
	7, // 6: auth.ResetPassword:input_type -> ResetRequest
	3, // 7: auth.Register:output_type -> AuthEmpty
	0, // 8: auth.Login:output_type -> AuthGet
	4, // 9: auth.GetTicket:output_type -> AuthTicket
	2, // 10: auth.Activate:output_type -> AuthResponse
	3, // 11: auth.ChangePassword:output_type -> AuthEmpty
	3, // 12: auth.RequestPasswordReset:output_type -> AuthEmpty
	0, // 13: auth.ResetPassword:output_type -> AuthGet
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_app_proto_init() }
func file_app_proto_init() {
	if File_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthGet); i {
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
		file_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthRequest); i {
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
		file_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResponse); i {
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
		file_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthEmpty); i {
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
		file_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthTicket); i {
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
		file_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivateRequest); i {
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
		file_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePasswordRequest); i {
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
		file_app_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
			RawDescriptor: file_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proto_goTypes,
		DependencyIndexes: file_app_proto_depIdxs,
		MessageInfos:      file_app_proto_msgTypes,
	}.Build()
	File_app_proto = out.File
	file_app_proto_rawDesc = nil
	file_app_proto_goTypes = nil
	file_app_proto_depIdxs = nil
}
