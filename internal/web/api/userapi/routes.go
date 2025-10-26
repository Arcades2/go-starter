package userapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/userservice"
	"app/internal/web/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := newUserHandler()

	router.Use(middlewares.RecoveryDomainError(httpStatusMap))
	router.GET("/users/:id", h.GetUserByID)
	router.GET("/users/me", h.GetMe)
}

type userHandler struct{}

func newUserHandler() *userHandler {
	return &userHandler{}
}

func GetUserServiceFromContext(ctx *gin.Context) *userservice.UserService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := services.NewContainer(tx)
	return container.GetUserService(&services.ServiceSettings{
		PanicOnError: true,
	})
}
