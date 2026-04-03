package common

type Validatable interface {
	Validate() error
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
