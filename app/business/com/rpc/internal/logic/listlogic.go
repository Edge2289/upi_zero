package logic

import (
	"context"

	"upi/app/business/com/rpc/internal/svc"
	"upi/app/business/com/rpc/pd"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *pd.SplitBillsSummaryListReq) (*pd.SplitBillsSummaryListResp, error) {
	// todo: add your logic here and delete this line

	return &pd.SplitBillsSummaryListResp{}, nil
}
