package authapi

import (
	"net/http"

	"app/internal/domain/services/authservice"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var request LoginRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
