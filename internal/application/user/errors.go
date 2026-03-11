package user

import "app/internal/domain/errors"

var registry = errors.NewRegistry()

var ErrUserNotFound = registry.Register("USER_NOT_FOUND", "user not found")

func AllUserErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
