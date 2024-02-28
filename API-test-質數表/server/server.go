package server

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DateBaseConnection *gorm.DB //調用一次後connect連接給它
)

func SetupDatabaseConnection() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/test2"), &gorm.Config{})
	if err != nil {
		fmt.Println("無法連接數據庫", err)
	}

	DateBaseConnection = database
	sqlDB, _ := DateBaseConnection.DB()
	//fmt.Println("1")
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
}
