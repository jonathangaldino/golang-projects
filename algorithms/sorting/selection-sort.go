package Sorting

import "fmt"

func getLower(items []int) int {
	lowerIndex := 0
	lower := items[lowerIndex]

	for i := 1; i < len(items); i++ {
		if items[i] < lower {
			lower = items[i]
			lowerIndex = i
		}

		// fmt.Println("Index:", i, "Value:", items[i])
	}

	return lowerIndex
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func SelectionSort(items []int) []int {
	// Copy the original array
	copyArr := items[:]

	// Save the size of the original array
	size := len(items)

	// Create the new supposed ordered array with the size of original, of course
	orderedArr := make([]int, 0, size)

	for i := 0; i < len(items); i++ {
		// Get the lowest number
		lowerIndex := getLower(copyArr)

		// Add it to the original array
		// Append at to the first empty position
		orderedArr = append(orderedArr, items[lowerIndex])

		// Remove the used number from the copied array
		copyArr = removeIndex(copyArr, lowerIndex)
	}

	return orderedArr
}

func BasicExample() {
	items := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	orderedItems := SelectionSort(items)
	fmt.Println(orderedItems)
}
