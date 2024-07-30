package Sorting

import (
	"slices"
)

func filter(n []int, pivot int, test func(int, int) bool) (arr []int) {
	for _, number := range n {
		if test(number, pivot) {
			arr = append(arr, number)
		}
	}

	return
}

func lowerEqualThan(n int, pivot int) bool {
	return n <= pivot
}

func greaterThan(n int, pivot int) bool {
	return n > pivot
}

func Quicksort(items []int) []int {
	// Stop when array is empty or has just one item
	if len(items) < 2 {
		return items
	}

	// Choose the pivot, which is the first item.
	pivot := items[0]

	// Grab all numbers that are lower or equal than the pivot and put them in the left.
	lowers := filter(items[1:], pivot, lowerEqualThan)

	// Grab all numbers that are greater than the pivot and put them in the right.
	greaters := filter(items[1:], pivot, greaterThan)

	// Run again:                for the left side           and right side -|
	return slices.Concat(Quicksort(lowers), []int{pivot}, Quicksort(greaters))
}

// Starting slice: [4 3 8 2 7 0 6 1 5 9]
// Pivot: 4
// Left side: [3 2 0 1]
// Right side: [8 7 6 5 9]
//
// Run QuickSort on [3 2 0 1]
//
// Partition: [3 2 0 1]
// Pivot: 3
// Left: [2 0 1]
// Right: []
//
// Run Quicksort on [2 0 1]
// Pivot: 2
// Left: [0 1]
// Right: [] <--- done
//
// // Run Quicksort on [0 1]
//
// Pivot: 0
// Left: [] <--- done
// Right: [1] <--- done
//
//
// We have: [0, 1, 2, 3] + [4] (first pivot) + Quicksort(right) <--- we still need to run this one
//
//
// Partition: [8 7 6 5 9]
// Pivot: 8
// Left: [7 6 5]
// Right: [9] <-- done
//
// Partition: [7 6 5]
// Pivot: 7
// Left: [6 5]
// Right: [] <-- done
//
// Partition: [6 5]
// Pivot: 6
// Left: [5] <--- done
// Right: [0] <--- done
//
// End ------\
//
// Result: [0, 1, 2, 3] + [4] + [5, 6, 7, 8, 9]
// Result: [0 1 2 3 4 5 6 7 8 9]
