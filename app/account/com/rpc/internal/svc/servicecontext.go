package svc

import (
	"gorm.io/gorm"
	"upi/app/account/com/rpc/internal/config"
	"upi/common/resources/db"
)

type ServiceContext struct {
	Config config.Config

	DbEngine *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config: c,

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
