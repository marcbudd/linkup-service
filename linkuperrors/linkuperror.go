package linkuperrors

type LinkupError struct {
	Message    string
	HTTPStatus int // HTTP status codes
}

func New(message string, httpStatus int) *LinkupError {
	return &LinkupError{
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

func (e LinkupError) Error() string {
	return e.Message
}

func (e LinkupError) HTTPStatusCode() int {
	return e.HTTPStatus
}
