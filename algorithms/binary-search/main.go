package BinarySearch

import "fmt"

func Search(items []int, item int) (position int, e error) {
	var low = 0
	var high = len(items) - 1

	for low <= high {
		var mid = (low + high) / 2
		var kick = items[mid]

		if kick == item {
			return mid, nil
		}

		if kick > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, fmt.Errorf("Item not present in the array of items.")
}

func ExampleOne() {
	var items = []int{1, 2, 3, 4, 5, 6, 7}

	var pos, err = Search(items, 2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Item found at position %d\n", pos)
	}
}
