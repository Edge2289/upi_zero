package splitBillsSummary

import (
	"context"

	"upi/app/business/com/api/internal/svc"
	"upi/app/business/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.SplitBillsSummaryListReq) (resp *types.SplitBillsSummaryListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
