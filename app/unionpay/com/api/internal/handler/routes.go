// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	notify "upi/app/unionpay/com/api/internal/handler/notify"
	trade "upi/app/unionpay/com/api/internal/handler/trade"
	"upi/app/unionpay/com/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/trade/pay",
				Handler: trade.PayHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/trade/query",
				Handler: trade.QueryHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/trade/modes",
				Handler: trade.ModesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/trade/cancel",
				Handler: trade.CancelHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/unionpay/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/notify/trade",
				Handler: notify.TradeNotifyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/unionpay/v1"),
	)
}
