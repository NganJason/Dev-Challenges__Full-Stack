package processor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/auth"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func SignupProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.SignupRequest)
	response := resp.(*vo.SignupResponse)

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

	hashedPassword, saltString := auth.CreatePasswordSHA(*p.req.Password, util.SaltSize)
	fmt.Println(hashedPassword, saltString)

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
