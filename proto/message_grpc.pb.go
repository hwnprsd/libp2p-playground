// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/message.proto

package proto

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

const (
	TransactionService_HandleTransaction_FullMethodName      = "/TransactionService/HandleTransaction"
	TransactionService_HandleSignatureRequest_FullMethodName = "/TransactionService/HandleSignatureRequest"
	TransactionService_HandleCreateRule_FullMethodName       = "/TransactionService/HandleCreateRule"
	TransactionService_HandleMetricsQuery_FullMethodName     = "/TransactionService/HandleMetricsQuery"
	TransactionService_HandleGenericRequest_FullMethodName   = "/TransactionService/HandleGenericRequest"
)

// TransactionServiceClient is the client API for TransactionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionServiceClient interface {
	HandleTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	HandleSignatureRequest(ctx context.Context, in *SolaceTx, opts ...grpc.CallOption) (*TransactionResponse, error)
	HandleCreateRule(ctx context.Context, in *CreateRuleData, opts ...grpc.CallOption) (*TransactionResponse, error)
	HandleMetricsQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetricsResponse, error)
	HandleGenericRequest(ctx context.Context, in *GenericRequestData, opts ...grpc.CallOption) (*TransactionResponse, error)
}

type transactionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionServiceClient(cc grpc.ClientConnInterface) TransactionServiceClient {
	return &transactionServiceClient{cc}
}

func (c *transactionServiceClient) HandleTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_HandleTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) HandleSignatureRequest(ctx context.Context, in *SolaceTx, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_HandleSignatureRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) HandleCreateRule(ctx context.Context, in *CreateRuleData, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_HandleCreateRule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) HandleMetricsQuery(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, TransactionService_HandleMetricsQuery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionServiceClient) HandleGenericRequest(ctx context.Context, in *GenericRequestData, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := c.cc.Invoke(ctx, TransactionService_HandleGenericRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServiceServer is the server API for TransactionService service.
// All implementations must embed UnimplementedTransactionServiceServer
// for forward compatibility
type TransactionServiceServer interface {
	HandleTransaction(context.Context, *Transaction) (*TransactionResponse, error)
	HandleSignatureRequest(context.Context, *SolaceTx) (*TransactionResponse, error)
	HandleCreateRule(context.Context, *CreateRuleData) (*TransactionResponse, error)
	HandleMetricsQuery(context.Context, *Empty) (*MetricsResponse, error)
	HandleGenericRequest(context.Context, *GenericRequestData) (*TransactionResponse, error)
	mustEmbedUnimplementedTransactionServiceServer()
}

// UnimplementedTransactionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionServiceServer struct {
}

func (UnimplementedTransactionServiceServer) HandleTransaction(context.Context, *Transaction) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleTransaction not implemented")
}
func (UnimplementedTransactionServiceServer) HandleSignatureRequest(context.Context, *SolaceTx) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSignatureRequest not implemented")
}
func (UnimplementedTransactionServiceServer) HandleCreateRule(context.Context, *CreateRuleData) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleCreateRule not implemented")
}
func (UnimplementedTransactionServiceServer) HandleMetricsQuery(context.Context, *Empty) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleMetricsQuery not implemented")
}
func (UnimplementedTransactionServiceServer) HandleGenericRequest(context.Context, *GenericRequestData) (*TransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleGenericRequest not implemented")
}
func (UnimplementedTransactionServiceServer) mustEmbedUnimplementedTransactionServiceServer() {}

// UnsafeTransactionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionServiceServer will
// result in compilation errors.
type UnsafeTransactionServiceServer interface {
	mustEmbedUnimplementedTransactionServiceServer()
}

func RegisterTransactionServiceServer(s grpc.ServiceRegistrar, srv TransactionServiceServer) {
	s.RegisterService(&TransactionService_ServiceDesc, srv)
}

func _TransactionService_HandleTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).HandleTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_HandleTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).HandleTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_HandleSignatureRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SolaceTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).HandleSignatureRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_HandleSignatureRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).HandleSignatureRequest(ctx, req.(*SolaceTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_HandleCreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).HandleCreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_HandleCreateRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).HandleCreateRule(ctx, req.(*CreateRuleData))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_HandleMetricsQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).HandleMetricsQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_HandleMetricsQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).HandleMetricsQuery(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionService_HandleGenericRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericRequestData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServiceServer).HandleGenericRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransactionService_HandleGenericRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServiceServer).HandleGenericRequest(ctx, req.(*GenericRequestData))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionService_ServiceDesc is the grpc.ServiceDesc for TransactionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TransactionService",
	HandlerType: (*TransactionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleTransaction",
			Handler:    _TransactionService_HandleTransaction_Handler,
		},
		{
			MethodName: "HandleSignatureRequest",
			Handler:    _TransactionService_HandleSignatureRequest_Handler,
		},
		{
			MethodName: "HandleCreateRule",
			Handler:    _TransactionService_HandleCreateRule_Handler,
		},
		{
			MethodName: "HandleMetricsQuery",
			Handler:    _TransactionService_HandleMetricsQuery_Handler,
		},
		{
			MethodName: "HandleGenericRequest",
			Handler:    _TransactionService_HandleGenericRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message.proto",
}
