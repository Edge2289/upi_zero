package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"upi/app/mqueue/cmd/mq/internal/config"
	"upi/app/unionpay/com/rpc/unionpay"
)

type ServiceContext struct {
	Config config.Config

	Unionpay unionpay.Unionpay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Unionpay: unionpay.NewUnionpay(zrpc.MustNewClient(c.UnionPayRpcConfig)),
	}
}
