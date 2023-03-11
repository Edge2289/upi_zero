package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"upi/app/unionpay/com/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	KqueuePay *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		KqueuePay: kq.NewPusher(c.KqueuePayConf.Brokers, c.KqueuePayConf.Topic),
	}
}
