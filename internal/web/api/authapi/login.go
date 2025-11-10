package authapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var request LoginRequestDTO

	c.BindJSON(&request)

	authService := GetAuthServiceFromContext(c)

	tokens, _ := authService.Login(request.Email, request.Password)

	c.JSON(http.StatusOK, tokens)
}

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
