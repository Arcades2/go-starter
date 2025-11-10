package repository

import (
	"app/internal/domain/model"
	"app/internal/domain/repository"

	"gorm.io/gorm"
)

type GormPostRepository struct {
	*GenericRepository[model.Post, repository.CreatePostInput]
}

func NewGormPostRepository(db *gorm.DB) *GormPostRepository {
	return &GormPostRepository{
		NewGenericRepository(
			db,
			func(input repository.CreatePostInput) *model.Post {
				return &model.Post{
					Title:    input.Title,
					Content:  input.Content,
					AuthorID: input.AuthorID,
				}
			},
		),
	}
}

func (r *GormPostRepository) UpdateTitle(id uint, updates repository.UpdatePostTitleInput) error {
	return r.UpdateByID(id, updates)
}
