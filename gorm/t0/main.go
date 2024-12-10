package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var user User

var db *gorm.DB

type User struct {
	*gorm.Model
	Name         string         `gorm:"column:name"`
	Email        *string        `gorm:"column:email"`
	Age          uint8          `gorm:"column:age"`
	Birthday     *time.Time     `gorm:"column:birthday"`
	MemberNumber sql.NullString `gorm:"column:member_number"`
	ActivatedAt  sql.NullTime   `gorm:"column:activated_at"`
}

type testTable struct {
	id   int    `gorm:"column:id`
	name string `gorm:"column:name"`
}

func main() {
	dsn := "root:123456@tcp(172.17.0.3:9999)/UnionTest"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	t := testTable{id: 1, name: "test_name"}
	res := db.Create(&t)
	fmt.Println("res: ", res)
	fmt.Println("t.Error: ", res.Error)

}

func deletes() {
	res := db.Where("name = ?", "test_name5").Delete(&User{})
	fmt.Printf("res: %v \n", res)
}

func updates() {
	db.First(&user)
	user.Name = "test_name11"
	db.Save(&user)

	db.Model(&user).Update("name", "test_name111")
	fmt.Println("user: ", user)

	db.Model(&user).Updates(User{Name: "test_name1111", Age: 22})
	fmt.Println("user: ", user)
}

func finds() {
	var user User
	var users []User
	db.First(&user)
	fmt.Printf("user: %v \n", user)
	db.Last(&user)
	fmt.Printf("user: %v \n", user)

	db.Where("name = ?", "test_name3").First(&user)
	fmt.Printf("user: %v \n", user)

	db.Select("name, email").Find(&users)
	fmt.Printf("users: %v \n", users)

	var user2 User
	db.Limit(2).Offset(1).Find(&user2).Limit(-1).Offset(-1).Find(&user)
	fmt.Printf("user: %v \n", user)
	fmt.Println("user2: ", user2)

	var count int64
	res := db.Model(&user).Count(&count)
	fmt.Printf("count: %v \n", res)
}

func create() {
	now := time.Now()
	email := "test5@gmail.com"
	user := User{Name: "test_name5", Email: &email, Age: 19, Birthday: &now}
	email3 := "test3@gmail.com"
	email4 := "test4@gmail.com"
	users := []User{
		{Name: "test_name3", Age: 20, Email: &email3, Birthday: &now},
		{Name: "test_name4", Age: 21, Email: &email4, Birthday: &now},
	}
	res := db.Create(&user)
	res = db.Create(&users)
	fmt.Printf("user: %v \n", user.ID)
	fmt.Printf("res: %v \n", res)
	fmt.Printf("res.AffectedRows: %v \n", res.RowsAffected)
	fmt.Printf("res.Error: %v \n", res.Error)
}
