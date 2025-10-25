package baseservice

type Panicable interface {
	SetPanicOnError(bool)
}

type BaseService struct {
	PanicOnError bool
}

func (b *BaseService) HandleError(err error) error {
	if err == nil {
		return nil
	}
	if b.PanicOnError {
		panic(err)
	}
	return err
}

func (b *BaseService) SetPanicOnError(enable bool) {
	b.PanicOnError = enable
}

type Option[T Panicable] func(T)

func WithPanicOnError[T Panicable](svc T) {
	svc.SetPanicOnError(true)
}

func NewBaseService() BaseService {
	return BaseService{
		PanicOnError: false,
	}
}
