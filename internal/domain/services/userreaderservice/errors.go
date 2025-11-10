package userreaderservice

import (
	"reflect"

	"app/internal/domain/errors"
)

type UserReaderError = errors.DomainError

var UserReaderErrors = struct {
	ErrUserNotFound UserReaderError
}{
	ErrUserNotFound: UserReaderError{Code: "USER_NOT_FOUND", Message: "user not found"},
}

func NewUserReaderError(err UserReaderError) *errors.DomainError {
	return &errors.DomainError{
		Code:    errors.ErrorCode(err.Code),
		Message: err.Message,
	}
}

func AllUserReaderErrorCodes() []errors.ErrorCode {
	val := reflect.ValueOf(UserReaderErrors)
	codes := make([]errors.ErrorCode, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if userErr, ok := field.Interface().(UserReaderError); ok {
			codes = append(codes, userErr.Code)
		}
	}

	return codes
}
