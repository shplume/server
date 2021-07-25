package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", "root:plume@tcp(localhost)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test.hello(name varchar(50), password varchar(50))")
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "account.html", "")
	})

	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, name)
		password := c.PostForm("password")
		c.String(200, password)
	})
	r.POST("/web", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(200, name)
		password := c.PostForm("password")
		c.String(200, password)
	})

	r.Run(":8080")
}
