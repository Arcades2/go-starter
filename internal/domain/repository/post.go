package repository

import "app/internal/domain/model"

type PostRepository interface {
	GenericRepository[model.Post, CreatePostInput]
}

type CreatePostInput struct {
	Title    string
	Content  string
	AuthorID uint
}
