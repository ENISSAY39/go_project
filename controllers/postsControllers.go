package controllers

import (
	"github.com/ENISSAY39/go_project/intializers"
	"github.com/ENISSAY39/go_project/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data off req body

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := intializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return post as response

	c.JSON(200, gin.H{
		"post": "pong ",
	})
}
