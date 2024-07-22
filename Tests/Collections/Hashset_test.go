package Collections

import (
	"github.com/RENCI/GoUtils/Collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewHashset(t *testing.T) {
	hs := Collections.NewHashset[int]()
	assert.Equal(t, 0, hs.Size())
}

func Test_NewFromSlice(t *testing.T) {
	hs := Collections.NewHashsetFromSlice[int](&([]int{100, 20, 3}))
	assert.Equal(t, 3, hs.Size())
}

func Test_NewFromIterable(t *testing.T) {
	hs1 := Collections.NewHashsetFromSlice[int](&([]int{100, 20, 3}))
	hs2 := Collections.NewHashsetFromIterable[int](hs1)
	assert.Equal(t, 3, hs2.Size())
}

func Test_Add_Size_Contains(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(1)
	hs.Add(2)
	assert.Equal(t, 2, hs.Size())
	assert.True(t, hs.Contains(1))
	assert.True(t, hs.Contains(2))

}

func Test_Remove(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(1)
	hs.Add(2)
	hs.Remove(1)
	assert.Equal(t, 1, hs.Size())
	assert.True(t, hs.Contains(2))
}

func Test_Clear(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(1)
	hs.Add(2)

	hs.Clear()
	assert.Equal(t, 0, hs.Size())
}

func Test_Iterate(t *testing.T) {
	hs := Collections.NewHashset[int]()
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
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs.ExceptWith(&([]int{100, 3}))
	assert.Equal(t, 1, hs.Size())
	assert.True(t, hs.Contains(20))

}

func Test_ExceptWithHS(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs2 := Collections.NewHashset[int]()
	hs2.Add(100)
	hs2.Add(3)

	hs.ExceptWithHS(hs2)
	assert.Equal(t, 1, hs.Size())
	assert.True(t, hs.Contains(20))

}

func Test_IntersectWith(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs.IntersectWith(&([]int{100, 3}))
	assert.Equal(t, 2, hs.Size())
	assert.True(t, hs.Contains(100))
	assert.True(t, hs.Contains(3))

}

func Test_IntersectWithHS(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs2 := Collections.NewHashset[int]()
	hs2.Add(100)
	hs2.Add(3)

	hs.IntersectWithHS(hs2)
	assert.Equal(t, 2, hs.Size())
	assert.True(t, hs.Contains(100))
	assert.True(t, hs.Contains(3))

}

func Test_UnionWith(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs.UnionWith(&([]int{33, 44}))
	assert.Equal(t, 5, hs.Size())
	assert.True(t, hs.Contains(33))
	assert.True(t, hs.Contains(44))

}

func Test_UnionWithHS(t *testing.T) {
	hs := Collections.NewHashset[int]()
	hs.Add(100)
	hs.Add(20)
	hs.Add(3)

	hs2 := Collections.NewHashset[int]()
	hs2.Add(33)
	hs2.Add(44)

	hs.UnionWithHS(hs2)
	assert.Equal(t, 5, hs.Size())
	assert.True(t, hs.Contains(33))
	assert.True(t, hs.Contains(44))

}

func Test_ToList(t *testing.T) {
	hs1 := Collections.NewHashset[int]()
	hs1.Add(100)
	hs1.Add(20)
	hs1.Add(3)

	li := hs1.ToList()

	assert.Equal(t, 3, li.Size())
	assert.Equal(t, 100, li.Get(0))
	assert.Equal(t, 20, li.Get(1))
	assert.Equal(t, 3, li.Get(2))

}

func Test_Clone(t *testing.T) {
	hs1 := Collections.NewHashset[int]()
	hs1.Add(100)
	hs1.Add(20)
	hs1.Add(3)

	hs2 := hs1.Clone()
	hs2.Add(5)
	assert.Equal(t, 3, hs1.Size())
	assert.Equal(t, 4, hs2.Size())
	assert.False(t, hs1.Contains(5))
	assert.True(t, hs2.Contains(5))

}
