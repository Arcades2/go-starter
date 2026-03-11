// Package common defines a generic Repository interface for CRUD operations on entities of any type T. This interface can be implemented by various data storage mechanisms (e.g., databases, in-memory storage) to provide a consistent way to manage entities across the application.
package common

type Validatable interface {
	Validate() []error
}

type Repository[T Validatable] interface {
	Create(entity T) error
	CreateMany(entities *[]T) error
	FindByID(id uint) (T, error)
	FindAll() ([]T, error)
	Update(entity T) error
	Delete(entity T) error
	DeleteByID(id uint) error
	Count() (int64, error)
}
