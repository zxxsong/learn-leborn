package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"learn-lebron/apps/app/api/internal/logic"
	"learn-lebron/apps/app/api/internal/svc"
)

func HomeBannerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHomeBannerLogic(r.Context(), svcCtx)
		resp, err := l.HomeBanner()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
