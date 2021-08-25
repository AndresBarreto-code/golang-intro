package main

import "fmt"

func main() {
	// Float const
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)

	// Int variables
	width := 123
	var height int = 14
	fmt.Println("width", width)
	fmt.Println("height", height)

	// Zero values
	var a bool
	var b int
	var c float32
	var d float64
	var e string
	fmt.Println("bool", a)
	fmt.Println("int", b)
	fmt.Println("float32", c)
	fmt.Println("float64", d)
	fmt.Println("string", e)

	// Area of a square
	const x = 100.3
	area := x * x
	fmt.Println("area", area)

}
