package errors

import "errors"

var (
	UserDoesNotExist    = errors.New("user does not exist")
	SessionDoesNotExist = errors.New("session does not exist")
	WrongUserPassword   = errors.New("wrong user password")
)

var (
	GatewayErrorMsgSessionNotFound = "сессия отсутствует"
)
