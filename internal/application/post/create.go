package post

import (
	"app/internal/domain/errors"
	"app/internal/domain/post"
)

func (s *postService) Create(cmd CreatePostCommand) (*post.Post, error) {
	_, err := s.userReader.GetByID(cmd.AuthorID)
	if err != nil {
		return nil, errors.WithMessage(ErrPostCreateInvalid, "author not found")
	}

	post := post.Post{
		Title:    cmd.Title,
		Content:  cmd.Content,
		AuthorID: cmd.AuthorID,
	}

	err = s.repository.Create(&post)
	if err != nil {
		return nil, ErrPostCreateFailed
	}

	return &post, nil
}

type CreatePostCommand struct {
	Title    string
	Content  string
	AuthorID uint
}
