package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/facebook"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/github"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/auth"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

type authHandler struct {
	ctx    context.Context
	fb     facebook.Service
	github github.Service
	authDM model.UserAuthInterface
}

func NewAuthHandler(ctx context.Context, authDM model.UserAuthInterface) *authHandler {
	return &authHandler{
		ctx:    ctx,
		authDM: authDM,
	}
}

func (h *authHandler) SetGithubService(service github.Service) {
	h.github = service
}

func (h *authHandler) SetFacebookService(service facebook.Service) {
	h.fb = service
}

func (h *authHandler) LoginGithub(accessCode *string) (userID *uint64, err error) {
	if accessCode == nil {
		return nil, cerr.New(
			"access code cannot be empty",
			http.StatusBadRequest,
		)
	}

	githubUserID, err := h.github.Login(*accessCode)
	if err != nil {
		return nil, err
	}

	userAuth, err := h.authDM.GetUserAuth(
		nil,
		nil,
		util.StrPtr(strconv.Itoa(int(githubUserID))),
		util.IntPtr(model.GithubAuthMethod),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get user auth err=%s", err.Error()),
			http.StatusBadRequest,
		)
	}

	if userAuth == nil {
		return nil, cerr.New(
			"cannot find user",
			http.StatusUnauthorized,
		)
	}

	userID = &userAuth.ID

	return userID, nil
}

func (h *authHandler) LoginFacebook(accessCode *string) (userID *uint64, err error) {
	if accessCode == nil {
		return nil, cerr.New(
			"access code cannot be empty",
			http.StatusBadRequest,
		)
	}

	fbUserID, _, err := h.fb.Login(*accessCode)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("login facebook via accessCode err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	userAuth, err := h.authDM.GetUserAuth(
		nil,
		nil,
		util.StrPtr(fbUserID),
		util.IntPtr(model.FacebookAuthMethod),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get user auth err=%s", err.Error()),
			http.StatusBadRequest,
		)
	}

	if userAuth == nil {
		return nil, cerr.New(
			"cannot find user",
			http.StatusUnauthorized,
		)
	}

	userID = &userAuth.ID

	return userID, nil
}

func (h *authHandler) LoginGoogle(subID *string) (userID *uint64, err error) {
	if subID == nil {
		return nil, cerr.New(
			"subID cannot be empty",
			http.StatusBadRequest,
		)
	}

	userAuth, err := h.authDM.GetUserAuth(
		nil,
		nil,
		subID,
		util.IntPtr(model.GoogleAuthMethod),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get user auth err=%s", err.Error()),
			http.StatusBadRequest,
		)
	}

	if userAuth == nil {
		return nil, cerr.New(
			"cannot find user",
			http.StatusUnauthorized,
		)
	}

	userID = &userAuth.ID

	return userID, nil
}

func (h *authHandler) DefaultLogin(username *string, password *string) (userID *uint64, err error) {
	if username == nil || password == nil {
		return nil, cerr.New(
			"username and password cannot be empty",
			http.StatusUnauthorized,
		)
	}

	userAuth, err := h.authDM.GetUserAuth(
		nil,
		username,
		nil,
		util.IntPtr(model.DefaultAuthMethod),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get user auth err=%s", err.Error()),
			http.StatusBadRequest,
		)
	}

	if userAuth == nil {
		return nil, cerr.New(
			"cannot find user",
			http.StatusUnauthorized,
		)
	}

	isPasswordMatch := auth.ComparePasswordSHA(
		*password,
		userAuth.HashedPassword,
		userAuth.Salt,
	)

	if !isPasswordMatch {
		return nil, cerr.New(
			"password is wrong",
			http.StatusUnauthorized,
		)
	}

	return &userAuth.ID, nil
}

func (h *authHandler) DefaultSignup(username, password *string) (userID *uint64, err error) {
	if username == nil || password == nil {
		return nil, cerr.New(
			"username and password cannot be empty",
			http.StatusBadRequest,
		)
	}

	existingUserAuth, err := h.authDM.GetUserAuth(
		nil, username, nil, util.IntPtr(model.DefaultAuthMethod),
	)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get existing user auth err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if existingUserAuth != nil {
		return nil, cerr.New(
			"username has already existed",
			http.StatusBadRequest,
		)
	}

	hashedPassword, salt := auth.CreatePasswordSHA(*password, 16)

	ID, err := h.authDM.CreateUserAuth(
		&model.CreateUserAuthRequest{
			Username:       *username,
			AuthMethod:     model.DefaultAuthMethod,
			HashedPassword: &hashedPassword,
			Salt:           &salt,
		},
	)
	if err != nil {
		return nil, err
	}

	return &ID, nil
}

func (h *authHandler) ValidateUser(userID *string) (bool, error) {
	if userID == nil {
		return false, cerr.New("userID cannot be empty", http.StatusUnauthorized)
	}

	userIDInt, err := strconv.Atoi(*userID)
	if err != nil {
		return false, cerr.New(
			fmt.Sprintf("invalid userID err=%s", err.Error()),
			http.StatusUnauthorized,
		)
	}

	userIDInt64 := uint64(userIDInt)

	userIDFromDB, err := h.authDM.GetUserAuth(&userIDInt64, nil, nil, nil)
	if err != nil {
		return false, cerr.New(
			fmt.Sprintf("get userID err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	if userIDFromDB == nil {
		return false, cerr.New(
			"cannot find user",
			http.StatusUnauthorized,
		)
	}

	return true, nil
}
