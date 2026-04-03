package post

import (
	"app/internal/domain/common"
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

func NewPost(title, content string, authorID uint) *Post {
	return &Post{
		Title:    title,
		Content:  content,
		AuthorID: authorID,
	}
}

func (e *Post) Validate() error {
	return nil
}
