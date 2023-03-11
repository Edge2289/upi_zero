package trade

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/unionpay/com/api/internal/logic/trade"
	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"
)

func CancelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CancelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := trade.NewCancelLogic(r.Context(), svcCtx)
		resp, err := l.Cancel(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
