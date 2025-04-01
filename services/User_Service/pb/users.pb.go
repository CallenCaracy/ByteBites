// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: users.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Role          string                 `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Address       string                 `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Phone         string                 `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	IsActive      string                 `protobuf:"bytes,8,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_users_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetIsActive() string {
	if x != nil {
		return x.IsActive
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type SignUpRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Role          string                 `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	Address       *string                `protobuf:"bytes,6,opt,name=address,proto3,oneof" json:"address,omitempty"`
	Phone         *string                `protobuf:"bytes,7,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpRequest) Reset() {
	*x = SignUpRequest{}
	mi := &file_users_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpRequest) ProtoMessage() {}

func (x *SignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpRequest.ProtoReflect.Descriptor instead.
func (*SignUpRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignUpRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SignUpRequest) GetAddress() string {
	if x != nil && x.Address != nil {
		return *x.Address
	}
	return ""
}

func (x *SignUpRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

type SignUpResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName     string                 `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Role          string                 `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignUpResponse) Reset() {
	*x = SignUpResponse{}
	mi := &file_users_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignUpResponse) ProtoMessage() {}

func (x *SignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignUpResponse.ProtoReflect.Descriptor instead.
func (*SignUpResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
}

func (x *SignUpResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SignUpResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SignUpResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SignUpResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type SignInRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInRequest) Reset() {
	*x = SignInRequest{}
	mi := &file_users_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInRequest) ProtoMessage() {}

func (x *SignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInRequest.ProtoReflect.Descriptor instead.
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInResponse) Reset() {
	*x = SignInResponse{}
	mi := &file_users_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResponse) ProtoMessage() {}

func (x *SignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResponse.ProtoReflect.Descriptor instead.
func (*SignInResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
}

func (x *SignInResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *SignInResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SignInOnlyEmployeeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInOnlyEmployeeRequest) Reset() {
	*x = SignInOnlyEmployeeRequest{}
	mi := &file_users_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInOnlyEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInOnlyEmployeeRequest) ProtoMessage() {}

func (x *SignInOnlyEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInOnlyEmployeeRequest.ProtoReflect.Descriptor instead.
func (*SignInOnlyEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
}

func (x *SignInOnlyEmployeeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInOnlyEmployeeRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInOnlyEmployeeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Error         string                 `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignInOnlyEmployeeResponse) Reset() {
	*x = SignInOnlyEmployeeResponse{}
	mi := &file_users_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignInOnlyEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInOnlyEmployeeResponse) ProtoMessage() {}

func (x *SignInOnlyEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInOnlyEmployeeResponse.ProtoReflect.Descriptor instead.
func (*SignInOnlyEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
}

func (x *SignInOnlyEmployeeResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInOnlyEmployeeResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *SignInOnlyEmployeeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SignOutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignOutRequest) Reset() {
	*x = SignOutRequest{}
	mi := &file_users_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignOutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutRequest) ProtoMessage() {}

func (x *SignOutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutRequest.ProtoReflect.Descriptor instead.
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
}

func (x *SignOutRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SignOutResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Error         string                 `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignOutResponse) Reset() {
	*x = SignOutResponse{}
	mi := &file_users_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignOutResponse) ProtoMessage() {}

func (x *SignOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignOutResponse.ProtoReflect.Descriptor instead.
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{8}
}

func (x *SignOutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SignOutResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetUserRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRoleRequest) Reset() {
	*x = GetUserRoleRequest{}
	mi := &file_users_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleRequest) ProtoMessage() {}

func (x *GetUserRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRoleRequest.ProtoReflect.Descriptor instead.
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{9}
}

func (x *GetUserRoleRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetUserRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Role          string                 `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRoleResponse) Reset() {
	*x = GetUserRoleResponse{}
	mi := &file_users_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleResponse) ProtoMessage() {}

func (x *GetUserRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRoleResponse.ProtoReflect.Descriptor instead.
func (*GetUserRoleResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserRoleResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *GetUserRoleResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	mi := &file_users_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{11}
}

func (x *TokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type TokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	mi := &file_users_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{12}
}

func (x *TokenResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TokenResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ForgotPasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForgotPasswordRequest) Reset() {
	*x = ForgotPasswordRequest{}
	mi := &file_users_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRequest) ProtoMessage() {}

func (x *ForgotPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{13}
}

func (x *ForgotPasswordRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ForgotPasswordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ForgotPasswordResponse) Reset() {
	*x = ForgotPasswordResponse{}
	mi := &file_users_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForgotPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordResponse) ProtoMessage() {}

func (x *ForgotPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForgotPasswordResponse.ProtoReflect.Descriptor instead.
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{14}
}

func (x *ForgotPasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ChangeUserPasswordRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	UserId          string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CurrentPassword string                 `protobuf:"bytes,2,opt,name=current_password,json=currentPassword,proto3" json:"current_password,omitempty"`
	NewPassword     string                 `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ChangeUserPasswordRequest) Reset() {
	*x = ChangeUserPasswordRequest{}
	mi := &file_users_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordRequest) ProtoMessage() {}

func (x *ChangeUserPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{15}
}

func (x *ChangeUserPasswordRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ChangeUserPasswordRequest) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangeUserPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ChangeUserPasswordResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangeUserPasswordResponse) Reset() {
	*x = ChangeUserPasswordResponse{}
	mi := &file_users_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeUserPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordResponse) ProtoMessage() {}

func (x *ChangeUserPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_users_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeUserPasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{16}
}

func (x *ChangeUserPasswordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_users_proto protoreflect.FileDescriptor

const file_users_proto_rawDesc = "" +
	"\n" +
	"\vusers.proto\x12\x04auth\"\x87\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\x12\x12\n" +
	"\x04role\x18\x05 \x01(\tR\x04role\x12\x18\n" +
	"\aaddress\x18\x06 \x01(\tR\aaddress\x12\x14\n" +
	"\x05phone\x18\a \x01(\tR\x05phone\x12\x1b\n" +
	"\tis_active\x18\b \x01(\tR\bisActive\x12\x1d\n" +
	"\n" +
	"created_at\x18\t \x01(\tR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_at\x18\n" +
	" \x01(\tR\tupdatedAt\"\xe1\x01\n" +
	"\rSignUpRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\x12\x12\n" +
	"\x04role\x18\x05 \x01(\tR\x04role\x12\x1d\n" +
	"\aaddress\x18\x06 \x01(\tH\x00R\aaddress\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\a \x01(\tH\x01R\x05phone\x88\x01\x01B\n" +
	"\n" +
	"\b_addressB\b\n" +
	"\x06_phone\"y\n" +
	"\x0eSignUpResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x1d\n" +
	"\n" +
	"first_name\x18\x02 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x03 \x01(\tR\blastName\x12\x12\n" +
	"\x04role\x18\x05 \x01(\tR\x04role\"A\n" +
	"\rSignInRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"n\n" +
	"\x0eSignInResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error\"M\n" +
	"\x19SignInOnlyEmployeeRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"z\n" +
	"\x1aSignInOnlyEmployeeResponse\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x12\x14\n" +
	"\x05error\x18\x03 \x01(\tR\x05error\")\n" +
	"\x0eSignOutRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"A\n" +
	"\x0fSignOutResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\x12\x14\n" +
	"\x05error\x18\x02 \x01(\tR\x05error\"*\n" +
	"\x12GetUserRoleRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"C\n" +
	"\x13GetUserRoleResponse\x12\x12\n" +
	"\x04role\x18\x01 \x01(\tR\x04role\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"$\n" +
	"\fTokenRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"5\n" +
	"\rTokenResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\"-\n" +
	"\x15ForgotPasswordRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"2\n" +
	"\x16ForgotPasswordResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\x82\x01\n" +
	"\x19ChangeUserPasswordRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12)\n" +
	"\x10current_password\x18\x02 \x01(\tR\x0fcurrentPassword\x12!\n" +
	"\fnew_password\x18\x03 \x01(\tR\vnewPassword\"6\n" +
	"\x1aChangeUserPasswordResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage2\xaa\x04\n" +
	"\vAuthService\x123\n" +
	"\x06SignUp\x12\x13.auth.SignUpRequest\x1a\x14.auth.SignUpResponse\x123\n" +
	"\x06SignIn\x12\x13.auth.SignInRequest\x1a\x14.auth.SignInResponse\x12W\n" +
	"\x12SignInOnlyEmployee\x12\x1f.auth.SignInOnlyEmployeeRequest\x1a .auth.SignInOnlyEmployeeResponse\x126\n" +
	"\aSignOut\x12\x14.auth.SignOutRequest\x1a\x15.auth.SignOutResponse\x12B\n" +
	"\vGetUserRole\x12\x18.auth.GetUserRoleRequest\x1a\x19.auth.GetUserRoleResponse\x126\n" +
	"\vVerifyToken\x12\x12.auth.TokenRequest\x1a\x13.auth.TokenResponse\x12K\n" +
	"\x0eForgotPassword\x12\x1b.auth.ForgotPasswordRequest\x1a\x1c.auth.ForgotPasswordResponse\x12W\n" +
	"\x12ChangeUserPassword\x12\x1f.auth.ChangeUserPasswordRequest\x1a .auth.ChangeUserPasswordResponseB\x06Z\x04./pbb\x06proto3"

var (
	file_users_proto_rawDescOnce sync.Once
	file_users_proto_rawDescData []byte
)

func file_users_proto_rawDescGZIP() []byte {
	file_users_proto_rawDescOnce.Do(func() {
		file_users_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_users_proto_rawDesc), len(file_users_proto_rawDesc)))
	})
	return file_users_proto_rawDescData
}

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_users_proto_goTypes = []any{
	(*User)(nil),                       // 0: auth.User
	(*SignUpRequest)(nil),              // 1: auth.SignUpRequest
	(*SignUpResponse)(nil),             // 2: auth.SignUpResponse
	(*SignInRequest)(nil),              // 3: auth.SignInRequest
	(*SignInResponse)(nil),             // 4: auth.SignInResponse
	(*SignInOnlyEmployeeRequest)(nil),  // 5: auth.SignInOnlyEmployeeRequest
	(*SignInOnlyEmployeeResponse)(nil), // 6: auth.SignInOnlyEmployeeResponse
	(*SignOutRequest)(nil),             // 7: auth.SignOutRequest
	(*SignOutResponse)(nil),            // 8: auth.SignOutResponse
	(*GetUserRoleRequest)(nil),         // 9: auth.GetUserRoleRequest
	(*GetUserRoleResponse)(nil),        // 10: auth.GetUserRoleResponse
	(*TokenRequest)(nil),               // 11: auth.TokenRequest
	(*TokenResponse)(nil),              // 12: auth.TokenResponse
	(*ForgotPasswordRequest)(nil),      // 13: auth.ForgotPasswordRequest
	(*ForgotPasswordResponse)(nil),     // 14: auth.ForgotPasswordResponse
	(*ChangeUserPasswordRequest)(nil),  // 15: auth.ChangeUserPasswordRequest
	(*ChangeUserPasswordResponse)(nil), // 16: auth.ChangeUserPasswordResponse
}
var file_users_proto_depIdxs = []int32{
	1,  // 0: auth.AuthService.SignUp:input_type -> auth.SignUpRequest
	3,  // 1: auth.AuthService.SignIn:input_type -> auth.SignInRequest
	5,  // 2: auth.AuthService.SignInOnlyEmployee:input_type -> auth.SignInOnlyEmployeeRequest
	7,  // 3: auth.AuthService.SignOut:input_type -> auth.SignOutRequest
	9,  // 4: auth.AuthService.GetUserRole:input_type -> auth.GetUserRoleRequest
	11, // 5: auth.AuthService.VerifyToken:input_type -> auth.TokenRequest
	13, // 6: auth.AuthService.ForgotPassword:input_type -> auth.ForgotPasswordRequest
	15, // 7: auth.AuthService.ChangeUserPassword:input_type -> auth.ChangeUserPasswordRequest
	2,  // 8: auth.AuthService.SignUp:output_type -> auth.SignUpResponse
	4,  // 9: auth.AuthService.SignIn:output_type -> auth.SignInResponse
	6,  // 10: auth.AuthService.SignInOnlyEmployee:output_type -> auth.SignInOnlyEmployeeResponse
	8,  // 11: auth.AuthService.SignOut:output_type -> auth.SignOutResponse
	10, // 12: auth.AuthService.GetUserRole:output_type -> auth.GetUserRoleResponse
	12, // 13: auth.AuthService.VerifyToken:output_type -> auth.TokenResponse
	14, // 14: auth.AuthService.ForgotPassword:output_type -> auth.ForgotPasswordResponse
	16, // 15: auth.AuthService.ChangeUserPassword:output_type -> auth.ChangeUserPasswordResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	file_users_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_users_proto_rawDesc), len(file_users_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_users_proto_goTypes,
		DependencyIndexes: file_users_proto_depIdxs,
		MessageInfos:      file_users_proto_msgTypes,
	}.Build()
	File_users_proto = out.File
	file_users_proto_goTypes = nil
	file_users_proto_depIdxs = nil
}
