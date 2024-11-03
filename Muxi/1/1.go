// // package main

// // import (
// // 	"fmt"
// // )

// // func move(nums []int, k int) {

// // 	//递归结束条件
// // 	if k > num {
// // 		fmt.Println("yes")

// // 	}
// // 	if k < 0 {
// // 		fmt.Println("no")
// // 	} else {
// // 		move(nums, nums[k+nums[k]])
// // 	}

// // }

// // var num int

// // func main() {

// // 	fmt.Scanln(&num) // 输入节点个数
// // 	path := []int{}
// // 	//fmt.Println(num_block)
// // 	//fmt.Println(len(num_block))
// // 	for i := 0; i < num; i++ {
// // 		var step int
// // 		fmt.Scanln(&step) //输入每步指定移动
// // 		path = append(path, step)
// // 		//fmt.Println(num_block)
// // 		//fmt.Println(len(num_block))
// // 	}

// // 	move(path, 0)
// // }

// package main

// //创造路径
// //计算移动，判断
// //输出结果
// import (
// 	"fmt"
// )

// func move(nums []int, k int) {

// 	//递归结束条件
// 	if k >= num {
// 		fmt.Println("Yes")
// 		return
// 	}
// 	if k < 0 {
// 		fmt.Println("No")
// 		return
// 	} else {
// 		move(nums, nums[k+nums[k]])
// 	}

// }

// var num int

// func main() {

// 	fmt.Scanln(&num) // 输入节点个数
// 	path := []int{}
// 	for i := 0; i < num; i++ {
// 		var step int
// 		fmt.Scanln(&step) //输入每个节点移动步数
// 		path = append(path, step)
// 	}

// 	move(num_block, 0) //进行移动
// }

package main

//创造路径
//计算移动，判断
//输出结果
import (
	"fmt"
)

func move(nums []int, k int) {

	//递归结束条件
	if k >= num {
		fmt.Println("Yes")
		return
	}
	if k < 0 {
		fmt.Println("No")
		return
	} else {
		move(nums, nums[k+nums[k]])
	}

}

var num int

func main() {

	fmt.Scanln(&num) // 输入节点个数
	path := []int{}
	for i := 0; i < num; i++ {
		var step int
		fmt.Scan(&step) //输入每个节点移动步数
		path = append(path, step)
	}

	move(path, 0) //进行移动
}
