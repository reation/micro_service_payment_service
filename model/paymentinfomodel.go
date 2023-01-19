package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ PaymentInfoModel = (*customPaymentInfoModel)(nil)

type (
	// PaymentInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPaymentInfoModel.
	PaymentInfoModel interface {
		paymentInfoModel
	}

	customPaymentInfoModel struct {
		*defaultPaymentInfoModel
	}
)

// NewPaymentInfoModel returns a model for the database table.
func NewPaymentInfoModel(conn sqlx.SqlConn) PaymentInfoModel {
	return &customPaymentInfoModel{
		defaultPaymentInfoModel: newPaymentInfoModel(conn),
	}
}
