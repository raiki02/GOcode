package main

import (
	"fmt"
	"os"
)



type student struct {
	id   int
	name string
	age  int
}

var students map[int]student = make(map[int]student, 50)

func Login() {
	fmt.Println("欢迎进入学生管理系统！")
	fmt.Print(
		`请选择您想要执行的操作：
	1： 查看所有学生信息
	2： 增加学生信息
	3： 修改学生信息
	4： 删除学生信息
	5： 退出管理系统`)

}

func Logout() {
	var sele string
	fmt.Println("您要退出该操作系统吗？\n ")
	fmt.Println("按y退出，按n取消 \n")
	fmt.Scanf("%s", &sele)
	if sele == "y" {
		Login()
	}

}

func showAllStu() {
	for _, j := range students {
		fmt.Printf("id： %d , 姓名： %s , 年龄： %d", j.id, j.name, j.age)
	}
}

func addStu() {
	var (
		stuid   int
		stuname string
		stuage  int
	)

	// fmt.Println("请输入学生id：")
	// fmt.Scanf("%d", &stuid)
	// fmt.Println("请输入学生姓名：")
	// fmt.Scanf("%s", &stuname)
	// fmt.Println("请输入学生年龄：")
	// fmt.Scanf("%d", &stuage)
	// fmt.Println("请输入")
	// fmt.Scanf("%d ,%s ,%d", &stuid, &stuname, &stuage)

	// students[stuid] = student{stuid, stuname, stuage}
	fmt.Print("请输入:")
	fmt.Scanln(&stuid)
	fmt.Scanln(&stuname)
	fmt.Scanln(&stuage)
}

func changeStu() {

}

func deleteStu() {

}

func main() {
	Login()
	for {
		var choice int = 0
		fmt.Println("\n 输入对应数字进入操作系统")
		fmt.Scanf("%d", &choice)
		fmt.Printf("您选择的是: %d \n", choice)

		switch choice {
		case 1:
			fmt.Print("您已进入学生信息查询系统！\n")
			showAllStu()
			Logout()
		

		case 2:
			fmt.Print("您已进入学生信息添加系统！\n")
			addStu()
			Logout()

		case 3:
			fmt.Print("您已进入学生信息修改系统！\n")
			changeStu()
			Logout()
		

		case 4:
			fmt.Print("您已进入学生信息删除系统！\n")
			deleteStu()
			Logout()
		

		case 5:
			fmt.Print("您已成功退出！\n")
			os.Exit(1)
		default:
			fmt.Print("您的输入有误，请重新输入! \n")
		

		}
	}

}
