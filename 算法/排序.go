package main

import (
	"fmt"
	"math/rand"
	"time"
)

func loop(n int) {
	if n == 101 {
		return
	}
	fmt.Println(n, "\n")
	loop(n + 1)

}

func binarysearch(arr []int, begin int, end int, val int) (pos int) {
	mid := (begin + end) / 2

	if begin > end {
		fmt.Println("val not found")
		return -1
	}

	if arr[mid] > val {
		binarysearch(arr, begin, mid-1, val)
	} else if arr[mid] < val {
		binarysearch(arr, mid+1, end, val)
	} else if arr[mid] == val {
		fmt.Println("val found:", mid)
		return mid
	}

	return
}

func show(arr []int) {
	for _, val := range arr {
		fmt.Print(" ", val)
	}
	fmt.Print("\n")
}
func bubblesort(arr []int, size int) {
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func selectsort(arr []int, size int) {
	for i := 0; i < size-1; i++ {
		k := i
		for j := i + 1; j < size; j++ {
			if arr[k] > arr[j] {
				k = j
			}

		}
		arr[k], arr[i] = arr[i], arr[k]

	}
}

/*func insertsort(arr []int , size int){
for i := 1 ; i < size ; i++{
	key := arr[i]
	j := i -1
	for; j >= 0 ; j --{
		if arr[j] < key{
			break}
		arr[j +1] = arr[j]
		}
		arr[j +1] = key
		}*/
// func insertsort(arr []int, size int) {
// 	for i := 1; i < size-1; i++ {
// 		k := i - 1
// 		j := i - 1
// 		for ; j >= 0; j-- {
// 			if arr[j] > arr[i] {
// 				break
// 			}
// 			arr[j+1] = arr[j]

//			}
//			arr[j+1] = arr[k]
//			show(arr)
//		}
//	}
//
//	func insertsort(arr []int, size int) {
//		for i := 1; i < size; i++ {
//			k := i
//			j := i
//			for ; j >= 0; j-- {
//				if arr[k] > arr[j] {
//					break
//				}
//				arr[j+1] = arr[j]
//			}
//			arr[j+1] = arr[k]
//			show(arr)
//		}
//	}
// func insertsort(arr []int, size int) {
// 	for i := 1; i < size; i++ {
// 		j := i - 1
// 		k := i
// 		for ; j >= 0; j-- {
// 			if arr[k] > arr[j] {
// 				break
// 			}
// 			arr[j+1] = arr[j]
// 		}
// 		arr[j+1] = arr[k]
// 		show(arr)
// 	}
// }

// func insertsort(arr []int, size int) {
// 	for i := 1; i < size; i++ {
// 		k := arr[i]
// 		j := i - 1
// 		for ; j >= 0; j-- {
// 			if k > arr[j] {
// 				break
// 			}
// 			arr[j+1] = arr[j]
// 		}
// 		arr[j+1] = k
// 		show(arr)
// 	}
// }

func insertsort(arr []int, size int) {
	for i := 0; i < size-1; i++ {
		j := i
		k := arr[i+1]
		for ; j >= 0; j-- {
			if arr[j] < k {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = k
	}

}

func partation(arr []int, begin int, end int) (pos int) {
	val := arr[begin]
	for begin < end {

		for begin < end && arr[end] > val {
			end--
		}
		if begin < end {
			arr[begin] = arr[end]
			begin++
		}
		for begin < end && arr[begin] < val {
			begin++
		}
		if begin < end {
			arr[end] = arr[begin]
			end--
		}
		//show(arr)

		// if arr[end] < val {
		// 	arr[begin] = arr[end]
		// 	begin++
		// }

		// if arr[begin] > val {
		// 	arr[end] = arr[begin]
		// 	end--
		// }
	}
	arr[end] = val
	return end
}

func quicksort(arr []int, begin int, end int) {
	if begin > end {
		return
	}
	pos := partation(arr, begin, end)

	quicksort(arr, begin, pos-1)
	quicksort(arr, pos+1, end)

}

func main() {
	arr := []int{12, 23, 42, 1, 53, 85, 99, 5, 2}

	show(arr)
	fmt.Println("====================================\n")
	//bubblesort(arr, len(arr))
	//selectsort(arr, len(arr))
	//insertsort(arr, len(arr))
	quicksort(arr, 0, 8)
	fmt.Println("====================================\n")
	show(arr)
	binarysearch(arr, 0, 8, 42)
	binarysearch(arr, 0, 8, 4)
	var myarr [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		myarr[i] = rand.Intn(100)
	}
	show(myarr[:])

	loop(1)
}
