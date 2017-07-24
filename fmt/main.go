package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	a := 123.321
	randA := rand.Intn(10)
	randB := rand.Intn(10)
	originA := 81.0
	sqrtA := math.Sqrt(originA)

	fmt.Println("Hello world")
	// width sets the minimum width of the field
	// precision sets the number of places after the decimal place if appropriate
	fmt.Printf("value of `a` with width=def, precision=def is %f\n", a)
	fmt.Printf("value of `a` with width=def, precision=3 is %.3f\n", a)
	fmt.Printf("value of `a` with width=def, precision=2 is %.2f\n", a)
	fmt.Printf("value of `a` with width=def, precision=1 is %.1f\n", a)
	fmt.Printf("type of `a` is %T\n", a)

	fmt.Println("Here is a random number:", randA)
	fmt.Printf("Same number, but with interpolation: %d\n", randA)
	fmt.Println("Here is another random number:", randB)
	fmt.Printf("Again, same number but with interpolation: %d\n", randB)

	fmt.Println("The square root of", originA, "is:", sqrtA)
	fmt.Printf("The square root of %.2f is %.2f\n", originA, sqrtA)
}
