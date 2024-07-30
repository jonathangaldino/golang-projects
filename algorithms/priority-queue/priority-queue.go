package PriorityQueue

import (
	"container/heap"
)

// Item represents a node in the Heap with it's distance
type Item struct {
	Node     int
	Priority int
	Index    int
}

// PriorityQueue implements the heap.Interface.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old) // get the size of the pq

	item := old[n-1] // grab the last item

	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // safety

	*pq = old[0 : n-1] // pq now is all the items but the last one
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.Priority = priority // update the item with the new priority

	// Fix re-establishes the heap ordering after the element at index i has changed its value.
	// Changing the value of the element at index i and then calling Fix is equivalent to,
	// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
	// The complexity is O(log n) where n = h.Len().
	heap.Fix(pq, item.Index)
}
