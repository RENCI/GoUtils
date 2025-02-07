package Misc

import "sync"

type Singleton[T any] struct {
	once sync.Once
	Val  Nullable[T]
	Init func() Result[T]
}

func NewSingleton[T any](f func() Result[T]) *Singleton[T] {
	return &Singleton[T]{
		Init: f,
		Val:  NewNullable[T](),
	}
}

func (this *Singleton[T]) Get() Result[T] {
	var err error
	this.once.Do(func() {
		if this.Val.IsNil {
			newObj := this.Init()
			this.Val.Value = newObj.Value
			err = newObj.Err
		}
	})
	return Result[T]{Value: this.Val.Value, Err: err}
}
