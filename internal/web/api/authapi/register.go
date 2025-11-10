package authapi

import (
	"net/http"

	"app/internal/domain/services/authservice"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Register(c *gin.Context) {
	var request RegisterRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := GetAuthServiceFromContext(c)

	command := authservice.RegisterCommand{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	}

	user, _ := authService.Register(command)

	c.JSON(http.StatusCreated, RegisterResponseDTO{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	})
}

type RegisterRequest struct {
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
