package userreaderservice

import (
	"app/internal/domain/repository"
	"app/internal/domain/services/baseservice"
)

type UserReaderService struct {
	baseservice.BaseService
	UserRepository repository.UserRepository
}

func NewUserReaderService(
	userRepository repository.UserRepository,
	opts ...baseservice.Option[*UserReaderService],
) *UserReaderService {
	s := &UserReaderService{
		UserRepository: userRepository,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}
