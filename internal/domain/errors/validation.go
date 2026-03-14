package errors

type ValidationError struct {
	Errors []*DomainError
}

func (v *ValidationError) Error() string {
	var errMsg string
	for _, err := range v.Errors {
		errMsg += err.Error() + "; "
	}
	return errMsg
}
