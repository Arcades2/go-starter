package authservice

import (
	"app/internal/domain/repository"
)

func (s *AuthService) Login(email, password string) (*LoginOutputDTO, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, NewAuthError(AuthErrors.ErrInvalidCredentials)
	}

	if !s.PasswordHasher.VerifyPassword(password, user.HashedPassword) {
		return nil, NewAuthError(AuthErrors.ErrInvalidCredentials)
	}

	accessToken, err := s.TokenGenerator.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, NewAuthError(AuthErrors.ErrFailedToGenerateToken)
	}

	refreshToken, err := s.TokenGenerator.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, NewAuthError(AuthErrors.ErrFailedToGenerateToken)
	}

	if err := s.UserRepo.UpdateRefreshToken(user.ID, repository.UpdateUserRefreshTokenInput{
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, NewAuthError(AuthErrors.ErrUpdatingRefreshToken)
	}

	return &LoginOutputDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Type:         "Bearer",
	}, nil
}

type LoginOutputDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Type         string `json:"type"`
}
