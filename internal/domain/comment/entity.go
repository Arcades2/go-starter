package comment

import (
	"app/internal/domain/common"
	"app/internal/domain/errors"
	"app/internal/domain/post"
	"app/internal/domain/user"
)

type Comment struct {
	ID       uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Content  string `json:"content" gorm:"size:1024;not null"`
	AuthorID *uint
	Author   user.User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PostID   uint
	Post     post.Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	common.TimestampTracking
}

func NewComment(content string, authorID *uint, postID uint) (*Comment, error) {
	comment := &Comment{
		Content:  content,
		AuthorID: authorID,
		PostID:   postID,
	}
	err := comment.Validate()

	return comment, err
}

func (e *Comment) Validate() error {
	var errs []*errors.DomainError

	if e.Content == "" || len(e.Content) > 1024 {
		errs = append(errs, ErrCommentInvalidContent)
	}

	if e.AuthorID != nil && *e.AuthorID <= 0 {
		errs = append(errs, ErrCommentInvalidAuthorID)
	}

	if e.PostID <= 0 {
		errs = append(errs, ErrCommentInvalidPostID)
	}

	if len(errs) > 0 {
		return &errors.ValidationError{
			Errors: errs,
		}
	}

	return nil
}
