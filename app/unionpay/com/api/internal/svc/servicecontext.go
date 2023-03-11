package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"upi/app/unionpay/com/api/internal/config"
	"upi/app/unionpay/com/rpc/unionpay"
)

type ServiceContext struct {
	Config config.Config

	Cache *redis.Redis

	UnionPay unionpay.Unionpay
	Greet unionpay.StreamGreeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		Cache: redis.New(c.RedisConf.Host, func(r *redis.Redis) {
			r.Type = c.RedisConf.Type
			r.Pass = c.RedisConf.Pass
		}),

		UnionPay: unionpay.NewUnionpay(zrpc.MustNewClient(c.UnionPayRpcConfig)),
		Greet: unionpay.NewStreamGreeter(zrpc.MustNewClient(c.UnionPayRpcConfig)),
	}
}
