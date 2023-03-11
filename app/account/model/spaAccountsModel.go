package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SpaAccountsModel = (*customSpaAccountsModel)(nil)

type (
	// SpaAccountsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSpaAccountsModel.
	SpaAccountsModel interface {
		spaAccountsModel
		customModel
	}


	customModel interface {
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SpaAccounts, error)
		RowBuilder() squirrel.SelectBuilder
	}

	customSpaAccountsModel struct {
		*defaultSpaAccountsModel
	}
)

// NewSpaAccountsModel returns a model for the database table.
func NewSpaAccountsModel(conn sqlx.SqlConn, c cache.CacheConf) SpaAccountsModel {
	return &customSpaAccountsModel{
		defaultSpaAccountsModel: newSpaAccountsModel(conn, c),
	}
}

func (m *defaultSpaAccountsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(spaAccountsRows).From(m.tableName())
}

// FindAll 获取所有的列表数据（带搜索条件）
func (m *defaultSpaAccountsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*SpaAccounts, error) {

	querySql, queryData, err := rowBuilder.From(m.tableName()).ToSql()
	if err != nil {
		return nil, err
	}

	var res []*SpaAccounts
	err = m.QueryRowsNoCacheCtx(ctx, &res, querySql, queryData...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
