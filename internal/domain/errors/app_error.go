package errors

type ErrorCode string

type DomainError struct {
	Code    ErrorCode
	Message string
}

func (e *DomainError) Error() string {
	return e.Message
}

func New(code ErrorCode, message string) *DomainError {
	return &DomainError{Code: code, Message: message}
}
