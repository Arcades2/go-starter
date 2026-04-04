package user

import (
	"app/internal/domain/errors"
)

var registry = errors.NewRegistry()

var ErrUserNotFound = registry.Register("USER_NOT_FOUND", "user not found")

// Internal/validation errors
var (
	ErrUserInvalidEmail          = errors.New("USER_INVALID_EMAIL", "invalid email")
	ErrUserInvalidFirstname      = errors.New("USER_INVALID_FIRSTNAME", "invalid firstname")
	ErrUserInvalidLastname       = errors.New("USER_INVALID_LASTNAME", "invalid lastname")
	ErrUserInvalidHashedPassword = errors.New("USER_INVALID_HASHED_PASSWORD", "invalid hashed password")
	ErrUserInvalidRefreshToken   = errors.New("USER_INVALID_REFRESH", "invalid refresh token")
)

func AllUserErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
