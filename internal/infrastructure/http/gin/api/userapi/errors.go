package userapi

import (
	"net/http"

	"app/internal/domain/user"
	"app/internal/infrastructure/http/gin/errors"
)

func init() {
	errors.HTTPStatusMap[user.ErrUserNotFound.Code] = http.StatusNotFound

	errors.EnsureAllErrorsMapped(user.AllUserErrorCodes())
}
