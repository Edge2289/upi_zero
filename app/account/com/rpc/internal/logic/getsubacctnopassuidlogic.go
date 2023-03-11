package logic

import (
	"context"
	"fmt"

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

	fmt.Println("==============")

	model := l.svcCtx.SpaAccountBindingsModel.RowBuilder()
	model = model.Where("user_id", in.Uid)
	data, err := l.svcCtx.SpaAccountBindingsModel.FindAll(l.ctx, model)
	if err != nil {
		return nil, err
	}
	var subAcctNo []string

	for _, c := range data {
		subAcctNo = append(subAcctNo, c.SubAcctNo)
	}

	return &pb.GetSubAcctNoPassUidResp{
		SubAcctNo: subAcctNo,
	}, nil
}
