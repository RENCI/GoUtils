package Interfaces

type Iterable[T any] interface {
	Iterate(foreach func(item *T) bool)
}

type Cloneable[T any] interface {
	Clone() T
}
