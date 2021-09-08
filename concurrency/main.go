package main

import (
	"fmt"
	"sync"
)

func say(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(message)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("0")

	wg.Add(3)
	go say("1", &wg)
	go say("2", &wg)
	go say("3", &wg)
	go say("4", &wg)
	wg.Wait()

	// time.Sleep(time.Millisecond * 1)
}
