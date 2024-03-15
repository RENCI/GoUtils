package Collections

import (
	"github.com/RENCI/GoUtils/Collections/Queue"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Queue_New(t *testing.T) {
	list := Queue.New[int]()
	assert.Equal(t, 0, list.Size())
}

func Test_Queue_NewFromSlice(t *testing.T) {
	hs := Queue.NewFromSlice[int](&([]int{1, 2, 3, 4}))
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, *hs.Peek())
}

func Test_Queue_Enqueue(t *testing.T) {
	hs := Queue.NewFromSlice[int](&([]int{1, 2, 3, 4}))
	hs.Enqueue(5)
	assert.Equal(t, 5, hs.Size())
	assert.Equal(t, 1, *hs.Peek())
}

func Test_Queue_Dequeue(t *testing.T) {
	hs := Queue.NewFromSlice[int](&([]int{1, 2, 3, 4}))
	hs.Enqueue(5)
	v := hs.Dequeue()
	assert.Equal(t, 4, hs.Size())
	assert.Equal(t, 1, *v)
	assert.Equal(t, 2, *hs.Peek())
}

func Test_Queue_Dequeue_2(t *testing.T) {
	hs := Queue.NewFromSlice[int](&([]int{1, 2, 3, 4}))
	v1 := hs.Dequeue()
	v2 := hs.Dequeue()
	v3 := hs.Dequeue()
	v4 := hs.Dequeue()
	v5 := hs.Dequeue()
	assert.Equal(t, 0, hs.Size())
	assert.Equal(t, 1, *v1)
	assert.Equal(t, 2, *v2)
	assert.Equal(t, 3, *v3)
	assert.Equal(t, 4, *v4)
	assert.Nil(t, v5)
}
