package main

import "fmt"

var a, b, c int
var e, f, g bool
var h, i, j float64 = 123.111, 234.222, 345.333

func main() {
	fmt.Printf("Zero values of a, b, and c are %d, %d, and %d\n", a, b, c)
	a = 1
	b = 2
	c = 3
	fmt.Printf("After assignment, they are %d, %d, and %d\n", a, b, c)
	fmt.Printf("Zero values of e, f, and g are %t, %t, and %t\n", e, f, g)
	e = true
	f = true
	g = true || false
	fmt.Printf("After assignment, they are %t, %t, and %t\n", e, f, g)
	fmt.Printf("h = %g\ni = %g\nj = %g\n", h, i, j)
}
