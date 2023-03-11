package logic

import (
	"context"
	"fmt"
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
	logx.Info("-------------x------")
	subAcctNo, err := l.svcCtx.AccountRpc.GetSubAcctNoPassUid(l.ctx, &account.GetSubAcctNoPassUidReq{
		Uid: in.UserId,
	})
	if err != nil {
		fmt.Println("---")
		return nil, err
	}

	var data []*pd.SplitBillsSummaryListMap

	fmt.Println("subAcctNo", subAcctNo)
	// capitalManager.SplitBillsSummary{} Where("buyer_sub_acct_no in (?)", subAcctNo).
	err = l.svcCtx.DbEngine.Model(&capitalManager.SplitBillsSummary{}).Select("*").Find(&data).Error
	logx.Info("data", data)
	if err != nil && err != model.ErrNotFound {
		return nil, xerr.NewErrCode(xerr.DB_ERROR)
	}

	return &pd.SplitBillsSummaryListResp{
		List: data,
	}, nil
}
