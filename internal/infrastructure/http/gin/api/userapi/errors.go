package userapi

import (
	"net/http"

	"app/internal/application/user"
	"app/internal/domain/errors"
	weberrors "app/internal/infrastructure/http/gin/errors"
)

var httpStatusMap = map[errors.ErrorCode]int{
	user.ErrUserNotFound.Code: http.StatusNotFound,
}

func init() {
	weberrors.EnsureAllErrorsMapped(user.AllUserErrorCodes(), httpStatusMap)
}
