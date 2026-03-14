package authapi

import (
	"app/internal/application/auth"
	"app/internal/infrastructure"
	"app/internal/infrastructure/gin/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	h := newAuthHandler()

	router.POST("/auth/login", h.Login)
	router.POST("/auth/register", h.Register)
}

type authHandler struct{}

func newAuthHandler() *authHandler {
	return &authHandler{}
}

func getAuthServiceFromContext(ctx *gin.Context) auth.AuthService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := infrastructure.NewContainer(tx)
	return container.GetAuthService()
}

var errorHandler = errors.NewErrorHandler(
	httpStatusMap,
)
