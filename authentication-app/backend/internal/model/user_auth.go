package model

type UserAuthInterface interface {
	GetUserAuth(id *uint64, username *string, externalID *string, authMethod *int) (*UserAuth, error)
	CreateUserAuth(req *CreateUserAuthRequest) (uint64, error)
}

const (
	DefaultAuthMethod = iota + 1
	GithubAuthMethod
	FacebookAuthMethod
	GoogleAuthMethod
)

type UserAuth struct {
	ID             uint64 `json:"id"`
	Username       string `json:"username"`
	ExternalID     string `json:"external_id"`
	AuthMethod     int    `json:"auth_method"`
	HashedPassword string `json:"hashed_password"`
	Salt           string `json:"salt"`
	CreatedAt      uint64 `json:"created_at"`
	UpdatedAt      uint64 `json:"updated_at"`
}

type CreateUserAuthRequest struct {
	Username       string
	AuthMethod     int
	ExternalID     *string
	HashedPassword *string
	Salt           *string
}
