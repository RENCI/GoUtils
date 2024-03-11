package List

import (
	"github.com/RENCI/GoUtils/Collections/Iterable"
	"slices"
)

// InternalList Internal structure. Do not use in your code directly.
type InternalList[T any] struct {
	_arr []T
}

// List is a wrapper, that provides number of common methods.
type List[T any] struct {
	_ilist *InternalList[T]
}

// New creates new instance of List.
func New[T any]() List[T] {
	ilist := &InternalList[T]{
		_arr: []T{},
	}
	list := List[T]{
		_ilist: ilist,
	}
	return list
}

func NewFromSlice[T any](items *[]T) List[T] {
	list := New[T]()
	for _, i := range *items {
		list.Add(i)
	}
	return list
}

func NewFromIterable[T any](items Iterable.Iterable[T]) List[T] {
	list := New[T]()
	items.Iterate(func(item *T) bool {
		list.Add(*item)
		return true
	})
	return list
}

// New creates new instance of List.
func NewList[T any](size int) List[T] {
	ilist := &InternalList[T]{
		_arr: make([]T, size),
	}
	list := List[T]{
		_ilist: ilist,
	}
	return list
}

// Add adds new item.
func (list List[T]) Add(item T) {
	list._ilist._arr = append(list._ilist._arr, item)
}

// Get gets an item by index.
func (list List[T]) Get(index int) T {
	return list._ilist._arr[index]
}

// Set sets an item to index.
func (list List[T]) Set(index int, value T) {
	list._ilist._arr[index] = value
}

// Size returns size of the list
func (list List[T]) Size() int {
	return len(list._ilist._arr)
}

// AddRange adds slice of items to the list
func (list List[T]) AddRange(items []T) {
	list._ilist._arr = append(list._ilist._arr, items...)
}

// Clear removes all items from the list
func (list List[T]) Clear() {
	list._ilist._arr = []T{}
}

// Any iterates all elements of the list and invokes a function until
// that function returns True for any element, then Any returns true, otherwise it returns False
func (list List[T]) Any(cmpfunc func(item1 T) bool) bool {
	for i := 0; i < list.Size(); i++ {
		if cmpfunc(list.Get(i)) {
			return true
		}
	}
	return false
}

// Sort sorts all elements of the list
func (list List[T]) Sort(cmpfunc func(item1 T, item2 T) int) {
	slices.SortFunc(list._ilist._arr, cmpfunc)
}

// All iterates all elements of the list and checks if it satisfies condition specified in the function f
func (list List[T]) All(f func(item1 T) bool) bool {
	for i := 0; i < list.Size(); i++ {
		if !f(list.Get(i)) {
			return false
		}
	}
	return true
}

// ForEach iterates all elements of the list and invokes specified function f for each element in the list.
func (list List[T]) ForEach(f func(item T)) {
	for i := 0; i < list.Size(); i++ {
		f(list.Get(i))
	}
}

// ForEachIndexed iterates all elements of the list and invokes specified function f for each
// element in the list.
func (list List[T]) ForEachIndexed(f func(item T, index int)) {
	for i := 0; i < list.Size(); i++ {
		f(list.Get(i), i)
	}
}

// IndexOf finds first occurrence of an element in the list.
func (list List[T]) IndexOf(f func(item T) bool) int {
	for i := 0; i < list.Size(); i++ {
		if f(list.Get(i)) {
			return i
		}
	}
	return -1
}

// IndexOfInRange finds first occurrence of an element in the list in specified index range
func (list List[T]) IndexOfInRange(f func(item T) bool, startIndex int, endIndex int) int {
	for i := startIndex; i < list.Size() || i <= endIndex; i++ {
		if f(list.Get(i)) {
			return i
		}
	}
	return -1
}

// RemoveAt removes an element in the list
func (list List[T]) RemoveAt(index int) {
	list._ilist._arr = slices.Delete(list._ilist._arr, index, index+1)
}

// RemoveRange removes specified number of elements in the list starting from index.
func (list List[T]) RemoveRange(index int, count int) {
	list._ilist._arr = slices.Delete(list._ilist._arr, index, index+count)
}

// GetRange returns new list with elements specified by range.
func (list List[T]) GetRange(index int, count int) List[T] {
	newList := New[T]()
	newList.AddRange(list._ilist._arr[index : index+count])
	return newList
}

// Clone returns new instance of the list.
func (list List[T]) Clone() List[T] {
	newList := New[T]()
	newList.AddRange(list._ilist._arr[:])
	return newList
}

// GetSlice returns a copy of the slice.
func (list List[T]) GetSlice() []T {
	res := make([]T, len(list._ilist._arr))
	copy(res, list._ilist._arr)
	return res
}

// First returns the first element.
func (list List[T]) First() T {
	return list.Get(0)
}

// Last returns the last element.
func (list List[T]) Last() T {
	return list.Get(list.Size() - 1)
}

// Insert inserts an element at specified index.
func (list List[T]) Insert(item T, index int) {
	list._ilist._arr = append(list._ilist._arr[:index+1], list._ilist._arr[index:]...)
	list._ilist._arr[index] = item
}

// Push adds an element to the end of the list.
func (list List[T]) Push(item T) {
	list.Add(item)
}

// Pop returns and removes an element at the end of the list.
func (list List[T]) Pop() T {
	res := list.Get(list.Size() - 1)
	list.RemoveAt(list.Size() - 1)
	return res
}

// Peek returns an element at the end of the list wothout removing it.
func (list List[T]) Peek() T {
	return list.Get(list.Size() - 1)
}

// Enqueue adds an element.
func (list List[T]) Enqueue(item T) {
	if list.Size() == 0 {
		list.Add(item)
	} else {
		list.Insert(item, 0)
	}
}

// Dequeue returns and removes an element.
func (list List[T]) Dequeue() T {
	res := list.Get(list.Size() - 1)
	list.RemoveAt(list.Size() - 1)
	return res
}

func (list *List[T]) Iterate(foreach func(item *T) bool) {
	for i := 0; i < list.Size(); i++ {
		if !foreach(&list._ilist._arr[i]) {
			return
		}
	}
}
