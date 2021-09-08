package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (yourPc pc) ping() {
	fmt.Println(yourPc.brand, "pong")
}

func (yourPc pc) String() string {
	return fmt.Sprintf("A pc with %d GB RAM, %d GB on diks from %s ", yourPc.ram, yourPc.disk, yourPc.brand)
}

func (yourPc *pc) duplicateRAM() {
	yourPc.ram = 2 * yourPc.ram
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	// & memory
	// * value in memory

	*b = 100

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 200, brand: "mac"}
	myPc.ping()
	myPc.duplicateRAM()
	fmt.Println(myPc)
	myPc.duplicateRAM()
	fmt.Println(myPc)
}
