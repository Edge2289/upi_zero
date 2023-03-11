package logic

import (
	"context"

	"upi/app/unionpay/com/rpc/internal/svc"
	"upi/app/unionpay/com/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SavePayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSavePayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SavePayOrderLogic {
	return &SavePayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SavePayOrderLogic) SavePayOrder(in *pb.SavePayOrderReq) (*pb.SavePayOrderResp, error) {
	// todo: add your logic here and delete this line
	logx.Info("xxxxx-xxxxxxxxxaaaaaaa")
	return &pb.SavePayOrderResp{}, nil
}
