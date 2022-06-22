package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/auth"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func LoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.LoginRequest)
	response := resp.(*vo.LoginResponse)

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

	p.resp.IsAuth = vo.BoolPtr(false)

	// Get username, hashedPassword and saltstring from DB
	userName := "jason"
	hashedPasswordInDB := "abc"
	saltInDB := "123"

	if userName == "" {
		return cerr.New("cannot find username from DB", http.StatusBadRequest)
	}

	isPasswordMatch := auth.ComparePasswordSHA(*p.req.Password, hashedPasswordInDB, saltInDB)

	if !isPasswordMatch {
		return cerr.New("password and username does not match", http.StatusBadRequest)
	}

	jwt, err := util.GenerateJWTToken(*p.req.Username)
	if err != nil {
		return cerr.New(
			fmt.Sprintf("generate jwt err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	c := cookies.CreateCookie(jwt)
	cookies.AddServerCookieToCtx(p.ctx, c)

	p.resp.JWT = &jwt
	p.resp.IsAuth = vo.BoolPtr(true)

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
