package Hashset

type Hashset[K any] struct {
	_imap *map[interface{}]bool
}

func New[K any]() *Hashset[K] {
	return &Hashset[K]{_imap: &map[interface{}]bool{}}
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
