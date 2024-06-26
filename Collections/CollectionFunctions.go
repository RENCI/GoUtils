package Collections

// This was supposed to be a method, but due to limitations of Go generics we are forced to have it as function
func GroupBy[T any, R comparable, M ~map[R]List[T]](col Iterable[T], groupFunc func(item1 T) R) M {

	var res = make(map[R]List[T])

	col.Iterate(func(item *T) bool {
		r := groupFunc(*item)
		g, exists := res[r]
		if !exists {
			g = NewList[T]()
			res[r] = g
		}
		g.Add(*item)

		return true
	})
	return res
}

func GetKeysFromMap[K comparable, V any](m map[K]V) List[K] {
	r := NewList[K]()
	for k := range m {
		r.Add(k)
	}
	return r
}

func CopyMap(m *map[any]bool) *map[any]bool {
	m2 := make(map[any]bool, len(*m))
	var id any
	for id, m2[id] = range *m {
	}
	return &m2
}
