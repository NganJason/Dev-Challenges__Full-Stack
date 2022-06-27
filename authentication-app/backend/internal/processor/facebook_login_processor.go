package processor

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/facebook"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func FacebookLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.FacebookLoginRequest)
	response := resp.(*vo.FacebookLoginResponse)

	if request.AccessCode == nil {
		return cerr.New("access_token cannot be empty", http.StatusBadRequest)
	}

	userAuthDM := model.NewUserAuthDM(ctx)
	userInfoDM := model.NewUserInfoDM(ctx)

	h := handler.NewAuthHandler(ctx, userAuthDM, userInfoDM)
	h.SetFacebookService(facebook.NewFacebookService(ctx))

	userInfo, err := h.LoginFacebook(request.AccessCode)
	if err != nil {
		return err
	}

	err = util.GenerateTokenAndAddCookies(ctx, strconv.Itoa(int(*userInfo.UserID)))
	if err != nil {
		return cerr.New(err.Error(), http.StatusBadGateway)
	}

	response.UserInfo = userInfo

	return nil
}
