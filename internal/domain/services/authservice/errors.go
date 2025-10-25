package authservice

import (
	"reflect"

	"app/internal/domain/errors"
)

type AuthError = errors.DomainError

var AuthErrors = struct {
	ErrInvalidCredentials    AuthError
	ErrFailedToGenerateToken AuthError
	ErrUpdatingRefreshToken  AuthError
	ErrRegisterInvalidInput  AuthError
	ErrHashingPassword       AuthError
}{
	ErrInvalidCredentials:    AuthError{Code: "INVALID_CREDENTIALS", Message: "invalid credentials"},
	ErrFailedToGenerateToken: AuthError{Code: "FAILED_GENERATE_TOKEN", Message: "failed to generate token"},
	ErrUpdatingRefreshToken:  AuthError{Code: "UPDATE_USER_FAILED", Message: "failed to update user"},
	ErrRegisterInvalidInput:  AuthError{Code: "REGISTER_INVALID_INPUT", Message: "invalid input"},
	ErrHashingPassword:       AuthError{Code: "HASHING_PASSWORD_FAILED", Message: "failed to hash password"},
}

func NewAuthError(err AuthError) *errors.DomainError {
	return &errors.DomainError{
		Code:    errors.ErrorCode(err.Code),
		Message: err.Message,
	}
}

func AllAuthErrorCodes() []errors.ErrorCode {
	val := reflect.ValueOf(AuthErrors)
	codes := make([]errors.ErrorCode, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if authErr, ok := field.Interface().(AuthError); ok {
			codes = append(codes, authErr.Code)
		}
	}

	return codes
}
