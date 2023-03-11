package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"upi/app/account/com/rpc/account"
	"upi/app/business/com/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	DbEngine               *gorm.DB
	AccountRpc             account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.DB.DataSource), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,

		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpcConfig)),
		DbEngine: db,
	}
}
