package user

import (
	"reflect"

	"app/internal/domain/errors"
)

type UserError = errors.DomainError

var UserErrors = struct {
	ErrUserNotFound UserError
	ErrInvalidUser  UserError
}{
	ErrUserNotFound: UserError{Code: "USER_NOT_FOUND", Message: "user not found"},
	ErrInvalidUser:  UserError{Code: "INVALID_USER", Message: "invalid user"},
}

func NewUserError(err UserError, message string) *errors.DomainError {
	error := &errors.DomainError{
		Code: errors.ErrorCode(err.Code),
	}

	if message != "" {
		error.Message = message
	} else {
		error.Message = err.Message
	}

	return error
}

func AllUserErrorCodes() []errors.ErrorCode {
	val := reflect.ValueOf(UserErrors)
	codes := make([]errors.ErrorCode, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if userErr, ok := field.Interface().(UserError); ok {
			codes = append(codes, userErr.Code)
		}
	}

	return codes
}
