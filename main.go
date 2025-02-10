package main

import (
	"backend/config"
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/folders", controllers.GetAllData)
	r.GET("/selected", controllers.GetSelectedData)
	r.Run(":1234")
}
