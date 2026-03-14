package authapi

import (
	"net/http"

	"app/internal/application/auth"
	"app/internal/domain/errors"
	weberrors "app/internal/infrastructure/http/gin/errors"
)

var httpStatusMap = map[errors.ErrorCode]int{
	auth.ErrInvalidCredentials.Code:    http.StatusUnauthorized,
	auth.ErrFailedToGenerateToken.Code: http.StatusInternalServerError,
	auth.ErrUpdatingUser.Code:          http.StatusInternalServerError,
	auth.ErrRegisterInvalidInput.Code:  http.StatusBadRequest,
	auth.ErrHashingPassword.Code:       http.StatusInternalServerError,
}

func init() {
	weberrors.EnsureAllErrorsMapped(auth.AllAuthErrorCodes(), httpStatusMap)
}
