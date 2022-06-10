package util

import (
	"net/http"
)

type cookieOption func(*http.Cookie)

func GetDefaultCookies(options ...cookieOption) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "auth-app",
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
