package main

import "fmt"

func main() {
	a := 123.321

	fmt.Println("Hello world")
	// width sets the minimum width of the field
	// precision sets the number of places after the decimal place if appropriate
	fmt.Printf("value of `a` with width=def, precision=def is %f\n", a)
	fmt.Printf("value of `a` with width=def, precision=3 is %.3f\n", a)
	fmt.Printf("value of `a` with width=def, precision=2 is %.2f\n", a)
	fmt.Printf("value of `a` with width=def, precision=1 is %.1f\n", a)
	fmt.Printf("type of `a` is %T\n", a)
}
