package main

import "fmt"

// func find_row(t [][]int)	{
// 	for t[i][]
// }

// func find_col(t [][]int) {
// 	r ,c := 0
// 	for t[r][c] != 0 {
// 		r++
// 	}
// }

func main() {

	//制作桌子
	l, w := 0, 0
	fmt.Scanln(&l, &w)
	table := [][]int{}

	//放布
	for i := 0; i < l; i++ {
		j := 0
		for ; j < w; j++ {
			ele_w := 0
			fmt.Scanln(&ele_w)
			table[i][j] = ele_w
		}
		ele_l := 0
		fmt.Scanln(&ele_l)
		table[i][j] = ele_l
	}

	//查找最大面积并输出

}
