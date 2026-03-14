package postapi

import (
	"net/http"

	"app/internal/application/post"

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

	postService := getPostServiceFromContext(ctx)

	post, err := postService.Create(post.CreatePostCommand{
		Title:    req.Title,
		Content:  req.Content,
		AuthorID: userID,
	})
	if err != nil {
		errorHandler(ctx, err)
	}

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
