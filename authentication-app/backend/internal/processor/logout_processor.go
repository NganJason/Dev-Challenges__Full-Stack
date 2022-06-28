package processor

import (
	"context"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
)

func LogoutProcessor(ctx context.Context, req, resp interface{}) error {
	deleteCookie := cookies.DeleteCookie()

	cookies.AddServerCookieToCtx(ctx, deleteCookie)

	return nil
}
