package postservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
	"app/internal/domain/services/userreaderservice"
)

type PostService interface {
	baseservice.Panicable
	Create(cmd CreatePostCommand) (*model.Post, error)
	GetByID(ID uint) (*model.Post, error)
	UpdateTitle(cmd UpdateTitleCommand) error
}

type postService struct {
	baseservice.BaseService
	PostRepository    repository.PostRepository
	UserReaderService userreaderservice.UserReader
}

func NewPostService(
	postRepository repository.PostRepository,
	userReaderService userreaderservice.UserReader,
	opts ...baseservice.Option[PostService],
) PostService {
	s := &postService{
		PostRepository:    postRepository,
		UserReaderService: userReaderService,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
