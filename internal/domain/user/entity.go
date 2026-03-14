// Package user defines the User entity and its fields, along with GORM annotations for database mapping.
package user

import (
	"net/mail"

	"app/internal/domain/common"
	"app/internal/domain/errors"
)

type User struct {
	ID             uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Email          string `json:"email" gorm:"uniqueIndex;not null"`
	Firstname      string `json:"firstname" gorm:"size:255;not null"`
	Lastname       string `json:"lastname" gorm:"size:255;not null"`
	HashedPassword string `json:"hashedPassword" gorm:"size:255;not null"`
	RefreshToken   string `json:"refreshToken" gorm:"size:255"`
	IsActive       bool   `json:"isActive" gorm:"default:true"`
	common.TimestampTracking
}

func (e *User) Validate() error {
	var errs []*errors.DomainError

	if e.Email == "" || len(e.Email) > 255 {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid email"))
	} else if _, err := mail.ParseAddress(e.Email); err != nil {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid email"))
	}

	if e.Firstname == "" || len(e.Firstname) > 255 {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid firstname"))
	}

	if e.Lastname == "" || len(e.Lastname) > 255 {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid lastname"))
	}

	if e.HashedPassword == "" || len(e.HashedPassword) > 255 {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid hashed password"))
	}

	if e.RefreshToken != "" && len(e.RefreshToken) > 255 {
		errs = append(errs, NewUserError(UserErrors.ErrInvalidUser, "invalid refresh token"))
	}

	if len(errs) > 0 {
		return &errors.ValidationError{
			Errors: errs,
		}
	}

	return nil
}
