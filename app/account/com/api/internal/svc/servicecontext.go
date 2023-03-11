package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"upi/app/account/com/api/internal/config"
	"upi/app/account/com/rpc/account"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpcConfig)),
	}
}
