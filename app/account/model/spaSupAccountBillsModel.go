package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaSupAccountBillsModel = (*customSpaSupAccountBillsModel)(nil)

type (
	// SpaSupAccountBillsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaSupAccountBillsModel.
	SpaSupAccountBillsModel interface {
		spaSupAccountBillsModel
	}

	customSpaSupAccountBillsModel struct {
		*defaultSpaSupAccountBillsModel
	}
)

// NewSpaSupAccountBillsModel returns a model for the database table.
func NewSpaSupAccountBillsModel(conn sqlx.SqlConn, c cache.CacheConf) SpaSupAccountBillsModel {
	return &customSpaSupAccountBillsModel{
		defaultSpaSupAccountBillsModel: newSpaSupAccountBillsModel(conn, c),
	}
}
