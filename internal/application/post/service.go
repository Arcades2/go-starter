package post

import (
	"app/internal/application/user"
	"app/internal/domain/post"
)

type PostService interface {
	Create(cmd CreatePostCommand) (*post.Post, error)
	UpdateTitle(cmd UpdateTitleCommand) (*post.Post, error)
}

type postService struct {
	reader     PostReaderService
	repository post.PostRepository
	userReader user.UserReader
}

func NewPostService(
	postRepository post.PostRepository,
	postReaderService PostReaderService,
	userReaderService user.UserReader,
) PostService {
	return &postService{
		reader:     postReaderService,
		repository: postRepository,
		userReader: userReaderService,
	}
}
