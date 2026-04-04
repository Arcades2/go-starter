package userapi

import (
	"app/internal/application/user"
	"app/internal/infrastructure"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := newUserHandler()

	router.GET("/users/:id", h.GetUserByID)
	router.GET("/users/me", h.GetMe)
}

type userHandler struct{}

func newUserHandler() *userHandler {
	return &userHandler{}
}

func getUserReaderServiceFromContext(ctx *gin.Context) user.UserReader {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := infrastructure.NewContainer(tx)
	return container.GetUserReaderService()
}
