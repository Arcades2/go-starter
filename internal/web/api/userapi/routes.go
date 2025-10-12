package userapi

import (
	"app/internal/domain/services/userservice"
	"app/internal/infrastructure/auth"
	"app/internal/infrastructure/gorm/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := repository.NewGormUserRepository(db)
	passwordHasher := auth.NewPasswordHasher()

	userService := userservice.NewUserService(
		userRepo,
		passwordHasher,
	)

	h := newUserHandler(
		userService,
	)

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
