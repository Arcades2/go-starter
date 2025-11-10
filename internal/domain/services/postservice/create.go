package postservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/pkg/validator"
)

func (s *postService) Create(cmd CreatePostCommand) (*model.Post, error) {
	input := repository.CreatePostInput{
		Title:    cmd.Title,
		Content:  cmd.Content,
		AuthorID: cmd.AuthorID,
	}

	if err := validator.Validate.Struct(cmd); err != nil {
		return nil, s.HandleError(
			NewPostError(PostErrors.ErrPostCreateInvalid),
		)
	}

	_, err := s.UserReaderService.GetByID(cmd.AuthorID)
	if err != nil {
		return nil, s.HandleError(
			err,
		)
	}

	post, err := s.PostRepository.Create(input)
	if err != nil {
		return nil, err
	}
	return post, nil
}

type CreatePostCommand struct {
	Title    string `validate:"required,min=3,max=100"`
	Content  string `validate:"required,min=10"`
	AuthorID uint   `validate:"required"`
}
