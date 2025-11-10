package postapi

import (
	"app/internal/domain/services/postservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *postHandler) CreatePost(ctx *gin.Context) {
	var req CreatePostRequest
	ctx.BindJSON(&req)

	userID := ctx.MustGet("userID").(uint)

	postService := GetPostServiceFromContext(ctx)

	command := postservice.CreatePostCommand{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	}
	post, _ := postService.CreatePost(command)

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
