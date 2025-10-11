package authservice

import "app/internal/domain/errors"

const (
	ErrInvalidCredentialsCode    errors.ErrorCode = "INVALID_CREDENTIALS"
	ErrFailedToGenerateTokenCode errors.ErrorCode = "FAILED_GENERATE_TOKEN"
	ErrUpdatingUserCode          errors.ErrorCode = "UPDATE_USER_FAILED"
)

func ErrInvalidCredentials() *errors.AppError {
	return errors.New(ErrInvalidCredentialsCode, "invalid credentials")
}

func ErrFailedToGenerateAccessToken() *errors.AppError {
	return errors.New(ErrFailedToGenerateTokenCode, "failed to generate access token")
}

func ErrFailedToGenerateRefreshToken() *errors.AppError {
	return errors.New(ErrFailedToGenerateTokenCode, "failed to generate refresh token")
}

func ErrUpdatingUser() *errors.AppError {
	return errors.New(ErrUpdatingUserCode, "failed to update user")
}
