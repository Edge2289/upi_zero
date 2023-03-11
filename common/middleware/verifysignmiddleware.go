package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type VerifySignMiddleware struct {
}

func NewVerifySignMiddleware() *VerifySignMiddleware {
	return &VerifySignMiddleware{}
}

func (m *VerifySignMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("---------xxxxx--------")
		// Passthrough to next handler if need
		next(w, r)
	}
}
