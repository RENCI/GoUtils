package Hashset

import (
	Iterable "github.com/RENCI/GoUtils/Collections/Interfaces"
)

type Hashset[K any] struct {
	_imap *map[interface{}]bool
}

func New[K any]() *Hashset[K] {
	return &Hashset[K]{_imap: &map[interface{}]bool{}}
}

func NewFromSlice[K any](items *[]K) *Hashset[K] {
	hs := New[K]()
	for _, i := range *items {
		hs.Add(i)
	}
	return hs
}

func NewFromIterable[K any](items Iterable.Iterable[K]) *Hashset[K] {
	hs := New[K]()

	items.Iterate(func(item *K) bool {
		hs.Add(*item)
		return true
	})
	return hs
}

func (this *Hashset[K]) Add(item K) {
	(*this._imap)[item] = false
}

func (this *Hashset[K]) Remove(item K) {
	delete((*this._imap), item)
}

func (this *Hashset[K]) Clear() {
	clear(*this._imap)
}

func (this *Hashset[K]) Size() int {
	return len(*this._imap)
}

func (this *Hashset[K]) Contains(key K) bool {
	_, ok := (*this._imap)[key]
	if ok {
		return true
	}

	return false
}

func (this *Hashset[K]) Iterate(foreach func(item *K) bool) {
	for k, _ := range *this._imap {
		switch t := k.(type) {
		default:
			return
		case K:
			if !foreach(&t) {
				return
			}
		}
	}
}

func (this *Hashset[K]) ExceptWith(other *[]K) {
	for _, k := range *other {
		this.Remove(k)
	}
}

func (this *Hashset[K]) IntersectWith(other *[]K) {
	res := map[interface{}]bool{}

	for _, k := range *other {
		if this.Contains(k) {
			res[k] = false
		}
	}

	this._imap = &res
}

func (this *Hashset[K]) UnionWith(other *[]K) {
	for _, k := range *other {
		this.Add(k)
	}
}

func (this *Hashset[K]) Clone() *Hashset[K] {
	res := New[K]()
	res._imap = copyMap(this._imap)
	return res
}

func copyMap(m *map[interface{}]bool) *map[interface{}]bool {
	m2 := make(map[interface{}]bool, len(*m))
	var id interface{}
	for id, m2[id] = range *m {
	}
	return &m2
}
