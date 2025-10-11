package authservice

func (s *AuthService) Login(email, password string) (*LoginOutputDTO, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return nil, ErrInvalidCredentials()
	}

	if !s.PasswordHasher.VerifyPassword(password, user.HashedPassword) {
		return nil, ErrInvalidCredentials()
	}

	accessToken, err := s.TokenGenerator.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, ErrFailedToGenerateAccessToken()
	}

	refreshToken, err := s.TokenGenerator.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, ErrFailedToGenerateRefreshToken()
	}

	if err := s.UserRepo.UpdateRefreshToken(user.ID, refreshToken); err != nil {
		return nil, ErrUpdatingUser()
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
