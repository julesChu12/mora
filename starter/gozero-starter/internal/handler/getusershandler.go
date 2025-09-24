package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	gozeroauth "mora/adapters/gozero"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func GetUsersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := gozeroauth.GetUserID(r.Context())

		// Mock users data
		users := []types.User{
			{ID: "user-123", Username: "admin", Role: "admin"},
			{ID: "user-456", Username: "user1", Role: "user"},
		}

		resp := &types.UsersResponse{
			Users:     users,
			Total:     len(users),
			RequestBy: userID,
		}

		httpx.OkJson(w, resp)
	}
}