package main

import "fmt"

func pass(n int, t int, last int) {

	//递归结束条件
	// if t == 0 && n != last {
	// 	return
	// } else if t == 0 && n == last {
	// 	Sum += 1
	// 	return
	// }

	if t == 0 {
		if n != last {
			return
		} else {
			Sum += 1
			return
		}
	}

	if n == last { //走到尾
		pass(n-1, t-1, last)
		pass(1, t-1, last)
	} else if n == 1 { //走到头
		pass(2, t-1, last)
		pass(last, t-1, last)
	} else {
		pass(n-1, t-1, last)
		pass(n+1, t-1, last)
	}

}

var n, t int

// 最后一个人
var Sum int = 0 // 存在次数
//const Minp = 1   // 第一个人
//var Begin = n    //起始位置

func main() {
	fmt.Scanln(&n, &t)
	var last int = n
	pass(n, t, last)
	fmt.Println(Sum)
}
