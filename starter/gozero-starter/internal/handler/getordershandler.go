package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	gozeroauth "mora/adapters/gozero"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func GetOrdersHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := gozeroauth.GetUserID(r.Context())

		// Mock orders data - in production, query from database
		orders := []types.Order{
			{ID: "order-1", UserID: userID, Amount: 100.00, Status: "completed"},
			{ID: "order-2", UserID: userID, Amount: 250.50, Status: "pending"},
		}

		resp := &types.OrdersResponse{
			Orders: orders,
			Total:  len(orders),
		}

		httpx.OkJson(w, resp)
	}
}