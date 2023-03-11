package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaAccountBindingsModel = (*customSpaAccountBindingsModel)(nil)

type (
	// SpaAccountBindingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaAccountBindingsModel.
	SpaAccountBindingsModel interface {
		spaAccountBindingsModel
	}

	customSpaAccountBindingsModel struct {
		*defaultSpaAccountBindingsModel
	}
)

// NewSpaAccountBindingsModel returns a model for the database table.
func NewSpaAccountBindingsModel(conn sqlx.SqlConn, c cache.CacheConf) SpaAccountBindingsModel {
	return &customSpaAccountBindingsModel{
		defaultSpaAccountBindingsModel: newSpaAccountBindingsModel(conn, c),
	}
}
