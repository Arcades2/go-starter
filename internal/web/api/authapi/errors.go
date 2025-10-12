package authapi

import (
	"app/internal/domain/errors"
	"app/internal/domain/services/authservice"
	weberrors "app/internal/web/errors"
	"net/http"
)

var httpStatusMap = map[errors.ErrorCode]int{
	authservice.AuthErrors.ErrInvalidCredentials.Code:    http.StatusUnauthorized,
	authservice.AuthErrors.ErrFailedToGenerateToken.Code: http.StatusInternalServerError,
	authservice.AuthErrors.ErrUpdatingRefreshToken.Code:  http.StatusInternalServerError,
	authservice.AuthErrors.ErrRegisterInvalidInput.Code:  http.StatusBadRequest,
	authservice.AuthErrors.ErrHashingPassword.Code:       http.StatusInternalServerError,
}

func init() {
	authserviceErrors := authservice.AllAuthErrorCodes()
	weberrors.EnsureAllErrorsMapped(authserviceErrors, httpStatusMap)
}
