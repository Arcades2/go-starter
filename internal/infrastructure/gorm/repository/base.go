package repository

import (
	"app/internal/domain/common"

	"gorm.io/gorm"
)

type Repository[T common.Validatable] struct {
	db *gorm.DB
}

func NewRepository[T common.Validatable](db *gorm.DB) *Repository[T] {
	return &Repository[T]{db: db}
}

func (r *Repository[T]) Create(entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}
	return r.db.Create(entity).Error
}

func (r *Repository[T]) CreateMany(entities *[]T) error {
	for i := range *entities {
		if err := (*entities)[i].Validate(); err != nil {
			return err
		}
	}
	return r.db.Create(entities).Error
}

func (r *Repository[T]) FindByID(id uint) (T, error) {
	var entity T

	err := r.db.First(&entity, id).Error
	if err != nil {
		var zero T
		return zero, err
	}

	return entity, nil
}

func (r *Repository[T]) FindAll() ([]T, error) {
	var entities []T

	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *Repository[T]) Update(entity T) error {
	if err := entity.Validate(); err != nil {
		return err
	}
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity T) error {
	return r.db.Delete(entity).Error
}

func (r *Repository[T]) DeleteByID(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}

func (r *Repository[T]) Count() (int64, error) {
	var count int64
	var entity T

	err := r.db.Model(&entity).Count(&count).Error
	return count, err
}
