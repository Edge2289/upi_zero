package trade

import (
	"context"
	"upi/app/unionpay/com/rpc/pb"

	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.QueryReq) (resp *types.QueryResp, err error) {
	respa, err := l.svcCtx.Greet.Greet(l.ctx, &pb.StreamReq{
		Name: req.TradeNo,
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryResp {
		Result: respa.Greet,
	}, nil
}
