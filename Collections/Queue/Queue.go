package Queue

import (
	Iterable "github.com/RENCI/GoUtils/Collections/Interfaces"
	"github.com/RENCI/GoUtils/Collections/LinkedList"
)

type Queue[T any] struct {
	_dlist *LinkedList.LinkedList[T]
}

func New[T any]() *Queue[T] {
	res := Queue[T]{
		_dlist: LinkedList.New[T](),
	}
	return &res
}

func NewFromSlice[T any](items *[]T) *Queue[T] {
	res := New[T]()
	for _, t := range *items {
		res.Enqueue(t)
	}
	return res
}

func NewFromIterable[T any](items Iterable.Iterable[T]) *Queue[T] {
	res := New[T]()
	items.Iterate(func(item *T) bool {
		res.Enqueue(*item)
		return true
	})
	return res
}

func (this *Queue[T]) Size() int {
	return this._dlist.Size()
}

func (this *Queue[T]) Peek() *T {
	if this._dlist.First() != nil {
		res := this._dlist.First().GetValue()
		return &res
	}
	return nil
}

func (this *Queue[T]) Enqueue(item T) {
	this._dlist.AddLast(item)
}

func (this *Queue[T]) Dequeue() *T {
	resNode := this._dlist.First()
	if resNode == nil {
		return nil
	}
	v := resNode.GetValue()
	this._dlist.RemoveFirst()
	return &v
}
