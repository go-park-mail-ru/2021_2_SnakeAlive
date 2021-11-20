package grpc_errors

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"snakealive/m/pkg/error_adapter"
	"snakealive/m/pkg/errors"
)

var (
	PreparedAuthServiceErrorMap = map[error]error{
		errors.UserDoesNotExist:    status.Error(codes.NotFound, "user not found"),
		errors.SessionDoesNotExist: status.Error(codes.NotFound, "session not found"),
		errors.WrongUserPassword:   status.Error(codes.PermissionDenied, "wrong password"),
	}
)

var (
	PreparedAuthErrors = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: error_adapter.HttpError{
			MSG:  "пользователь не авторизован",
			Code: http.StatusUnauthorized,
		},
	}
	CommonAuthError = error_adapter.HttpError{
		MSG:  "произошла ошибка обращения во внутренний сервис",
		Code: http.StatusBadRequest,
	}

	UserGatewayError = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: error_adapter.HttpError{
			MSG:  "запись отсутствует",
			Code: http.StatusNotFound,
		},
		codes.PermissionDenied: {
			MSG:  "доступ к данным запрешен",
			Code: http.StatusBadRequest,
		},
	}
	CommonError = error_adapter.HttpError{
		MSG:  "произошла ошибка обращения во внутренний сервис",
		Code: http.StatusBadRequest,
	}
)
