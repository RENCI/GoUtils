package Collections

import (
	"github.com/RENCI/GoUtils/Collections/LinkedList"
	"github.com/RENCI/GoUtils/Collections/List"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	list := LinkedList.New[int]()
	assert.Equal(t, 0, list.Size())
}

func Test_LLNewFromSlice(t *testing.T) {
	hs := LinkedList.NewFromSlice[int](&([]int{1, 2, 3, 4}))
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
}

func Test_LLNewFromIterable(t *testing.T) {
	list := List.NewFromSlice(&([]int{1, 2, 3, 4}))
	hs := LinkedList.NewFromIterable[int](&list)
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
}

func Test_AddLast(t *testing.T) {
	list := LinkedList.New[int]()
	list.AddLast(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.First().GetValue())
	assert.Equal(t, 1, list.Last().GetValue())
}

func Test_AddLast_2(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.Last().GetValue())
}

func Test_AddLast_3(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	hs.AddLast(3)
	assert.Equal(t, 3, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 2, hs.Last().GetBefore().GetValue())
	assert.Equal(t, 3, hs.Last().GetValue())
}

func Test_Remove_1(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	hs.AddLast(3)
	hs.Remove(func(a int, b int) bool {
		return a == b
	}, 3)
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 2, hs.Last().GetValue())
}

func Test_AddBefore(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	n := hs.AddLast(3)
	hs.AddLast(4)

	hs.AddBefore(2, n)
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 3, hs.First().GetNext().GetNext().GetValue())
	assert.Equal(t, 4, hs.Last().GetValue())
}

func Test_AddAfter(t *testing.T) {
	hs := LinkedList.New[int]()
	n := hs.AddLast(1)
	hs.AddLast(3)
	hs.AddLast(4)

	hs.AddAfter(2, n)
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 3, hs.First().GetNext().GetNext().GetValue())
	assert.Equal(t, 4, hs.Last().GetValue())
}

func Test_GetNext(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Nil(t, hs.First().GetNext().GetNext())
}

func Test_GetBefore(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	assert.Equal(t, 1, hs.Last().GetBefore().GetValue())
	assert.Nil(t, hs.Last().GetBefore().GetBefore())
}

func Test_AddFirst(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	hs.AddFirst(3)
	assert.Equal(t, 3, hs.First().GetValue())
	assert.Equal(t, 1, hs.First().GetNext().GetValue())
	assert.Equal(t, 3, hs.Size())
}

func Test_AddAfter_2(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	n := hs.AddLast(2)
	hs.AddLast(3)
	hs.AddAfter(4, n)
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 4, hs.First().GetNext().GetNext().GetValue())
	assert.Equal(t, 4, hs.Size())
}

func Test_Find(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	n := hs.AddLast(2)
	hs.AddLast(3)
	hs.AddAfter(4, n)

	n3 := hs.Find(func(a int) bool {
		return a == 3
	})

	assert.Equal(t, 3, n3.GetValue())
	assert.True(t, n3.GetBefore().GetValue() == 4)

	n_nil := hs.Find(func(a int) bool {
		return a == 100
	})

	assert.Nil(t, n_nil)
}

func Test_RemoveFirst(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	hs.AddLast(3)
	hs.AddLast(4)

	assert.Equal(t, 4, hs.Size())
	hs.RemoveFirst()

	assert.Equal(t, 3, hs.Size())
	assert.Equal(t, 2, hs.First().GetValue())

	hs.RemoveFirst()
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 3, hs.First().GetValue())

	hs.RemoveFirst()
	assert.Equal(t, 1, hs.Size())
	assert.Equal(t, 4, hs.First().GetValue())

	hs.RemoveFirst()
	assert.Equal(t, 0, hs.Size())
	assert.Nil(t, hs.First())

	hs.RemoveFirst()
}

func Test_RemoveLast(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.AddLast(1)
	hs.AddLast(2)
	hs.AddLast(3)
	hs.AddLast(4)

	assert.Equal(t, 4, hs.Size())
	hs.RemoveLast()

	assert.Equal(t, 3, hs.Size())
	assert.Equal(t, 3, hs.Last().GetValue())

	hs.RemoveLast()
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 2, hs.Last().GetValue())

	hs.RemoveLast()
	assert.Equal(t, 1, hs.Size())
	assert.Equal(t, 1, hs.Last().GetValue())

	hs.RemoveLast()
	assert.Equal(t, 0, hs.Size())
	assert.Nil(t, hs.First())

	hs.RemoveLast()
}
