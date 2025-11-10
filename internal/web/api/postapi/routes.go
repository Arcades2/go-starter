package postapi

import (
	"app/internal/domain/services"
	"app/internal/domain/services/postservice"
	"app/internal/web/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPostRoutes(router *gin.RouterGroup, db *gorm.DB) {
	h := newPostHandler()

	router.Use(middlewares.RecoveryDomainError(httpStatusMap))

	router.POST("/posts", h.CreatePost)
}

type postHandler struct{}

func newPostHandler() *postHandler {
	return &postHandler{}
}

func GetPostServiceFromContext(ctx *gin.Context) *postservice.PostService {
	tx := ctx.MustGet("tx").(*gorm.DB)
	container := services.NewContainer(tx)
	return container.GetPostService(&services.ServiceSettings{
		PanicOnError: true,
	})
}
