package trade

import (
	"fmt"
	"net/http"
	"upi/common/result"

	"upi/app/unionpay/com/api/internal/logic/trade"
	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PayReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		fmt.Println("--------------------------------")
		l := trade.NewPayLogic(r.Context(), svcCtx)
		resp, err := l.Pay(&req)
		fmt.Println("--------------------------------")
		result.HttpResult(r, w, resp, err)
	}
}
