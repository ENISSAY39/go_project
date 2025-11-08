package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func RequireAuth(c *gin.Context) {
	fmt.Println("Hello from require auth middleware")

	
	c.Next()

}