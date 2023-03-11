package splitBillsSummary

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"upi/app/business/com/api/internal/svc"
	"upi/app/business/com/api/internal/types"
	"upi/app/business/com/rpc/splitbillssummary"
	"upi/common/ctxdata"

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

// List 列表查询
func (l *ListLogic) List(req *types.SplitBillsSummaryListReq) (resp *types.SplitBillsSummaryListResp, err error) {

	uid := ctxdata.GetUidFromCtx(l.ctx)

	respData, err := l.svcCtx.SplitBillsSummaryRpc.List(l.ctx, &splitbillssummary.SplitBillsSummaryListReq{
		Page: req.Page,
		PerPage: req.PerPage,
		TheMonthStart: req.TheMonthStart,
		TheMonthEnd: req.TheMonthEnd,
		BuyerCompanyName: req.BuyerCompanyName,
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	//
	// 格式转换
	//
	var respMap []types.SplitBillsSummaryListmap

	if len(respData.List) > 0 {
		for _, r := range respData.List {
			fmt.Printf(r.BuyerSubAcctNo)
			var b types.SplitBillsSummaryListmap
			err = copier.Copy(r, b)
			if err != nil {
				fmt.Println("-------------err-------------", err.Error())
			}
			fmt.Println("b", b)

			respMap = append(respMap, b)
		}
	}

	return &types.SplitBillsSummaryListResp{
		List: respMap,
	}, nil
}
