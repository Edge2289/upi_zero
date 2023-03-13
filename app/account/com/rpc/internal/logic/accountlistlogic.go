package logic

import (
	"context"
	"upi/app/account/com/rpc/internal/svc"
	"upi/app/account/com/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccountListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAccountListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccountListLogic {
	return &AccountListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AccountListLogic) AccountList(in *pb.AccountListReq) (*pb.AccountListResp, error) {

	// 拼接搜索条件
	//  limit 20,20 --- 分页
	//rowBuilder := l.svcCtx.SpaAccountModel.RowBuilder()
	//if in.GetMemberName() != "" {
	//	rowBuilder = rowBuilder.Where("member_name = ?", in.GetMemberName())
	//}
	//
	//if in.GetSubAcctNo() != "" {
	//	rowBuilder = rowBuilder.Where("sub_acct_no = ?", in.GetSubAcctNo())
	//}
	//
	//if in.GetStartTime() != "" {
	//	rowBuilder = rowBuilder.Where("start_time = ?", in.GetStartTime())
	//}
	//
	//if in.GetEndTime() != "" {
	//	rowBuilder = rowBuilder.Where("start_time = ?", in.GetEndTime())
	//}
	//
	//
	//// limit = per_page, offset (page_number-1) * per_page
	//rowBuilder = rowBuilder.Limit(uint64(in.GetPerPage())).Offset(uint64((in.Page - 1) * in.GetPerPage())).OrderBy("id DESC")
	//
	//list, err := l.svcCtx.SpaAccountModel.FindAll(l.ctx, rowBuilder)
	//if err != nil {
	//	return nil, err
	//}
	var accountList []*pb.AccountData
	//if list != nil {
	//	_ = copier.Copy(&accountList, list)
	//}
	return &pb.AccountListResp{
		List: accountList,
	}, nil
}
