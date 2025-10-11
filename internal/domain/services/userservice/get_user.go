package userservice

import "app/internal/domain/model"

func (s *UserService) GetUserByID(ID uint) (*model.User, error) {
	user, err := s.UserRepository.FindByID(ID)
	if err != nil {
		return nil, ErrUserNotFound()
	}
	return user, nil
}
