package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "awsdb.clg9pxaxypoz.us-east-2.rds.amazonaws.com"
	database = "gorm"
	user     = "awsuser"
	password = "quddnr!2"
)

type User struct {
	gorm.Model
	ID   uint `gorm:"primary_key"`
	Name string
}

func main() {
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, database)
	mysqlDB, err := sql.Open("mysql", connectionString)
	defer mysqlDB.Close()

	gormDb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: mysqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	gormDb.AutoMigrate(&User{})
	gormDb.Create(&User{Name: "byungwook2"})
	fmt.Println("eof")
}
