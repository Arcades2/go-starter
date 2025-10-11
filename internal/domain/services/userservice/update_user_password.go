package userservice

import "app/internal/domain/repository"

func (s *UserService) UpdateUserPassword(ID uint, newPassword string) error {
	hashedPassword, err := s.PasswordHasher.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.UserRepository.UpdateUserPassword(ID, repository.UpdateUserPasswordInput{
		HashedPassword: hashedPassword,
	})
}
