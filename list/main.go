package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	for i := range text {
		if text[i] != text[len(text)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	res := isPalindrome("Amor a roma")
	fmt.Println(res)

	// Array (Inmutable)
	var array [4]int
	array[0] = 10
	array[1] = 20
	fmt.Println(array, len(array), cap(array))

	// Slice (Mutable)
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// Methods

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, []int{8, 9, 10}...)

	fmt.Println(slice)

	slice0 := []string{"hello", "what", "is", "up"}

	for i, valor := range slice0 {
		fmt.Println(i, valor)
	}

	for _, valor := range slice0 {
		fmt.Println(valor)
	}

	for i := range slice0 {
		fmt.Println(i)
	}

}
