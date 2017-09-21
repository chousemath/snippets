// main.go
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of
// two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func isPalindrome(target int) bool {
	// in Go, a string is a read-only slice of bytes
	// Runes literals in Go are just integer values
	// Runes are mapped to their unicode codepoint
	// This is why you can perform integer addition on runes to get other runes
	strArr := []rune(fmt.Sprintf("%v", target))
	// The built-in `len` functions returns the length of its target
	// Array: the number of elements in v.
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
