package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func HealthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &types.HealthResponse{
			Status: "ok",
			Time:   time.Now().Format(time.RFC3339),
		}

		httpx.OkJson(w, resp)
	}
}