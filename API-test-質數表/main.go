package main

import (
	"test/api"

	"github.com/gin-gonic/gin"
)

func init() {
	//server.SetupDatabaseConnection()
}

func main() {

	router := gin.Default()
	router.POST("/isPrimesToApi", api.CheckAndInsertPrimeValues) // 設置路由處理函式
	router.Run(":8080")                                          //將計算結果回傳給API伺服器
}
