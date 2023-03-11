package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaAccountBindingsModel = (*customSpaAccountBindingsModel)(nil)

type (
	// SpaAccountBindingsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaAccountBindingsModel.
	SpaAccountBindingsModel interface {
		spaAccountBindingsModel

		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SpaAccountBindings, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customSpaAccountBindingsModel struct {
		*defaultSpaAccountBindingsModel
	}
)

func (c *customSpaAccountBindingsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SpaAccountBindings, error) {
	panic("implement me")
}

func (c *customSpaAccountBindingsModel) RowBuilder() squirrel.SelectBuilder {
	panic("implement me")
}

// NewSpaAccountBindingsModel returns a model for the database table.
func NewSpaAccountBindingsModel(conn sqlx.SqlConn, c cache.CacheConf) SpaAccountBindingsModel {
	return &customSpaAccountBindingsModel{
		defaultSpaAccountBindingsModel: newSpaAccountBindingsModel(conn, c),
	}
}
