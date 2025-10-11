package errors

type ErrorCode string

type AppError struct {
	Code    ErrorCode
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func New(code ErrorCode, message string) *AppError {
	return &AppError{Code: code, Message: message}
}
