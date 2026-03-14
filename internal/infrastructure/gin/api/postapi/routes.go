package postapi

import (
	"app/internal/application/post"
	"app/internal/infrastructure"
	"app/internal/infrastructure/gin/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPostRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := newPostHandler()

	router.POST("/posts", h.CreatePost)
	router.GET("/posts/:id", h.GetPostByID)
	router.PATCH("/posts/:id/title", h.UpdatePostTitle)
}

type postHandler struct{}

func newPostHandler() *postHandler {
	return &postHandler{}
}

func getPostServiceFromContext(ctx *gin.Context) post.PostService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := infrastructure.NewContainer(tx)
	return container.GetPostService()
}

func getPostReaderServiceFromContext(ctx *gin.Context) post.PostReaderService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := infrastructure.NewContainer(tx)
	return container.GetPostReaderService()
}

var errorHandler = errors.NewErrorHandler(
	httpStatusMap,
)
