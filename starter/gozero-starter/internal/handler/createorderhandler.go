package handler

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
	gozeroauth "mora/adapters/gozero"
	"mora/starter/gozero-starter/internal/svc"
	"mora/starter/gozero-starter/internal/types"
)

func CreateOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		userID := gozeroauth.GetUserID(r.Context())

		// Mock order creation
		order := types.Order{
			ID:     "order-" + time.Now().Format("20060102150405"),
			UserID: userID,
			Amount: req.Amount,
			Status: "created",
		}

		resp := &types.CreateOrderResponse{
			Order: order,
		}

		httpx.WriteJson(w, http.StatusCreated, resp)
	}
}