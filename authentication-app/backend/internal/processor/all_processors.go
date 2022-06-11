package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
)

type ProcessorConfig struct {
	Path      string
	Processor func(ctx context.Context, req, resp interface{}) error
	Req       interface{}
	Resp      interface{}
	Cookie    *http.Cookie
}

func boolPtr(b bool) *bool {
	return &b
}

func GetAllProcessors() []ProcessorConfig {
	return []ProcessorConfig{
		{
			Path:      "/api/login/github",
			Processor: GithubLoginProcessor,
			Req:       &vo.GithubLoginRequest{},
			Resp:      &vo.GithubLoginResponse{},
			Cookie:    util.GetDefaultCookies(),
		},
	}
}
