package middlewares

import (
	"app/internal/domain/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryDomainError(statusMap map[errors.ErrorCode]int) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if domainErr, ok := r.(*errors.DomainError); ok {
					status := http.StatusInternalServerError
					if statusMap != nil {
						if s, ok := statusMap[domainErr.Code]; ok {
							status = s
						}
					}

					c.JSON(status, gin.H{
						"code":    domainErr.Code,
						"message": domainErr.Message,
					})
					return
				}

				panic(r)
			}
		}()

		c.Next()
	}
}
