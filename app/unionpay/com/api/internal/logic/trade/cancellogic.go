package trade

import (
	"context"

	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelLogic {
	return &CancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelLogic) Cancel(req *types.CancelReq) (resp *types.CancelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
