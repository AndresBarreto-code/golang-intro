package main

import (
	"fmt"
)

func read(message string, c chan<- string) {
	c <- message
}

func speak(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	c := make(chan string, 2)

	fmt.Println("0")

	go read("1", c)
	go read("2", c)
	go read("3", c)
	go read("4", c)
	speak(c)
	speak(c)
	speak(c)
	speak(c)

}
