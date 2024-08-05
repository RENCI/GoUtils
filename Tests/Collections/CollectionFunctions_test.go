package Collections

import (
	"github.com/RENCI/GoUtils/Collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GroupBy(t *testing.T) {
	list := Collections.NewList[int]()
	list.AddRange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})

	grouped := Collections.GroupBy[int, bool, map[bool]Collections.List[int]](&list, func(item int) bool {
		return item%2 == 0
	})

	assert.Equal(t, 2, len(grouped))
	assert.Equal(t, 5, grouped[true].Size())
	assert.Equal(t, 5, grouped[false].Size())
}
