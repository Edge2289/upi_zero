package splitBillsSummary

import (
	"context"

	"upi/app/business/com/api/internal/svc"
	"upi/app/business/com/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExportLogic {
	return &ExportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExportLogic) Export(req *types.SplitBillsSummaryListReq) (resp *types.NullResp, err error) {
	// todo: add your logic here and delete this line

	return
}
