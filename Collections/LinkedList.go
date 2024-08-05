package Collections

import "iter"

type LinkedList[T any] struct {
	_root *LinkedListNode[T]
	_end  *LinkedListNode[T]
	_size int
}

type LinkedListNode[T any] struct {
	_item     T
	_before   *LinkedListNode[T]
	_next     *LinkedListNode[T]
	_readonly bool
}

func (this *LinkedListNode[T]) GetValue() T {
	return this._item
}

func (this *LinkedListNode[T]) GetNext() *LinkedListNode[T] {
	if !this._next._readonly {
		return this._next
	}
	return nil
}

func (this *LinkedListNode[T]) GetBefore() *LinkedListNode[T] {
	if !this._before._readonly {
		return this._before
	}
	return nil
}

func _newNode[T any](item T) *LinkedListNode[T] {
	res := LinkedListNode[T]{
		_item: item,
	}
	return &res
}

func _newReadOnlyNode[T any]() *LinkedListNode[T] {
	res := LinkedListNode[T]{
		_item:     *new(T),
		_readonly: true,
	}
	return &res
}

func NewLinkedList[T any]() *LinkedList[T] {
	var zeroNodeRoot = _newReadOnlyNode[T]()
	var zeroNodeEnd = _newReadOnlyNode[T]()

	zeroNodeRoot._next = zeroNodeEnd
	zeroNodeEnd._before = zeroNodeRoot

	res := LinkedList[T]{
		_root: zeroNodeRoot,
		_end:  zeroNodeEnd,
	}
	return &res
}

func NewLinkedListFromSlice[T any](items *[]T) *LinkedList[T] {
	list := NewLinkedList[T]()
	for _, i := range *items {
		list.AddLast(i)
	}
	return list
}

func NewLinkedListFromIterable[T any](items Iterable[T]) *LinkedList[T] {
	list := NewLinkedList[T]()
	for item := range items.GetSeq() {
		list.AddLast(item)
	}
	return list
}

func (this *LinkedList[T]) Size() int {
	return this._size
}

func (this *LinkedList[T]) First() *LinkedListNode[T] {
	if !this._root._next._readonly {
		return this._root._next
	}
	return nil
}

func (this *LinkedList[T]) Last() *LinkedListNode[T] {
	if !this._end._before._readonly {
		return this._end._before
	}
	return nil
}

func (this *LinkedList[T]) AddFirst(item T) *LinkedListNode[T] {
	return this.AddBefore(item, this._root._next)
}

func (this *LinkedList[T]) AddLast(item T) *LinkedListNode[T] {
	return this.AddBefore(item, this._end)
}

func (this *LinkedList[T]) AddAfter(item T, after *LinkedListNode[T]) *LinkedListNode[T] {
	return this.AddBefore(item, after._next)
}

func (this *LinkedList[T]) AddBefore(item T, before *LinkedListNode[T]) *LinkedListNode[T] {
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
	for c._readonly == false {
		if eq(value, c.GetValue()) {
			this.RemoveNode(c)
			return
		}
		c = c._next
	}
}

func (this *LinkedList[T]) Find(chk func(a T) bool) *LinkedListNode[T] {
	c := this._root._next
	for c._readonly == false {
		if chk(c.GetValue()) {
			return c
		}
		c = c._next
	}

	return nil
}

func (this *LinkedList[T]) RemoveFirst() {
	if !this._root._next._readonly {
		this.RemoveNode(this._root._next)
	}
}

func (this *LinkedList[T]) RemoveLast() {
	if !this._end._before._readonly {
		this.RemoveNode(this._end._before)
	}
}

func (this *LinkedList[T]) RemoveNode(node *LinkedListNode[T]) {
	if !node._readonly {
		n1 := node._before
		n2 := node._next
		n1._next = n2
		n2._before = n1
		this._size -= 1
	}
}

func (this *LinkedList[K]) GetSeq() iter.Seq[K] {
	return func(yield func(K) bool) {
		c := this._root._next
		for c._readonly == false {
			val := c.GetValue()
			do_next := yield(val)
			if !do_next {
				return
			}
			c = c._next
		}
	}
}

func (this *LinkedList[T]) Clone() *LinkedList[T] {
	r := NewLinkedList[T]()
	for item := range this.GetSeq() {
		r.AddLast(item)
	}
	return r
}
