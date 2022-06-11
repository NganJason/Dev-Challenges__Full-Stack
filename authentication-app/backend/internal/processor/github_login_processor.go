package processor

import (
	"context"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
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

	response.UserID = &userID
	return nil
}
