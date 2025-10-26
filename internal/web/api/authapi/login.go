package authapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var input LoginInputDTO

	c.BindJSON(&input)

	authService := GetAuthServiceFromContext(c)

	tokens, _ := authService.Login(input.Email, input.Password)

	c.JSON(http.StatusOK, tokens)
}

type LoginInputDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
