package comment

import (
	"app/internal/domain/errors"
)

var registry = errors.NewRegistry()

var ErrCommentNotFound = registry.Register("COMMENT_NOT_FOUND", "comment not found")

var (
	ErrCommentInvalidContent  = errors.New("COMMENT_INVALID_CONTENT", "comment invalid content")
	ErrCommentInvalidAuthorID = errors.New("COMMENT_INVALID_AUTHOR_ID", "comment invalid author id")
	ErrCommentInvalidPostID   = errors.New("COMMENT_INVALID_POST_ID", "comment invalid post id")
)

func AllCommentErrorCodes() []errors.ErrorCode {
	return registry.AllCodes()
}
