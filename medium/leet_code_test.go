package medium

import (
	"fmt"
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

func Test_StockSpanner(t *testing.T) {
	obj := Constructor()
	assert.Equal(t, 1, obj.Next(100))
	assert.Equal(t, 1, obj.Next(80))
	assert.Equal(t, 1, obj.Next(60))
	assert.Equal(t, 2, obj.Next(70))
	assert.Equal(t, 1, obj.Next(60))
	assert.Equal(t, 4, obj.Next(75))
	assert.Equal(t, 6, obj.Next(85))
}

func Test_ReverseWords(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{
			str:      "the sky is blue",
			expected: "blue is sky the",
		}, {
			str:      "  hello world  ",
			expected: "world hello",
		}, {
			str:      "a good   example",
			expected: "example good a",
		}, {
			str:      "a",
			expected: "a",
		}, {
			str:      " a",
			expected: "a",
		}, {
			str:      " a ",
			expected: "a",
		}, {
			str:      " ",
			expected: "",
		}, {
			str:      "           item          some",
			expected: "some item",
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			assert.Equal(t, tt.expected, reverseWords(tt.str))
		})
	}
}

func Test_RemoveStones(t *testing.T) {
	tests := []struct {
		stones   [][]int
		expected int
	}{
		{
			stones:   [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}},
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, removeStones(tt.stones))
		})
	}
}

func Test_CountNodes(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}},
			expected: 6,
		}, {
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}},
			expected: 5,
		}, {
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}},
			expected: 3,
		}, {
			root:     &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			expected: 2,
		}, {
			root:     &TreeNode{Val: 1},
			expected: 1,
		}, {
			root:     nil,
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, countNodes(tt.root))
		})
	}
}

func Test_ComputeArea(t *testing.T) {
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := -3, 0, 3, 4, 0, -1, 9, 2
	a := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	assert.Equal(t, 45, a)

	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 = -2, -2, 2, 2, -2, 2, 2, 4
	a = computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	assert.Equal(t, 24, a)

	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 = -2, -2, 2, 2, -3, 1, -1, 3
	a = computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	assert.Equal(t, 19, a)

	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 = -2, -2, 2, 2, -2, -2, 2, 2
	a = computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	assert.Equal(t, 16, a)
}
