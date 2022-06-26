package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/facebook"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service/github"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/auth"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

type authHandler struct {
	ctx        context.Context
	fb         facebook.Service
	github     github.Service
	authDM     model.UserAuthInterface
	userInfoDM model.UserInfoInterface
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

func (h *authHandler) SetUserInfoDM(userInfoDM model.UserInfoInterface) {
	h.userInfoDM = userInfoDM
}

func (h *authHandler) LoginGithub(accessCode *string) (userInfo *model.UserInfo, err error) {
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
		githubUserIDStr := strconv.Itoa(int(githubUserID))

		id, err := h.authDM.CreateUserAuth(
			&model.CreateUserAuthRequest{
				LoginID:    &githubUserIDStr,
				AuthMethod: model.GithubAuthMethod,
			},
		)
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf(
					"cannot find and create user err=%s", err.Error(),
				),
				http.StatusUnauthorized,
			)
		}

		userInfo, err := h.userInfoDM.CreateUserInfo(&model.CreateUserInfoRequest{
			UserID: &id,
		})
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf("create userInfo for new user err=%s", err.Error()),
				cerr.Code(err),
			)
		}

		return userInfo, nil
	}

	userInfo, err = h.userInfoDM.GetUserInfo(nil, userAuth.ID)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get userInfo err=%s", err.Error()),
			cerr.Code(err),
		)
	}

	return userInfo, nil
}

func (h *authHandler) LoginFacebook(accessCode *string) (userInfo *model.UserInfo, err error) {
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
		id, err := h.authDM.CreateUserAuth(
			&model.CreateUserAuthRequest{
				LoginID:    &fbUserID,
				AuthMethod: model.FacebookAuthMethod,
			},
		)
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf(
					"cannot find and create user err=%s", err.Error(),
				),
				http.StatusUnauthorized,
			)
		}

		userInfo, err := h.userInfoDM.CreateUserInfo(&model.CreateUserInfoRequest{
			UserID: &id,
		})
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf("create userInfo for new user err=%s", err.Error()),
				cerr.Code(err),
			)
		}

		return userInfo, nil
	}

	userInfo, err = h.userInfoDM.GetUserInfo(nil, userAuth.ID)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get userInfo err=%s", err.Error()),
			cerr.Code(err),
		)
	}

	return userInfo, nil
}

func (h *authHandler) LoginGoogle(subID *string) (userInfo *model.UserInfo, err error) {
	if subID == nil {
		return nil, cerr.New(
			"subID cannot be empty",
			http.StatusBadRequest,
		)
	}

	userAuth, err := h.authDM.GetUserAuth(
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
		id, err := h.authDM.CreateUserAuth(
			&model.CreateUserAuthRequest{
				LoginID:    subID,
				AuthMethod: model.GoogleAuthMethod,
			},
		)
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf(
					"cannot find and create user err=%s", err.Error(),
				),
				http.StatusUnauthorized,
			)
		}

		userInfo, err := h.userInfoDM.CreateUserInfo(&model.CreateUserInfoRequest{
			UserID: &id,
		})
		if err != nil {
			return nil, cerr.New(
				fmt.Sprintf("create userInfo for new user err=%s", err.Error()),
				cerr.Code(err),
			)
		}

		return userInfo, nil
	}

	userInfo, err = h.userInfoDM.GetUserInfo(nil, userAuth.ID)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get userInfo err=%s", err.Error()),
			cerr.Code(err),
		)
	}

	return userInfo, nil
}

func (h *authHandler) DefaultLogin(username *string, password *string) (userInfo *model.UserInfo, err error) {
	if username == nil || password == nil {
		return nil, cerr.New(
			"username and password cannot be empty",
			http.StatusUnauthorized,
		)
	}

	userAuth, err := h.authDM.GetUserAuth(
		nil,
		username,
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

	var realHashedPasswordStr = strings.Replace(string(userAuth.HashedPassword), "\"", "", -1)
	isPasswordMatch := auth.ComparePasswordSHA(
		*password,
		realHashedPasswordStr,
		*userAuth.Salt,
	)

	if !isPasswordMatch {
		return nil, cerr.New(
			"wrong password",
			http.StatusUnauthorized,
		)
	}

	userInfo, err = h.userInfoDM.GetUserInfo(nil, userAuth.ID)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("get userInfo err=%s", err.Error()),
			cerr.Code(err),
		)
	}

	return userInfo, nil
}

func (h *authHandler) DefaultSignup(username, password *string) (userInfo *model.UserInfo, err error) {
	if username == nil || password == nil {
		return nil, cerr.New(
			"username and password cannot be empty",
			http.StatusBadRequest,
		)
	}

	existingUserAuth, err := h.authDM.GetUserAuth(
		nil, username, util.IntPtr(model.DefaultAuthMethod),
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
	hashedBytes, err := json.Marshal(hashedPassword)
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("marshal hashedpassword err=%s", err.Error()),
			http.StatusBadGateway,
		)
	}

	ID, err := h.authDM.CreateUserAuth(
		&model.CreateUserAuthRequest{
			LoginID:        username,
			AuthMethod:     model.DefaultAuthMethod,
			HashedPassword: hashedBytes,
			Salt:           &salt,
		},
	)
	if err != nil {
		return nil, err
	}

	userInfo, err = h.userInfoDM.CreateUserInfo(&model.CreateUserInfoRequest{
		UserID: &ID,
	})
	if err != nil {
		return nil, cerr.New(
			fmt.Sprintf("create userInfo for new user err=%s", err.Error()),
			cerr.Code(err),
		)
	}

	return userInfo, nil
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
	userIDFromDB, err := h.authDM.GetUserAuth(&userIDInt64, nil, nil)
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
