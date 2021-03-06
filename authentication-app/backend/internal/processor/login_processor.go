package processor

import (
	"context"
	"net/http"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func LoginProcessor(ctx context.Context, req, resp interface{}) error {
	request, ok := req.(*vo.LoginRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.LoginResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := loginProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type loginProcessor struct {
	ctx  context.Context
	req  *vo.LoginRequest
	resp *vo.LoginResponse
}

func (p *loginProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return err
	}

	userAuthDM := model.NewUserAuthDM(p.ctx)
	userInfoDM := model.NewUserInfoDM(p.ctx)

	h := handler.NewAuthHandler(p.ctx, userAuthDM, userInfoDM)

	userInfo, err := h.DefaultLogin(p.req.Username, p.req.Password)
	if err != nil {
		return err
	}

	err = util.GenerateTokenAndAddCookies(p.ctx, strconv.Itoa(int(*userInfo.UserID)))
	if err != nil {
		return cerr.New(err.Error(), http.StatusBadGateway)
	}

	p.resp.UserInfo = userInfo

	return nil
}

func (p *loginProcessor) validateReq() error {
	if p.req.Password == nil {
		return cerr.New("password cannot be empty", http.StatusBadRequest)
	}

	if p.req.Username == nil {
		return cerr.New("username cannot be empty", http.StatusBadRequest)
	}

	return nil
}
