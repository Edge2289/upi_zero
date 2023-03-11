package trade

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/unionpay/com/api/internal/logic/trade"
	"upi/app/unionpay/com/api/internal/svc"
)

func ModesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := trade.NewModesLogic(r.Context(), svcCtx)
		resp, err := l.Modes()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
