package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

type person struct {
	User_id  int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func init() {
	dsn := "root:lost9725ost@tcp(127.0.0.1:3306)/go_mysql_new"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("err open database: ", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("err ping database: ", err)
		return
	}
	fmt.Println("Connected to the database !")
}

func insert(user_id, username, email string) {

	/*var (
		user_id  string
		username string
		email    string
	)
	fmt.Scanln(&user_id, &username, &email)
	insert(user_id, username, email)*/

	sql := `insert into person (user_id , username, email) values(?,?,?)`
	r, err := db.Exec(sql, user_id, username, email)
	if err != nil {
		fmt.Println("err insert database: ", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("err insert database: ", err)
		return
	}
	fmt.Println("inserted id: ", id)
}

func showrow(id int) {
	var persons []person
	sql := `select user_id, username , email from person where user_id = ?`
	err := db.Select(&persons, sql, id)
	if err != nil {
		fmt.Println("err select database: ", err)
		return
	}
	fmt.Println("select success: ", persons)
}

func update(key, val string, id int) {
	sql := `update person set ? = ? where user_id=?`
	ret, err := db.Exec(sql, key, val, id)
	if err != nil {
		fmt.Println("err update database: ", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("err update database: ", err)
		return
	}
	fmt.Println("update success: ", num)
}

func delete_row(id int) {
	sql := `delete from person where user_id = ?`
	res, err := db.Exec(sql, id)
	if err != nil {
		fmt.Println("err delete database: ", err)
		return
	}
	num, err := res.RowsAffected()
	if err != nil {
		fmt.Println("err delete database: ", err)
		return
	}
	fmt.Println("delete success: ", num)
}

func sqlInjectiion() {

}

func main() {
	sql := "update person set username = ? where user_id = ? "
	//drop table person
	inject := "or 1=1; #"
	//inject := " `drop table person; --"
	res, err := db.Exec(sql, "delete_table", inject)
	if err != nil {
		fmt.Println("err insert database: ", err)
		return

	}
	fmt.Println("id inserted")

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("err insert database: ", err)
		return
	}
	fmt.Println("inserted id: ", id)
}
