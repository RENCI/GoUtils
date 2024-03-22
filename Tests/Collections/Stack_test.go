package Collections

import (
	"github.com/RENCI/GoUtils/Collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Stack_New(t *testing.T) {
	list := Collections.NewStack[int]()
	assert.Equal(t, 0, list.Size())
}

func Test_Stack_NewFromSlice(t *testing.T) {
	hs := Collections.NewStackFromSlice[int](&([]int{1, 2, 3, 4}))
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 4, *hs.Peek())
}

func Test_Stack_Push(t *testing.T) {
	hs := Collections.NewStackFromSlice[int](&([]int{1, 2, 3, 4}))
	hs.Push(5)
	assert.Equal(t, 5, hs.Size())
	assert.Equal(t, 5, *hs.Peek())
}

func Test_Stack_Pop(t *testing.T) {
	hs := Collections.NewStackFromSlice[int](&([]int{1, 2, 3, 4}))
	hs.Push(5)
	v := hs.Pop()
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 5, *v)
	assert.Equal(t, 4, *hs.Peek())
}

func Test_Stack_Pop_2(t *testing.T) {
	hs := Collections.NewStackFromSlice[int](&([]int{1, 2, 3, 4}))
	v1 := hs.Pop()
	v2 := hs.Pop()
	v3 := hs.Pop()
	v4 := hs.Pop()
	v5 := hs.Pop()
	assert.Equal(t, 0, hs.Size())
	assert.Equal(t, 4, *v1)
	assert.Equal(t, 3, *v2)
	assert.Equal(t, 2, *v3)
	assert.Equal(t, 1, *v4)
	assert.Nil(t, v5)

}
