package userapi

import (
	"net/http"
	"strconv"

	"app/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	id_param := ctx.Param("id")

	id, err := strconv.Atoi(id_param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	user, err := h.UserService.GetUserByID(uint(id))
	if appErr, ok := err.(*errors.AppError); ok {
		status := httpStatusMap[appErr.Code]
		ctx.JSON(status, gin.H{"code": appErr.Code, "message": appErr.Message})
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

type UserResponse struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
