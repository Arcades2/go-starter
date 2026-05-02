package post

import (
	"app/internal/domain/common"
	"app/internal/domain/errors"
	"app/internal/domain/user"
)

type Post struct {
	ID       uint   `json:"id" gorm:"primaryKey,autoIncrement"`
	Title    string `json:"title" gorm:"size:255;not null"`
	Content  string `json:"content" gorm:"type:text;not null"`
	AuthorID uint
	Author   user.User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	common.TimestampTracking
}

func NewPost(title, content string, authorID uint) (*Post, error) {
	post := &Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}
	error := post.Validate()

	return post, error
}

func (e *Post) Validate() error {
	var errs []*errors.DomainError

	if e.Title == "" || len(e.Title) > 255 {
		errs = append(errs, ErrPostInvalidTitle)
	}

	if e.Content == "" {
		errs = append(errs, ErrPostInvalidContent)
	}

	if e.AuthorID <= 0 {
		errs = append(errs, ErrPostInvalidAuthorID)
	}

	if len(errs) > 0 {
		return &errors.ValidationError{
			Errors: errs,
		}
	}

	return nil
}
