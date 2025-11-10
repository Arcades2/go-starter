package repository

import "app/internal/domain/model"

type PostRepository interface {
	GenericRepository[model.Post, CreatePostInput]
	UpdateTitle(id uint, updates UpdatePostTitleInput) error
}

type CreatePostInput struct {
	Title    string
	Content  string
	AuthorID uint
}

type UpdatePostTitleInput struct {
	Title string
}
