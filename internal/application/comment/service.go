package comment

import (
	"app/internal/application/post"
	"app/internal/application/user"
	"app/internal/domain/comment"
)

type CommentService interface {
	Create(cmd CreateCommentCommand) (*comment.Comment, error)
}

type commentService struct {
	postReader post.PostReaderService
	repository comment.CommentRepository
	userReader user.UserReader
}

func NewCommentService(
	commentRepository comment.CommentRepository,
	postReaderService post.PostReaderService,
	userReaderService user.UserReader,
) CommentService {
	return &commentService{
		postReader: postReaderService,
		repository: commentRepository,
		userReader: userReaderService,
	}
}
