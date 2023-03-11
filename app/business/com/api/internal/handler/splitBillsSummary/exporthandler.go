package splitBillsSummary

import (
	"net/http"
	"upi/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/business/com/api/internal/logic/splitBillsSummary"
	"upi/app/business/com/api/internal/svc"
	"upi/app/business/com/api/internal/types"
)

func ExportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SplitBillsSummaryListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := splitBillsSummary.NewExportLogic(r.Context(), svcCtx)
		resp, err := l.Export(&req)

		result.HttpResult(r, w, resp, err)
	}
}
