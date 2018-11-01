package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, int64(200),
		arrayManipulation(5, [][]int32{
			[]int32{1, 2, 100},
			[]int32{2, 5, 100},
			[]int32{3, 4, 100},
		}))
	assert.Equal(t, int64(10),
		arrayManipulation(10, [][]int32{
			[]int32{1, 5, 3},
			[]int32{4, 8, 7},
			[]int32{6, 9, 1},
		}))
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	r := make([]int, n)
	max := -1
	for i := 0; i < int(n); i++ {
		var sum int
		for _, query := range queries {
			a, b, k := int(query[0])-1, int(query[1])-1, query[2]
			if i >= a && i <= b {
				sum += int(k)
			}
		}
		r[i] = sum
		if max < sum {
			max = sum
		}
	}
	return int64(max)
}

func arrayManipulationSimple(n int32, queries [][]int32) int64 {
	r := make([]int, n)
	max := int64(-1)
	for _, query := range queries {
		a, b, k := query[0], query[1], query[2]
		for i := a - 1; i < b; i++ {
			r[i] += int(k)
			if max < int64(r[i]) {
				max = int64(r[i])
			}
		}
	}
	return int64(max)
}
