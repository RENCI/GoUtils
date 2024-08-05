package Collections

import "iter"

type Hashset[K any] struct {
	_imap *map[interface{}]bool
}

func NewHashset[K any]() *Hashset[K] {
	return &Hashset[K]{_imap: &map[interface{}]bool{}}
}

func NewHashsetFromSlice[K any](items *[]K) *Hashset[K] {
	hs := NewHashset[K]()
	for _, i := range *items {
		hs.Add(i)
	}
	return hs
}

func NewHashsetFromIterable[K any](items Iterable[K]) *Hashset[K] {
	hs := NewHashset[K]()

	for item := range items.GetSeq() {
		hs.Add(item)
	}
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

func (this *Hashset[K]) GetSeq() iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range *this._imap {
			switch t := k.(type) {
			default:
				return
			case K:
				if !yield(t) {
					return
				}
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

func (this *Hashset[K]) UnionWithHS(other *Hashset[K]) {
	for k, _ := range *other._imap {
		this.Add(k.(K))
	}
}

func (this *Hashset[K]) ExceptWithHS(other *Hashset[K]) {
	for k, _ := range *other._imap {
		this.Remove(k.(K))
	}
}

func (this *Hashset[K]) IntersectWithHS(other *Hashset[K]) {
	res := map[interface{}]bool{}

	for k, _ := range *other._imap {
		if this.Contains(k.(K)) {
			res[k] = false
		}
	}

	this._imap = &res
}

func (this *Hashset[K]) ToList() List[K] {
	r := NewList[K]()
	for k := range *this._imap {
		r.Add(k.(K))
	}
	return r
}

func (this *Hashset[K]) Clone() *Hashset[K] {
	res := NewHashset[K]()
	res._imap = CopyMap(this._imap)
	return res
}
