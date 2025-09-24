package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	gozeroauth "mora/adapters/gozero"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func ProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := gozeroauth.GetUserID(r.Context())
		claims := gozeroauth.GetClaims(r.Context())

		if claims == nil {
			httpx.WriteJson(w, http.StatusInternalServerError, map[string]string{
				"error": "failed to get user claims",
			})
			return
		}

		resp := &types.ProfileResponse{
			UserID:   userID,
			Username: claims.Username,
			Subject:  claims.Subject,
			Exp:      claims.ExpiresAt.Time.Format(time.RFC3339),
			Iat:      claims.IssuedAt.Time.Format(time.RFC3339),
		}

		httpx.OkJson(w, resp)
	}
}