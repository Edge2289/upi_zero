package trade

import (
	"context"

	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ModesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModesLogic {
	return &ModesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModesLogic) Modes() (resp *types.ModesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
