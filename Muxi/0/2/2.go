// package main

// import "fmt"

// func main() {
//         T := 0 //测试组的个数

//         fmt.Scanln(&T)
//         /* if T <= 0 {
//                 fmt.Println("err : T <= 0")
//         }*/
//         n := 0 //元素个数
//         for i := 0; i < T; i++ {
//                 arr := []int{}
//                 fmt.Scanln(&n)
//                 if n == 1 {
//                         fmt.Println("syj")
//                         break
//                 } else if n == 2 {
//                         fmt.Println("cc")
//                         break
//                 } else if n > 2 {
//                         for j := 0; j < n; j++ {
//                                 ele := 0 //每次插入元素值
//                                 fmt.Scanln(&ele)
//                                 arr = append(arr, ele)
//                         }

//                         if (arr[0]+arr[n-1])%2 != 0 {
//                                 fmt.Println("cc") // 一开始首尾和为奇，syj取不了数(?
//                         } else {
//                                 //和为偶，取数  syj判断取头还是尾
//                                 last := n - 1
//                                 begin := 0
//                                 for begin <= last {
//                                         // if (arr[begin + 1] + arr[last])  % 2 == 0{

//                                         // } else if (arr[begin]+ arr[last - 1 ] % 2 == 0{

//                                         // }
//                                 }

//                         }
//                 }

//         }
//         /*else {
//                 fmt.Println("err : n <= 0")
//         }*/
// }

// ============================================================================================
// 第二版(
package main

import "fmt"

func main() {
	T := 0
	n := 0
	lace := []int{}
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {
		fmt.Scanln(&n)
		for j := 0; j < n; j++ {
			ele := 0
			fmt.Scanln(&ele)
			lace = append(lace, ele)
		}
	}

	if n%2 == 0 {
		fmt.Println("cc")
	} else {
		fmt.Println("syj")
	}
}
