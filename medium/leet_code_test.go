package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindBall(t *testing.T) {
	tests := []struct {
		name     string
		in       [][]int
		expected []int
	}{
		{
			name:     "bord from example 1",
			in:       [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}},
			expected: []int{1, -1, -1, -1, -1},
		}, {
			name:     "bord from example 2",
			in:       [][]int{{-1}},
			expected: []int{-1},
		}, {
			name:     "bord from example 3",
			in:       [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}},
			expected: []int{0, 1, 2, 3, 4, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, findBall(tt.in))
		})
	}
}

func Test_MinMutation(t *testing.T) {
	tests := []struct {
		start    string
		end      string
		bank     []string
		expected int
	}{
		{
			start:    "AACCGGTT",
			end:      "AACCGGTA",
			bank:     []string{"AACCGGTA"},
			expected: 1,
		}, {
			start:    "AACCGGTT",
			end:      "AAACGGTA",
			bank:     []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			expected: 2,
		}, {
			start:    "AAAAACCC",
			end:      "AACCCCCC",
			bank:     []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			expected: 3,
		}, {
			start:    "AACCGGTT",
			end:      "AAACGGTA",
			bank:     []string{"AACCGATT", "AACCGATA", "AAACGATA", "AAACGGTA"},
			expected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.start, func(t *testing.T) {
			assert.Equal(t, tt.expected, minMutation(tt.start, tt.end, tt.bank))
		})
	}
}

func Test_LongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected int
	}{
		{
			name:     "Example 1",
			words:    []string{"lc", "cl", "gg"},
			expected: 6,
		}, {
			name:     "Example 2",
			words:    []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			expected: 8,
		}, {
			name:     "Example 3",
			words:    []string{"cc", "ll", "xx"},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, longestPalindrome(tt.words))
		})
	}
}
