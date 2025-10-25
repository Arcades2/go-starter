package authapi

import (
	"net/http"

	"app/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Login(c *gin.Context) {
	var input LoginInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	tokens, err := h.AuthService.Login(input.Email, input.Password)
	if appErr, ok := err.(*errors.AppError); ok {
		status := httpStatusMap[appErr.Code]
		c.JSON(status, gin.H{"code": appErr.Code, "message": appErr.Message})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

type LoginInputDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
