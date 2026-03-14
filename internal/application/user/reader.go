package user

import (
	"app/internal/domain/user"
)

type UserReader interface {
	GetByID(ID uint) (*user.User, error)
}

type userReaderService struct {
	repository user.UserRepository
}

func NewUserReaderService(
	userRepository user.UserRepository,
) UserReader {
	return &userReaderService{
		repository: userRepository,
	}
}
