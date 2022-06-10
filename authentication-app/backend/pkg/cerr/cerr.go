package cerr

import (
	"errors"
	"net/http"
)

type cerr struct {
	error
	code int
}

func New(msg string, code int) error {
	return &cerr{
		error: errors.New(msg),
		code:  code,
	}
}

func Code(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if cerr, ok := err.(*cerr); ok {
		return cerr.code
	} else {
		return http.StatusBadGateway
	}
}
