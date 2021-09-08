package main

import (
	"fmt"
	"sync"
	"time"
)

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("0")

	wg.Add(4)
	go say("1", &wg)
	go say("2", &wg)
	go say("3", &wg)
	go say("4", &wg)
	wg.Wait()

	go func(text string) {
		fmt.Println(text)
	}("Anonymous function")

	time.Sleep(time.Millisecond * 1)

}
