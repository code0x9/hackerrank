package hackerrank

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, int32(0), sherlockAndAnagrams("abcd"))
	assert.Equal(t, int32(4), sherlockAndAnagrams("abba"))
	assert.Equal(t, int32(3), sherlockAndAnagrams("ifailuhkqq"))
	assert.Equal(t, int32(10), sherlockAndAnagrams("kkkk"))
}

func sherlockAndAnagrams(s string) int32 {
	toMap := func(s string) map[rune]int {
		m := map[rune]int{}
		for _, r := range s {
			m[r]++
		}
		return m
	}

	isEqual := func(m1, m2 map[rune]int) bool {
		if len(m1) != len(m2) {
			return false
		}
		for k, v := range m1 {
			if m2[k] != v {
				return false
			}
		}
		return true
	}

	var c int32
	// i is part length
	for i := 1; i < len(s)+1; i++ {
		// j is part 1 idx
		for j := 0; j < len(s)-i; j++ {
			p1 := toMap(s[j : j+i])
			// k is part 2 idx
			for k := j + 1; k < len(s)-i+1; k++ {
				p2 := toMap(s[k : k+i])
				if isEqual(p1, p2) {
					c++
				}
			}
		}
	}
	return c
}
