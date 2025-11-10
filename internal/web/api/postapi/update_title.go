package postapi

import (
	"net/http"
	"strconv"

	"app/internal/domain/services/postservice"

	"github.com/gin-gonic/gin"
)

func (h *postHandler) UpdatePostTitle(ctx *gin.Context) {
	var request UpdatePostTitleRequest

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post id"})
		return
	}

	postService := GetPostServiceFromContext(ctx)

	_ = postService.UpdateTitle(postservice.UpdateTitleCommand{
		ID:    uint(id),
		Title: request.Title,
	})

	ctx.Status(http.StatusNoContent)
}

type UpdatePostTitleRequest struct {
	Title string `json:"title" binding:"required"`
}
