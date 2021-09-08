package mypackage

import "fmt"

// Private
type carPrivate struct {
	brand string
	year  int
	km    int
}

// CarPublic to create a car
type CarPublic struct {
	Brand string
	year  int // private
	Km    int
}

// PrintMessage to print message
func PrintMessage(message string) {
	fmt.Println(message)
}

// private
func printMessage(message string) {
	fmt.Println(message)
}
