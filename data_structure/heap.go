package data_structure

import "sync"

type Heap interface {
	Add(value int)
	Delete(index int) int
	Len() int
	Peek(index int) int
}

type heap struct {
	arr []int
	sync.RWMutex
	compare func(x, y int) bool
}

func (h *heap) Peek(index int) int {
	h.RLock()
	defer h.RUnlock()
	return h.arr[index]
}

func (h *heap) Len() int {
	h.RLock()
	defer h.RUnlock()
	return len(h.arr)
}

func NewHeap(desc bool) Heap {
	h := &heap{}
	if desc {
		h.compare = func(x, y int) bool {
			return x > y
		}
	} else {
		h.compare = func(x, y int) bool {
			return x < y
		}
	}
	return h
}
func (h *heap) Add(value int) {
	h.Lock()
	defer h.Unlock()
	h.arr = append(h.arr, value)
	h.heapify()
}
func (h *heap) Delete(index int) int {
	h.Lock()
	defer h.Unlock()
	if index > len(h.arr)-1 {
		panic("out if heap")
	}
	val := h.arr[index]
	if index != len(h.arr)-1 {
		h.swap(index, len(h.arr)-1)
	}
	h.arr = h.arr[:len(h.arr)-1]
	h.heapify()
	return val
}
func (h *heap) heapify() {
	for cur := len(h.arr)/2 - 1; cur >= 0; cur-- {
		l := 2*cur + 1
		r := 2*cur + 2
		if l < len(h.arr) && h.compare(h.arr[l], h.arr[cur]) {
			h.swap(cur, l)
		}
		if r < len(h.arr) && h.compare(h.arr[r], h.arr[cur]) {
			h.swap(cur, r)
		}
	}
}
func (h *heap) swap(x, y int) {
	h.arr[x], h.arr[y] = h.arr[y], h.arr[x]
}
