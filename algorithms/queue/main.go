package Queue

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	items []T
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: []T{}}
}

func (queue *Queue[T]) Enqueue(item T) {
	queue.items = append(queue.items, item)
	fmt.Printf("After enqueing: %v\n", queue.items)
}

func (queue *Queue[T]) Dequeue() (item T, err error) {
	if len(queue.items) < 1 {
		// can't return nil as T
		var zeroValue T

		return zeroValue, errors.New("Queue is empty.")
	}

	i := queue.items[0]

	queue.items = queue.items[1:]

	return i, nil
}
