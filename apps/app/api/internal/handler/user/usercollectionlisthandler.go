package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"learn-lebron/apps/app/api/internal/logic/user"
	"learn-lebron/apps/app/api/internal/svc"
	"learn-lebron/apps/app/api/internal/types"
)

func UserCollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCollectionListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUserCollectionListLogic(r.Context(), svcCtx)
		resp, err := l.UserCollectionList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
