package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("html/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})
	r.POST("/form", func(c *gin.Context) {
		user := c.PostForm("user")
		c.String(200, user)
		password := c.PostForm("password")
		c.String(200, password)
	})
	r.Run(":8080")
}
