package userapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	userService := GetUserServiceFromContext(ctx)
	id_param := ctx.Param("id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, _ := userService.GetUserByID(uint(id))

	response := UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
