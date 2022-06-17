package cookies

import (
	"context"
	"net/http"
)

type cookieOption func(*http.Cookie)

type CookieKey string

const (
	authKey = CookieKey("authApp")
)

func AddCookieToCtx(ctx context.Context, cookie *http.Cookie) context.Context {
	return context.WithValue(ctx, GetCookieKey(), cookie)
}

func GetCookieFromCtx(ctx context.Context) *http.Cookie {
	c := ctx.Value(GetCookieKey())
	if c == nil {
		return nil
	}

	return c.(*http.Cookie)
}

func GetCookieKey() CookieKey {
	return authKey
}

func GetDefaultCookies(options ...cookieOption) *http.Cookie {
	cookie := &http.Cookie{
		Name:     string(GetCookieKey()),
		Value:    "auth-app-value",
		Path:     "/",
		MaxAge:   10 * 60,
		HttpOnly: true,
		Secure:   true,
	}

	for _, opt := range options {
		opt(cookie)
	}

	return cookie
}

func WithMaxAge(seconds int) cookieOption {
	return func(c *http.Cookie) {
		c.MaxAge = seconds
	}
}

func WithName(name string) cookieOption {
	return func(c *http.Cookie) {
		c.Name = name
	}
}

func WithValue(value string) cookieOption {
	return func(c *http.Cookie) {
		c.Value = value
	}
}
