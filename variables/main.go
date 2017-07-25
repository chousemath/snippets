package main

import "fmt"

// should just use int unless there is a specific reason to use something different
var a, b, c int
var e, f, g bool
var h, i, j float64 = 123.111, 234.222, 345.333
var aa, bb string

func main() {
	var k, l, m = true, false, "aloha"
	n := "var 1"
	o := "var 2"
	p := "var 3"
	q := 10
	r := 3

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
	fmt.Printf("k = %t\nl = %t\nm = %s\n", k, l, m)
	fmt.Printf("n = %s\no = %s\np = %s\n", n, o, p)

	fmt.Printf("\nThe string zero values are\naa = %s, bb = %s\n", aa, bb)
	fmt.Printf("q/r = %d / %d = %.3f\n", q, r, divFloat(float64(q), float64(r)))
}

func divFloat(a, b float64) float64 {
	return a / b
}
