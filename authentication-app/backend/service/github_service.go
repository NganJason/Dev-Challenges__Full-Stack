package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/vo"
)

const (
	loginURL = "https://github.com/login/oauth/access_token"
)

type githubAccessTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectURI  string `json:"redirectURI"`
}

type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Error string `json:"error"`
}

type GithubService struct{}

func NewGithubService() *GithubService {
	return &GithubService{}
}

func (s *GithubService) Login(code, redirectURI string) (string, error) {
	accessToken, err := s.getAccessToken(code)	
	if err != nil {
		return "", fmt.Errorf("error getting access token err=%s", err.Error())
	}

	return accessToken, nil
}

func (s *GithubService) getAccessToken(code string) (string, error) {
	var resp githubAccessTokenResponse

	req := &githubAccessTokenRequest{
		ClientID:     vo.GithubClientID,
		ClientSecret: vo.GithubClientSecret,
		Code:         code,
	}

	respBytes, err := httpPost(loginURL, req)
	if err != nil {
		return "", fmt.Errorf("error posting gitlab req err=%s", err.Error())
	}

	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling respBytes err=%s", err.Error())
	}

	return resp.AccessToken, nil
}

func httpPost(url string, req interface{}) ([]byte, error) {
	client := http.Client{}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling reqBytes err=%s", err.Error())
	}

	reqBody := bytes.NewBuffer(reqBytes)

	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return nil, err
	}

	httpReq.Header.Add("Accept", "application/json")
	
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
