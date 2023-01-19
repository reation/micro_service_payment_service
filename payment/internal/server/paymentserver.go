// Code generated by goctl. DO NOT EDIT!
// Source: payment.proto

package server

import (
	"context"

	"github.com/reation/micro_service_payment_service/payment/internal/logic"
	"github.com/reation/micro_service_payment_service/payment/internal/svc"
	"github.com/reation/micro_service_payment_service/protoc"
)

type PaymentServer struct {
	svcCtx *svc.ServiceContext
	protoc.UnimplementedPaymentServer
}

func NewPaymentServer(svcCtx *svc.ServiceContext) *PaymentServer {
	return &PaymentServer{
		svcCtx: svcCtx,
	}
}

func (s *PaymentServer) PaymentOrder(ctx context.Context, in *protoc.PaymentOrderRequest) (*protoc.PaymentOrderResponse, error) {
	l := logic.NewPaymentOrderLogic(ctx, s.svcCtx)
	return l.PaymentOrder(in)
}

func (s *PaymentServer) GetPaymentInfoByOrderNum(ctx context.Context, in *protoc.GetPaymentInfoByOrderNumRequest) (*protoc.GetPaymentInfoByOrderNumResponse, error) {
	l := logic.NewGetPaymentInfoByOrderNumLogic(ctx, s.svcCtx)
	return l.GetPaymentInfoByOrderNum(in)
}