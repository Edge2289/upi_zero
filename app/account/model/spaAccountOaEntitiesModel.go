package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaAccountOaEntitiesModel = (*customSpaAccountOaEntitiesModel)(nil)

type (
	// SpaAccountOaEntitiesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaAccountOaEntitiesModel.
	SpaAccountOaEntitiesModel interface {
		spaAccountOaEntitiesModel
	}

	customSpaAccountOaEntitiesModel struct {
		*defaultSpaAccountOaEntitiesModel
	}
)

// NewSpaAccountOaEntitiesModel returns a model for the database table.
func NewSpaAccountOaEntitiesModel(conn sqlx.SqlConn, c cache.CacheConf) SpaAccountOaEntitiesModel {
	return &customSpaAccountOaEntitiesModel{
		defaultSpaAccountOaEntitiesModel: newSpaAccountOaEntitiesModel(conn, c),
	}
}
