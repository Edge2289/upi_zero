package notify

import (
	"context"

	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TradeNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTradeNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TradeNotifyLogic {
	return &TradeNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TradeNotifyLogic) TradeNotify(req *types.NotifyTradeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
