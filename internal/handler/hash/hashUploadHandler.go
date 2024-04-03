package hash

import (
	"net/http"

	"connection/internal/logic/hash"
	"connection/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func HashUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := hash.NewHashUploadLogic(r.Context(), svcCtx)
		resp, err := l.HashUpload()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
