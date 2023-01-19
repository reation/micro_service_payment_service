package logic

import (
	"context"
	"github.com/reation/micro_service_payment_service/config"
	"github.com/reation/micro_service_payment_service/model"

	"github.com/reation/micro_service_payment_service/payment/internal/svc"
	"github.com/reation/micro_service_payment_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentInfoByOrderNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentInfoByOrderNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentInfoByOrderNumLogic {
	return &GetPaymentInfoByOrderNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentInfoByOrderNumLogic) GetPaymentInfoByOrderNum(in *protoc.GetPaymentInfoByOrderNumRequest) (*protoc.GetPaymentInfoByOrderNumResponse, error) {
	payInfo, err := l.svcCtx.PaymentModel.GetPaymentInfoByOrderNum(l.ctx, in.GetOrderNum())
	if err == model.ErrNotFound {
		return &protoc.GetPaymentInfoByOrderNumResponse{States: config.RETURN_STATES_EMPTY, PaymentInfo: nil}, nil
	}
	if err != nil {
		return &protoc.GetPaymentInfoByOrderNumResponse{States: config.RETURN_STATES_ERROR, PaymentInfo: nil}, nil
	}

	var resp = protoc.PaymentInfo{
		Id:            payInfo.Id,
		OrderNum:      payInfo.OrderNum,
		Prices:        payInfo.Prices,
		OrderUserId:   payInfo.OrderUserId,
		PaymentUserId: payInfo.PaymentUserId,
		PaymentType:   payInfo.PaymentType,
		PaymentState:  payInfo.PaymentState,
		UpdateTime:    payInfo.UpdateTime.Format("2006-01-02 15:04:05"),
		CreateTime:    payInfo.CreateTime.Format("2006-01-02 15:04:05"),
	}

	return &protoc.GetPaymentInfoByOrderNumResponse{States: config.RETURN_STATES_NORMAL, PaymentInfo: &resp}, nil
}
