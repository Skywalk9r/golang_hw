package main

import "fmt"

func main() {

	strs1 := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs1))

	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs2))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	commonPrefix := ""

	for i := 0; i < minLen; i++ {
		char := strs[0][i]
		for _, s := range strs {
			if s[i] != char {
				return commonPrefix
			}
		}
		commonPrefix += string(char)
	}

	return commonPrefix
}
