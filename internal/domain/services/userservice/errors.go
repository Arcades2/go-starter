package userservice

import (
	"reflect"

	"app/internal/domain/errors"
)

type UserError = errors.AppError

var UserErrors = struct {
	ErrUserNotFound UserError
}{
	ErrUserNotFound: UserError{Code: "USER_NOT_FOUND", Message: "user not found"},
}

func NewUserError(err UserError) *errors.AppError {
	return &errors.AppError{
		Code:    errors.ErrorCode(err.Code),
		Message: err.Message,
	}
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
