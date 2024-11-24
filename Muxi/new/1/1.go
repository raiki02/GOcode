package main

import "fmt"

func main() {
	in := ""
	fmt.Scanln(&in)
	arr := make([]int, 0)
	for i := 0; i < len(in); i++ {
		arr[i] = int(in[i])
	}

}
