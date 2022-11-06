package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func main() {
	//db.ConnectDB()
	r := gin.Default()

	r.POST("/create", controllers.Creates)
	r.POST("/delete", controllers.Deletes)
	r.POST("/getall", controllers.GetAll)
	//r.POST("/update", controllers.Update)

	r.Run(":8080")
}
