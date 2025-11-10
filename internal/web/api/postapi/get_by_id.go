package postapi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *postHandler) GetPostByID(ctx *gin.Context) {
	postService := GetPostServiceFromContext(ctx)
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	post, _ := postService.GetByID(uint(id))

	response := PostResponseDTO{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}

	ctx.JSON(http.StatusOK, response)
}

type PostResponseDTO struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
}
