package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"upi/common/resources/db"

	"upi/app/account/com/rpc/account"
	"upi/app/business/com/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	DbEngine               *gorm.DB
	AccountRpc             account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpcConfig)),
		DbEngine: db.Instance(db.GormConfig{
			Host:        c.Db.Host,
			User:        c.Db.User,
			PassWord:    c.Db.Password,
			DbName:      c.Db.DbName,
			DeBug:       c.Db.DeBug,
			MaxLifetime: c.Db.MaxLifetime,
			MaxOpenConn: c.Db.MaxOpenConn,
			MaxIdleConn: c.Db.MaxIdleConn,
		}),
	}
}
