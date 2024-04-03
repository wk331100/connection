package middleware

import (
	"connection/internal"
	"connection/internal/types"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RequestHeaderInterceptorMiddleware struct {
}

func NewRequestHeaderInterceptorMiddleware() *RequestHeaderInterceptorMiddleware {
	return &RequestHeaderInterceptorMiddleware{}
}

func (m *RequestHeaderInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 将 header 中 requestId 传递到 context 中，方便 logic 那边获取
		requestId := r.Header.Get(internal.XRequestID)
		if requestId == "" {
			resp := types.BaseResponse{
				Code: 400,
				Msg:  "not specified header X-Request-ID",
			}
			httpx.OkJsonCtx(r.Context(), w, resp)
		} else {
			r = r.WithContext(context.WithValue(r.Context(), internal.XRequestID, requestId)) // nolint:staticcheck
			next(w, r)
		}
	}
}
