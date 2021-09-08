package main

import (
	"fmt"
)

type car struct {
	brand string
	year  int
	km    int
}

func main() {
	m := make(map[string]int)

	m["Name0"] = 14
	m["Name1"] = 20
	m["Name2"] = 25

	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	// Find values for
	value3, ok3 := m["Name3"]
	value1, ok1 := m["Name1"]
	fmt.Println(value1, ok1)
	fmt.Println(value3, ok3) // Zero value

	myCar := car{brand: "Ford", year: 2020, km: 5000}
	fmt.Println(myCar, myCar.brand)

	var otherCar car
	otherCar.brand = "Mazda"
	fmt.Println(otherCar)
	year := otherCar.year
	fmt.Println(year)

}
