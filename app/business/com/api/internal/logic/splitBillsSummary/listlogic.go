package splitBillsSummary

import (
	"context"
	"github.com/jinzhu/copier"
	"upi/app/business/com/api/internal/svc"
	"upi/app/business/com/api/internal/types"
	"upi/app/business/com/rpc/splitbillssummary"

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

	//subAcctNos := ctxdata.GetUidFromSubAcctNoCtx(l.ctx)
	respData, err := l.svcCtx.SplitBillsSummaryRpc.List(l.ctx, &splitbillssummary.SplitBillsSummaryListReq{

	})
	if err != nil {
		return nil, err
	}

	//
	// 格式转换
	//
	var respMap []types.SplitBillsSummaryListmap

	if len(respData.List) > 0 {
		for r := range respData.List {
			var b types.SplitBillsSummaryListmap
			_ = copier.Copy(r, b)

			respMap = append(respMap, b)
		}
	}

	return &types.SplitBillsSummaryListResp{
		List: respMap,
	}, nil
}
