package userservice

import "app/internal/domain/repository"

func (s *userService) UpdatePassword(cmd UpdatePasswordCommand) error {
	hashedPassword, err := s.PasswordHasher.HashPassword(cmd.NewPassword)
	if err != nil {
		return err
	}

	return s.UserRepository.UpdatePassword(cmd.UserID, repository.UpdateUserPasswordInput{
		HashedPassword: hashedPassword,
	})
}

type UpdatePasswordCommand struct {
	UserID      uint
	NewPassword string
}
