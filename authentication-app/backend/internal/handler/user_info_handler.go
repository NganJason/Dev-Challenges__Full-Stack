package handler

import (
	"context"
	"fmt"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/clog"
)

type userInfoHandler struct {
	ctx context.Context
	dm  model.UserInfoInterface
}

func NewUserInfoHandler(ctx context.Context, dm model.UserInfoInterface) *userInfoHandler {
	return &userInfoHandler{
		ctx: ctx,
		dm:  dm,
	}
}

func (h *userInfoHandler) GetUserInfo(userID uint64) (*model.UserInfo, error) {
	userInfo, err := h.dm.GetUserInfo(nil, &userID)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (h *userInfoHandler) UpdateUserInfo(req *vo.UpdateUserInfoRequest) (*model.UserInfo, error) {
	userInfo, err := h.dm.UpdateUserInfo(
		&model.UpdateUserInfoRequest{
			UserID:   req.UserID,
			Username: req.Username,
			Bio:      req.Bio,
			Phone:    req.Phone,
			Email:    req.Email,
		},
	)
	if err != nil {
		clog.Error(
			h.ctx,
			fmt.Sprintf("update userID=%d err=%s", req.UserID, err.Error()),
		)
		return nil, err
	}

	clog.Info(
		h.ctx,
		fmt.Sprintf("updated userID=%d", req.UserID),
	)

	return userInfo, nil
}

func (h *userInfoHandler) CreateUserInfo(req *vo.CreateUserInfoRequest) (*model.UserInfo, error) {
	userInfo, err := h.dm.CreateUserInfo(
		&model.CreateUserInfoRequest{
			UserID:   req.UserID,
			Username: req.Username,
			Bio:      req.Bio,
			Phone:    req.Phone,
			Email:    req.Email,
		},
	)
	if err != nil {
		clog.Error(
			h.ctx,
			fmt.Sprintf("create userID=%d err=%s", req.UserID, err.Error()),
		)
		return nil, err
	}

	clog.Info(
		h.ctx,
		fmt.Sprintf("created userID=%d", req.UserID),
	)

	return userInfo, nil
}
