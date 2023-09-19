package main

import "fmt"

func main() {
	intSlice := []int{7, 3, 6, 5, 9, 2, 4, 1, 0, 8}

	fmt.Println("Before sorting:", intSlice)
	BubbleSort(intSlice)
	fmt.Println("After sorting:", intSlice)
}

func BubbleSort(arr []int) {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
	}
}
