package post

import (
	"app/internal/domain/errors"
)

var registry = errors.NewRegistry()

var (
	ErrPostNotFound      = registry.Register("POST_NOT_FOUND", "post not found")
	ErrPostCreateInvalid = registry.Register("POST_CREATE_INVALID", "post create invalid")
	ErrPostUpdateInvalid = registry.Register("POST_UPDATE_INVALID", "post update invalid")
)

func AllPostErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
