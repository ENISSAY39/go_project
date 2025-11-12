package controllers

import (
	"github.com/ENISSAY39/go_project/initializers"
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
	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return post as response

	c.JSON(200, gin.H{
		"post": "pong ",
	})
}

func PostsIndex(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// Return posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get post
	var post models.Post
	initializers.DB.First(&post, id)
	// Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)
	// Find post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Delete post
	initializers.DB.Delete(&models.Post{}, id)
	// Return status
	c.Status(200)
}
