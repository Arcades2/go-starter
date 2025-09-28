package repository

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

// GenericRepositoryInterface is a generic CRUD interface.
// Type parameters:
//   - T: entity type (struct)
//   - I: input type for Create (struct)
type GenericRepository[T any, I any] struct {
	DB          *gorm.DB
	Constructor func(I) *T
}

func NewGenericRepository[T any, I any](db *gorm.DB, constructor func(I) *T) *GenericRepository[T, I] {
	return &GenericRepository[T, I]{
		DB:          db,
		Constructor: constructor,
	}
}

func (r *GenericRepository[T, I]) Create(input I) (*T, error) {
	entity := r.Constructor(input)
	if err := r.DB.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *GenericRepository[T, I]) Delete(id uint) error {
	return r.DB.Delete(new(T), id).Error
}

func (r *GenericRepository[T, I]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *GenericRepository[T, I]) UpdateByID(id uint, updates any) error {
	if reflect.TypeOf(updates).Kind() != reflect.Struct {
		return fmt.Errorf("UpdateByID expects a struct, got %T", updates)
	}

	return r.DB.Model(new(T)).Where("id = ?", id).Updates(updates).Error
}
