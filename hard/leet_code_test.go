package hard

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SolveSudoku(t *testing.T) {
	tests := []struct {
		name     string
		in       [][]byte
		expected [][]byte
	}{
		{
			name: "bord from example 1",
			in: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			expected: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.in)
			assert.Equal(t, tt.expected, tt.in)
		})
	}
}

func Test_Trap(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "foundation",
			in:       []int{4, 2, 3},
			expected: 1,
		}, {
			name:     "starts with 0",
			in:       []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		}, {
			name:     "bowl",
			in:       []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		}, {
			name:     "starts with 0 + foundation + bowl",
			in:       []int{0, 7, 1, 4, 6},
			expected: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, trap(tt.in))
		})
	}
}

func Test_FirstMissingPositive(t *testing.T) {
	tests := []struct {
		name     string
		in       []int
		expected int
	}{
		{
			name:     "from example 1",
			in:       []int{1, 2, 0},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, firstMissingPositive(tt.in))
		})
	}
}

func Test_IsMatch(t *testing.T) {
	tests := []struct {
		str      string
		pattern  string
		expected bool
	}{
		{
			str:      "",
			pattern:  "***?",
			expected: false,
		}, {
			str:      "",
			pattern:  "***",
			expected: true,
		}, {
			str:      "",
			pattern:  "*",
			expected: true,
		}, {
			str:      "",
			pattern:  "?",
			expected: false,
		}, {
			str:      "adda",
			pattern:  "ad?",
			expected: false,
		}, {
			str:      "adda",
			pattern:  "ad?ad",
			expected: false,
		}, {
			str:      "",
			pattern:  "",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "a",
			expected: false,
		}, {
			str:      "aa",
			pattern:  "*",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "?a",
			expected: true,
		}, {
			str:      "aa",
			pattern:  "?",
			expected: false,
		}, {
			str:      "cb",
			pattern:  "?a",
			expected: false,
		}, {
			str:      "caecfvacd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "caecfvacdcdcd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "caacd",
			pattern:  "?a*?cd",
			expected: true,
		}, {
			str:      "bbbbbbbabaab",
			pattern:  "b*b*ab**ba",
			expected: false,
		}, {
			str:      "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab",
			pattern:  "b*b*ab**ba*b**b***bba",
			expected: false,
		}, {
			str:      "bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabbba",
			pattern:  "b*b*ab**ba*b**b***bba",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%s)/(%s)(%v)", tt.str, tt.pattern, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, isMatch(tt.str, tt.pattern))
		})
	}
}

func Test_SolveNQueens(t *testing.T) {
	tests := []struct {
		in       int
		expected [][]string
	}{
		{
			in: 4,
			expected: [][]string{{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			}, {
				"..Q.",
				"Q...",
				"...Q",
				".Q.."}},
		},
		{
			in:       1,
			expected: [][]string{{"Q"}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N count: %d", tt.in), func(t *testing.T) {
			assert.Equal(t, tt.expected, solveNQueens(tt.in))
		})
	}
}

func Test_TotalNQueens(t *testing.T) {
	tests := []struct {
		in       int
		expected int
	}{
		{
			in:       5,
			expected: 10,
		}, {
			in:       4,
			expected: 2,
		}, {
			in:       1,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N count: %d", tt.in), func(t *testing.T) {
			assert.Equal(t, tt.expected, totalNQueens(tt.in))
		})
	}
}

func Test_GetPermutation(t *testing.T) {
	tests := []struct {
		inN      int
		inK      int
		expected string
	}{
		{
			inN:      3,
			inK:      2,
			expected: "132",
		}, {
			inN:      3,
			inK:      3,
			expected: "213",
		}, {
			inN:      2,
			inK:      2,
			expected: "21",
		}, {
			inN:      8,
			inK:      6593,
			expected: "24186735",
		}, {
			inN:      4,
			inK:      9,
			expected: "2314",
		}, {
			inN:      3,
			inK:      1,
			expected: "123",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("N:%d,K:%d,(%s)", tt.inN, tt.inK, tt.expected), func(t *testing.T) {
			assert.Equal(t, tt.expected, getPermutation(tt.inN, tt.inK))
		})
	}
}

func Test_IsNumber(t *testing.T) {
	valid := []string{"46.e3", "0", "2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	for _, tt := range valid {
		t.Run(tt, func(t *testing.T) {
			assert.True(t, isNumber(tt))
		})
	}
	notValid := []string{".e1", "+.", "4e+", "abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", "e", "."}
	for _, tt := range notValid {
		t.Run(tt, func(t *testing.T) {
			assert.False(t, isNumber(tt))
		})
	}
}