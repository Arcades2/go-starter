package userapi

import (
	"net/http"
	"strconv"

	"app/internal/infrastructure/http/gin/errors"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	userService := getUserReaderServiceFromContext(ctx)
	id_param := ctx.Param("id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := userService.GetByID(uint(id))
	if err != nil {
		errors.ErrorHandler(ctx, err)
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) GetMe(ctx *gin.Context) {
	userService := getUserReaderServiceFromContext(ctx)
	userID := ctx.MustGet("userID").(uint)

	user, err := userService.GetByID(userID)
	if err != nil {
		errors.ErrorHandler(ctx, err)
	}

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
