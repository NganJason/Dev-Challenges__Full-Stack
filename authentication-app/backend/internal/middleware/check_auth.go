package middleware

import (
	"fmt"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/handler"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/model"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c := cookies.ExtractCookie(r)
		if c == nil {
			err := cerr.New(
				"cookies not found",
				http.StatusUnauthorized,
			)
			handleErr(next, w, r, err)
			return
		}

		jwt := c.Value
		auth, err := util.ParseJWTToken(jwt)
		if err != nil || auth == nil {
			err = cerr.New(
				fmt.Sprintf("parse jwt token err=%s", err.Error()),
				http.StatusBadGateway,
			)
			handleErr(next, w, r, err)
			return
		}

		if auth.Valid() != nil {
			err = cerr.New(
				fmt.Sprintf("token is not valid err=%s", auth.Valid().Error()),
				http.StatusUnauthorized,
			)

			handleErr(next, w, r, err)
			return
		}

		userID := auth.Value
		userAuthDM := model.NewUserAuthDM(r.Context())
		userInfoDM := model.NewUserInfoDM(r.Context())

		h := handler.NewAuthHandler(r.Context(), userAuthDM, userInfoDM)

		isAuth, err := h.ValidateUser(&userID)
		if err != nil {
			handleErr(next, w, r, err)
			return
		}

		if !isAuth {
			err = cerr.New(
				"user is not authenticated",
				http.StatusUnauthorized,
			)
			handleErr(next, w, r, err)
			return
		}

		r = r.WithContext(cookies.AddClientCookieValToCtx(r.Context(), &userID))

		next(w, r)
	}

	return fn
}
