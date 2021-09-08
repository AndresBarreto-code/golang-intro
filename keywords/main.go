package main

import "fmt"

func main() {
	// defer when have to be invoked mandatory
	defer fmt.Println("Bye")
	fmt.Println("World")
	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("It's two")
			continue
		}
		fmt.Println(i)
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
