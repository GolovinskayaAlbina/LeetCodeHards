package data_structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxHeap(t *testing.T) {
	maxHeap := NewHeap(true)
	maxHeap.Add(3)
	maxHeap.Add(9)
	maxHeap.Add(2)
	maxHeap.Add(1)
	maxHeap.Add(4)
	maxHeap.Add(5)
	maxHeap.Add(7)
	assert.Equal(t, []int{9, 4, 7, 1, 3, 2, 5}, maxHeap.(*heap).arr)
	maxHeap.Delete(0)
	assert.Equal(t, []int{7, 4, 5, 1, 3, 2}, maxHeap.(*heap).arr)
}

func Test_MinHeap(t *testing.T) {
	maxHeap := NewHeap(false)
	maxHeap.Add(3)
	maxHeap.Add(9)
	maxHeap.Add(2)
	maxHeap.Add(1)
	maxHeap.Add(4)
	maxHeap.Add(5)
	maxHeap.Add(7)
	assert.Equal(t, []int{1, 2, 3, 9, 4, 5, 7}, maxHeap.(*heap).arr)
}
