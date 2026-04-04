package authapi

import (
	"net/http"

	"app/internal/application/auth"
	"app/internal/infrastructure/http/gin/errors"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var request LoginRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := getAuthServiceFromContext(c)

	tokens, err := authService.Login(auth.LoginCommand{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		errors.ErrorHandler(c, err)

		return
	}

	c.JSON(http.StatusOK, tokens)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
