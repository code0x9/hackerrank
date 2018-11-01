package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, int64(200),
		arrayManipulation(5, [][]int32{
			{1, 2, 100},
			{2, 5, 100},
			{3, 4, 100},
		}))
	assert.Equal(t, int64(10),
		arrayManipulation(10, [][]int32{
			{1, 5, 3},
			{4, 8, 7},
			{6, 9, 1},
		}))
}

// use "stateful" x value
func arrayManipulation(n int32, queries [][]int32) int64 {
	r := make([]int, n)
	for _, query := range queries {
		a, b, k := int(query[0])-1, int(query[1])-1, int(query[2])
		r[a] += k
		if b+1 < int(n) {
			r[b+1] -= k
		}
	}
	max, x := -1, 0
	for _, v := range r {
		x += v
		if max < x {
			max = x
		}
	}
	return int64(max)
}
