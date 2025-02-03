package Misc

import (
	"github.com/RENCI/GoUtils/Collections"
	"time"
)

type BenchmarkResult struct {
	Name       string
	Iterations Collections.List[time.Duration]
	Average    time.Duration
	Median     time.Duration
}

func Benchmark(name string, iterations int, f func()) *BenchmarkResult {
	var br = &BenchmarkResult{
		Name:       name,
		Iterations: Collections.NewList[time.Duration](),
	}

	for i := 0; i < iterations; i++ {
		start := time.Now()
		f()
		br.Iterations.Add(time.Since(start))
	}

	return br
}

func (this *BenchmarkResult) Calculate() {
	if this.Iterations.Size() == 0 {
		return
	}
	var sum time.Duration

	this.Iterations.ForEach(func(v time.Duration) {
		sum += v
	})

	this.Average = sum / time.Duration(this.Iterations.Size())
	this.Iterations.Sort(func(item1 time.Duration, item2 time.Duration) int {
		if item1 > item2 {
			return 1
		} else if item1 == item2 {
			return 0
		} else {
			return -1
		}
	})
	this.Median = this.Iterations.Get(this.Iterations.Size() / 2)
}
