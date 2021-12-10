package http

import (
	"context"
	"encoding/json"

	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"snakealive/m/pkg/common"
	cnst "snakealive/m/pkg/constants"
	"snakealive/m/pkg/error_adapter"
	auth_service "snakealive/m/pkg/services/auth"
)

type sessionChecker interface {
	ValidateSession(ctx context.Context, in *auth_service.Session, opts ...grpc.CallOption) (*auth_service.ValidateSessionResponse, error)
}

type sessionValidator struct {
	sessionChecker sessionChecker
}

func NewSessionValidatorMiddleware(
	checker sessionChecker, grpcErrorAdapter error_adapter.HttpAdapter,
) Middleware {
	sessionValidator := &sessionValidator{sessionChecker: checker}

	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			if err := sessionValidator.ValidateSession(ctx); err != nil {
				httpError := grpcErrorAdapter.AdaptError(err)

				ctx.Response.SetStatusCode(httpError.Code)
				b, _ := json.Marshal(common.ErrorMessage{Message: httpError.MSG})
				ctx.Response.SetBody(b)
				return
			}

			handler(ctx)
		}
	}
}

func (s *sessionValidator) ValidateSession(ctx *fasthttp.RequestCtx) (err error) {
	response, err := s.sessionChecker.ValidateSession(ctx, &auth_service.Session{
		Token:  "??",
		Cookie: string(ctx.Request.Header.Cookie(cnst.CookieName)),
	})
	if err != nil {
		return err
	}

	ctx.SetUserValue(cnst.UserIDContextKey, int(response.UserId))
	return nil
}
