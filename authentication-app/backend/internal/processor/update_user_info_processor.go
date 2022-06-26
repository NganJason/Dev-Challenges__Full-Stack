package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func UpdateUserInfoProcessor(ctx context.Context, req, resp interface{}) error {
	request, ok := req.(*vo.UpdateUserInfoRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.UpdateUserInfoResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	dm := model.NewUserInfoDM(ctx)
	h := handler.NewUserInfoHandler(ctx, dm)

	userInfo, err := h.UpdateUserInfo(request)
	if err != nil {
		return err
	}

	response.UserInfo = userInfo

	return nil
}
