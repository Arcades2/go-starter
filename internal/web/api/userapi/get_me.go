package userapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userHandler) GetMe(ctx *gin.Context) {
	userService := GetUserServiceFromContext(ctx)
	userID := ctx.MustGet("userID").(uint)

	user, _ := userService.GetUserByID(userID)

	response := UserResponse{
		ID:        user.ID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	ctx.JSON(http.StatusOK, response)
}
