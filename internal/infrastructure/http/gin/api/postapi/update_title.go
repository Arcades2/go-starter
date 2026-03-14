package postapi

import (
	"net/http"
	"strconv"

	"app/internal/application/post"

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

	postService := getPostServiceFromContext(ctx)

	_, err = postService.UpdateTitle(post.UpdateTitleCommand{
		ID:    uint(id),
		Title: request.Title,
	})
	if err != nil {
		errorHandler(ctx, err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

type UpdatePostTitleRequest struct {
	Title string `json:"title" binding:"required"`
}
