// package main

// import "fmt"

// type Stack struct {
// 	arr []int
// 	len int
// }

// func (s Stack) push(val int) {
// 	s.arr[len] = val
// 	s.len++

// }

// func (s Stack) pop() {
// 	s.len--
// }

// func (s Stack) show() {
// 	for i := 0; i < s.len; i++ {
// 		fmt.Print(s.arr[i], " ")
// 	}
// 	fmt.Print("\n")
// }

// func (s Stack) empty() (isempty bool) {
// 	if s.len == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	s := Stack{arr[1]int{}, 0}

// 	s.push(12)
// 	s.push(24)
// 	s.push(5)
// 	s.push(26)
// 	s.show()
// 	s.pop()
// 	s.pop()
// 	s.show()
// }



package main()

import "fmt"

type stack struct {
	arr [...]int
}

