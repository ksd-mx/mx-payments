package presenter

type Presenter[T any] interface {
	Show() ([]byte, error)
	Bind(T) error
}
