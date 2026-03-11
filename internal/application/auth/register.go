package auth

import (
	"app/internal/domain/user"
)

func (s *authService) Register(command RegisterCommand) (*user.User, error) {
	hashedPassword, err := s.passwordHasher.HashPassword(command.Password)
	if err != nil {
		return nil, s.HandleError(ErrHashingPassword)
	}

	newUser := user.User{
		Email:          command.Email,
		Firstname:      command.Firstname,
		Lastname:       command.Lastname,
		HashedPassword: hashedPassword,
	}

	err = s.userRepo.Create(&newUser)

	return &newUser, s.HandleError(err)
}

type RegisterCommand struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}
