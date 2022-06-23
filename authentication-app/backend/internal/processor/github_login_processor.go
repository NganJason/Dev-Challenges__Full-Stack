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

	h := handler.NewAuthHandler(ctx, userAuthDM)
	h.SetGithubService(github.NewGithubService(ctx))

	userID, err := h.LoginGithub(request.AccessCode)
	if err != nil {
		return err
	}

	err = util.GenerateTokenAndAddCookies(ctx, strconv.Itoa(int(*userID)))
	if err != nil {
		return cerr.New(err.Error(), http.StatusBadGateway)
	}

	response.UserID = userID
	return nil
}
