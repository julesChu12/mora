package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	gozeroauth "mora/adapters/gozero"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func ProtectedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := gozeroauth.GetUserID(r.Context())

		resp := &types.ProtectedResponse{
			Message: "This is a protected endpoint",
			UserID:  userID,
			Time:    time.Now().Format(time.RFC3339),
		}

		httpx.OkJson(w, resp)
	}
}