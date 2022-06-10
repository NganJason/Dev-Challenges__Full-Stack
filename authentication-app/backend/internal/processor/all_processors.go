package processor

import (
	"context"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
)

type ProcessorConfig struct {
	Path      string
	Processor func(ctx context.Context, req, resp interface{}) error
	Req       interface{}
	Resp      interface{}
}

func GetAllProcessors() []ProcessorConfig {
	return []ProcessorConfig{
		{
			Path:      "/api/login/github",
			Processor: GithubLoginProcessor,
			Req:       &vo.GithubLoginRequest{},
			Resp:      &vo.GithubLoginResponse{},
		},
	}
}
