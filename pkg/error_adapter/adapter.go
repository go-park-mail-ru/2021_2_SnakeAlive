package error_adapter

type ErrorAdapter interface {
	AdaptError(err error) error
}
type errorAdapter struct {
	errorMap map[error]error
}

func (e *errorAdapter) AdaptError(err error) error {
	if adapted, exist := e.errorMap[err]; exist {
		return adapted
	}

	return err
}

func NewErrorAdapter(errorMap map[error]error) ErrorAdapter {
	return &errorAdapter{errorMap: errorMap}
}
