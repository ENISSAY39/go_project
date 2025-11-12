package main

import (
	"github.com/ENISSAY39/go_project/initializers"
	"github.com/ENISSAY39/go_project/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})

}
