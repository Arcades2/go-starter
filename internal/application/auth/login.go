package auth

import (
	"app/internal/domain/errors"
)

func (s *authService) Login(cmd LoginCommand) (*LoginOutput, error) {
	user, err := s.userRepo.FindByEmail(cmd.Email)
	if err != nil {
		return nil, s.HandleError(ErrUserNotFound)
	}

	if !s.passwordHasher.VerifyPassword(cmd.Password, user.HashedPassword) {
		return nil, s.HandleError(ErrInvalidCredentials)
	}

	accessToken, err := s.tokenGenerator.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, s.HandleError(errors.WithMessage(ErrFailedToGenerateToken, "failed to generate access token"))
	}

	refreshToken, err := s.tokenGenerator.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, s.HandleError(errors.WithMessage(ErrFailedToGenerateToken, "failed to generate refresh token"))
	}

	user.RefreshToken = refreshToken

	err = s.userRepo.Update(user)
	if err != nil {
		return nil, s.HandleError(errors.WithMessage(ErrUpdatingUser, "failed to save refresh token"))
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
