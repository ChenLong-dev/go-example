package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hello01/internal/logic"
	"hello01/internal/svc"
	"hello01/internal/types"
)

func Hello01Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewHello01Logic(r.Context(), svcCtx)
		resp, err := l.Hello01(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
