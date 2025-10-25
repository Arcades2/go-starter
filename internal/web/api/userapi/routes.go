package userapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/userservice"
	"app/internal/web/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userService := services.NewContainer(db).GetUserService(&services.ServiceSettings{
		PanicOnError: true,
	})

	h := newUserHandler(
		userService,
	)

	router.Use(middlewares.RecoveryDomainError(httpStatusMap))
	router.GET("/users/:id", h.GetUserByID)
}

type userHandler struct {
	UserService *userservice.UserService
}

func newUserHandler(userService *userservice.UserService) *userHandler {
	return &userHandler{
		UserService: userService,
	}
}
