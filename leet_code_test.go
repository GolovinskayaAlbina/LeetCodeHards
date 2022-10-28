package main

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
