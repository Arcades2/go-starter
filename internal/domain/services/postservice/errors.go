package postservice

import (
	"reflect"

	"app/internal/domain/errors"
)

type PostError = errors.DomainError

var PostErrors = struct {
	ErrPostNotFound             PostError
	ErrPostCreateInvalid        PostError
	ErrPostCreateAuthorNotFound PostError
	ErrPostUpdateInvalid        PostError
}{
	ErrPostNotFound:             PostError{Code: "POST_NOT_FOUND", Message: "post not found"},
	ErrPostCreateInvalid:        PostError{Code: "POST_CREATE_INVALID", Message: "post create invalid"},
	ErrPostCreateAuthorNotFound: PostError{Code: "POST_CREATE_AUTHOR_NOT_FOUND", Message: "post create author not found"},
	ErrPostUpdateInvalid:        PostError{Code: "POST_UPDATE_INVALID", Message: "post update invalid"},
}

func NewPostError(err PostError) *errors.DomainError {
	return &errors.DomainError{
		Code:    errors.ErrorCode(err.Code),
		Message: err.Message,
	}
}

func AllPostErrorCodes() []errors.ErrorCode {
	val := reflect.ValueOf(PostErrors)
	codes := make([]errors.ErrorCode, 0, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if postErr, ok := field.Interface().(PostError); ok {
			codes = append(codes, postErr.Code)
		}
	}

	return codes
}
