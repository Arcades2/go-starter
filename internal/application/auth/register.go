package auth

import (
	"app/internal/domain/auth"
	"app/internal/domain/user"
)

func (s *authService) Register(command RegisterCommand) (*user.User, error) {
	hashedPassword, err := s.passwordHasher.HashPassword(command.Password)
	if err != nil {
		return nil, auth.ErrHashingPassword
	}

	newUser, err := user.NewUser(
		command.Email,
		command.Firstname,
		command.Lastname,
		hashedPassword,
	)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.Create(newUser)

	return newUser, err
}

type RegisterCommand struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}
