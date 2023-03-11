package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"upi/app/account/com/rpc/internal/config"
	"upi/app/account/model"
)

type ServiceContext struct {
	Config config.Config

	SpaAccountModel           model.SpaAccountsModel
	SpaAccountBindingsModel   model.SpaAccountBindingsModel
	SpaAccountOaEntitiesModel model.SpaAccountOaEntitiesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:                    c,
		SpaAccountModel:           model.NewSpaAccountsModel(sqlConn, c.Cache),
		SpaAccountBindingsModel:   model.NewSpaAccountBindingsModel(sqlConn, c.Cache),
		SpaAccountOaEntitiesModel: model.NewSpaAccountOaEntitiesModel(sqlConn, c.Cache),
	}
}
