// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsServiceClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error)
}

type accountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsServiceClient(cc grpc.ClientConnInterface) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountsService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountsService/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) RefreshSession(ctx context.Context, in *RefreshSessionRequest, opts ...grpc.CallOption) (*RefreshSessionResponse, error) {
	out := new(RefreshSessionResponse)
	err := c.cc.Invoke(ctx, "/protobuf.AccountsService/RefreshSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
// All implementations must embed UnimplementedAccountsServiceServer
// for forward compatibility
type AccountsServiceServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error)
	mustEmbedUnimplementedAccountsServiceServer()
}

// UnimplementedAccountsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServiceServer struct {
}

func (UnimplementedAccountsServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountsServiceServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedAccountsServiceServer) RefreshSession(context.Context, *RefreshSessionRequest) (*RefreshSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshSession not implemented")
}
func (UnimplementedAccountsServiceServer) mustEmbedUnimplementedAccountsServiceServer() {}

// UnsafeAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServiceServer will
// result in compilation errors.
type UnsafeAccountsServiceServer interface {
	mustEmbedUnimplementedAccountsServiceServer()
}

func RegisterAccountsServiceServer(s grpc.ServiceRegistrar, srv AccountsServiceServer) {
	s.RegisterService(&AccountsService_ServiceDesc, srv)
}

func _AccountsService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountsService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountsService/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_RefreshSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).RefreshSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AccountsService/RefreshSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).RefreshSession(ctx, req.(*RefreshSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountsService_ServiceDesc is the grpc.ServiceDesc for AccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountsService_CreateAccount_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _AccountsService_CreateSession_Handler,
		},
		{
			MethodName: "RefreshSession",
			Handler:    _AccountsService_RefreshSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/infrastructure/grpc/proto/index.proto",
}
