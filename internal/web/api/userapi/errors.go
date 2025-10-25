package userapi

import (
	"net/http"

	"app/internal/domain/errors"
	"app/internal/domain/services/userservice"
	weberrors "app/internal/web/errors"
)

var httpStatusMap = map[errors.ErrorCode]int{
	userservice.UserErrors.ErrUserNotFound.Code: http.StatusNotFound,
}

func init() {
	userserviceErrors := userservice.AllUserErrorCodes()
	weberrors.EnsureAllErrorsMapped(userserviceErrors, httpStatusMap)
}
