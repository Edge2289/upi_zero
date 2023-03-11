package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaAccountHistoriesModel = (*customSpaAccountHistoriesModel)(nil)

type (
	// SpaAccountHistoriesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaAccountHistoriesModel.
	SpaAccountHistoriesModel interface {
		spaAccountHistoriesModel
	}

	customSpaAccountHistoriesModel struct {
		*defaultSpaAccountHistoriesModel
	}
)

// NewSpaAccountHistoriesModel returns a model for the database table.
func NewSpaAccountHistoriesModel(conn sqlx.SqlConn, c cache.CacheConf) SpaAccountHistoriesModel {
	return &customSpaAccountHistoriesModel{
		defaultSpaAccountHistoriesModel: newSpaAccountHistoriesModel(conn, c),
	}
}
