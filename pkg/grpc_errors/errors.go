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
	PreparedSightServiceErrorMap = map[error]error{
		errors.SightDoesNotExist: status.Error(codes.NotFound, "sight not found"),
	}
)

var (
	PreparedAuthErrors = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: {
			MSG:  "пользователь не авторизован",
			Code: http.StatusUnauthorized,
		},
		codes.InvalidArgument: {
			MSG:  "пользователь не авторизован",
			Code: http.StatusUnauthorized,
		},
	}
	CommonAuthError = error_adapter.HttpError{
		MSG:  "произошла ошибка обращения во внутренний сервис",
		Code: http.StatusBadRequest,
	}

	UserGatewayError = map[codes.Code]error_adapter.HttpError{
		codes.NotFound: {
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

	PreparedCountryErrors = map[error]error_adapter.HttpError{
		errors.CountryDoesNotExist: {
			MSG:  "запись отсутствует",
			Code: http.StatusNotFound,
		},
	}
)
