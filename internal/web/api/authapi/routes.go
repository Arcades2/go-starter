package authapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/authservice"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authService := services.NewContainer(db).GetAuthService()

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
