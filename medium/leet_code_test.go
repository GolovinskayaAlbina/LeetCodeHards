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
