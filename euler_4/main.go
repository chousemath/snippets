package main

import "fmt"

func isPalindrome(target int) bool {
	// in Go, a string is a read-only slice of bytes
	strArr := []rune(fmt.Sprintf("%v", target))
	strLength := len(strArr)
	for k := 0; k <= strLength/2-1; k++ {
		if strArr[k] != strArr[strLength-k-1] {
			return false
		}
	}
	return true
}

func findPalindrome() int {
	largest := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			if isPalindrome(i * j) {
				if (i * j) > largest {
					largest = i * j
				}
			}
		}
	}
	return largest
}

func main() {
	fmt.Println(findPalindrome())
}
