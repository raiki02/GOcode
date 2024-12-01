package main

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:lost9725ost@tcp(127.0.0.1:3306)/go_mysql" //...
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func QueryRow() {
	sqlStr := "select id, name , age from user where id=1;"
	var u user
	err := db.Get(&u, sqlStr) //...
	if err != nil {
		panic(err)
	}
	fmt.Println(u)

}

func QueryRows() {
	sqlStr := `select id ,name, age from user;`
	var userlist = []user{}
	err := db.Select(&userlist, sqlStr)
	if err != nil {
		panic(err)
	}

	//fmt.Println(userlist)
	for _, user := range userlist {
		fmt.Println(user)
	}
}

func main() {
	err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected! ")

	QueryRow()
	fmt.Println("===")
	QueryRows()
}
