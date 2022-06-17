package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

type ProcessorConfig struct {
	Path      string
	Processor func(ctx context.Context, req, resp interface{}) error
	Req       interface{}
	Resp      interface{}
	Cookie    *http.Cookie
}

func GetAllProcessors() []ProcessorConfig {
	return []ProcessorConfig{
		{
			Path:      "/api/login/github",
			Processor: GithubLoginProcessor,
			Req:       &vo.GithubLoginRequest{},
			Resp:      &vo.GithubLoginResponse{},
			Cookie:    cookies.GetDefaultCookies(),
		},
		{
			Path:      "/api/login/verify_auth",
			Processor: VerityAuthProcessor,
			Req:       &vo.VerifyAuthRequest{},
			Resp:      &vo.VerifyAuthResponse{},
		},
	}
}
