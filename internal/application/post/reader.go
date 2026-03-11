package post

import (
	"app/internal/application/common"
	"app/internal/domain/post"
)

type PostReaderService interface {
	common.Panicable
	GetByID(ID uint) (*post.Post, error)
}

type postReaderService struct {
	common.BaseService
	PostRepository post.PostRepository
}

func NewPostReaderService(
	postRepository post.PostRepository,
	opts ...common.Option[PostReaderService],
) PostReaderService {
	s := &postReaderService{
		PostRepository: postRepository,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
