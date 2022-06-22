package service

import (
	"context"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/http_util"
)

type FacebookService struct {
	ctx context.Context
}

const (
	getUserURL = "https://graph.facebook.com/me"
)

type getUserResponse struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

func NewFacebookService(ctx context.Context) *FacebookService {
	return &FacebookService{
		ctx: ctx,
	}
}

func (s *FacebookService) Login(accessToken string) (userID string, userName string, err error) {
	var resp getUserResponse

	url := getUserURL + "?access_token=" + accessToken
	err = http_util.Get(
		url,
		&resp,
	)
	if err != nil {
		clog.Error(s.ctx, err.Error())
		return
	}

	return resp.Id, resp.Name, nil
}
