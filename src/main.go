package main

import "fmt"

func floatConst() {
	// Float const
	const pi float64 = 3.14
	const pi2 = 3.1416
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)
}

func intVariables(width, height int, message string) {
	// Int variables
	fmt.Println(message, width)
	fmt.Println(message, height)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturnValue(a, b int) (c, d int) {
	return a, a * b
}

func main() {

	floatConst()
	intVariables(10, 12, "int")
	value := returnValue(2)
	fmt.Println("Value: ", value)

	value1, value2 := doubleReturnValue(2, 3)
	fmt.Println("Value1: ", value1)
	fmt.Println("Value2: ", value2)

	_, value3 := doubleReturnValue(2, 3) // escape vars
	fmt.Println("Value3: ", value3)

	value4, _ := doubleReturnValue(2, 3) // escape vars
	fmt.Println("Value4: ", value4)

	// Complex variables
	const h = 123 + 123i
	h1 := 10 + 10i
	var h2 = 2i
	fmt.Println("h", h)
	fmt.Println("h1", h1)
	fmt.Println("h2", h2)

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

	// Arithmetic overflow
	var x = 12
	const y = 15
	// Area of a square
	area := x * x
	fmt.Println("area: ", area)
	fmt.Println("sum: ", x+y)
	fmt.Println("subs: ", x-y)
	fmt.Println("multiplication: ", x*y)
	fmt.Println("division: ", x/y)
	fmt.Println("Modulo: ", x%y)
	x++
	fmt.Println("Incremental: ", x)
	x--
	fmt.Println("Decremental: ", x)

	// // FMT Module
	hello := "Hello"
	world := "World"

	// Printf
	fmt.Printf("%s some aditional text %s, and a number %d\n", hello, world, 500)
	// %s for string
	// %d for number
	// %v for unkown value type
	fmt.Printf("%s some aditional text %v, and a number %d\n", hello, world, 500)

	// Sprintf
	message := fmt.Sprintf("%s some aditional text %v, and a number %d\n", hello, world, 500)
	fmt.Println(message)

	// Data type
	fmt.Printf("Data type: %T\n", hello)
	fmt.Printf("Data type: %T\n", 500)
	fmt.Printf("Data type: %T\n", 500.0)
	fmt.Printf("Data type: %T\n", true)

}
