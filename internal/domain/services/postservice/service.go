package postservice

import (
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
	"app/internal/domain/services/userreaderservice"
)

type PostService struct {
	baseservice.BaseService
	PostRepository    repository.PostRepository
	UserReaderService *userreaderservice.UserReaderService
}

func NewPostService(
	postRepository repository.PostRepository,
	userReaderService *userreaderservice.UserReaderService,
	opts ...baseservice.Option[*PostService],
) *PostService {
	s := &PostService{
		PostRepository:    postRepository,
		UserReaderService: userReaderService,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
