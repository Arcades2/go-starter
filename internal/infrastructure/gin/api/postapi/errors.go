package postapi

import (
	"net/http"

	"app/internal/application/post"
	"app/internal/domain/errors"
	weberrors "app/internal/web/errors"
)

var httpStatusMap = map[errors.ErrorCode]int{
	post.ErrPostNotFound.Code:      http.StatusNotFound,
	post.ErrPostCreateInvalid.Code: http.StatusBadRequest,
	post.ErrPostCreateFailed.Code:  http.StatusInternalServerError,
	post.ErrPostUpdateInvalid.Code: http.StatusBadRequest,
}

func init() {
	weberrors.EnsureAllErrorsMapped(post.AllPostErrorCodes(), httpStatusMap)
}
