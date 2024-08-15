package main

import (
	"Go_rest_now_android/database"
	"Go_rest_now_android/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	database.Init()
	router := gin.Default()

	router.POST("/reg", handlers.Registration)

	router.Run(":8081")

}
