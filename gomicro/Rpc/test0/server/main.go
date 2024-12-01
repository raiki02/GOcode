package main

import (
	"net/http"
	"net/rpc"
)

type Nums struct {
	Num1, Num2 int
}

type Calc struct{}

func (c Calc) Multiply(n Nums, res *int) error {
	*res = n.Num1 * n.Num2
	return nil
}

func (c Calc) divide(n Nums, res *float64) error {
	*res = float64(n.Num1) / float64(n.Num2)
	return nil
}

func (c Calc) Remainder(n Nums, res *int) error {
	*res = n.Num1 % n.Num2
	return nil
}

func main() {
	calc := new(Calc)

	rpc.Register(calc)

	rpc.HandleHTTP()

	http.ListenAndServe(":8080", nil)
}
