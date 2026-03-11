package post

import (
	"app/internal/application/common"
	"app/internal/application/user"
	"app/internal/domain/post"
)

type PostService interface {
	common.Panicable
	Create(cmd CreatePostCommand) (*post.Post, error)
	UpdateTitle(cmd UpdateTitleCommand) (*post.Post, error)
}

type postService struct {
	common.BaseService
	reader     PostReaderService
	repository post.PostRepository
	userReader user.UserReader
}

func NewPostService(
	postRepository post.PostRepository,
	postReaderService PostReaderService,
	userReaderService user.UserReader,
	opts ...common.Option[PostService],
) PostService {
	s := &postService{
		reader:     postReaderService,
		repository: postRepository,
		userReader: userReaderService,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
