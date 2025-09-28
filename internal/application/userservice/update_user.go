package userservice

import "app/internal/domain/repository"

func (s *UserService) UpdateUserPassword(ID uint, newPassword string) error {
	hashedPassword, err := s.passwordHasher.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.userRepository.UpdateUserPassword(ID, repository.UpdateUserPasswordInput{
		HashedPassword: hashedPassword,
	})
}

func (s *UserService) UpdateUserInfo(ID uint, firstname, lastname string) error {
	return s.userRepository.UpdateUserInfo(ID, repository.UpdateUserInfoInput{
		Firstname: firstname,
		Lastname:  lastname,
	})
}
