package user

import "app/internal/domain/user"

func (s *userReaderService) GetByID(ID uint) (*user.User, error) {
	u, err := s.repository.FindByID(ID)
	if err != nil {
		return nil, user.ErrUserNotFound
	}

	return u, nil
}
