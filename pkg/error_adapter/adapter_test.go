package error_adapter

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	firstError  = errors.New("some text")
	secondError = errors.New("some text too")

	notExistingError = errors.New("does not exist")

	defaultHttpError = HttpError{
		MSG:  "some msg",
		Code: http.StatusTeapot,
	}
	existError = HttpError{
		MSG:  "exist",
		Code: http.StatusBadRequest,
	}
)

func TestCommonAdapter(t *testing.T) {
	adapter := NewErrorAdapter(map[error]error{
		firstError: secondError,
	})

	assert.Equal(t, adapter.AdaptError(firstError), secondError)
	assert.Equal(t, adapter.AdaptError(notExistingError), notExistingError)
}

func TestGRPCAdapter(t *testing.T) {
	adapter := NewGrpcToHttpAdapter(map[codes.Code]HttpError{
		codes.InvalidArgument: existError,
	}, defaultHttpError)

	assert.Equal(t, adapter.AdaptError(status.Error(codes.InvalidArgument, "")), existError)
	assert.Equal(t, adapter.AdaptError(notExistingError), defaultHttpError)
}

func TestHttpAdapter(t *testing.T) {
	adapter := NewErrorToHttpAdapter(map[error]HttpError{
		firstError: existError,
	}, defaultHttpError)

	assert.Equal(t, adapter.AdaptError(firstError), existError)
	assert.Equal(t, adapter.AdaptError(notExistingError), defaultHttpError)
}
