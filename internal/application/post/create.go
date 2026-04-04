package post

import (
	"app/internal/domain/post"
)

func (s *postService) Create(cmd CreatePostCommand) (*post.Post, error) {
	_, err := s.userReader.GetByID(cmd.AuthorID)
	if err != nil {
		return nil, err
	}

	post := post.NewPost(cmd.Title, cmd.Content, cmd.AuthorID)

	err = s.repository.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

type CreatePostCommand struct {
	Title    string
	Content  string
	AuthorID uint
}
