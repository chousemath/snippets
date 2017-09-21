// main_test.go

package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(1001) {
		t.Errorf("1001 should be a palindrome...")
	}

	if !isPalindrome(100020001) {
		t.Errorf("100020001 should be a palindrome...")
	}

	if isPalindrome(123419238479213847) {
		t.Errorf("123419238479213847 should NOT be a palindrome...")
	}

	if isPalindrome(100000034000001) {
		t.Errorf("100000034000001 should NOT be a palindrome...")
	}

	if findPalindrome() != 906609 {
		t.Errorf("findPalindrome is not working...")
	}
}
