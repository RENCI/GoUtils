package Collections

import (
	"github.com/RENCI/GoUtils/Collections/Hashset"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewHashset(t *testing.T) {
	hs := Hashset.New[int]()
	assert.Equal(t, 0, hs.Size())
}

func Test_NewFromSlice(t *testing.T) {
	hs := Hashset.NewFromSlice[int](&([]int{100, 20, 3}))
	assert.Equal(t, 3, hs.Size())
}

func Test_NewFromIterable(t *testing.T) {
	hs1 := Hashset.NewFromSlice[int](&([]int{100, 20, 3}))
	hs2 := Hashset.NewFromIterable[int](hs1)
	assert.Equal(t, 3, hs2.Size())
}

func Test_Add_Size_Contains(t *testing.T) {
	hs := Hashset.New[int]()
	hs.Add(1)
	hs.Add(2)
	assert.Equal(t, 2, hs.Size())
	assert.True(t, hs.Contains(1))
	assert.True(t, hs.Contains(2))

}

func Test_Remove(t *testing.T) {
	hs := Hashset.New[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Remove(1)
	assert.Equal(t, 1, hs.Size())
	assert.True(t, hs.Contains(2))
}

func Test_Clear(t *testing.T) {
	hs := Hashset.New[int]()
	hs.Add(1)
	hs.Add(2)

	hs.Clear()
	assert.Equal(t, 0, hs.Size())
}

func Test_Iterate(t *testing.T) {
	hs := Hashset.New[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)
	s := 0
	hs.Iterate(func(item *int) bool {
		s += *item
		return true
	})
	assert.Equal(t, 123, s)
}

func Test_ExceptWith(t *testing.T) {
	hs := Hashset.New[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs.ExceptWith(&([]int{100, 3}))
	assert.Equal(t, 1, hs.Size())
	assert.True(t, hs.Contains(20))

}
