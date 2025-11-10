package postservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/pkg/validator"
)

func (s *PostService) CreatePost(command CreatePostCommand) (*model.Post, error) {
	input := repository.CreatePostInput{
		Title:    command.Title,
		Content:  command.Content,
		AuthorID: command.AuthorID,
	}

	if err := validator.Validate.Struct(command); err != nil {
		return nil, s.HandleError(
			NewPostError(PostErrors.ErrPostCreateInvalid),
		)
	}

	_, err := s.UserReaderService.GetUserByID(command.AuthorID)
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
	Content  string `validate:"required"`
	AuthorID uint   `validate:"required"`
}
