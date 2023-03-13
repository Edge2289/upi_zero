package logic

import (
	"context"
	"upi/app/account/model"

	"upi/app/account/com/rpc/internal/svc"
	"upi/app/account/com/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSubAcctNoPassUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSubAcctNoPassUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSubAcctNoPassUidLogic {
	return &GetSubAcctNoPassUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetSubAcctNoPassUid 通过uid获取子账户列表
func (l *GetSubAcctNoPassUidLogic) GetSubAcctNoPassUid(in *pb.GetSubAcctNoPassUidReq) (*pb.GetSubAcctNoPassUidResp, error) {

	var info []model.SpaAccountBindingsInfo
	err := l.svcCtx.DbEngine.WithContext(l.ctx).Model(&info).Where("user_id = ?", in.Uid).Select("sub_acct_no").Group("sub_acct_no").Find(&info).Error
	if err != nil {
		return nil, err
	}
	var subAcctNo []string
	for _, c := range info {
		subAcctNo = append(subAcctNo, c.SubAcctNo)
	}

	return &pb.GetSubAcctNoPassUidResp{
		SubAcctNo: subAcctNo,
	}, nil
}
