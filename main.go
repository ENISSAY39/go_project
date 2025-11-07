package main

import (
	"github.com/ENISSAY39/go_project/controllers"
	"github.com/ENISSAY39/go_project/intializers"
	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.Run() // listens on 0.0.0.0:8080 by default
}
