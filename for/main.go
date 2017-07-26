package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0

	for i := 0; i <= 5; i++ {
		sum += i
	}

	fmt.Println("the sum is:", sum)

	for j := 0; j < 10; j += 2 {
		fmt.Println("This should be an even number:", j)
	}

	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			fmt.Println("This should also be an even number:", k)
		}
	}

	// init and post statements of a for loop are optional
	belowOneHundred := 0
	// in this case, the `for` statement stands in for the traditional `while`
	// an example of an infinite while loop is -> `for {}`
	for belowOneHundred < 100 {
		fmt.Println("This should be below 100 ->", belowOneHundred)
		belowOneHundred += 10
	}

	multipleOfTen := 0
	for multipleOfTen < 100 {
		if multipleOfTen%10 == 0 {
			fmt.Println("Should be multiple of 10 ->", multipleOfTen)
		}
		multipleOfTen++
	}

	valA := -100.0
	sqrtA, msgA := sqrtCheck(valA)
	if sqrtA >= 0 {
		fmt.Println(msgA)
		fmt.Printf("sqrt of %.1f is %.3f\n", valA, sqrtA)
	} else {
		fmt.Println(msgA)
	}

	if aa := sumInt(5, 5); aa > 10 {
		fmt.Println(aa, "is greater than 10")
	} else {
		fmt.Println(aa, "is not greater than ten")
	}
}

func sqrtCheck(x float64) (float64, string) {
	if x < 0 {
		return -1, "Input should not be negative"
	}
	return math.Sqrt(x), "Valid input"
}

func sumInt(x, y int) int {
	return x + y
}
