package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sp string
	for i := 0; i < len(os.Args); i++ {
		s += sp + os.Args[i]
		fmt.Println(s)
	}
}
