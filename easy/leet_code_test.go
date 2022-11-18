package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsOrderPrinter(t *testing.T) {
	nums := []int{0, 2}
	res := removeDuplicates(nums)

	assert.Equal(t, 2, res)
}

func Test_MakeGood(t *testing.T) {
	tests := []struct {
		str      string
		expected string
	}{
		{
			str:      "leEeetcode",
			expected: "leetcode",
		}, {
			str:      "leEeetTcode",
			expected: "leecode",
		}, {
			str:      "abBAcC",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			assert.Equal(t, tt.expected, makeGood(tt.str))
		})
	}
}
