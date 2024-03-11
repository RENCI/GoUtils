package Iterable

type Iterable[T any] interface {
	Iterate(foreach func(item T) bool)
}
