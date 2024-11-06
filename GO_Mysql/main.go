package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func InitDB() (err error) {
	dsn := "root:lost9725ost@tcp(127.0.0.1:3306)/GO_Mysql"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("init db failed: ", dsn, err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping db faileed", err)
		return
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	return err
}

func query() {
	var u1 user
	sqlStr := `select id, name, age from user where id=1;`
	rowObj := db.QueryRow(sqlStr)
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	fmt.Println(u1)
}

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("yes")

}
