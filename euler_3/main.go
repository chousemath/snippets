package main

import "fmt"

// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	target := 600851475143
	counter := 2
	for ; counter <= target; counter++ {
		if isPrime(target) {
			fmt.Println("Largest prime factor:", target)
			break
		}
		if target%counter == 0 {
			target = target / counter
			counter = 2
		}
	}
}
