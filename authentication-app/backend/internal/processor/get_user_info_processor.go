package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func GetUserInfoProcessor(ctx context.Context, req, resp interface{}) error {
	request, ok := req.(*vo.GetUserInfoRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.GetUserInfoResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	userInfoDM := model.NewUserInfoDM(ctx)
	h := handler.NewUserInfoHandler(ctx, userInfoDM)

	userInfo, err := h.GetUserInfo(*request.UserID)
	if err != nil {
		return err
	}

	response.UserInfo = userInfo

	return nil
}
