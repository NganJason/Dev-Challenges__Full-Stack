package processor

import (
	"context"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cerr"
)

func GoogleLoginProcessor(ctx context.Context, req, resp interface{}) error {
	request := req.(*vo.GoogleLoginRequest)

	if request.Email == nil || request.SubID == nil {
		return cerr.New("email or sub_id cannot be empty", http.StatusBadRequest)
	}

	return nil
}
