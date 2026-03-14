package post

import (
	"app/internal/domain/post"
)

type PostReaderService interface {
	GetByID(ID uint) (*post.Post, error)
}

type postReaderService struct {
	PostRepository post.PostRepository
}

func NewPostReaderService(
	postRepository post.PostRepository,
) PostReaderService {
	return &postReaderService{
		PostRepository: postRepository,
	}
}
