package processor

import (
	"context"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/pkg/cookies"
	"google.golang.org/protobuf/proto"
)

func VerityAuthProcessor(ctx context.Context, req, resp interface{}) error {
	response := resp.(*vo.VerifyAuthResponse)

	cookieVal := cookies.GetClientCookieValFromCtx(ctx)
	if cookieVal == nil {
		response.IsAuth = proto.Bool(false)
	} else {
		response.IsAuth = proto.Bool(true)
	}

	return nil
}
