package spaAccount

import (
	"net/http"
	"upi/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"upi/app/account/com/api/internal/logic/spaAccount"
	"upi/app/account/com/api/internal/svc"
	"upi/app/account/com/api/internal/types"
)

// AccountListHandler 子账户列表
func AccountListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AccountListReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := spaAccount.NewAccountListLogic(r.Context(), svcCtx)
		resp, err := l.AccountList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
