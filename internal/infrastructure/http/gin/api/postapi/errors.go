package postapi

import (
	"net/http"

	"app/internal/domain/post"
	"app/internal/infrastructure/http/gin/errors"
)

func init() {
	errors.HTTPStatusMap[post.ErrPostNotFound.Code] = http.StatusNotFound

	errors.EnsureAllErrorsMapped(post.AllPostErrorCodes())
}
