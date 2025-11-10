package userreaderservice

import "app/internal/domain/model"

func (s *UserReaderService) GetUserByID(ID uint) (*model.User, error) {
	user, err := s.UserRepository.FindByID(ID)
	if err != nil {
		return nil, s.HandleError(NewUserReaderError(UserReaderErrors.ErrUserNotFound))
	}
	return user, nil
}
