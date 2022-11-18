package hard

import (
	"leet-code/data_structure"
)

type MedianFinder struct {
	maxHeap data_structure.Heap
	minHeap data_structure.Heap
}

func Constructor() MedianFinder {
	minHeap := data_structure.NewHeap(false)
	maxHeap := data_structure.NewHeap(true)
	return MedianFinder{maxHeap, minHeap}
}

func (this *MedianFinder) AddNum(num int) {
	this.maxHeap.Add(num)
	this.minHeap.Add(this.maxHeap.Delete(0))
	for this.maxHeap.Len() < this.minHeap.Len() {
		this.maxHeap.Add(this.minHeap.Delete(0))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxHeap.Len() > this.minHeap.Len() {
		return float64(this.maxHeap.Peek(0))
	}
	return float64(this.maxHeap.Peek(0)+this.minHeap.Peek(0)) / 2
}
