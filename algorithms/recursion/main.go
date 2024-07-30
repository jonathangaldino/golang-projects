package Recursion

import "fmt"

func fat(x int) int {
	if x == 1 {
		return 1
	}

	return x * fat(x-1)
}

func SlowSum(items []int) int {
	total := 0

	for i := 0; i < len(items); i++ {
		total += items[i]
	}

	return total
}

func Count(items []int) int {
	if len(items) == 0 {
		return 0
	}

	return 1 + Count(items[1:])
}

func Sum(items []int) int {
	if len(items) == 0 {
		return 0
	}

	return items[0] + Sum(items[1:])
}

func Max(items []int, highest int) int {
	if len(items) == 0 {
		return highest
	}

	if items[0] > highest {
		return Max(items[1:], items[0])
	}

	return Max(items[1:], highest)
}

func Run() {
	arr := []int{1, 4, 39, 2, 10}

	// fmt.Println(fat(3))
	fmt.Printf("The Slow sum is %d\n", SlowSum(arr))
	fmt.Printf("The Recursion sum is %d\n", Sum(arr))
	fmt.Printf("There is %d items on the list.\n", Count(arr))
	fmt.Printf("This is the highest number on the list: %d\n", Max(arr, 0))
}
