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

func mutipleQuery() {
	sqlStr := "select id , name , age from user where id > 0;"
	rows, err := db.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	defer db.Close() //...

	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			panic(err)
			return
		}

		fmt.Printf("u = %v \n", u)
	}
}

func prepareQuery() {
	sqlStr := `select id,name ,age from user where id > ?;`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(0)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			panic(err)
		}
		fmt.Println("u = ", u)
	}
	defer stmt.Close()
}

func insert() {
	sqlStr := `insert into user(name, age) values(?, ?);`
	ret, err := db.Exec(sqlStr, "Diana", 15)
	if err != nil {
		panic(err)
		return
	}

	id, err := ret.LastInsertId()
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("id = ", id)
}

func update(id int) {
	sqlStr := `update user set age=114514 where id > ?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		panic(err)
		return
	}
	row, err := ret.RowsAffected()
	if err != nil {
		panic(nil)
		return
	}

	fmt.Println("affected row =", row)
}

func delete(id int) {
	sqlStr := `delete from user where id=?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		panic(err)
		return
	}
	row, err := ret.RowsAffected()
	if err != nil {
		panic(nil)
		return
	}

	fmt.Println("affected row =", row)
}

func transanction() {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Begin err: ", err)
		return
	}
	sqlStr1 := `update user set age=age-3 where id=2;`
	sqlStr2 := `update user set age=age+3 where id=4;`

	_, err = tx.Exec(sqlStr1)
	if err != nil {
		fmt.Println("exec sql1 err:", err)
		tx.Rollback() //...
		return
	}
	_, err = tx.Exec(sqlStr2)
	if err != nil {
		fmt.Println("exec sql2 err:", err)
		tx.Rollback() //...
		return
	}

	fmt.Println("sql1 , sql2 exec succeeded")

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit err :", err)
	}

	fmt.Println("sql exec , commit suceeded")
}

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("database connected")

	//update(8)
	//insert()
	//mutipleQuery()
	//delete(6)
	//prepareQuery()
	transanction()
}
