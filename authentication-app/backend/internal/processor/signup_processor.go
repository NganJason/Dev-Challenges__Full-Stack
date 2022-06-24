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

func SignupProcessor(ctx context.Context, req, resp interface{}) error {
	request, ok := req.(*vo.SignupRequest)
	if !ok {
		return cerr.New(
			"convert request body error",
			http.StatusBadRequest,
		)
	}

	response, ok := resp.(*vo.SignupResponse)
	if !ok {
		return cerr.New(
			"convert response body error",
			http.StatusBadRequest,
		)
	}

	p := signupProcessor{
		ctx:  ctx,
		req:  request,
		resp: response,
	}

	return p.process()
}

type signupProcessor struct {
	ctx  context.Context
	req  *vo.SignupRequest
	resp *vo.SignupResponse
}

func (p *signupProcessor) process() error {
	err := p.validateReq()
	if err != nil {
		return err
	}

	authDM := model.NewUserAuthDM(p.ctx)
	h := handler.NewAuthHandler(p.ctx, authDM)

	userID, err := h.DefaultSignup(p.req.Username, p.req.Password)
	if err != nil {
		return err
	}

	err = util.GenerateTokenAndAddCookies(p.ctx, strconv.Itoa(int(*userID)))
	if err != nil {
		return cerr.New(err.Error(), http.StatusBadGateway)
	}

	p.resp.UserID = userID

	return nil
}

func (p *signupProcessor) validateReq() error {
	if p.req.Password == nil {
		return cerr.New("password cannot be empty", http.StatusBadRequest)
	}

	if p.req.Username == nil {
		return cerr.New("username cannot be empty", http.StatusBadRequest)
	}

	return nil
}
