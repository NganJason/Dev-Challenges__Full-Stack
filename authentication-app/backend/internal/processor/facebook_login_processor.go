package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
)

func FacebookLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.FacebookLoginRequest)
	// response := resp.(*vo.FacebookLoginResponse)

	if request.AccessCode == nil {
		return cerr.New("access_token cannot be empty", http.StatusBadRequest)
	}

	s := service.NewFacebookService(ctx)
	userID, userName, err := s.Login(*request.AccessCode)
	if err != nil {
		clog.Error(ctx, err.Error())
		return err
	}

	clog.Info(ctx, fmt.Sprintf("userID=%s, userName=%s", userID, userName))
	return nil
}
