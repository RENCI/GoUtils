package Stack

import (
	Iterable "github.com/RENCI/GoUtils/Collections/Interfaces"
	"github.com/RENCI/GoUtils/Collections/LinkedList"
)

type Stack[T any] struct {
	_dlist *LinkedList.LinkedList[T]
}

func New[T any]() *Stack[T] {
	res := Stack[T]{
		_dlist: LinkedList.New[T](),
	}
	return &res
}

func NewFromSlice[T any](items *[]T) *Stack[T] {
	res := New[T]()
	for _, t := range *items {
		res.Push(t)
	}
	return res
}

func NewFromIterable[T any](items Iterable.Iterable[T]) *Stack[T] {
	res := New[T]()
	items.Iterate(func(item *T) bool {
		res.Push(*item)
		return true
	})
	return res
}

func (this *Stack[T]) Size() int {
	return this._dlist.Size()
}

func (this *Stack[T]) Peek() *T {
	if this._dlist.First() != nil {
		res := this._dlist.First().GetValue()
		return &res
	}
	return nil
}

func (this *Stack[T]) Push(item T) {
	this._dlist.AddFirst(item)
}

func (this *Stack[T]) Pop() *T {
	resNode := this._dlist.First()
	if resNode == nil {
		return nil
	}
	v := resNode.GetValue()
	this._dlist.RemoveFirst()
	return &v
}
