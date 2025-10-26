package api

import (
	"app/internal/web/api/authapi"
	"app/internal/web/api/middlewares"
	"app/internal/web/api/userapi"

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
	}

	return router
}
