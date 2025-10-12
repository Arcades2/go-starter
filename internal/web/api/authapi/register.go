package authapi

import (
	"net/http"

	"app/internal/domain/errors"
	"app/internal/domain/services/authservice"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Register(c *gin.Context) {
	var input RegisterInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	command := authservice.RegisterCommand{
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Email:     input.Email,
		Password:  input.Password,
	}

	user, err := h.AuthService.Register(command)
	if appErr, ok := err.(*errors.AppError); ok {
		status := httpStatusMap[appErr.Code]
		c.JSON(status, gin.H{"code": appErr.Code, "message": appErr.Message})
		return
	}

	c.JSON(http.StatusCreated, RegisterResponseDTO{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	})
}

type RegisterInputDTO struct {
	Firstname string `json:"firstname" binding:"required,min=1,max=255"`
	Lastname  string `json:"lastname" binding:"required,min=1,max=255"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=8,max=100"`
}

type RegisterResponseDTO struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
