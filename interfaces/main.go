package main

import (
	"fmt"
)

type figure interface {
	area() float64
}

type square struct {
	width float64
}

type rectangle struct {
	width  float64
	height float64
}

func (c square) area() float64 {
	return c.width * c.width
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func calculate(f figure) {
	fmt.Println("Area: ", f.area())
}

func main() {
	mySquare := square{width: 2}
	myRectangle := rectangle{width: 2, height: 3}
	calculate(mySquare)
	calculate(myRectangle)

	// Interfaces List
	myInterfaces := []interface{}{"Hola", 12, 4.9, true}
	fmt.Println(myInterfaces)
	for i, v := range myInterfaces {
		fmt.Println(i, v)
		fmt.Printf("Data type: %T\n", v)
	}

}
