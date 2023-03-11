package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaCustAccountsModel = (*customSpaCustAccountsModel)(nil)

type (
	// SpaCustAccountsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaCustAccountsModel.
	SpaCustAccountsModel interface {
		spaCustAccountsModel
	}

	customSpaCustAccountsModel struct {
		*defaultSpaCustAccountsModel
	}
)

// NewSpaCustAccountsModel returns a model for the database table.
func NewSpaCustAccountsModel(conn sqlx.SqlConn, c cache.CacheConf) SpaCustAccountsModel {
	return &customSpaCustAccountsModel{
		defaultSpaCustAccountsModel: newSpaCustAccountsModel(conn, c),
	}
}
