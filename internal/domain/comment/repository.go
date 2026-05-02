package comment

import "app/internal/domain/common"

type CommentRepository interface {
	common.Repository[*Comment]
}
