package Collections

import "iter"

type Iterable[T any] interface {
	GetSeq() iter.Seq[T]
}

type Cloneable[T any] interface {
	Clone() T
}
