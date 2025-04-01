package main

import "fmt"

type Heap []int

var size = 10
var capacity = 10

func (h Heap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h Heap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h Heap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h Heap) hasLeftChild(index int) bool {
	return h.getLeftChildIndex(index) < size
}
func (h Heap) hasRightChild(index int) bool {
	return h.getRightChildIndex(index) < size
}
func (h Heap) hasParent(index int) bool {
	return h.getParentIndex(index) >= 0
}

func (h Heap) leftChild(index int) int {
	return h[h.getLeftChildIndex(index)]
}
func (h Heap) rightChild(index int) int {
	return h[h.getRightChildIndex(index)]
}

func (h Heap) parent(index int) int {
	return h[h.getParentIndex(index)]
}

func (h Heap) swap(indexOne, indexTwo int) {
	temp := h[indexOne]

	h[indexOne] = h[indexTwo]
	h[indexTwo] = temp
}

func (h Heap) ensureExtraCapacity() {
	if size == capacity {
		size = size * 2
		capacity = capacity * 2

		newHeap := make([]int, size, capacity)
		h = append(newHeap, h...)
	}
}

func main() {
	heap := Heap{2, 4, 8}
	heap.ensureExtraCapacity()
	fmt.Println(cap(heap))
}
