package Collections

import (
	"github.com/RENCI/GoUtils/Collections/LinkedList"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_New(t *testing.T) {
	list := LinkedList.New[int]()
	assert.Equal(t, 0, list.Size())
}

func Test_Add(t *testing.T) {
	list := LinkedList.New[int]()
	list.Add(1)
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 1, list.First().GetValue())
	assert.Equal(t, 1, list.Last().GetValue())
}

func Test_Add_2(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.Add(1)
	hs.Add(2)
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.Last().GetValue())
}

func Test_Add_3(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	assert.Equal(t, 3, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 2, hs.Last().GetBefore().GetValue())
	assert.Equal(t, 3, hs.Last().GetValue())
}

func Test_Remove_1(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Add(3)
	hs.Remove(func(a int, b int) bool {
		return a == b
	}, 3)
	assert.Equal(t, 2, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 2, hs.Last().GetValue())
}

func Test_InsertBefore(t *testing.T) {
	hs := LinkedList.New[int]()
	hs.Add(1)
	n := hs.Add(3)
	hs.Add(4)

	hs.InsertBefore(2, n)
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 3, hs.First().GetNext().GetNext().GetValue())
	assert.Equal(t, 4, hs.Last().GetValue())
}

func Test_InsertAfter(t *testing.T) {
	hs := LinkedList.New[int]()
	n := hs.Add(1)
	hs.Add(3)
	hs.Add(4)

	hs.InsertAfter(2, n)
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, hs.First().GetValue())
	assert.Equal(t, 2, hs.First().GetNext().GetValue())
	assert.Equal(t, 3, hs.First().GetNext().GetNext().GetValue())
	assert.Equal(t, 4, hs.Last().GetValue())
}
