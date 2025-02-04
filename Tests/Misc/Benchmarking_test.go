package Misc

import (
	"github.com/RENCI/GoUtils/Misc"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func Test_Benchmark(t *testing.T) {

	res := Misc.Benchmark("random sleep", 10, func() {
		r := rand.Intn(400)
		time.Sleep(time.Duration(r) * time.Millisecond)
	})
	res.Calculate()
	assert.NotNil(t, res)
}
