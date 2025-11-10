package authapi

import (
	"app/internal/domain/services/authservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var request LoginRequest

	c.BindJSON(&request)

	authService := GetAuthServiceFromContext(c)

	tokens, _ := authService.Login(authservice.LoginCommand{
		Email:    request.Email,
		Password: request.Password,
	})

	c.JSON(http.StatusOK, tokens)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
