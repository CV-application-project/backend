// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.13.0
// source: cv_service/api/data.proto

package api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type Gender int32

const (
	Gender_MALE   Gender = 0
	Gender_FEMALE Gender = 1
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "MALE",
		1: "FEMALE",
	}
	Gender_value = map[string]int32{
		"MALE":   0,
		"FEMALE": 1,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_cv_service_api_data_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_cv_service_api_data_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{0}
}

type RegisterCICForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Front  []byte `protobuf:"bytes,2,opt,name=front,proto3" json:"front,omitempty"`
	Back   []byte `protobuf:"bytes,3,opt,name=back,proto3" json:"back,omitempty"`
}

func (x *RegisterCICForUserRequest) Reset() {
	*x = RegisterCICForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCICForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCICForUserRequest) ProtoMessage() {}

func (x *RegisterCICForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCICForUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterCICForUserRequest) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterCICForUserRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterCICForUserRequest) GetFront() []byte {
	if x != nil {
		return x.Front
	}
	return nil
}

func (x *RegisterCICForUserRequest) GetBack() []byte {
	if x != nil {
		return x.Back
	}
	return nil
}

type RegisterCICForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterCICForUserResponse) Reset() {
	*x = RegisterCICForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterCICForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterCICForUserResponse) ProtoMessage() {}

func (x *RegisterCICForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterCICForUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterCICForUserResponse) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterCICForUserResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterCICForUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCICByUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetCICByUserIdRequest) Reset() {
	*x = GetCICByUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCICByUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCICByUserIdRequest) ProtoMessage() {}

func (x *GetCICByUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCICByUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetCICByUserIdRequest) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{2}
}

func (x *GetCICByUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetCICByUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64                `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Card   *CitizenIdentityCard `protobuf:"bytes,2,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *GetCICByUserIdResponse) Reset() {
	*x = GetCICByUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCICByUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCICByUserIdResponse) ProtoMessage() {}

func (x *GetCICByUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCICByUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetCICByUserIdResponse) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{3}
}

func (x *GetCICByUserIdResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetCICByUserIdResponse) GetCard() *CitizenIdentityCard {
	if x != nil {
		return x.Card
	}
	return nil
}

type CitizenIdentityCard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RegisterDate int64  `protobuf:"varint,2,opt,name=register_date,json=registerDate,proto3" json:"register_date,omitempty"`
	ExpireDate   int64  `protobuf:"varint,3,opt,name=expire_date,json=expireDate,proto3" json:"expire_date,omitempty"`
	Provider     string `protobuf:"bytes,4,opt,name=provider,proto3" json:"provider,omitempty"`
	Birthday     int64  `protobuf:"varint,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	BornProvince string `protobuf:"bytes,6,opt,name=born_province,json=bornProvince,proto3" json:"born_province,omitempty"`
	Name         string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Gender       Gender `protobuf:"varint,8,opt,name=gender,proto3,enum=cv_service.api.Gender" json:"gender,omitempty"`
	Country      string `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	Location     string `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CitizenIdentityCard) Reset() {
	*x = CitizenIdentityCard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitizenIdentityCard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitizenIdentityCard) ProtoMessage() {}

func (x *CitizenIdentityCard) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitizenIdentityCard.ProtoReflect.Descriptor instead.
func (*CitizenIdentityCard) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{4}
}

func (x *CitizenIdentityCard) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CitizenIdentityCard) GetRegisterDate() int64 {
	if x != nil {
		return x.RegisterDate
	}
	return 0
}

func (x *CitizenIdentityCard) GetExpireDate() int64 {
	if x != nil {
		return x.ExpireDate
	}
	return 0
}

func (x *CitizenIdentityCard) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CitizenIdentityCard) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *CitizenIdentityCard) GetBornProvince() string {
	if x != nil {
		return x.BornProvince
	}
	return ""
}

func (x *CitizenIdentityCard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CitizenIdentityCard) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_MALE
}

func (x *CitizenIdentityCard) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CitizenIdentityCard) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

type RegisterUserFaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Image  []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *RegisterUserFaceRequest) Reset() {
	*x = RegisterUserFaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserFaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserFaceRequest) ProtoMessage() {}

func (x *RegisterUserFaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserFaceRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserFaceRequest) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterUserFaceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RegisterUserFaceRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type RegisterUserFaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RegisterUserFaceResponse) Reset() {
	*x = RegisterUserFaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterUserFaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserFaceResponse) ProtoMessage() {}

func (x *RegisterUserFaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserFaceResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserFaceResponse) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterUserFaceResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RegisterUserFaceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AuthorizeUserFaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Image  []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *AuthorizeUserFaceRequest) Reset() {
	*x = AuthorizeUserFaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserFaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserFaceRequest) ProtoMessage() {}

func (x *AuthorizeUserFaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserFaceRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeUserFaceRequest) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{7}
}

func (x *AuthorizeUserFaceRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthorizeUserFaceRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type AuthorizeUserFaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	UserId  int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthorizeUserFaceResponse) Reset() {
	*x = AuthorizeUserFaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cv_service_api_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserFaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserFaceResponse) ProtoMessage() {}

func (x *AuthorizeUserFaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cv_service_api_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserFaceResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeUserFaceResponse) Descriptor() ([]byte, []int) {
	return file_cv_service_api_data_proto_rawDescGZIP(), []int{8}
}

func (x *AuthorizeUserFaceResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AuthorizeUserFaceResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthorizeUserFaceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cv_service_api_data_proto protoreflect.FileDescriptor

var file_cv_service_api_data_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x2a, 0x6c, 0x69, 0x62,
	0x72, 0x61, 0x72, 0x79, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x19, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x49, 0x43, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x61, 0x63, 0x6b,
	0x22, 0x4a, 0x0a, 0x1a, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x49, 0x43, 0x46,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3b, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x49, 0x43, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x22, 0x04, 0x20, 0x00, 0x40,
	0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x49, 0x43, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x04,
	0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x76, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x69, 0x74, 0x69,
	0x7a, 0x65, 0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x52,
	0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0xf2, 0x02, 0x0a, 0x13, 0x43, 0x69, 0x74, 0x69, 0x7a, 0x65,
	0x6e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a,
	0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x2f, 0x0a, 0x0d, 0x62, 0x6f, 0x72, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x72, 0x05, 0x10, 0x00, 0x18, 0xff, 0x01, 0x52, 0x0c, 0x62, 0x6f, 0x72, 0x6e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x17, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x7a, 0x04, 0x10, 0x00, 0x70,
	0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x5d, 0x0a, 0x18, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x7a, 0x04, 0x10, 0x00, 0x70, 0x00, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x22, 0x62, 0x0a, 0x19, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1e, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12,
	0x08, 0x0a, 0x04, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x45, 0x4d,
	0x41, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cv_service_api_data_proto_rawDescOnce sync.Once
	file_cv_service_api_data_proto_rawDescData = file_cv_service_api_data_proto_rawDesc
)

func file_cv_service_api_data_proto_rawDescGZIP() []byte {
	file_cv_service_api_data_proto_rawDescOnce.Do(func() {
		file_cv_service_api_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_cv_service_api_data_proto_rawDescData)
	})
	return file_cv_service_api_data_proto_rawDescData
}

var file_cv_service_api_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cv_service_api_data_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cv_service_api_data_proto_goTypes = []interface{}{
	(Gender)(0),                        // 0: cv_service.api.Gender
	(*RegisterCICForUserRequest)(nil),  // 1: cv_service.api.RegisterCICForUserRequest
	(*RegisterCICForUserResponse)(nil), // 2: cv_service.api.RegisterCICForUserResponse
	(*GetCICByUserIdRequest)(nil),      // 3: cv_service.api.GetCICByUserIdRequest
	(*GetCICByUserIdResponse)(nil),     // 4: cv_service.api.GetCICByUserIdResponse
	(*CitizenIdentityCard)(nil),        // 5: cv_service.api.CitizenIdentityCard
	(*RegisterUserFaceRequest)(nil),    // 6: cv_service.api.RegisterUserFaceRequest
	(*RegisterUserFaceResponse)(nil),   // 7: cv_service.api.RegisterUserFaceResponse
	(*AuthorizeUserFaceRequest)(nil),   // 8: cv_service.api.AuthorizeUserFaceRequest
	(*AuthorizeUserFaceResponse)(nil),  // 9: cv_service.api.AuthorizeUserFaceResponse
}
var file_cv_service_api_data_proto_depIdxs = []int32{
	5, // 0: cv_service.api.GetCICByUserIdResponse.card:type_name -> cv_service.api.CitizenIdentityCard
	0, // 1: cv_service.api.CitizenIdentityCard.gender:type_name -> cv_service.api.Gender
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cv_service_api_data_proto_init() }
func file_cv_service_api_data_proto_init() {
	if File_cv_service_api_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cv_service_api_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCICForUserRequest); i {
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
		file_cv_service_api_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterCICForUserResponse); i {
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
		file_cv_service_api_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCICByUserIdRequest); i {
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
		file_cv_service_api_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCICByUserIdResponse); i {
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
		file_cv_service_api_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitizenIdentityCard); i {
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
		file_cv_service_api_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserFaceRequest); i {
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
		file_cv_service_api_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterUserFaceResponse); i {
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
		file_cv_service_api_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserFaceRequest); i {
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
		file_cv_service_api_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserFaceResponse); i {
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
			RawDescriptor: file_cv_service_api_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cv_service_api_data_proto_goTypes,
		DependencyIndexes: file_cv_service_api_data_proto_depIdxs,
		EnumInfos:         file_cv_service_api_data_proto_enumTypes,
		MessageInfos:      file_cv_service_api_data_proto_msgTypes,
	}.Build()
	File_cv_service_api_data_proto = out.File
	file_cv_service_api_data_proto_rawDesc = nil
	file_cv_service_api_data_proto_goTypes = nil
	file_cv_service_api_data_proto_depIdxs = nil
}
