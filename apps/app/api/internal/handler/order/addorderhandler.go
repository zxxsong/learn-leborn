package order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"learn-lebron/apps/app/api/internal/logic/order"
	"learn-lebron/apps/app/api/internal/svc"
	"learn-lebron/apps/app/api/internal/types"
)

func AddOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewAddOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
