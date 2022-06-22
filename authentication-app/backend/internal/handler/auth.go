package handler

type Auth struct{}

func NewAuthHandler() *Auth {
	return &Auth{}
}

func (a *Auth) CheckIfUserExist(userName string) bool {
	return true
}
