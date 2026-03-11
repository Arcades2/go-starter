package post

import "app/internal/domain/common"

type PostRepository interface {
	common.Repository[*Post]
}
