package vo

const (
	GithubClientSecret = "2d894c196dcba01892957b3759d9d7ddde4e5531"
	GithubClientID     = "85de73f0c04a2f06d9d5"
)

type GithubLoginRequest struct {
	AccessCode  string `json:"access_code"`
	RedirectURI string `json:"redirect_uri"`
}

type GithubLoginResponse struct {
	DebugMsg string
}
