package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	//assert.Equal(t, int64(2), countTriplets([]int64{1, 2, 2, 4}, 2))
	//assert.Equal(t, int64(3), countTriplets([]int64{1, 2, 2, 2, 4}, 2))
	//assert.Equal(t, int64(4), countTriplets([]int64{1, 2, 2, 4, 4}, 2))
	//assert.Equal(t, int64(6), countTriplets([]int64{1, 3, 9, 9, 27, 81}, 3))
	//assert.Equal(t, int64(4), countTriplets([]int64{1, 5, 5, 25, 125}, 5))
	assert.Equal(t, int64(1), countTriplets([]int64{1, 1, 1}, 1))
	//assert.Equal(t, int64(161700), countTriplets([]int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 1))
}

func BenchmarkTest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		countTriplets([]int64{1, 3, 9, 9, 27, 81}, 3)
	}
}

func countTriplets(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return int64(0)
	}
	am := make(map[int64]int)
	am[arr[0]] = 1

	for ai := 1; ai < len(arr); ai++ {
		am[arr[ai]]++
	}
	result := int64(1)
	for b := arr[0]; ; b *= r {
		if _, ok := am[b]; !ok {
			break
		}
		result *= int64(am[b])
	}
	return result
}

func countTripletsBad(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return int64(0)
	}
	am := make(map[int64]int)
	am[arr[0]] = 1
	base := int64(1)

	for ai := 1; ai < len(arr); ai++ {
		am[arr[ai]]++
		if arr[ai-1] != arr[ai] {
			base *= int64(am[arr[ai-1]])
		}
	}
	return int64(base * int64(len(am)-2))
}
