package main

import "fmt"

func main() {
	multiplyA := 5
	multiplyB := 6
	divA := 100
	divB := 25
	fmt.Println("5 - 3 =", subtractInt(5, 3))
	fmt.Printf("%d * %d = %d\n", multiplyA, multiplyB, multiplyInt(multiplyA, multiplyB))
	fmt.Printf("%d / %d = %d\n", divA, divB, divideInt(divA, divB))
}

func subtractInt(a int, b int) int {
	return a - b
}

func multiplyInt(a int, b int) int {
	return a * b
}

func divideInt(a int, b int) int {
	return a / b
}
