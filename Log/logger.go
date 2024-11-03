package main

import (
	"fmt"
	"io"
	"os"
	//"time"
)

func f1() {
	reader, err := os.Open("./logger.go")
	if err != nil {
		fmt.Println("open file failed : ", err)
		return //er
	}
	b := [1024]byte{}

	for {
		n, err := reader.Read(b[:])
		if err == io.EOF {
			fmt.Println("read flie over")
			return
		}

		if err != nil {
			fmt.Println("read file failed : ", err)
			return
		}

		fmt.Println(string(b[:n]))
		// for _, x := range b[:n] {
		// 	fmt.Println(x)
		// }
	}

}

func main() {
	f1()
}
