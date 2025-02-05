package Misc

type Result[T any] struct {
	Value T
	Err   error
}

func NewResult[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Err: err}
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

func (r Result[T]) Unwrap() (T, error) {
	return r.Value, r.Err
}

type Nullable[T any] struct {
	Value T
	IsNil bool
}

func NewNullable[T any]() Nullable[T] {
	return Nullable[T]{IsNil: true}
}
