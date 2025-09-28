package repository

type GenericRepository[T any, I any] interface {
	Create(input I) (*T, error)
	Delete(id uint) error
	FindByID(id uint) (*T, error)
	UpdateByID(id uint, updates any) error
}
