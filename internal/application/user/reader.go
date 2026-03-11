package user

import (
	"app/internal/application/common"
	"app/internal/domain/user"
)

type UserReader interface {
	common.Panicable
	GetByID(ID uint) (*user.User, error)
}

type userReaderService struct {
	common.BaseService
	repository user.UserRepository
}

func NewUserReaderService(
	userRepository user.UserRepository,
	opts ...common.Option[UserReader],
) UserReader {
	s := &userReaderService{
		repository: userRepository,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
