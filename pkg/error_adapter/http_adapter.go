package error_adapter

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HttpAdapter interface {
	AdaptError(err error) (error HttpError)
}

type HttpError struct {
	MSG  string
	Code int
}

type grpcToHttpAdapter struct {
	errorMap     map[codes.Code]HttpError
	defaultError HttpError
}

func (h *grpcToHttpAdapter) AdaptError(err error) (adapted HttpError) {
	st, ok := status.FromError(err)
	if !ok {
		return h.defaultError
	}

	var exist bool
	adapted, exist = h.errorMap[st.Code()]
	if !exist {
		adapted = h.defaultError
	}

	return
}

func NewGrpcToHttpAdapter(errorMap map[codes.Code]HttpError, defaultError HttpError) HttpAdapter {
	return &grpcToHttpAdapter{
		errorMap:     errorMap,
		defaultError: defaultError,
	}
}
