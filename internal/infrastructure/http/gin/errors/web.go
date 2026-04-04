package errors

import (
	"errors"
	"log"
	"net/http"

	domainerrors "app/internal/domain/errors"

	"github.com/gin-gonic/gin"
)

func EnsureAllErrorsMapped(all []domainerrors.ErrorCode) {
	for _, code := range all {
		if _, ok := HTTPStatusMap[code]; !ok {
			log.Panicf("missing HTTP status mapping for error code: %s", code)
		}
	}
}

func ErrorHandler(c *gin.Context, err error) {
	if err == nil {
		return
	}

	var vErr *domainerrors.ValidationError
	if errors.As(err, &vErr) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"errors": vErr.Errors,
		})
		return
	}

	var dErr *domainerrors.DomainError
	if errors.As(err, &dErr) {
		status := http.StatusInternalServerError

		if s, ok := HTTPStatusMap[dErr.Code]; ok {
			status = s
		}

		c.AbortWithStatusJSON(status, gin.H{
			"code":    dErr.Code,
			"message": dErr.Message,
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": "internal server error",
	})
}

var HTTPStatusMap = map[domainerrors.ErrorCode]int{}
