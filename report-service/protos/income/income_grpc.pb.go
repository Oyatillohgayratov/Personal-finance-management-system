// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.21.12
// source: protos/income/income.proto

package income

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	IncomeExpensesService_LogIncome_FullMethodName       = "/IncomeExpensesService/LogIncome"
	IncomeExpensesService_LogExpense_FullMethodName      = "/IncomeExpensesService/LogExpense"
	IncomeExpensesService_GetTransactions_FullMethodName = "/IncomeExpensesService/GetTransactions"
)

// IncomeExpensesServiceClient is the client API for IncomeExpensesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IncomeExpensesServiceClient interface {
	LogIncome(ctx context.Context, in *IncomeRequest, opts ...grpc.CallOption) (*IncomeResponse, error)
	LogExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error)
	GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
}

type incomeExpensesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIncomeExpensesServiceClient(cc grpc.ClientConnInterface) IncomeExpensesServiceClient {
	return &incomeExpensesServiceClient{cc}
}

func (c *incomeExpensesServiceClient) LogIncome(ctx context.Context, in *IncomeRequest, opts ...grpc.CallOption) (*IncomeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncomeResponse)
	err := c.cc.Invoke(ctx, IncomeExpensesService_LogIncome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeExpensesServiceClient) LogExpense(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*ExpenseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpenseResponse)
	err := c.cc.Invoke(ctx, IncomeExpensesService_LogExpense_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *incomeExpensesServiceClient) GetTransactions(ctx context.Context, in *GetTransactionsRequest, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, IncomeExpensesService_GetTransactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IncomeExpensesServiceServer is the server API for IncomeExpensesService service.
// All implementations must embed UnimplementedIncomeExpensesServiceServer
// for forward compatibility
type IncomeExpensesServiceServer interface {
	LogIncome(context.Context, *IncomeRequest) (*IncomeResponse, error)
	LogExpense(context.Context, *ExpenseRequest) (*ExpenseResponse, error)
	GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error)
	mustEmbedUnimplementedIncomeExpensesServiceServer()
}

// UnimplementedIncomeExpensesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIncomeExpensesServiceServer struct {
}

func (UnimplementedIncomeExpensesServiceServer) LogIncome(context.Context, *IncomeRequest) (*IncomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIncome not implemented")
}
func (UnimplementedIncomeExpensesServiceServer) LogExpense(context.Context, *ExpenseRequest) (*ExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogExpense not implemented")
}
func (UnimplementedIncomeExpensesServiceServer) GetTransactions(context.Context, *GetTransactionsRequest) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (UnimplementedIncomeExpensesServiceServer) mustEmbedUnimplementedIncomeExpensesServiceServer() {}

// UnsafeIncomeExpensesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IncomeExpensesServiceServer will
// result in compilation errors.
type UnsafeIncomeExpensesServiceServer interface {
	mustEmbedUnimplementedIncomeExpensesServiceServer()
}

func RegisterIncomeExpensesServiceServer(s grpc.ServiceRegistrar, srv IncomeExpensesServiceServer) {
	s.RegisterService(&IncomeExpensesService_ServiceDesc, srv)
}

func _IncomeExpensesService_LogIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomeExpensesServiceServer).LogIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncomeExpensesService_LogIncome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomeExpensesServiceServer).LogIncome(ctx, req.(*IncomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncomeExpensesService_LogExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomeExpensesServiceServer).LogExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncomeExpensesService_LogExpense_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomeExpensesServiceServer).LogExpense(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IncomeExpensesService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IncomeExpensesServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IncomeExpensesService_GetTransactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IncomeExpensesServiceServer).GetTransactions(ctx, req.(*GetTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IncomeExpensesService_ServiceDesc is the grpc.ServiceDesc for IncomeExpensesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IncomeExpensesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "IncomeExpensesService",
	HandlerType: (*IncomeExpensesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LogIncome",
			Handler:    _IncomeExpensesService_LogIncome_Handler,
		},
		{
			MethodName: "LogExpense",
			Handler:    _IncomeExpensesService_LogExpense_Handler,
		},
		{
			MethodName: "GetTransactions",
			Handler:    _IncomeExpensesService_GetTransactions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/income/income.proto",
}
