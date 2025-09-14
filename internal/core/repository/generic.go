package repository

import (
	"gorm.io/gorm"
)

type GenericRepositoryInterface[T any] interface {
	Create(entity *T) error
	Save(entity *T) error
	Delete(id uint) error
	FindByID(id uint) (*T, error)
	UpdateByID(id uint, updates any) error
}

type GenericRepository[T any] struct {
	DB *gorm.DB
}

func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{DB: db}
}

func (r *GenericRepository[T]) Create(entity *T) error {
	return r.DB.Create(entity).Error
}

func (r *GenericRepository[T]) Save(entity *T) error {
	return r.DB.Save(entity).Error
}

func (r *GenericRepository[T]) Delete(id uint) error {
	return r.DB.Delete(new(T), id).Error
}

func (r *GenericRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T]) UpdateByID(id uint, updates any) error {
	return r.DB.Model(new(T)).Where("id = ?", id).Updates(updates).Error
}
