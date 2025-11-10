package postapi

import (
	"net/http"

	"app/internal/domain/errors"
	"app/internal/domain/services/postservice"
	weberrors "app/internal/web/errors"
)

var httpStatusMap = map[errors.ErrorCode]int{
	postservice.PostErrors.ErrPostNotFound.Code:             http.StatusNotFound,
	postservice.PostErrors.ErrPostCreateInvalid.Code:        http.StatusBadRequest,
	postservice.PostErrors.ErrPostCreateAuthorNotFound.Code: http.StatusBadRequest,
	postservice.PostErrors.ErrPostUpdateInvalid.Code:        http.StatusBadRequest,
}

func init() {
	postserviceErrors := postservice.AllPostErrorCodes()
	weberrors.EnsureAllErrorsMapped(postserviceErrors, httpStatusMap)
}
