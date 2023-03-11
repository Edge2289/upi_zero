package spaAccount

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"upi/app/account/com/api/internal/svc"
	"upi/app/account/com/api/internal/types"
	"upi/app/account/com/rpc/pb"
)

type AccountListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountListLogic {
	return &AccountListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AccountListLogic) AccountList(req *types.AccountListReq) (resp *types.AccountListResps, err error) {
	data, err := l.svcCtx.AccountRpc.AccountList(l.ctx, &pb.AccountListReq{
		Page:       int64(req.Page),
		PerPage:    int64(req.PerPage),
		SubAcctNo:  req.SubAcctNo,
		MemberName: req.MemberName,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
	})
	if err != nil {
		return nil, err
	}

	var accountList []types.AccountListResp
	if data != nil {
		for _, value := range data.GetList() {
			//homestay := value.(*model.SpaAccounts)
			var tyHomestay types.AccountListResp
			_ = copier.Copy(&tyHomestay, value)
			accountList = append(accountList, tyHomestay)
		}
	}
	return &types.AccountListResps{
		Data: accountList,
	}, nil
}
