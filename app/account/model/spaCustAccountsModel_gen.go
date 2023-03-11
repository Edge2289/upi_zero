// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	spaCustAccountsFieldNames          = builder.RawFieldNames(&SpaCustAccounts{})
	spaCustAccountsRows                = strings.Join(spaCustAccountsFieldNames, ",")
	spaCustAccountsRowsExpectAutoSet   = strings.Join(stringx.Remove(spaCustAccountsFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	spaCustAccountsRowsWithPlaceHolder = strings.Join(stringx.Remove(spaCustAccountsFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"

	cacheSpaCustAccountsIdPrefix = "cache:spaCustAccounts:id:"
)

type (
	spaCustAccountsModel interface {
		Insert(ctx context.Context, data *SpaCustAccounts) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SpaCustAccounts, error)
		Update(ctx context.Context, data *SpaCustAccounts) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSpaCustAccountsModel struct {
		sqlc.CachedConn
		table string
	}

	SpaCustAccounts struct {
		Id         int64        `db:"id"`
		SubAcctNo  string       `db:"sub_acct_no"` // 见证功能子账号
		MemberName string       `db:"member_name"` // 见证功能子账名
		CreatedAt  sql.NullTime `db:"created_at"`
		UpdatedAt  sql.NullTime `db:"updated_at"`
	}
)

func newSpaCustAccountsModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSpaCustAccountsModel {
	return &defaultSpaCustAccountsModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`spa_cust_accounts`",
	}
}

func (m *defaultSpaCustAccountsModel) Delete(ctx context.Context, id int64) error {
	spaCustAccountsIdKey := fmt.Sprintf("%s%v", cacheSpaCustAccountsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, spaCustAccountsIdKey)
	return err
}

func (m *defaultSpaCustAccountsModel) FindOne(ctx context.Context, id int64) (*SpaCustAccounts, error) {
	spaCustAccountsIdKey := fmt.Sprintf("%s%v", cacheSpaCustAccountsIdPrefix, id)
	var resp SpaCustAccounts
	err := m.QueryRowCtx(ctx, &resp, spaCustAccountsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", spaCustAccountsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSpaCustAccountsModel) Insert(ctx context.Context, data *SpaCustAccounts) (sql.Result, error) {
	spaCustAccountsIdKey := fmt.Sprintf("%s%v", cacheSpaCustAccountsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, spaCustAccountsRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SubAcctNo, data.MemberName)
	}, spaCustAccountsIdKey)
	return ret, err
}

func (m *defaultSpaCustAccountsModel) Update(ctx context.Context, data *SpaCustAccounts) error {
	spaCustAccountsIdKey := fmt.Sprintf("%s%v", cacheSpaCustAccountsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, spaCustAccountsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.SubAcctNo, data.MemberName, data.Id)
	}, spaCustAccountsIdKey)
	return err
}

func (m *defaultSpaCustAccountsModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSpaCustAccountsIdPrefix, primary)
}

func (m *defaultSpaCustAccountsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", spaCustAccountsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSpaCustAccountsModel) tableName() string {
	return m.table
}