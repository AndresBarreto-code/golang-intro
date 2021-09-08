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

func message(message string, c chan<- string) {
	c <- message
}

func main() {
	c0 := make(chan string, 2)

	fmt.Println("0")

	go read("1", c0)
	go read("2", c0)
	go read("3", c0)
	go read("4", c0)
	speak(c0)
	speak(c0)
	speak(c0)
	speak(c0)

	fmt.Println("Channels")

	c1 := make(chan string, 3)

	c1 <- "0"
	fmt.Println(len(c1), cap(c1))
	c1 <- "1"
	c1 <- "2"
	fmt.Println(len(c1), cap(c1))
	close(c1)

	for v := range c1 {
		fmt.Println(v)
	}

	// select

	email1 := make(chan string, 1)
	email2 := make(chan string)
	go message("message 1", email1)
	go message("message 2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email from email1", m1)
		case m2 := <-email2:
			fmt.Println("Email from email2", m2)
		}
	}

}
