package Tests_test

import (
	"testing"

	. "github.com/RENCI/GoUtils/Collections"
)
import "github.com/stretchr/testify/assert"

func Test_NewGoList(t *testing.T) {
	list := NewList[int]()
	assert.Equal(t, 0, list.Size())
}

func Test_Add_Get(t *testing.T) {
	list := NewList[int]()
	list.Add(1)
	list.Add(2)
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, 2, list.Get(1))
}

func Test_AddRange(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 5, list.Size())
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, 5, list.Get(4))
}

func Test_Clean(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	list.Clear()
	assert.Equal(t, 0, list.Size())
}

func Test_Any_Pos(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.Any(func(item1 int) bool {
		return item1 == 3
	})
	assert.Equal(t, true, res)
}

func Test_Has_Neg(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.Any(func(item1 int) bool {
		return item1 == 10
	})
	assert.Equal(t, false, res)
}

func Test_Sort(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{10, 2, 13, 3, 30, 25})
	list.Sort(func(item1 int, item2 int) int {
		return item1 - item2
	})
	assert.Equal(t, 2, list.Get(0))
	assert.Equal(t, 3, list.Get(1))
	assert.Equal(t, 10, list.Get(2))
	assert.Equal(t, 13, list.Get(3))
	assert.Equal(t, 25, list.Get(4))
	assert.Equal(t, 30, list.Get(5))
}

func Test_All_Neg(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.All(func(item1 int) bool {
		return item1 == 10
	})
	assert.Equal(t, false, res)
}

func Test_All_Pos(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.All(func(item1 int) bool {
		return item1 > 0
	})
	assert.Equal(t, true, res)
}

func Test_ForEach(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	sum := 0
	list.ForEach(func(item1 int) {
		sum += item1
	})
	assert.Equal(t, 15, sum)
}

func Test_ForEachIndexed(t *testing.T) {
	list := NewList[int]()
	ints := []int{1, 2, 3, 4, 5}
	indexes := []int{0, 1, 2, 3, 4}
	list.AddRange(ints)

	list.ForEachIndexed(func(item1 int, index int) {
		assert.Equal(t, ints[index], item1)
		assert.Equal(t, indexes[index], index)
	})

}

func Test_IndexOf(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.IndexOf(func(item1 int) bool {
		return item1 == 3
	})
	assert.Equal(t, 2, res)
}

func Test_IndexOf_neg(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.IndexOf(func(item1 int) bool {
		return item1 == 10
	})
	assert.Equal(t, -1, res)
}

func Test_IndexOfRange(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.IndexOfInRange(func(item1 int) bool {
		return item1 == 3
	}, 1, 4)
	assert.Equal(t, 2, res)
}

func Test_IndexOfRange_neg(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	res := list.IndexOfInRange(func(item1 int) bool {
		return item1 == 3
	}, 3, 4)
	assert.Equal(t, -1, res)
}

func Test_RemoveAt(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	list.RemoveAt(0)
	assert.Equal(t, 4, list.Size())
	assert.Equal(t, 2, list.Get(0))
}

func Test_RemoveRange(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5})
	list.RemoveRange(1, 2)
	assert.Equal(t, 3, list.Size())
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, 4, list.Get(1))
}

func Test_GetRange(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	res := list.GetRange(3, 3)
	assert.Equal(t, 3, res.Size())
	assert.Equal(t, 4, res.Get(0))
	assert.Equal(t, 5, res.Get(1))
	assert.Equal(t, 6, res.Get(2))
}

func Test_GetRange_no_impact_to_original(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	res := list.GetRange(3, 3)
	res.Set(0, -1)

	assert.Equal(t, 10, list.Size())
	assert.Equal(t, 4, list.Get(3))
}

func Test_GetClone(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	res := list.Clone()
	assert.Equal(t, 10, res.Size())
	assert.Equal(t, 4, res.Get(3))
	assert.Equal(t, 5, res.Get(4))
	assert.Equal(t, 6, res.Get(5))
}

func Test_GetClone_no_impact_to_original(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	res := list.Clone()
	list.Set(3, -1)
	assert.Equal(t, 10, res.Size())
	assert.Equal(t, 4, res.Get(3))
	assert.Equal(t, 5, res.Get(4))
	assert.Equal(t, 6, res.Get(5))
}

func Test_GetSlice_no_impact_to_original(t *testing.T) {
	list := NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})
	res := list.GetSlice()
	assert.Equal(t, 10, len(res))
	assert.Equal(t, 4, res[3])
	assert.Equal(t, 5, res[4])
	assert.Equal(t, 6, res[5])
}

func Test_PushPopPeek(t *testing.T) {
	list := NewList[int]()
	list.Push(1)
	list.Push(2)
	list.Push(3)

	assert.Equal(t, 3, list.Peek())
	assert.Equal(t, 3, list.Pop())
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Pop())
	assert.Equal(t, 1, list.Pop())
	assert.Equal(t, 0, list.Size())
}

func Test_EnqueueDequeue(t *testing.T) {
	list := NewList[int]()
	list.Enqueue(1)
	list.Enqueue(2)
	list.Enqueue(3)

	assert.Equal(t, 1, list.Dequeue())
	assert.Equal(t, 2, list.Size())
	assert.Equal(t, 2, list.Dequeue())
	assert.Equal(t, 3, list.Dequeue())
	assert.Equal(t, 0, list.Size())
}

func Test_slices_ref(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}

	f := func(s *[]int) {
		for i := 0; i < 100; i++ {
			*s = append(*s, 1)
		}
	}

	f(&s1)

	assert.Equal(t, 106, len(s1))
}

func Test_slices_val(t *testing.T) {
	s1 := []int{1, 2, 3, 4, 5, 6}

	f := func(s []int) {
		for i := 0; i < 100; i++ {
			s = append(s, 1)
		}
	}

	f(s1)

	assert.Equal(t, 6, len(s1))
}
