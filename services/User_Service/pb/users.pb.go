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

type GetUserRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRoleRequest) Reset() {
	*x = GetUserRoleRequest{}
	mi := &file_users_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleRequest) ProtoMessage() {}

func (x *GetUserRoleRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserRoleRequest.ProtoReflect.Descriptor instead.
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{0}
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
	mi := &file_users_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleResponse) ProtoMessage() {}

func (x *GetUserRoleResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserRoleResponse.ProtoReflect.Descriptor instead.
func (*GetUserRoleResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{1}
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
	mi := &file_users_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{2}
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
	mi := &file_users_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{3}
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
	mi := &file_users_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForgotPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordRequest) ProtoMessage() {}

func (x *ForgotPasswordRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ForgotPasswordRequest.ProtoReflect.Descriptor instead.
func (*ForgotPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{4}
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
	mi := &file_users_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ForgotPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForgotPasswordResponse) ProtoMessage() {}

func (x *ForgotPasswordResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ForgotPasswordResponse.ProtoReflect.Descriptor instead.
func (*ForgotPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{5}
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
	mi := &file_users_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeUserPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordRequest) ProtoMessage() {}

func (x *ChangeUserPasswordRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ChangeUserPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordRequest) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{6}
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
	mi := &file_users_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangeUserPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeUserPasswordResponse) ProtoMessage() {}

func (x *ChangeUserPasswordResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ChangeUserPasswordResponse.ProtoReflect.Descriptor instead.
func (*ChangeUserPasswordResponse) Descriptor() ([]byte, []int) {
	return file_users_proto_rawDescGZIP(), []int{7}
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
	"\vusers.proto\x12\x04auth\"*\n" +
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
	"\amessage\x18\x01 \x01(\tR\amessage2\xaf\x02\n" +
	"\vAuthService\x12B\n" +
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

var file_users_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_users_proto_goTypes = []any{
	(*GetUserRoleRequest)(nil),         // 0: auth.GetUserRoleRequest
	(*GetUserRoleResponse)(nil),        // 1: auth.GetUserRoleResponse
	(*TokenRequest)(nil),               // 2: auth.TokenRequest
	(*TokenResponse)(nil),              // 3: auth.TokenResponse
	(*ForgotPasswordRequest)(nil),      // 4: auth.ForgotPasswordRequest
	(*ForgotPasswordResponse)(nil),     // 5: auth.ForgotPasswordResponse
	(*ChangeUserPasswordRequest)(nil),  // 6: auth.ChangeUserPasswordRequest
	(*ChangeUserPasswordResponse)(nil), // 7: auth.ChangeUserPasswordResponse
}
var file_users_proto_depIdxs = []int32{
	0, // 0: auth.AuthService.GetUserRole:input_type -> auth.GetUserRoleRequest
	2, // 1: auth.AuthService.VerifyToken:input_type -> auth.TokenRequest
	4, // 2: auth.AuthService.ForgotPassword:input_type -> auth.ForgotPasswordRequest
	6, // 3: auth.AuthService.ChangeUserPassword:input_type -> auth.ChangeUserPasswordRequest
	1, // 4: auth.AuthService.GetUserRole:output_type -> auth.GetUserRoleResponse
	3, // 5: auth.AuthService.VerifyToken:output_type -> auth.TokenResponse
	5, // 6: auth.AuthService.ForgotPassword:output_type -> auth.ForgotPasswordResponse
	7, // 7: auth.AuthService.ChangeUserPassword:output_type -> auth.ChangeUserPasswordResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_users_proto_init() }
func file_users_proto_init() {
	if File_users_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_users_proto_rawDesc), len(file_users_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
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
