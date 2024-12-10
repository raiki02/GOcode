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

func sqlInject(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Println("sql = ", sqlStr)
	var u user
	//sql =  select id, name, age from user where name='xxx' union select * from user #'
	ret, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}
	defer ret.Close()
	for ret.Next() {
		err := ret.Scan(&u.id, &u.name, &u.age)

		if err != nil {
			panic(err)
		}
		fmt.Println(u)
	}
}

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected")
	//sqlInject("Yuan Liu")
	//sqlInject("xxx' or 1=1 #")
	sqlInject("xxx' union select * from user #")
}
