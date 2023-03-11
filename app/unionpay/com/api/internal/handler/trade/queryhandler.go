package trade

import (
	"fmt"
	"net/http"
	"upi/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/unionpay/com/api/internal/logic/trade"
	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"
)

func QueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryReq
		fmt.Println("xxxxxx")
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println("xxxxxx12312312", req.TradeNo)

		l := trade.NewQueryLogic(r.Context(), svcCtx)
		resp, err := l.Query(&req)
		result.HttpResult(r, w, resp, err)
	}
}
