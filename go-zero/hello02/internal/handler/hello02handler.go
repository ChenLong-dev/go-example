package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hello02/internal/logic"
	"hello02/internal/svc"
	"hello02/internal/types"
)

func Hello02Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHello02Logic(r.Context(), svcCtx)
		resp, err := l.Hello02(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
