package vo

import (
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
)

const (
	GithubClientSecret = "2d894c196dcba01892957b3759d9d7ddde4e5531"
	GithubClientID     = "85de73f0c04a2f06d9d5"
)

type GithubLoginRequest struct {
	AccessCode  *string `json:"access_code"`
	RedirectURI *string `json:"redirect_uri"`
}

type GithubLoginResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type VerifyAuthRequest struct{}

type VerifyAuthResponse struct {
	DebugMsg *string `json:"debug_msg"`
	IsAuth   *bool   `json:"is_auth"`
}

type GoogleLoginRequest struct {
	Email *string `json:"email"`
	SubID *string `json:"sub_id"`
}

type GoogleLoginResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type FacebookLoginRequest struct {
	AccessCode *string `json:"access_code"`
}

type FacebookLoginResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type SignupRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type SignupResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type LoginRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type LoginResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type GetUserInfoRequest struct {
	UserID *uint64 `json:"user_id"`
}

type GetUserInfoResponse struct {
	UserInfo *model.UserInfo `json:"user_info"`
}

type UpdateUserInfoRequest struct {
	UserID   *uint64 `json:"user_id"`
	Username *string `json:"username"`
	Bio      *string `json:"bio"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}

type UpdateUserInfoResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type CreateUserInfoRequest struct {
	UserID   *uint64 `json:"user_id"`
	Username *string `json:"username"`
	Bio      *string `json:"bio"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}

type CreateUserInfoResponse struct {
	DebugMsg *string         `json:"debug_msg"`
	UserInfo *model.UserInfo `json:"user_info"`
}

type LogoutRequest struct{}

type LogoutResponse struct {
	DebugMsg *string `json:"debug_msg"`
}
