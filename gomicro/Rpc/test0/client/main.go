package main

import (
	"fmt"
	"net/rpc"
)

type Nums struct {
	Num1, Num2 int
}

func main() {
	conn, _ := rpc.DialHTTP("tcp", ":8080")

	ret := 0

	_ = conn.Call("Calc.Multiply", Nums{2, 3}, &ret)
	fmt.Println("Calc.Multiply: ", ret)

	_ = conn.Call("Calc.divide", Nums{2, 3}, &ret)
	fmt.Println("Calc.divide: ", ret)

	_ = conn.Call("Calc.Remainder", Nums{2, 3}, &ret)
	fmt.Println("Calc.Remainder: ", ret)

}
