package main

import (
	"fmt"
	"test/protocol"
)

func main() {
	protocol.F1()
	fmt.Println("helo in main")

	fmt.Println("test in main")

	// m1 := make(map[int]byte)
	// // fmt.Println(m1, len(m1))
	// // m1[0] = "tom"
	// // fmt.Println(m1, len(m1))
	// // m1[1] = "jery"
	// // m1[2] = "fcuk"
	// // fmt.Println(m1, len(m1))
	// s := "abcdefghijklmnopqrstuvwxyz"
	// for i := 0; i < 20; i++ {
	// 	m1[i] = byte(s[i])
	// 	fmt.Println(m1, len(m1))
	// }

	// fmt.Println("===============================")
	// s1 := make([]int, 0, 3)
	// fmt.Println(s1, len(s1), cap(s1))
	// old := cap(s1)
	// for i := 0; i < 10000; i++ {
	// 	s1 = append(s1, i*100+i*1+i)
	// 	if old != cap(s1) {
	// 		fmt.Println("old cap(s1) = ", cap(s1))
	// 		old = cap(s1)
	// 	}

	// }

	fmt.Println("WHAT")

}
