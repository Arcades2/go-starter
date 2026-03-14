package authapi

import (
	"net/http"

	"app/internal/application/auth"

	"github.com/gin-gonic/gin"
)

func (h *authHandler) Register(c *gin.Context) {
	var request RegisterRequest

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authService := getAuthServiceFromContext(c)

	user, err := authService.Register(auth.RegisterCommand{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		Password:  request.Password,
	})
	if err != nil {
		errorHandler(c, err)
		return
	}

	c.JSON(http.StatusCreated, RegisterResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	})
}

type RegisterRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type RegisterResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
