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
	NeedAuth  bool
}

func GetAllProcessors() []ProcessorConfig {
	return []ProcessorConfig{
		{
			Path:      "/api/login",
			Processor: LoginProcessor,
			Req:       &vo.LoginRequest{},
			Resp:      &vo.LoginResponse{},
		},
		{
			Path:      "/api/signup",
			Processor: SignupProcessor,
			Req:       &vo.SignupRequest{},
			Resp:      &vo.SignupResponse{},
		},
		{
			Path:      "/api/login/github",
			Processor: GithubLoginProcessor,
			Req:       &vo.GithubLoginRequest{},
			Resp:      &vo.GithubLoginResponse{},
		},
		{
			Path:      "/api/login/google",
			Processor: GoogleLoginProcessor,
			Req:       &vo.GoogleLoginRequest{},
			Resp:      &vo.GoogleLoginResponse{},
		},
		{
			Path:      "/api/login/facebook",
			Processor: FacebookLoginProcessor,
			Req:       &vo.FacebookLoginRequest{},
			Resp:      &vo.FacebookLoginResponse{},
		},
		{
			Path:      "/api/login/verify_auth",
			Processor: VerityAuthProcessor,
			Req:       &vo.VerifyAuthRequest{},
			Resp:      &vo.VerifyAuthResponse{},
			NeedAuth:  true,
		},
	}
}
