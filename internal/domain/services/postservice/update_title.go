package postservice

import (
	"app/internal/domain/repository"
	"app/internal/pkg/validator"
)

func (h *postService) UpdateTitle(cmd UpdateTitleCommand) error {

	err := validator.Validate.Struct(cmd)
	if err != nil {
		return h.HandleError(NewPostError(PostErrors.ErrPostUpdateInvalid))
	}

	return h.PostRepository.UpdateTitle(cmd.ID, repository.UpdatePostTitleInput{
		Title: cmd.Title,
	})
}

type UpdateTitleCommand struct {
	ID    uint   `validate:"required"`
	Title string `validate:"required,min=3,max=100"`
}
