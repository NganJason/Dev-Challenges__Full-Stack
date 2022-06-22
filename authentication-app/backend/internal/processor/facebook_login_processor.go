package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func FacebookLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.FacebookLoginRequest)
	// response := resp.(*vo.FacebookLoginResponse)

	if request.AccessCode == nil {
		return cerr.New("access_token cannot be empty", http.StatusBadRequest)
	}

	s := service.NewFacebookService(ctx)
	userID, _, err := s.Login(*request.AccessCode)
	if err != nil {
		clog.Error(ctx, err.Error())
		return err
	}

	jwt, err := util.GenerateJWTToken(userID)
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

	return nil
}
