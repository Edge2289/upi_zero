package logic

import (
	"context"

	"upi/app/unionpay/com/rpc/internal/svc"
	"upi/app/unionpay/com/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GreetLogic) Greet(in *pb.StreamReq) (*pb.StreamResp, error) {
	// todo: add your logic here and delete this line

	_ = l.svcCtx.KqueuePay.Push(in.Name)
	logx.Info("----------haaaaa-----------")
	logx.Error("----------haaaaa-----------")
	return &pb.StreamResp{Greet: in.Name}, nil
}
