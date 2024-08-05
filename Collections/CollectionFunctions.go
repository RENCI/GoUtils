package Collections

// This was supposed to be a method, but due to limitations of Go generics we are forced to have it as function
func GroupBy[T any, R comparable, M ~map[R]List[T]](col Iterable[T], groupFunc func(item1 T) R) M {

	var res = make(map[R]List[T])

	for item := range col.GetSeq() {
		r := groupFunc(item)
		g, exists := res[r]
		if !exists {
			g = NewList[T]()
			res[r] = g
		}
		g.Add(item)
	}

	return res
}

func CopyMap(m *map[any]bool) *map[any]bool {
	m2 := map[any]bool{}
	for k, v := range *m {
		m2[k] = v
	}
	return &m2
}
