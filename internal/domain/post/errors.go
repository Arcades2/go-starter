package post

import (
	"app/internal/domain/errors"
)

var registry = errors.NewRegistry()

var ErrPostNotFound = registry.Register("POST_NOT_FOUND", "post not found")

var (
	ErrPostInvalidTitle    = errors.New("POST_INVALID_TITLE", "post invalid title")
	ErrPostInvalidContent  = errors.New("POST_INVALID_CONTENT", "post invalid content")
	ErrPostInvalidAuthorID = errors.New("POST_INVALID_AUTHOR_ID", "post invalid author id")
)

func AllPostErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
