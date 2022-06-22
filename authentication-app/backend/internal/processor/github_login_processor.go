package processor

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func GithubLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.GithubLoginRequest)
	response := resp.(*vo.GithubLoginResponse)

	s := service.NewGithubService(ctx)

	clog.Info(ctx, "received access token, processing")
	userID, err := s.Login(request.AccessCode, "")
	if err != nil {
		return err
	}

	jwt, err := util.GenerateJWTToken(strconv.Itoa(int(userID)))
	if err != nil {
		err = cerr.New(
			fmt.Sprintf("generate jwt token err=%s", err.Error()),
			http.StatusBadGateway,
		)

		clog.Error(ctx, err.Error())

		return err
	}

	c := cookies.CreateCookie(jwt)
	cookies.AddServerCookieToCtx(ctx, c)

	response.UserID = &userID
	return nil
}
