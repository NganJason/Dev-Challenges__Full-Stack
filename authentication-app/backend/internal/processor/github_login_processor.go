package processor

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/github"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func GithubLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.GithubLoginRequest)
	response := resp.(*vo.GithubLoginResponse)

	userAuthDM := model.NewUserAuthDM(ctx)
	userInfoDM := model.NewUserInfoDM(ctx)

	h := handler.NewAuthHandler(ctx, userAuthDM, userInfoDM)
	h.SetGithubService(github.NewGithubService(ctx))

	userInfo, err := h.LoginGithub(request.AccessCode)
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
