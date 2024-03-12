package LinkedList

type LinkedList[T any] struct {
	_root *LinkedListNode[T]
	_end  *LinkedListNode[T]
	_size int
	_zero T
}

type LinkedListNode[T any] struct {
	_item   T
	_before *LinkedListNode[T]
	_next   *LinkedListNode[T]
}

func (this *LinkedListNode[T]) GetValue() T {
	return this._item
}

func (this *LinkedListNode[T]) GetNext() *LinkedListNode[T] {
	return this._next
}

func (this *LinkedListNode[T]) GetBefore() *LinkedListNode[T] {
	return this._before
}

func _newNode[T any](item T) *LinkedListNode[T] {
	res := LinkedListNode[T]{
		_item: item,
	}
	return &res
}

func New[T any]() *LinkedList[T] {
	var zero = *new(T)
	var zeroNodeRoot = _newNode(zero)
	var zeroNodeEnd = _newNode(zero)

	zeroNodeRoot._next = zeroNodeEnd
	zeroNodeEnd._before = zeroNodeRoot

	res := LinkedList[T]{
		_root: zeroNodeRoot,
		_zero: zero,
		_end:  zeroNodeEnd,
	}
	return &res
}

func (this *LinkedList[T]) Size() int {
	return this._size
}

func (this *LinkedList[T]) First() *LinkedListNode[T] {
	if this._root._next != this._end {
		return this._root._next
	}
	return nil
}

func (this *LinkedList[T]) Last() *LinkedListNode[T] {
	if this._end._before != this._root {
		return this._end._before
	}
	return nil
}

func (this *LinkedList[T]) IsZero(item *T) bool {
	return &this._zero == item
}

func (this *LinkedList[T]) Add(item T) *LinkedListNode[T] {
	return this.InsertBefore(item, this._end)
}

func (this *LinkedList[T]) InsertAfter(item T, after *LinkedListNode[T]) *LinkedListNode[T] {
	return this.InsertBefore(item, after._next)
}

func (this *LinkedList[T]) InsertBefore(item T, before *LinkedListNode[T]) *LinkedListNode[T] {
	var node = _newNode(item)
	node._before = before._before
	node._before._next = node
	node._next = before
	before._before = node
	this._size += 1
	return node
}

func (this *LinkedList[T]) Remove(eq func(a T, b T) bool, value T) {
	c := this._root._next
	for c != this._end {
		if eq(value, c.GetValue()) {
			n1 := c._before
			n2 := c._next
			n1._next = n2
			n2._before = n1
			this._size -= 1
			return
		}
		c = c._next
	}
}
