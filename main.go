package main

import (
	"github.com/ENISSAY39/go_project/controllers"
	"github.com/ENISSAY39/go_project/initializers"
	"github.com/ENISSAY39/go_project/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDataBase()

}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id", controllers.PostsDelete)
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	router.Run() // listens on 0.0.0.0:8080 by default
}
