package main

import (
	"fmt"
	"reflect"
)

// func reflectType(x interface{}) {
// 	v := reflect.TypeOf(x)
// 	fmt.Printf("type:%v \n", v)
// }

// type myInt int64

// func reflectType(x interface{}) {
// 	t := reflect.TypeOf(x)
// 	fmt.Printf("type:%v kind:%v\n", t.Name(), t.Kind())
// }

// func main() {
// 	// var a float32 = 3.14
// 	// reflectType(a)
// 	// var b int64 = 100
// 	// reflectType(b)

// 	var a *float32 // 指针
// 	var b myInt    // 自定义类型
// 	var c rune     // 类型别名
// 	reflectType(a) // type: kind:ptr
// 	reflectType(b) // type:myInt kind:int64
// 	reflectType(c) // type:int32 kind:int32
// 	type person struct {
// 		name string
// 		age  int
// 	}
// 	type book struct{ title string }
// 	var d = person{
// 		name: "沙河小王子",
// 		age:  18,
// 	}
// 	var e = book{title: "《跟小王子学Go语言》"}
// 	reflectType(d) // type:person kind:struct
// 	reflectType(e) // type:book kind:struct
// }

// func reflectValue(x interface{}) {
// 	v := reflect.ValueOf(x)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
// 		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
// 	case reflect.Float32:
// 		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
// 		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
// 	case reflect.Float64:
// 		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
// 		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
// 	}
// }
// func main() {
// 	var a float32 = 3.14
// 	var b int64 = 100
// 	reflectValue(a) // type is float32, value is 3.140000
// 	reflectValue(b) // type is int64, value is 100
// 	// 将int类型的原始值转换为reflect.Value类型
// 	c := reflect.ValueOf(10)
// 	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
// }

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v, kind: %v \n", v, v.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Println("v = ", v)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type: int64, value :%d \n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type: Float32, value : %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type: Float64, value :%f \n", float64(v.Float()))
	}
}

func reflectValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
}

func reflectValue2(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(2000)
	}
}

func main() {
	// s := make([]int, 3)
	// reflectType(s)
	// a := 1
	// p := &a
	// reflectType(p)
	// m := make(map[int]string, 1)
	// reflectType(m)

	// var a float32 = 3.14
	// var b int64 = 100
	// reflectValue(a)
	// reflectValue(b)

	// c := reflect.ValueOf(10)
	// fmt.Println("type c : ", c)

	//var a int64 = 100
	//reflectValue1(a)
	fmt.Println("===============")
	var b int64 = 10

	reflectValue2(&b)
	fmt.Println(b)
}
