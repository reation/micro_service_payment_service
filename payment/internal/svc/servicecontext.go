package svc

import (
	"fmt"
	"github.com/reation/micro_service_payment_service/model"
	"github.com/reation/micro_service_payment_service/payment/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	PaymentModel model.PaymentInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.PaymentTable.User,
		c.Mysql.PaymentTable.Passwd,
		c.Mysql.PaymentTable.Addr,
		c.Mysql.PaymentTable.Port,
		c.Mysql.PaymentTable.DBName,
	)
	return &ServiceContext{
		Config:       c,
		PaymentModel: model.NewPaymentInfoModel(sqlx.NewMysql(dataSourceUrl)),
	}
}
