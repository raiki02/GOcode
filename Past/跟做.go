// // package main

// // import "fmt"

// // // func bianliang(){
// // //     var a int = 1 //1.
// // //     b := 2          //2.
// // //     var c int       //3.
// // //     c = 3

// // //     var d , e int =  12, 23
// // //     _ , f := 122 , 2333
// // //     var(
// // //         g int = 11
// // //         h string = "hhh"
// // //     )
// // // }

// // // func changliang(){
// // //     const a int = 1
// // //     const(
// // //         b int = 2
// // //         c int = 3
// // //         d string = "aaa"
// // //     )

// // //     const(
// // //         a = iota //0
// // //         b //= iota
// // //         c //= iota
// // //     )

// // //     const(
// // //         d = iota  //0
// // //         e
// // //         f
// // //     )
// // // }

// // // func fenlei(){
// // //     var a bool = false
// // //     var b byte = 'a'
// // //     var c rune = 12// ==uint32
// // //     var d int = 23// uint
// // //     var e float32 = 3.13//64
// // //     var f complex64 = 13 + 12.3i  //128
// // //     var g uintptr = &d
// // //     var i string = "hhhh"

// // //     /* 格式化
// // //      %T 类型
// // //      \n 换行
// // //      %d 10进制
// // //      %c 字符
// // //      %f 浮点
// // //      %s 字符串
// // //      %v 自动匹配*/
// // // }

// // // func shurushuchu(){
// // //     var a int
// // //     fmt.Println("请输入一个数值: ")
// // //     fmt.Scan(&a)
// // //     fmt.Println("您输入的是：" , a)
// // // }

// // // func leixingzhuanhuan(){
// // //     var a int = 12
// // //     var(
// // //         b string = "hlasjg"
// // //         c int = 666
// // //     )

// // //     fmt.Printf("输入一个整形: %f" , float32(a))
// // //     //int，float 都不能和bool类型转换
// // // }

// // // func leixingbieming(){
// // //     type char byte //把byte 改成char
// // //     var c char = 'g'
// // // }

// // // func xunhuan(){
// // //     str := "hellowolrdhahahaha"
// // //     for i , j := range str{
// // //        fmt.Printf("str[%d] = %c" , i ,str[i])
// // //     }

// // //     for i , _ /*等同于" i " */ := range str{
// // //         fmt.Printf("str[%d] = %c" , i ,str[i])
// // //     }

// // //     var choice int = 0
// // //     switch choice {
// // //     case 1:
// // //           fmt.Println("您输入的是：" , 1)
// // //           break//可要可不要
// // //     case 2:
// // //           fmt.Println("您输入的是：" , 2)
// // //           fallthrough//继续向下
// // //     case 3:
// // //           fmt.Println("您输入的是：" , 3)
// // //     case 4:
// // //           fmt.Println("您输入的是：" , 4)
// // //     default:
// // //          fmt.Println("您输入的是：" , 111)

// // //     }
// // // }

// // // func budingcanshu(args ...int){
// // //     for i int = 0 ; i < len(args) ; i++ {
// // //          fmt.Printf("args[%d] = %d" , i , args[i])
// // //     }
// // // }

// // // func hanshuchuandi1(temp... int){
// // //     for _ , i := range temp{
// // //         fmt.Println(" " , i)
// // //     }

// // // func hanshuchuandi2(args... int){
// // //     hanshuchuandi1( args[:2]...) // <2
// // //     hanshuchuandi1( args[2:]...) // >= 2
// // // }

// // // func fanhuiyige() (res int){
// // // res = 466
// // // return res

// // // }

// // // func fanhuiduoge() (a int, b int, c int){
// // //     a , b , c = 123 ,234 , 356
// // //     return a , b, c
// // // }

// // // func hanshudigui(a int) {
// // //     if a > 0 {
// // //         hanshudigui( a - 1)
// // //     }

// // //     fmt.Println("a = " , a)
// // //     if a == 0{
// // //         fmt.Println(a)
// // //         return

// // //     }
// // // }

// // // func hello(a int, b string) (c int){
// // //     c = a
// // //     fmt.Println("b = " , b)
// // //     return c
// // // }

// // // type hellotype func( int , string) (int)

// // // func hanshuleixing(){
// // //     var functest hellotype
// // //     functest = hello
// // //     var res int
// // //     res = functest(12 , "ijng")
// // // }

// // // //huidiaohanshu
// // // func add(a , b int) (c int){
// // //     c = a + b
// // //     return c
// // // }

// // // func mul(a , b int ) (c int){
// // //     c = a * b
// // //     return c
// // // }

// // // type cal func(int , int) (int )

// // // func Calc(a ,b int , fun cal) (c int){
// // //     c = fun(a , b)
// // //     return c
// // // }
// // // //niminghanshu

// // // func niminghanshu (){
// // //    var a int =  234
// // //    var v string = "gdsy"

// // //    f1 := func(){
// // //     Printf("a = %d , v = %s" , a ,v)
// // //    }
// // //    f1()

// // //    Println("====================")
// // //    func(){
// // //     Printf("a = %d , v = %s" , a ,v)
// // //    }()
// // //     Println("====================")
// // //   type functype func
// // //   var niming functype
// // //   niming = func(){
// // //     Printf("a = %d , v = %s" , a ,v)
// // //   }
// // //   niming()
// // // }

// // // func nimingtedian(){
// // //     var(
// // //         a := 10
// // //         b := " aaaaa"
// // //     )

// // //     func niminghanshu2(){
// // //         a = 1000
// // //         b ="asdf"
// // //         fmt.Printf("a = %d , b = %s" , a ,b)
// // //     }
// // //       fmt.Printf("a = %d , b = %s" , a ,b)
// // // }

// // func test() func() int {
// // x := 0

// //     return func() int{
// //         x++
// //         fmt.Println("这个是")
// //         return x * x
// //     }
// // }

// // func main() {

// //     // f := test()
// //     // fmt.Println( f())
// //     //   fmt.Println( f())
// //     //     fmt.Println( f())
// //     //       fmt.Println( f())
// // //  defer fmt.Println("aaaaaaaaaaa")
// // //  fmt.Println("bbbbbbbbbbbb")  打印出来还是bbbb，aaaaa

// // }

// package main

// import (
//    . "fmt"
//   _  "os"
// )
// // import(
// //     "time"
// //     "math/rand"
// // )

// // import _ "fmt" // 忽略 ， 不用fmt. 只调用里面的 init函数
// // import aaa "fmt" //别名
// // import . "fmt" 不用fmt. 就用

// // func shuzu(){
// // //     var a  = []int {1,23,4}
// // //    Println(a)

// // //     b := [11]int {123,235,1356,15,5135,84,1361,95769,315,135,63}
// // //    Println(b)

// // //     c := []string { "a" , " asd" , "sdg"}
// // //    Println(c)

// // //     slice := b[0:2:4]
// // //    Println(slice)

// // //     s2 := make([]int , 12 ,24)
// // //     s2 = make([]int , 55)
// // //    Println(s2)
// // //     s3 := b[:]
// // //     Println(s3)

// //     // s4 := make([]int , 0 ,1)
// //     // oldCap := cap(s4)
// //     // for i := 0; i < 100 ; i++{
// //     //     s4 = append(s4 , i)
// //     //     if newCap := cap(s4); newCap > oldCap{
// //     //         Printf( " %d ===> %d \n" , i , newCap)
// //     //         oldCap = newCap
// //     //     }
// //     // }

// //     // s := make([]int , 12)
// //     // s1 := qiepianchuandi(s)
// //     // Println(s1)

// // // n := 10
// // // s := make([]int ,n)
// // // InitData(s)
// // // Println(s)

// //     m := make(map[int]string)
// //     m[1] = "aaa"
// //     m[2] = "bbb"
// //     Println(m)

// //     m2 := make(map[string] bool)
// //     m2["mike"] = true
// //     m2["jert"] = false
// //     Println(m2)
// // }

// // func InitData(s [] int){
// //     rand.Seed(time.Now() .UnixNano())

// //     for i := 0 ; i < len(s) ; i++ {
// //         s[i] = rand.Intn(100)
// //     }
// // }
// // func qiepianchuandi(ab []int) (ab){

// //     for i:= 0; i < 10 ; i++{
// //         ab = append(ab , i)
// //     }

// //     return ab
// // }

// // type Student struct{
// //     age int
// //     sex bool
// //     addr string
// //     name string
// //     id int
// // }

// // func main(){
// // //    Println("helo")
// // //     shuzu()
// //     var p1 *Student = &Student{18 , false ,"3 321" , "mkea" , 01}
// //     s2 := Student{12 , true , "2 123 " , "kley" , 02}
// //     Println(*p1)
// //     Println(s2)

// //     var s3 Student
// //     s3.age = 15
// //     s3.sex = false
// //     s3.addr = "aaa"
// //     s3.name = " asd v"
// //     s3.id = 22
// //     Println(s3)

// //     var s4 Student
// //     var ps4 *Student
// //     ps4 = &s4
// //     ps4.age = 1511
// //     ps4.sex = true
// //     ps4.addr = "aaaa sdadv"
// //     ps4.name = " asd vbdhfne"
// //     ps4.id = 22123
// //     Println(*ps4)

// // }

// type Person struct{
//     id int
//     name string
//     age int
// }

// type Student struct{
//     *Person
//     addr string
//     grade string
// }

// func main(){

//     // s1 := Student{Person{01,"mika",16},"japna","A"}
//     // Println(s1)
//     // s2 := Student{grade: "F"}
//     //   Printf("%v \n" ,s2)
//     // s3 := Student{Person:Person{id: 02}}
//     // Println(s3)
//     // s3.Person. name = "AAA"
//     // Println(s3)
//     // s3.Person = Person{03 , "BBB" , 2}
//     // Println(s3)
//   func () {
//     Println("this is niminghanshu")
// }()

//     // a := make(map[int]string(1, "LKJ")) //err
//     //a[1] = "LIH"
//     Println(a[1])
//     s1 := Student{ &Person{13 , "AAAA" , 24}, "japna", "A"}
//     Println(s1)
//     var s2 Student

//     s2.Person = new(Person)
//     s2.id = 12
//     s2.name = "ASDW"
//     s2.age = 123
//    Println(s2)

// }

// package main

// import ."fmt"

// type Person struct{
//     name string
//     age int
//     id int
// }

// type Student struct{
//     Person
//     grade string

// }

// func
// type long int

// func (a long) mianxiangduixiang (b long) long {
//     return a + b
// }

// func qiepian(){
//     //定义+初始化
//     var(
//         a []int
//         // b = []int{}
//         // c = []int{1 ,2 ,3}
//         // d = c[:2]
//         // e = c[0: 2 : cap(c)]
//         // f = c[ : 0] //默认cap是cap（c）
//         // g = make([]int , 3)   //默认len == cap
//         // h = make([]int , 2 ,3)
//         // i = make([]int , 0 ,3)
//     )

//     //追加+删除
//     a = append(a ,1)
//     a = append(a , 1 , 2 , 3 ,4)
//     a = append(a , []int{11,22,33}...)
//     a = append([]int{0},a...)
//     a = append(a[:3], append([]int{111,222,333} , a[3:]...)...)

//     a = a[1:]
//     a = a[:len(a) - 3]
// // }

// func main(){
// //   var (
// //     a long = 23
// //     b long = 22
// //     e long = 1
// //   )
// //   // a := 23 : err a.mianxiangduixiang undefined (type int has no field or method mianxiangduixiang
// //     c := a.mianxiangduixiang(b)
// //     d := e.mianxiangduixiang(114)
// //     Println(c," " , d)
// //qiepian()

// }

//语义

// type Person struct{
//     name string
//     sex byte
//     age int
// }

// func (p Person) SetInfoValue(){
//     p.name = " AAA "
//     p.sex = 'm'
//     p.age = 21

// }

// func (p * Person) SetInfoPtr(){
//     p.name = " AAA "
//     p.sex = 'm'
//     p.age = 21

// }

// func main() {
//     p1 := Person{"JM" , 'f' , 12}
//     Println(p1)
//     p1.SetInfoValue()
//     Println(p1)
//     (&p1).SetInfoPtr()
//     Println(p1)

// }

//方法集

// type Person struct{
//     name string
//     sex byte
//     age int
// }

// func (p Person) ValFunc(){
//     p.name = "aa"
//     p.sex = 'a'
//     p.age = 11
// }

// func (p *Person) PtrFunc(){
//     p.name = "aaa"
//     p.sex = 'b'
//     p.age = 10
// }

// func main(){
//     p1 := Person{"nima" , 'z' , 123}
//     p2 := Person{"cao" , 'a' , 234}
//     ptr := &p2
//     p1.ValFunc()
//     Printf("p1 = %v ; p2 = %v " , p1 , p2)
//     p1.PtrFunc()
//     Printf("p1 = %v ; p2 = %v " , p1 , p2)
//    ptr.ValFunc()
//     Printf("p1 = %v ; p2 = %v " , p1 , p2)
//     ptr.PtrFunc()
//     Printf("p1 = %v ; p2 = %v " , p1 , p2)

// }

//方法继承 + 重写

// type Person struct{
//     name string
//     sex byte
//     age int
// }

// func (tmp *Person) PrintInfo(){
//     Println(tmp)
// }

// type Student struct{
//     Person
//     id int
//     addr string
// }

// func (tmp *Student) PrintInfo(){
//     Println(tmp)
// }

// func main(){
//     s := Student{Person{"ads" , 'a' , 12} , 123 , "jp"}
//     Println(s)
//     s.PrintInfo()
//     s.Person.PrintInfo()
// }

//方法值
// type Person struct{
//     name string
//     sex byte
//     age int

// }

// func (p Person) InfoVal(){
//     Printf("Val: %p , %v\n" , & p, p)
// }

// func (p * Person) InfoPtr(){
//     Printf("Ptr: %p , %v\n" , p , p)
// }

// func main(){
//     p1 := Person{"a" , 's' , 12}
//     Printf("main: %p , %v\n" , &p1 , p1)

//     p1.InfoVal()
//     p1.InfoPtr()

//     vfunc := p1.InfoVal
//     vfunc()

//     pfunc := p1.InfoPtr
//     pfunc()
// }

//方法表达式

// type Person struct{
//     name string
//     sex byte
//     age int

// }

// func (p Person) InfoVal(){
//     Printf("Val: %p , %v\n" , & p, p)
// }

// func (p * Person) InfoPtr(){
//     Printf("Ptr: %p , %v\n" , p , p)
// }

// func main(){
//     p1 := Person{"mike" , 'm' , 12}

//     f := (Person).InfoVal
//     f(p1)
// }

//接口

// type Phone interface{
//     greet()
// }

// type Iphone struct{
//      a int
// }

// type Mi struct{
//     a int
// }

// type Oppo struct{
//     a int
// }

// func greet (p Phone) {
//     p.greet()
// }
// func (i Iphone) greet (){
//     Println("我是Iphone")
// }

// func (m Mi) greet (){
//     Println("我是Mi")
// }
// func (o Oppo) greet (){
//     Println("我是Oppo")
// }

// func main(){
//    ip :=  Iphone{01}
//    m := Mi{02}
//    o := Oppo{03}

//    greet(ip)
//    greet(m)
//    greet(o)
// }

//----------------------------------------
// type Humaner interface{
//     sayhi()
// }

// type Student struct{
//     name string
//     id int
// }

// func (s Student) sayhi(){
//     Println("student say hi")
// }

// type Teacher struct{
//     addr string
//     group string
// }

// func (t Teacher) sayhi(){
//     Println("teacher say hi")
// }

// type MyStr string

// func (tmp *MyStr) sayhi(){
//     Println( " mystr say hi")
// }
// // type i Humaner err

// func Whosayhi( u Humaner){
//     u.sayhi()
// }

// func main(){
//     s :=  Student{"mike", 01}
//     t := Teacher{"jp" , "cpp"}
//     var str MyStr = "halo"

//     Whosayhi(s)
//       Whosayhi(t)
//         Whosayhi(&str)

//     x := make([]Humaner , 3)
//     x[0] = s
//     x[1] = t
//     x[2] = &str

//     for _ , i := range x{
//         i.sayhi()
//     }
// }

// func main01(){

//     var i Humaner

//    s :=  Student{"mike", 01}
//    i = s
//     i.sayhi()
//    t := Teacher{"jp" , "cpp"}
//     i = t
//     i.sayhi()
//   var str MyStr = "halo"
//     i = &str
//     i.sayhi()

// }

//接口继承 + 转换

// type Humaner interface{
//     sayhi()
// }

// type Personer interface{
//     Humaner
//     sing(lrc string)
// }

// type Student struct{
//     name string
//     id int
// }

// func (s *Student) sayhi (){
//     Println("学生在say hi")
// }

// func (s *Student) sing(lrc string){
//     Println("学生在唱：" , lrc)
// }

// func main(){
//     var iPro Personer
//     var i Humaner

//     i = iPro
//     //iPro = i //err
// }

// func main(){
//     s := &Student{"mike " , 12}
//     var i Personer
//     i = s
//     i.sayhi()
//     i.sing("ocanada")
// }

//空接口
// func fname(arg ...interface{}){
//     //接受任意且多个类型
// }

// func main(){
//     var i interface{} = 1
//     Println(i)

//     i = " asdb"
//     Println(i)
// }

//类型断言 if ， switch

// type Student struct{
//         name string
//         id int
// }

// func main(){
//     i := make([]interface{} , 3)
//     i[0] = 1
//     i[1] = "stringaa"
//     i[2] = Student{"kaie" , 11}

// // for index , data := range i {
// //     if value , ok := data.(int); ok{
// //         Printf("i[%d] 类型为 int , 内容为%d\n" , index , value)
// //     }else if value , ok := data.(string); ok {
// //         Printf("i[%d] 类型为 string , 内容为%s\n" , index , value)
// //     }else if value , ok := data.(Student); ok{
// //         Printf("i[%d] 类型为 Student , 内容为name = %s , id = %d" , index , value.name, value.id)
// //     }
// // }

//     for index , data := range i{
//         switch val := data.(type) {
//         case int :
//             Printf(" i[%d] 的类型为 int , 内容为%d \n" , index , data)
//             break
//         case string :
//             Printf(" i[%d] 的类型为 string , 内容为%s \n", index , data)
//             break
//         case Student :
//             Printf(" i[%d] 的类型为 student , 内容为name = %s , id = %d \n", index , val.name,val.id)
//             break
//         default:
//             Println("未录入")
//         }
//     }

// // }

// //异常处理

// package main

// import (
// 	"fmt"
// 	"errors"
// )
//实现方法
// type errorString struct{
// 	text string
// }

// func New(text string) error {
// 	return &errorString{text}
// }

// func (e *errorString) Error() string {
// 	return e.text
// }

// func Errorf(format string , args ...interface{}) error {
// 	return errors.New(Sprintf(format , args...))
// }

// func main(){
// 	 err1 := errors.New( " a normal err1")
// 	 fmt.Println(err1)

// 	 var err2 error = fmt.Errorf("%s" , "a normal err2")
// 	  fmt.Println(err2)
// }
// func Divide(a ,b float64) (result float64 , err error){
// 	if b == 0 {
// 		result = 0.0
// 		err = errors.New("runtime error : divide by zero")
// 		return
// 	}
// 	result = a / b
// 	err = nil
// 	return

// }

// func main(){
// 	r , err := Divide(10.0 , 0.0)
// 	if err != nil{
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(r)
// 	}
// }

// func test01(){
// 	i := 1
// 	for ;  i < 101 ; {

// 		fmt.Printf(" i = %d \n" , i)
// 		i += 1
// 	}
// }

// func main (){
// 	test01()
// }

//字符串操作+转换

// strings.Contains,Replace,Repeat,Join,Index ,Spilt ,Fields
//           存在    替换     重复  加进 ，位置 ， 分裂 ， 删空格
// strconv.Append, Format, Parse
// 追尾 ，
// package main

// import (
// 	"fmt"
// 	"strings"
// 	"strconv"
// )

// func zifuchuan(){
// 	fmt.Println(strings.Contains("seafood" , "foo"))
// 	fmt.Println(strings.Contains("seafood" , "shit"))
// 	fmt.Println(strings.Contains("seafood" , " "))
// 	fmt.Println(strings.Contains("  " , " "))

// 	s := []string{"a" , "b" , "c", "d"}
// 	s2 := "chicken"
// 	fmt.Println(strings.Join(s , " and "))

// 	fmt.Println(strings.Index(s2 , "ck"))
// 	fmt.Println(strings.Index(s2 , "asd"))

// 	fmt.Println("fuck" + strings.Repeat("you and you" , 5))

// 	fmt.Println(strings.Replace(s2 , "ck" , "CK" ,5))

// 	fmt.Printf("%q\n" , strings.Split(s2 , "i"))

// 	fmt.Printf("%q\n" , strings.Split(" xyz  " , ""))
// 	fmt.Printf("%q\n" , strings.Split("" , "i"))

// 	fmt.Printf("fiels : %q" ,strings.Fields("   asd  asdw w wfsd vc sdgf e e    ") )

// 	str := make([]byte , 0 , 100)
// 	str = strconv.AppendInt(str , 4567 , 10)
// 	str = strconv.AppendBool(str , false)
// 	str = strconv.AppendQuote(str , "QWEQWRQRQWTQTSF")
// 	str = strconv.AppendQuoteRune(str, '单')
// 	fmt.Println(string(str))//类型从[]byte 转换为 string

// }

// func main(){
// 	zifuchuan()
// }

package main

import (
	"fmt"
	//"runtime"
	"time"
)

func forloop() {
	for i := 0; i < 6; i++ {
		fmt.Println("func: ", i)
		//time.Sleep(1 * time.Second)
	}

}

func main() {

	go forloop()
	for i := 0; i < 6; i++ {
		fmt.Println("main: ", i)
		time.Sleep(1 * time.Second)
	}

	//fmt.Println(" main .exit")

	// 	go func(s string){
	// 		for i := 0 ; i < 3; i++{
	// 			fmt.Println(s)
	// 		}
	// 	}("WORLD")

	// 	for i := 0 ; i < 3 ; i++{
	// 		//runtime.Gosched()
	// 		fmt.Println("HELLO")
	// 	}
}
