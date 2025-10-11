package userservice

import "app/internal/domain/errors"

const (
	ErrUserNotFoundCode errors.ErrorCode = "USER_NOT_FOUND"
)

func ErrUserNotFound() *errors.AppError {
	return errors.New(ErrUserNotFoundCode, "user not found")
}
