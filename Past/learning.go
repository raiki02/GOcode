// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"io/ioutil"

// 	//Test "test"
// 	"os"
// )

// // // 包的调用  main包放在src文件
// // func main() {
// // 	Test.Test01()
// // }

// ////切片
// // func main() {
// // 	s := [][2]int{{1, 2}, {2, 3}}
// // 	fmt.Print(s)

// // 	var s1 []int
// // 	if len(s1) == 0 {
// // 		fmt.Print("len: s1为空\n")
// // 	}

// // 	// if s1 == nil {
// // 	// 	fmt.Print("nil: s1为空\n")
// // 	// } x
// // 	s2 := []int{1, 2, 3}
// // 	s3 := append(s2, 4)
// // 	fmt.Print(s3, "\n")
// // 	fmt.Println(s3, len(s3), cap(s3))
// // 	fmt.Println(s2, len(s2), cap(s2))

// // 	s3[2] = 300
// // 	fmt.Println(s3, len(s3), cap(s3))
// // 	fmt.Println(s2, len(s2), cap(s2))

// // 	s2[1] = 2000
// // 	fmt.Println(s3, len(s3), cap(s3))
// // 	fmt.Println(s2, len(s2), cap(s2))

// // 	x1 := [...]int{1, 3, 5}
// // 	s11 := x1[:]
// // 	fmt.Println(s11)
// // 	s11 = append(s11[:1], s11[2:]...)
// // 	fmt.Println(s11)
// // }

// // 接口

// // type mover interface {
// // 	move()
// // }

// // type dog struct{}

// // type cat struct{}

// // func (c *cat) move() {
// // 	fmt.Println("cat moves ")
// // }

// // func (d dog) move() {
// // 	fmt.Println("dog moves ")
// // }

// // func main() {
// // 	var x mover
// // 	d_val := dog{}
// // 	d_ptr := &dog{}
// // 	//c_val := cat{}
// // 	c_ptr := &cat{}

// // 	x = d_ptr
// // 	x.move()
// // 	x = d_val
// // 	x.move()

// // 	fmt.Println("=============")

// // 	x = c_ptr
// // 	x.move()
// // 	//x = c_val
// // 	//x.move()
// // }

// // type people interface {
// // 	speak(string) string
// // }

// // type student struct {
// // }

// // func (stu student) speak(think string) (talk string) {
// // 	if think == "sb" {
// // 		talk = "帅比"
// // 	} else {
// // 		talk = "您好"
// // 	}
// // 	return
// // }

// // func meet(peo people, think string) (response string) {
// // 	response = peo.speak(think)
// // 	return
// // }
// // func main() {
// // 	stu := student{}
// // 	var peo people = stu
// // 	think := "sb"
// // 	fmt.Println(meet(peo, think))
// // }

// // type beginner interface {
// // 	caller
// // 	ring()
// // }

// // type caller interface {
// // 	called()
// // }

// // type MI struct {
// // 	name string
// // }

// // func (m MI) called() {
// // 	fmt.Println("Mi is called")
// // }

// // func (m MI) ring() {
// // 	fmt.Println("dududu")
// // }

// // type OPPO struct {
// // 	name string
// // }

// // func (o OPPO) called() {
// // 	fmt.Println("Oppo is called")
// // }

// // // func (cd caller) start() {
// // // 	cd.called()
// // // }

// // //	func (bg beginner) motivate() {
// // //		bg.ring()
// // //		bg.called()
// // //	}

// // func nilitf(a ...interface{}) {
// // 	fmt.Println("nilitf called.")
// // }
// // func main() {
// // 	// m := MI{"Redmi"}
// // 	// o := OPPO{"Op_promax"}
// // 	// var (
// // 	// 	bg  beginner = m
// // 	// 	cd1 caller   = m
// // 	// 	cd2 caller   = o
// // 	// )
// // 	// fmt.Println("m.called")
// // 	// m.called()
// // 	// fmt.Println("m.ring")
// // 	// m.ring()
// // 	// fmt.Println("o.called")
// // 	// o.called()

// // 	// fmt.Println("bg.c")
// // 	// bg.called()
// // 	// fmt.Println("bg.r")
// // 	// bg.ring()
// // 	// fmt.Println("cd1.c")
// // 	// cd1.called()
// // 	// fmt.Println("cd2.c")
// // 	// cd2.called()
// // 	/*// bg.motivate()
// // 	// cd1.start()
// // 	// cd2.start()

// // 	// m.start()
// // 	// o.srart()err*/
// // 	flag := true
// // 	nilitf(1)
// // 	nilitf("hello")
// // 	nilitf(flag)
// // 	nilitf('a')

// // }

// // 闭包
// // func main() {
// // 	n := 0
// // 	fmt.Println("main:n =", n)
// // 	func() {
// // 		n = 1
// // 		fmt.Println("func:n = ", n)
// // 	}()
// // 	fmt.Println("main:n =", n)
// // }

// // 函数类型
// // type functype func(a int, b int) (c int)

// // func Calculator(a int, b int, cal functype) (c int) {
// // 	c = cal(a, b)
// // 	return c
// // }

// // func add(a, b int) (c int) {
// // 	c = a + b
// // 	return c
// // }

// // func minus(a, b int) (c int) {
// // 	c = a - b
// // 	return c
// // }

// // func multiply(a, b int) (c int) {
// // 	c = a * b
// // 	return c
// // }

// // func divide(a, b int) (c int) {
// // 	if b == 0 {
// // 		err := "b is 0 !"
// // 		fmt.Println("err: ", err)
// // 	}
// // 	c = a / b
// // 	return c
// // }

// // func main() {
// // 	a, b := 2000, 100
// // 	fmt.Println("Calculator(a, b, add) = ", Calculator(a, b, add))
// // 	fmt.Println("Calculator(a, b, minus) = ", Calculator(a, b, minus))
// // 	fmt.Println("Calculator(a, b, multiply) = ", Calculator(a, b, multiply))
// // 	fmt.Println("Calculator(a, b, divide) = ", Calculator(a, b, divide))
// // }

// //输入和输出

// // func main() {
// // 	var (
// // 		name  string
// // 		age   int
// // 		class string
// // 	)
// // 	// fmt.Scanln(&name, &age, &class)
// // 	// fmt.Println("输入:", name, age, class)
// // 	// fmt.Scanf("name, age , class: \n", &name, &age, &class) err:""中写%占位符
// // 	fmt.Printf("name, age , class:\n", name, age, class)
// // }

// func usebufio() {
// 	// objfile, _ := os.OpenFile("./xxx.txt", os.O_APPEND|os.O_RDWR|os.O_TRUNC,0666)
// 	// reader := bufio.NewReader(objfile)
// 	var s string
// 	fmt.Println("buf:")
// 	reader := bufio.NewReader(os.Stdin)
// 	s, _ = reader.ReadString('\n')
// 	fmt.Println("system read:", s)
// }

// func usescanf() {
// 	var s string
// 	fmt.Println("scanf:")
// 	fmt.Println("write: ")
// 	fmt.Scanln(&s, "\n")
// 	fmt.Println("read: ", s)

// }

// func main() {
// 	usescanf()
// 	//usebufio()
// }

// //文件读写
// //调用函数生成对象 并判断是否出错
// //对象调用函数，返回结果
// //打印结果

// // 一次规定读多少
// func readfromfile1() {
// 	ObjFile, err := os.Open("./learning.go")
// 	if err != nil {
// 		fmt.Println("file open failed , err: %v", err)
// 		return
// 	}

// 	defer ObjFile.Close()

// 	temp := make([]byte, 256)
// 	//	temp := [256]byte{}
// 	for {
// 		n, err := ObjFile.Read(temp)
// 		if err == io.EOF {
// 			fmt.Println("file read over")
// 			return
// 		}
// 		if err != nil {
// 			fmt.Println("file read failed , err: %v", err)
// 			return
// 		}

// 		fmt.Printf("%d bytes read \n", n)
// 		fmt.Println("content: ", string(temp[:n]))
// 		if n < 256 {
// 			return
// 		}
// 	}

// }

// // 一次读一行
// func readfrombufio() {
// 	objfile, err := os.Open("./learning.go")
// 	if err != nil {
// 		fmt.Println("file open failed, err: ", err)
// 		return
// 	}
// 	defer objfile.Close()

// 	reader := bufio.NewReader(objfile)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			fmt.Println("read over")
// 			return
// 		}
// 		if err != nil {
// 			fmt.Println("file read failed , err: ", err)
// 			return
// 		}

// 		fmt.Print(line)
// 	}
// }

// // 一次读完
// func readfilebyioutil() {

// 	ret, err := ioutil.ReadFile("./learning.go")
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 	}
// 	fmt.Println("content: ", string(ret))

// }

// // func main() {
// // 	//readfrombufio()
// // 	//readfilebyioutil()
// // }

// // 写文件
// func writedemo1() {
// 	objfile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return
// 	}

// 	objfile.Write([]byte("sbly\n"))
// 	objfile.WriteString("ly is asshole")
// 	objfile.Close()
// }

// func writedemo2() {
// 	objfile, err := os.OpenFile("./log.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return
// 	}

// 	defer objfile.Close()

// 	writer := bufio.NewWriter(objfile)
// 	writer.Write([]byte("sb ly\n"))
// 	writer.WriteString("liu yuan is sb")

// 	writer.Flush()

// }

// func writedemo3() {
// 	ioutil.WriteFile("./log.txt", []byte("caonima"), 0666)
// 	//nil...
// }

// func main_write() {
// 	//writedemo1()
// 	//writedemo2()
// 	//writedemo3()
// }

//goroutine
// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func f() {
// 	rand.Seed(time.Now().UnixNano())
// 	for i := 0; i < 10; i++ {
// 		r1 := rand.Int()
// 		r2 := rand.Intn(10)
// 		fmt.Println(r1, r2)
// 	}
// }

// var wg sync.WaitGroup

// func f1(i int) {
// 	defer wg.Done()

// 	time.Sleep(time.Second * time.Duration(rand.Intn(4)))
// 	fmt.Println(i)
// }

// func main() {
// 	// for i := 0; i < 10; i++ {
// 	// 	wg.Add(1)
// 	// 	go f1(i)
// 	// }

// 	//wg.Wait()

// 	//f()
// }

// 管道
// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func f1(a chan int) {
// 	for i := 1; i < 101; i++ {
// 		a <- i
// 	}
// 	defer close(a)
// }

// func f2(a chan int, b chan int) {
// 	for x, ok := range a {
// 		if ok != nil {
// 			break
// 		}
// 		b <- x * x
// 	}

// 	defer close(a)
// 	defer close(b)

// }