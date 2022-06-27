package model

import "time"

type UserInfo struct {
	ID        *uint64    `json:"id"`
	UserID    *uint64    `json:"user_id"`
	Username  *string    `json:"username"`
	Bio       *string    `json:"bio"`
	Phone     *string    `json:"phone"`
	Email     *string    `json:"email"`
	Image     *string    `json:"image"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserInfoInterface interface {
	GetUserInfo(ID *uint64, userID *uint64) (*UserInfo, error)
	CreateUserInfo(req *CreateUserInfoRequest) (*UserInfo, error)
	UpdateUserInfo(req *UpdateUserInfoRequest) (*UserInfo, error)
}

type CreateUserInfoRequest struct {
	UserID   *uint64 `json:"user_id"`
	Username *string `json:"username"`
	Bio      *string `json:"bio"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}

type UpdateUserInfoRequest struct {
	UserID   *uint64 `json:"userID"`
	Username *string `json:"username"`
	Bio      *string `json:"bio"`
	Phone    *string `json:"phone"`
	Email    *string `json:"email"`
}
