package user

import (
	"app/internal/domain/common"
)

type UserRepository interface {
	common.Repository[*User]
	FindByEmail(email string) (*User, error)
}
