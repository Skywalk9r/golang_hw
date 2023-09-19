package main

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{4, 3, 2, 1}
	slice3 := []int{1, 2, 3, 5}

	result1 := areSlicesEqual(slice1, slice2)
	result2 := areSlicesEqual(slice1, slice3)

	println(result1)
	println(result2)
}

func areSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	counts := make(map[int]int)

	for _, num := range slice1 {
		counts[num]++
	}

	for _, num := range slice2 {
		if count, ok := counts[num]; !ok || count == 0 {
			return false
		}
		counts[num]--
	}

	return true
}
