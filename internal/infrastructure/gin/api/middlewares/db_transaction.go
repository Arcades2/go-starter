package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := db.Begin()
		if tx.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
			return
		}

		ctx.Set("tx", tx)

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}

			if ctx.IsAborted() {
				tx.Rollback()
				return
			}

			tx.Commit()
		}()

		ctx.Next()
	}
}
