package main

import (
	"fmt"
	"strconv"
)

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			result += symbol[i]
		}
	}

	return result
}

func main() {
	var inputStr string
	fmt.Print("Enter a number: ")
	fmt.Scanln(&inputStr)

	num, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return
	}

	fmt.Printf("Input: num = %d\n", num)
	fmt.Printf("Output: %s\n", intToRoman(num))
}
