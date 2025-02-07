package Misc

import (
	"github.com/RENCI/GoUtils/Misc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Singleton(t *testing.T) {

	r := 1
	s := Misc.NewSingleton(func() Misc.Result[int] {
		r += 1
		return Misc.NewResult(r, nil)
	})
	res := s.Get()
	res = s.Get()
	assert.Equal(t, 2, res.Value)
}
