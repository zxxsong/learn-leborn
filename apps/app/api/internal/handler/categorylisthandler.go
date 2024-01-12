package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"learn-lebron/apps/app/api/internal/logic"
	"learn-lebron/apps/app/api/internal/svc"
	"learn-lebron/apps/app/api/internal/types"
)

func CategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CategoryListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.CategoryList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
