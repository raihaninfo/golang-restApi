package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// DB name:	 		golang-restapi
// DB username: 	root
// DB password: 	mysql
// DB port: 		localhost:3306

var urlDSN = "root:mysql@tcp(localhost:3306)/golang-restapi?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("connection failed :(")
	}
	Database.AutoMigrate(&Employee{})
}
