package main

import (
	"github.com/ENISSAY39/go_project/intializers"
	"github.com/ENISSAY39/go_project/models"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()

}

func main() {
	intializers.DB.AutoMigrate(&models.Post{})

}
