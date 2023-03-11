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
	spaAccountHistoriesFieldNames          = builder.RawFieldNames(&SpaAccountHistories{})
	spaAccountHistoriesRows                = strings.Join(spaAccountHistoriesFieldNames, ",")
	spaAccountHistoriesRowsExpectAutoSet   = strings.Join(stringx.Remove(spaAccountHistoriesFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), ",")
	spaAccountHistoriesRowsWithPlaceHolder = strings.Join(stringx.Remove(spaAccountHistoriesFieldNames, "`id`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`", "`create_at`"), "=?,") + "=?"

	cacheSpaAccountHistoriesIdPrefix = "cache:spaAccountHistories:id:"
)

type (
	spaAccountHistoriesModel interface {
		Insert(ctx context.Context, data *SpaAccountHistories) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SpaAccountHistories, error)
		Update(ctx context.Context, data *SpaAccountHistories) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSpaAccountHistoriesModel struct {
		sqlc.CachedConn
		table string
	}

	SpaAccountHistories struct {
		Id                 int64        `db:"id"`
		SubAcctNo          string       `db:"sub_acct_no"`           // 见证子账户的账号
		TranNetMemberCode  string       `db:"tran_net_member_code"`  // 交易网会员代码
		MemberName         string       `db:"member_name"`           // 会员名称
		MemberGlobalType   string       `db:"member_global_type"`    // 会员证件类型 1:身份证 3:港澳台居民通行证（即回乡证） 4:中国护照 5:台湾居民来往大陆通行证（即台胞证） 19:外国护照 52:组织机构代码证 68:营业执照 73:统一社会信用代码
		MemberGlobalId     string       `db:"member_global_id"`      // 会员证件号码
		UserNickname       string       `db:"user_nickname"`         // 用户昵称
		Mobile             string       `db:"mobile"`                // 手机号码
		MemberProperty     string       `db:"member_property"`       // 会员属性 SH-商户子账户（默认）
		ReservedMsg        string       `db:"reserved_msg"`          // 保留域
		Status             string       `db:"status"`                // 状态：1:开立成功
		PassTime           string       `db:"pass_time"`             // 账户开立时间
		Type               string       `db:"type"`                  // 账户类型 1:经销商 2:供应商 3:经销商和供应商
		EiconBankBranchId  string       `db:"eicon_bank_branch_id"`  // 会员绑定账户的开户行的超级网银行号
		MemberAcctNo       string       `db:"member_acct_no"`        // 会员绑定账户的账号 提现的银行卡
		CnapsBranchId      string       `db:"cnaps_branch_id"`       // 会员绑定账户的开户行的联行号
		AcctOpenBranchName string       `db:"acct_open_branch_name"` // 会员绑定账户的开户行名称
		BankType           string       `db:"bank_type"`             // 会员绑定账户的本他行类型 1：本行 2：他行
		DeleteTime         sql.NullTime `db:"delete_time"`           // 删除时间
		CreatedAt          sql.NullTime `db:"created_at"`
		UpdatedAt          sql.NullTime `db:"updated_at"`
	}
)

func newSpaAccountHistoriesModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSpaAccountHistoriesModel {
	return &defaultSpaAccountHistoriesModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`spa_account_histories`",
	}
}

func (m *defaultSpaAccountHistoriesModel) Delete(ctx context.Context, id int64) error {
	spaAccountHistoriesIdKey := fmt.Sprintf("%s%v", cacheSpaAccountHistoriesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, spaAccountHistoriesIdKey)
	return err
}

func (m *defaultSpaAccountHistoriesModel) FindOne(ctx context.Context, id int64) (*SpaAccountHistories, error) {
	spaAccountHistoriesIdKey := fmt.Sprintf("%s%v", cacheSpaAccountHistoriesIdPrefix, id)
	var resp SpaAccountHistories
	err := m.QueryRowCtx(ctx, &resp, spaAccountHistoriesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", spaAccountHistoriesRows, m.table)
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

func (m *defaultSpaAccountHistoriesModel) Insert(ctx context.Context, data *SpaAccountHistories) (sql.Result, error) {
	spaAccountHistoriesIdKey := fmt.Sprintf("%s%v", cacheSpaAccountHistoriesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, spaAccountHistoriesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.SubAcctNo, data.TranNetMemberCode, data.MemberName, data.MemberGlobalType, data.MemberGlobalId, data.UserNickname, data.Mobile, data.MemberProperty, data.ReservedMsg, data.Status, data.PassTime, data.Type, data.EiconBankBranchId, data.MemberAcctNo, data.CnapsBranchId, data.AcctOpenBranchName, data.BankType, data.DeleteTime)
	}, spaAccountHistoriesIdKey)
	return ret, err
}

func (m *defaultSpaAccountHistoriesModel) Update(ctx context.Context, data *SpaAccountHistories) error {
	spaAccountHistoriesIdKey := fmt.Sprintf("%s%v", cacheSpaAccountHistoriesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, spaAccountHistoriesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.SubAcctNo, data.TranNetMemberCode, data.MemberName, data.MemberGlobalType, data.MemberGlobalId, data.UserNickname, data.Mobile, data.MemberProperty, data.ReservedMsg, data.Status, data.PassTime, data.Type, data.EiconBankBranchId, data.MemberAcctNo, data.CnapsBranchId, data.AcctOpenBranchName, data.BankType, data.DeleteTime, data.Id)
	}, spaAccountHistoriesIdKey)
	return err
}

func (m *defaultSpaAccountHistoriesModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheSpaAccountHistoriesIdPrefix, primary)
}

func (m *defaultSpaAccountHistoriesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", spaAccountHistoriesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSpaAccountHistoriesModel) tableName() string {
	return m.table
}