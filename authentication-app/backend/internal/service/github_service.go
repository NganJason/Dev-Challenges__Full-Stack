package service

import (
	"context"
	"fmt"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/http_util"
)

const (
	accessTokenURL = "https://github.com/login/oauth/access_token"
	loginURL       = "https://api.github.com/user"
)

type githubAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirect_uri"`
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Error       string `json:"error"`
}

type githubGetUserResponse struct {
	ID    int64  `json:"id"`
	Error string `json:"error"`
}

type GithubService struct {
	ctx context.Context
}

func NewGithubService(ctx context.Context) *GithubService {
	return &GithubService{
		ctx: ctx,
	}
}

func (s *GithubService) Login(code, redirectURI string) (int64, error) {
	accessToken, err := s.getAccessToken(code)
	if err != nil {
		clog.Error(s.ctx, fmt.Sprintf("get access token err=%s", err.Error()))
		return 0, err
	}

	userID, err := s.getUserID(accessToken)
	if err != nil {
		clog.Error(s.ctx, "get userID error")
		return 0, err
	}

	return userID, nil
}

func (s *GithubService) getAccessToken(code string) (string, error) {
	req := &githubAccessTokenRequest{
		ClientID:     vo.GithubClientID,
		ClientSecret: vo.GithubClientSecret,
		Code:         code,
	}

	var resp githubAccessTokenResponse

	clog.Info(s.ctx, "posting access code to github")
	err := http_util.Post(
		accessTokenURL,
		req,
		&resp,
		http_util.WithAccept("application/json"),
	)
	if err != nil {
		return "", fmt.Errorf("post gitlab req err=%s", err.Error())
	}

	if resp.Error != "" {
		return "", fmt.Errorf("get access token err=%s", resp.Error)
	}

	return resp.AccessToken, nil
}

func (s *GithubService) getUserID(accessToken string) (int64, error) {
	var resp githubGetUserResponse

	clog.Info(s.ctx, "posting access token to github")
	err := http_util.Get(
		loginURL,
		&resp,
		http_util.WithAccept("application/json"),
		http_util.WithBearer(accessToken),
	)
	if err != nil {
		return 0, err
	}

	if resp.Error != "" {
		return 0, fmt.Errorf("get userID err=%s", resp.Error)
	}

	return resp.ID, nil
}
