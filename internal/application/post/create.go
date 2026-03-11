package post

import (
	"app/internal/domain/errors"
	"app/internal/domain/post"
)

func (s *postService) Create(cmd CreatePostCommand) (*post.Post, error) {
	_, err := s.reader.GetByID(cmd.AuthorID)
	if err != nil {
		return nil, s.HandleError(
			errors.WithMessage(ErrPostCreateInvalid, "author not found"),
		)
	}

	post := post.Post{
		Title:    cmd.Title,
		Content:  cmd.Content,
		AuthorID: cmd.AuthorID,
	}

	s.repository.Create(&post)

	return &post, nil
}

type CreatePostCommand struct {
	Title    string
	Content  string
	AuthorID uint
}
