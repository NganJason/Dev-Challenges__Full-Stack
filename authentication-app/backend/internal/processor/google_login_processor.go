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

func GoogleLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.GoogleLoginRequest)
	response := resp.(*vo.GoogleLoginResponse)

	if request.SubID == nil {
		return cerr.New("email or sub_id cannot be empty", http.StatusBadRequest)
	}

	userAuthDM := model.NewUserAuthDM(ctx)
	h := handler.NewAuthHandler(ctx, userAuthDM)

	userInfo, err := h.LoginGoogle(request.SubID)
	if err != nil {
		return err
	}

	err = util.GenerateTokenAndAddCookies(ctx, strconv.Itoa(int(*userInfo.UserID)))
	if err != nil {
		return cerr.New(err.Error(), http.StatusBadGateway)
	}

	response.UserInfo = userInfo

	return nil
}
