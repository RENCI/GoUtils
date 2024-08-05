package Collections

type Queue[T any] struct {
	_dlist *LinkedList[T]
}

func NewQueue[T any]() *Queue[T] {
	res := Queue[T]{
		_dlist: NewLinkedList[T](),
	}
	return &res
}

func NewQueueFromSlice[T any](items *[]T) *Queue[T] {
	res := NewQueue[T]()
	for _, t := range *items {
		res.Enqueue(t)
	}
	return res
}

func NewQueueFromIterable[T any](items Iterable[T]) *Queue[T] {
	res := NewQueue[T]()
	for item := range items.GetSeq() {
		res.Enqueue(item)
	}
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
