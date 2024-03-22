package Collections

type Stack[T any] struct {
	_dlist *LinkedList[T]
}

func NewStack[T any]() *Stack[T] {
	res := Stack[T]{
		_dlist: NewLinkedList[T](),
	}
	return &res
}

func NewStackFromSlice[T any](items *[]T) *Stack[T] {
	res := NewStack[T]()
	for _, t := range *items {
		res.Push(t)
	}
	return res
}

func NewStackFromIterable[T any](items Iterable[T]) *Stack[T] {
	res := NewStack[T]()
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
