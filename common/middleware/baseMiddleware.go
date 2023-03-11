package middleware

import (
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"upi/common/result"
	"upi/common/xerr"
)

type NotFound struct {
}

// NewServer returns a new Server using the given operations.
func NewServer() *NotFound {
	return &NotFound{}
}
func (h NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(h))
	result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.NOT_ROUTE))
}

func BaseMiddleware() []rest.RunOption {
	var option []rest.RunOption
	// 不存在的路由
	option = append(option, rest.WithNotFoundHandler(NewServer()))
	// 不存在的路由
	option = append(option, rest.WithNotAllowedHandler(NewServer()))
	// 登陆认证失败
	option = append(option, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		result.HttpResult(r, w, nil, xerr.NewErrCode(xerr.TOKEN_EXPIRE_ERROR))
	}))

	return option
}
