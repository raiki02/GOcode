package main

import (
	"fmt"
	"t/pkg"
)

func main() {
	p1 := pkg.InitPerson("raiki", 3)
	fmt.Println("p1 = ", p1)
	pkg.EditPerson(p1, "raiki", 15)
	fmt.Println("p1 = ", p1)
	p := pkg.C1.GetPerson()
	fmt.Println("p = ", p)
}
