package authapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/authservice"
	"app/internal/web/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	authService := services.NewContainer(db).GetAuthService(
		&services.ServiceSettings{
			PanicOnError: true,
		},
	)

	h := newAuthHandler(
		authService,
	)

	router.Use(middlewares.RecoveryDomainError(httpStatusMap))

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
