package postapi

import (
	"net/http"

	"app/internal/domain/services/postservice"

	"github.com/gin-gonic/gin"
)

func (h *postHandler) CreatePost(ctx *gin.Context) {
	var req CreatePostRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := ctx.MustGet("userID").(uint)

	postService := GetPostServiceFromContext(ctx)

	command := postservice.CreatePostCommand{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}
	post, _ := postService.Create(command)

	response := CreatePostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
	}

	ctx.JSON(http.StatusCreated, response)
}

type CreatePostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CreatePostResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint   `json:"author_id"`
}
