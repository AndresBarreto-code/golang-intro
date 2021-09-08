package main

import (
	alias "CursoGolangPlatzi/HelloWorld/access/mypackage"
	"fmt"
)

func main() {
	fmt.Println("1")
	var myCar alias.CarPublic
	myCar.Brand = "Ford"
	fmt.Println(myCar)

	alias.PrintMessage("hello")

}
