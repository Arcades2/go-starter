package authapi

import (
	"net/http"

	"app/internal/domain/auth"
	"app/internal/infrastructure/http/gin/errors"
)

func init() {
	errors.HTTPStatusMap[auth.ErrInvalidCredentials.Code] = http.StatusUnauthorized
	errors.HTTPStatusMap[auth.ErrFailedToGenerateToken.Code] = http.StatusInternalServerError
	errors.HTTPStatusMap[auth.ErrRegisterInvalidInput.Code] = http.StatusBadRequest
	errors.HTTPStatusMap[auth.ErrLoginInvalidInput.Code] = http.StatusBadRequest
	errors.HTTPStatusMap[auth.ErrHashingPassword.Code] = http.StatusInternalServerError

	errors.EnsureAllErrorsMapped(auth.AllAuthErrorCodes())
}
