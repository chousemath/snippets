package main

import "fmt"

func main() {
	multiplyA := 5
	multiplyB := 6
	divA := 100
	divB := 25
	modA := 11
	modB := 10
	a := 123.333
	b := 111.111
	aa := 13
	bb := 10
	fmt.Println("5 - 3 =", subtractInt(5, 3))
	fmt.Printf("%d * %d = %d\n", multiplyA, multiplyB, multiplyInt(multiplyA, multiplyB))
	fmt.Printf("%d / %d = %d\n", divA, divB, divideInt(divA, divB))
	fmt.Printf("%d `mod` %d is %d\n", modA, modB, moduloInt(modA, modB))
	fmt.Println("1 + 2 + 3 =", tripleSum(1, 2, 3))
	sumAb, subAb := sumAndSubtractFloat(a, b)
	fmt.Printf(
		"sum of %g and %g is %.3f, sub of %g and %g is %.3f\n",
		a, b, sumAb, a, b, subAb,
	)
	multAb, divAb, modAb := multAndDivideAndModInt(aa, bb)
	fmt.Printf(
		"%d * %d = %d, %d / %d = %d, %d `mod` %d = %d\n",
		aa, bb, multAb, aa, bb, divAb, aa, bb, modAb,
	)
}

func subtractInt(a, b int) int {
	return a - b
}

func multiplyInt(a, b int) int {
	return a * b
}

func divideInt(a, b int) int {
	return a / b
}

func moduloInt(a, b int) int {
	return a % b
}

func tripleSum(a, b, c int) int {
	return a + b + c
}

func sumAndSubtractFloat(a, b float64) (float64, float64) {
	return a + b, a - b
}

func multAndDivideAndModInt(a, b int) (int, int, int) {
	return a * b, a / b, a % b
}
