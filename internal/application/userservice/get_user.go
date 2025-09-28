package userservice

import "app/internal/domain/model"

func (s *UserService) GetUserByID(ID uint) (*model.User, error) {
	return s.userRepository.FindByID(ID)
}
