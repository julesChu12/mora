package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mora/pkg/auth"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// Mock authentication - in production, validate against UserService
		if req.Username == "admin" && req.Password == "password" {
			// Generate access token using Mora auth
			tokenTTL := time.Duration(svcCtx.Config.JWT.TTL) * time.Second
			token, err := auth.GenerateToken("user-123", req.Username, svcCtx.Config.JWT.Secret, tokenTTL)
			if err != nil {
				httpx.Error(w, err)
				return
			}

			resp := &types.LoginResponse{
				AccessToken: token,
				TokenType:   "Bearer",
				ExpiresIn:   int(tokenTTL.Seconds()),
				UserID:      "user-123",
				Username:    req.Username,
			}

			httpx.OkJson(w, resp)
			return
		}

		// Authentication failed
		httpx.WriteJson(w, http.StatusUnauthorized, map[string]string{
			"error":   "authentication failed",
			"message": "invalid username or password",
		})
	}
}