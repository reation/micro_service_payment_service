// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: payment.proto

package protoc

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

// PaymentClient is the client API for Payment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentClient interface {
	PaymentOrder(ctx context.Context, in *PaymentOrderRequest, opts ...grpc.CallOption) (*PaymentOrderResponse, error)
	GetPaymentInfoByOrderNum(ctx context.Context, in *GetPaymentInfoByOrderNumRequest, opts ...grpc.CallOption) (*GetPaymentInfoByOrderNumResponse, error)
}

type paymentClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentClient(cc grpc.ClientConnInterface) PaymentClient {
	return &paymentClient{cc}
}

func (c *paymentClient) PaymentOrder(ctx context.Context, in *PaymentOrderRequest, opts ...grpc.CallOption) (*PaymentOrderResponse, error) {
	out := new(PaymentOrderResponse)
	err := c.cc.Invoke(ctx, "/payment.payment/paymentOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentClient) GetPaymentInfoByOrderNum(ctx context.Context, in *GetPaymentInfoByOrderNumRequest, opts ...grpc.CallOption) (*GetPaymentInfoByOrderNumResponse, error) {
	out := new(GetPaymentInfoByOrderNumResponse)
	err := c.cc.Invoke(ctx, "/payment.payment/getPaymentInfoByOrderNum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServer is the server API for Payment service.
// All implementations must embed UnimplementedPaymentServer
// for forward compatibility
type PaymentServer interface {
	PaymentOrder(context.Context, *PaymentOrderRequest) (*PaymentOrderResponse, error)
	GetPaymentInfoByOrderNum(context.Context, *GetPaymentInfoByOrderNumRequest) (*GetPaymentInfoByOrderNumResponse, error)
	mustEmbedUnimplementedPaymentServer()
}

// UnimplementedPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentServer struct {
}

func (UnimplementedPaymentServer) PaymentOrder(context.Context, *PaymentOrderRequest) (*PaymentOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PaymentOrder not implemented")
}
func (UnimplementedPaymentServer) GetPaymentInfoByOrderNum(context.Context, *GetPaymentInfoByOrderNumRequest) (*GetPaymentInfoByOrderNumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentInfoByOrderNum not implemented")
}
func (UnimplementedPaymentServer) mustEmbedUnimplementedPaymentServer() {}

// UnsafePaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServer will
// result in compilation errors.
type UnsafePaymentServer interface {
	mustEmbedUnimplementedPaymentServer()
}

func RegisterPaymentServer(s grpc.ServiceRegistrar, srv PaymentServer) {
	s.RegisterService(&Payment_ServiceDesc, srv)
}

func _Payment_PaymentOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).PaymentOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/paymentOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).PaymentOrder(ctx, req.(*PaymentOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Payment_GetPaymentInfoByOrderNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentInfoByOrderNumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServer).GetPaymentInfoByOrderNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/payment.payment/getPaymentInfoByOrderNum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServer).GetPaymentInfoByOrderNum(ctx, req.(*GetPaymentInfoByOrderNumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Payment_ServiceDesc is the grpc.ServiceDesc for Payment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Payment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.payment",
	HandlerType: (*PaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "paymentOrder",
			Handler:    _Payment_PaymentOrder_Handler,
		},
		{
			MethodName: "getPaymentInfoByOrderNum",
			Handler:    _Payment_GetPaymentInfoByOrderNum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
