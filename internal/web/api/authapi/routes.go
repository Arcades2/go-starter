package authapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/authservice"
	"app/internal/web/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	h := newAuthHandler()

	router.Use(middlewares.RecoveryDomainError(httpStatusMap))

	router.POST("/auth/login", h.Login)
	router.POST("/auth/register", h.Register)
}

type authHandler struct {
}

func newAuthHandler() *authHandler {
	return &authHandler{}
}

func GetAuthServiceFromContext(ctx *gin.Context) authservice.AuthService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := services.NewContainer(tx)
	return container.GetAuthService(&services.ServiceSettings{
		PanicOnError: true,
	})
}
