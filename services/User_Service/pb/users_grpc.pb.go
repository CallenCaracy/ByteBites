// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: users.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_SignUp_FullMethodName             = "/auth.AuthService/SignUp"
	AuthService_SignIn_FullMethodName             = "/auth.AuthService/SignIn"
	AuthService_SignInOnlyEmployee_FullMethodName = "/auth.AuthService/SignInOnlyEmployee"
	AuthService_SignOut_FullMethodName            = "/auth.AuthService/SignOut"
	AuthService_GetUserRole_FullMethodName        = "/auth.AuthService/GetUserRole"
	AuthService_VerifyToken_FullMethodName        = "/auth.AuthService/VerifyToken"
	AuthService_ForgotPassword_FullMethodName     = "/auth.AuthService/ForgotPassword"
	AuthService_ChangeUserPassword_FullMethodName = "/auth.AuthService/ChangeUserPassword"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignInOnlyEmployee(ctx context.Context, in *SignInOnlyEmployeeRequest, opts ...grpc.CallOption) (*SignInOnlyEmployeeResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error)
	VerifyToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	// rpc GetUserInfo(GetUserInfoRequest) returns (GetUserInfoResponse); Moved to gqlgen
	// rpc UpdateUserInfo(UpdateUserInfoRequest) returns (UpdateUserInfoResponse); Moved to gqlgen
	ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error)
	ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, AuthService_SignUp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, AuthService_SignIn_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignInOnlyEmployee(ctx context.Context, in *SignInOnlyEmployeeRequest, opts ...grpc.CallOption) (*SignInOnlyEmployeeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignInOnlyEmployeeResponse)
	err := c.cc.Invoke(ctx, AuthService_SignInOnlyEmployee_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, AuthService_SignOut_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*GetUserRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserRoleResponse)
	err := c.cc.Invoke(ctx, AuthService_GetUserRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyToken(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, AuthService_VerifyToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ForgotPassword(ctx context.Context, in *ForgotPasswordRequest, opts ...grpc.CallOption) (*ForgotPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForgotPasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ForgotPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ChangeUserPassword(ctx context.Context, in *ChangeUserPasswordRequest, opts ...grpc.CallOption) (*ChangeUserPasswordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChangeUserPasswordResponse)
	err := c.cc.Invoke(ctx, AuthService_ChangeUserPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignInOnlyEmployee(context.Context, *SignInOnlyEmployeeRequest) (*SignInOnlyEmployeeResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error)
	VerifyToken(context.Context, *TokenRequest) (*TokenResponse, error)
	// rpc GetUserInfo(GetUserInfoRequest) returns (GetUserInfoResponse); Moved to gqlgen
	// rpc UpdateUserInfo(UpdateUserInfoRequest) returns (UpdateUserInfoResponse); Moved to gqlgen
	ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error)
	ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthServiceServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthServiceServer) SignInOnlyEmployee(context.Context, *SignInOnlyEmployeeRequest) (*SignInOnlyEmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInOnlyEmployee not implemented")
}
func (UnimplementedAuthServiceServer) SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (UnimplementedAuthServiceServer) GetUserRole(context.Context, *GetUserRoleRequest) (*GetUserRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserRole not implemented")
}
func (UnimplementedAuthServiceServer) VerifyToken(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedAuthServiceServer) ForgotPassword(context.Context, *ForgotPasswordRequest) (*ForgotPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForgotPassword not implemented")
}
func (UnimplementedAuthServiceServer) ChangeUserPassword(context.Context, *ChangeUserPasswordRequest) (*ChangeUserPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserPassword not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SignUp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SignIn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignInOnlyEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInOnlyEmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignInOnlyEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SignInOnlyEmployee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignInOnlyEmployee(ctx, req.(*SignInOnlyEmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_SignOut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_GetUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyToken(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ForgotPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForgotPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ForgotPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ForgotPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ForgotPassword(ctx, req.(*ForgotPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ChangeUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ChangeUserPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ChangeUserPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ChangeUserPassword(ctx, req.(*ChangeUserPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthService_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthService_SignIn_Handler,
		},
		{
			MethodName: "SignInOnlyEmployee",
			Handler:    _AuthService_SignInOnlyEmployee_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthService_SignOut_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _AuthService_GetUserRole_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _AuthService_VerifyToken_Handler,
		},
		{
			MethodName: "ForgotPassword",
			Handler:    _AuthService_ForgotPassword_Handler,
		},
		{
			MethodName: "ChangeUserPassword",
			Handler:    _AuthService_ChangeUserPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
