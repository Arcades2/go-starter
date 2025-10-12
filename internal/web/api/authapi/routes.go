package authapi

import (
	"app/internal/domain/services/authservice"
	"app/internal/infrastructure/auth"
	"app/internal/infrastructure/gorm/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(db)
	passwordHasher := auth.NewPasswordHasher()
	tokenGenerator := auth.NewTokenGenerator()

	authService := authservice.NewAuthService(
		userRepo,
		passwordHasher,
		tokenGenerator,
	)

	h := newAuthHandler(
		authService,
	)

	router.POST("/auth/login", h.Login)
	router.POST("/auth/register", h.Register)
}

type authHandler struct {
	AuthService *authservice.AuthService
}

func newAuthHandler(authService *authservice.AuthService) *authHandler {
	return &authHandler{
		AuthService: authService,
	}
}
