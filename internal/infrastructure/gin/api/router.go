package api

import (
	"app/internal/infrastructure/gin/api/authapi"
	"app/internal/infrastructure/gin/api/middlewares"
	"app/internal/infrastructure/gin/api/postapi"
	"app/internal/infrastructure/gin/api/userapi"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.TransactionMiddleware(db))

	authapi.RegisterAuthRoutes(router, db)

	api := router.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		userapi.RegisterUserRoutes(api, db)
		postapi.RegisterPostRoutes(api, db)
	}

	return router
}
