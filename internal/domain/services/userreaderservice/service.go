package userreaderservice

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
)

type UserReader interface {
	baseservice.Panicable
	GetByID(ID uint) (*model.User, error)
}

type userReaderService struct {
	baseservice.BaseService
	UserRepository repository.UserRepository
}

func NewUserReaderService(
	userRepository repository.UserRepository,
	opts ...baseservice.Option[UserReader],
) UserReader {
	s := &userReaderService{
		UserRepository: userRepository,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
