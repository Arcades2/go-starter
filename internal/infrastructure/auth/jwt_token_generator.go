package auth

import (
	"time"

	"app/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

type TokenGenerator struct{}

func NewTokenGenerator() *TokenGenerator {
	return &TokenGenerator{}
}

type JwtClaims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

func (t *TokenGenerator) GenerateAccessToken(userID uint) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(config.AccessTokenDuration),
			),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}

func (t *TokenGenerator) GenerateRefreshToken(userID uint) (string, error) {
	claims := JwtClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(config.RefreshTokenDuration),
			),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JWTSecret))
}
