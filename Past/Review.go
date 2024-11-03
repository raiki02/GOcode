
package main
import (
	"fmt"
	_"time"
	_"strings"
	_"strconv"
	_"os"
)


//2.2变量
//声明变量
// var v1 int //全局
// var (
// 	v2 int
// 	v3 string
// )//多个异类
// var v9 , v10 ,v11 int //多个同类

// func test01(){
// 	var v1 int//局部
// 	var (
// 		v2 int
// 		v3 byte
// 	)
// }

// //初始化

// var v4 int = 10 //完整
// var v5 = 20  //自动推导1
// v6 := 30 //自动推导2  //这么写之前不能声明

// //赋值
// var v7 int //声明
// v7 = 10   //赋值

// v7 = v6 
// v1 , v2 ,v3 =  1 , 2 ,"string" //多个赋值

// v1 ,v2 = v2 ,v1 //可以实现交换，与cpp不同

// //匿名变量

// var v12 ,v13 int
// _ , v12 ,_ , v13 = 11, 22, 33 ,44 //_被赋值但是无法被调用，传入的值被废弃

// //2.3常量
// //定义+初始化
// const c1 int = 1
// const c2 = 2
// const PI = 3.14
// const (
// 	c3 byte	= 'A'
// 	c4 string = "it's c4"
// )

// //iota枚举
// const (
// 	c5 = iota// 0
// 	c6 //1
// 	c7 //2
// 	c8 = iota //0   //   4   my err
// )

// const (
// 	c9 ,c10 ,c11 //all 0
// 	c12 = "it's c12" 
// 	c13 //2
// 	c14 = iota * 13 // 39
// )

// //2.4数据类型
// var (
// 	t1 bool
// 	b1 := (1 == 2) // v2也会被推导为bool类型

// 	t2 byte //相当于char
// 	t3 int 8/16/32/64
// 	t4 float 32/64
// 	t8 string //``不会有字符转义 输入即输出

// 	t5 complex 64/128 //real(), imag()
// 	t6 rune
// 	t7 uintptr
// )
// //2.5格式
// //c ,d ,p ,v ,T , s 

// //2.6转换
// var ch byte = 'a'
// var a int = int(ch)

// //2.7别名    新 + 原
// type char type 
// var ch1 char

//5 函数

// func factorial (n int) (res int){
// 	if n == 1||n ==0{
// 		return 1
// 	} else{
// 		return factorial(n-1) * n
// 	}
// }

// func main(){

// 	res := factorial(1)
// 	fmt.Println(res)
// }

//5.4函数类型 !!!
 
//type Functype func(int , int ) int //声明一个函数类型f， 它接收两个参数返回一个整形

// func Calc(a int , b int) (c int) {
// 	c = a + b

// 	return c  
// }  //写成一坨

// func Calc (a int , b int, f Functype) {
// 	res := f(a ,b )
// 	fmt.Println("res = " , res )

// }

// func ADD(a int , b int ) (c int) {
// 	c = a + b
// 	return c
// }

// func MUL(a int , b int ) (c int) {
// 	c = a * b
// 	return c
// }

// func MIN(a int , b int) (c int ) {
// 	c = a - b
// 	return c
// }

// func DIV(a int , b int) (c int) {
// 	c = a / b
// 	return c
// }

// func test01(){
// 	a , b := 10 , 5
	
// 	Calc(a , b , ADD)
// 	Calc(a , b , MIN)
// 	Calc(a , b , DIV)
// 	Calc(a , b , MUL)
// }

// func main(){
// 	test01()
// }

//写一个函数类，该类型可以包含所有传入参数，返回值类型相同的函数

//数组

// func test(){
// 	var arr1 [...]int 
// 	arr1 = {1 ,2 ,3 ,4 }

// 	arr2 := [...]int{1 , 2 3, 4,45}

// }

//映射

// func test01(){
// 	m := make(map[int]string)
// 	m[0] = "tom"
// 	m[1] = "jery"
// 	fmt.Println(m)

// 	for i , j := range m{
// 		fmt.Printf("m[%d] = %s\n" , i ,j)
// 	}
// 	fmt.Println("=========")
// 	for _ , j := range m{
// 		fmt.Printf("%s\n" ,j)
// 	}
// 	delete(m, 0)
// 	fmt.Println("=========")
// 	for i , j := range m{
// 		fmt.Printf("m[%d] = %s\n" , i ,j)
// 	}
	
// }

// func test01()
// {
// 	m1 := make(map[string][]int , 10)
// 	m1["北京"] = []int{1,2,3}

// 	s1 := make([]map[string]int, 10, 10)
// 	s1[0]["上海"] = 123


// }

// func main(){

// }