package notify

import (
	"net/http"
	"upi/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/unionpay/com/api/internal/logic/notify"
	"upi/app/unionpay/com/api/internal/svc"
	"upi/app/unionpay/com/api/internal/types"
)

func TradeNotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyTradeReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := notify.NewTradeNotifyLogic(r.Context(), svcCtx)
		err := l.TradeNotify(&req)
		result.HttpResult(r, w, nil, err)
	}
}
