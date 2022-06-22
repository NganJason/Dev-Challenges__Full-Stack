package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func GoogleLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.GoogleLoginRequest)

	if request.Email == nil || request.SubID == nil {
		return cerr.New("email or sub_id cannot be empty", http.StatusBadRequest)
	}

	jwt, err := util.GenerateJWTToken(*request.SubID)
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
