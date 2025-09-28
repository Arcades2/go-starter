package userservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/domain/service"

	"app/internal/application/validation"
)

func (s *UserService) CreateUser(command service.CreateUserCommand) (*model.User, error) {

	if err := validation.Validate.Struct(command); err != nil {
		return nil, err
	}

	hashedPassword, err := s.passwordHasher.HashPassword(command.Password)
	if err != nil {
		return nil, err
	}

	userData := repository.CreateUserInput{
		Email:          command.Email,
		Firstname:      command.Firstname,
		Lastname:       command.Lastname,
		HashedPassword: hashedPassword,
	}

	return s.userRepository.Create(userData)
}
