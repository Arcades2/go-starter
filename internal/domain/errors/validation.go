package errors

import "strings"

type ValidationError struct {
	Errors []*DomainError
}

func (v *ValidationError) Error() string {
	var sb strings.Builder
	for _, err := range v.Errors {
		sb.WriteString(err.Error())
		sb.WriteString("; ")
	}
	return sb.String()
}
