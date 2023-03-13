package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"upi/app/account/com/rpc/account"
	"upi/app/account/model"
	"upi/app/business/com/rpc/internal/svc"
	"upi/app/business/com/rpc/pd"
	"upi/app/business/model/capitalManager"
	"upi/common/xerr"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pd.SplitBillsSummaryListReq) (*pd.SplitBillsSummaryListResp, error) {

	_, err := l.svcCtx.AccountRpc.GetSubAcctNoPassUid(l.ctx, &account.GetSubAcctNoPassUidReq{
		Uid: in.UserId,
	})
	if err != nil {
		return nil, err
	}

	var data []*pd.SplitBillsSummaryListMap
	// capitalManager.SplitBillsSummary{} Where("buyer_sub_acct_no in (?)", subAcctNo).
	err = l.svcCtx.DbEngine.Model(&capitalManager.SplitBillsSummary{}).Select("*").Find(&data).Error
	if err != nil && err != model.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &pd.SplitBillsSummaryListResp{
		List: data,
	}, nil
}
