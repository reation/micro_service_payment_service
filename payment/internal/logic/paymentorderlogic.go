package logic

import (
	"context"
	"github.com/reation/micro_service_payment_service/config"
	"github.com/reation/micro_service_payment_service/model"

	"github.com/reation/micro_service_payment_service/payment/internal/svc"
	"github.com/reation/micro_service_payment_service/protoc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentOrderLogic {
	return &PaymentOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaymentOrderLogic) PaymentOrder(in *protoc.PaymentOrderRequest) (*protoc.PaymentOrderResponse, error) {
	payInfo, err := l.svcCtx.PaymentModel.GetPaymentInfoByOrderNum(l.ctx, in.GetOrderNum())
	if err == model.ErrNotFound {
		state, lastID := l.newPaymentOrder(in)
		return &protoc.PaymentOrderResponse{States: state, OrderNum: in.GetOrderNum(), LastId: lastID}, nil
	}
	if err != nil {
		return &protoc.PaymentOrderResponse{States: config.RETURN_STATES_ERROR, OrderNum: in.GetOrderNum(), LastId: 0}, nil
	}

	if payInfo.PaymentType == in.GetPaymentType() {
		return &protoc.PaymentOrderResponse{States: config.RETURN_STATES_NORMAL, OrderNum: in.GetOrderNum(), LastId: payInfo.Id}, nil
	}

	err = l.svcCtx.PaymentModel.EditPaymentTypeByID(l.ctx, payInfo.Id, in.GetPaymentType())
	if err != nil {
		return &protoc.PaymentOrderResponse{States: config.RETURN_STATES_ERROR, OrderNum: in.GetOrderNum(), LastId: 0}, nil
	}

	return &protoc.PaymentOrderResponse{States: config.RETURN_STATES_NORMAL, OrderNum: in.GetOrderNum(), LastId: payInfo.Id}, nil
}

func (l *PaymentOrderLogic) newPaymentOrder(in *protoc.PaymentOrderRequest) (state, lastID int64) {
	var info model.PaymentInfo
	info.OrderNum = in.GetOrderNum()
	info.Prices = in.GetPrices()
	info.OrderUserId = in.GetOrderUserId()
	info.PaymentUserId = in.GetPaymentUserId()
	info.PaymentType = in.GetPaymentType()
	info.PaymentState = model.STATE_SUCCESS

	result, err := l.svcCtx.PaymentModel.Insert(l.ctx, &info)
	if err != nil {
		return config.RETURN_STATES_ERROR, 0
	}

	lastID, _ = result.LastInsertId()
	return config.RETURN_STATES_NORMAL, lastID
}
