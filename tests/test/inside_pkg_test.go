package test

import (
	"fmt"
	"t/pkg"
	"testing"
)

func TestFunc(t *testing.T) {
	p := pkg.InitPerson("ly", 111)
	fmt.Println("before edit:", p)
	pkg.EditPerson(p, "yl", 222)
	fmt.Println("after edit:", p)
}
