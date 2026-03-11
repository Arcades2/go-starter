package user

import "app/internal/domain/user"

func (s *userReaderService) GetByID(ID uint) (*user.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return nil, s.HandleError(ErrUserNotFound)
	}

	return user, nil
}
