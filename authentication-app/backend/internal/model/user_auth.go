package model

import "time"

type UserAuthInterface interface {
	GetUserAuth(userID *uint64, loginID *string, authMethod *int) (*UserAuth, error)
	CreateUserAuth(req *CreateUserAuthRequest) (userID uint64, err error)
}

const (
	DefaultAuthMethod = iota + 1
	GithubAuthMethod
	FacebookAuthMethod
	GoogleAuthMethod
)

type UserAuth struct {
	ID             *uint64    `json:"id"`
	LoginID        *string    `json:"external_id"`
	AuthMethod     *int       `json:"auth_method"`
	HashedPassword []byte     `json:"hashed_password"`
	Salt           *string    `json:"salt"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type CreateUserAuthRequest struct {
	AuthMethod     int
	LoginID        *string
	HashedPassword []byte
	Salt           *string
}
