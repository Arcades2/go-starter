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

func WithMessage(err *DomainError, message string) *DomainError {
	return &DomainError{
		Code:    err.Code,
		Message: message,
	}
}

type Registry struct {
	errors map[ErrorCode]*DomainError
}

func NewRegistry() *Registry {
	return &Registry{errors: make(map[ErrorCode]*DomainError)}
}

func (r *Registry) Register(code ErrorCode, message string) *DomainError {
	err := New(code, message)
	r.errors[code] = err
	return err
}

func (r *Registry) AllCodes() []ErrorCode {
	out := make([]ErrorCode, 0, len(r.errors))
	for code := range r.errors {
		out = append(out, code)
	}
	return out
}
