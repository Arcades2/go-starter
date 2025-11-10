package authservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"

	"app/internal/pkg/validator"
)

func (s *authService) Register(command RegisterCommand) (*model.User, error) {
	if err := validator.Validate.Struct(command); err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrRegisterInvalidInput))
	}

	hashedPassword, err := s.PasswordHasher.HashPassword(command.Password)
	if err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrHashingPassword))
	}

	userData := repository.CreateUserInput{
		Email:          command.Email,
		Firstname:      command.Firstname,
		Lastname:       command.Lastname,
		HashedPassword: hashedPassword,
	}

	return s.UserRepo.Create(userData)
}

type RegisterCommand struct {
	Firstname string `validate:"required,min=1,max=255"`
	Lastname  string `validate:"required,min=1,max=255"`
	Email     string `validate:"required,email,max=255"`
	Password  string `validate:"required,min=8,max=100"`
}
