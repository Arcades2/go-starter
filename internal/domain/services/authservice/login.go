package authservice

import (
	"app/internal/domain/repository"
)

func (s *authService) Login(cmd LoginCommand) (*LoginOutput, error) {
	user, err := s.UserRepo.FindByEmail(cmd.Email)
	if err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrInvalidCredentials))
	}

	if !s.PasswordHasher.VerifyPassword(cmd.Password, user.HashedPassword) {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrInvalidCredentials))
	}

	accessToken, err := s.TokenGenerator.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrFailedToGenerateToken))
	}

	refreshToken, err := s.TokenGenerator.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrFailedToGenerateToken))
	}

	if err := s.UserRepo.UpdateRefreshToken(user.ID, repository.UpdateUserRefreshTokenInput{
		RefreshToken: refreshToken,
	}); err != nil {
		return nil, s.HandleError(NewAuthError(AuthErrors.ErrUpdatingRefreshToken))
	}

	return &LoginOutput{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Type:         "Bearer",
	}, nil
}

type LoginCommand struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Type         string `json:"type"`
}
