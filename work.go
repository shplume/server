package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type account struct {
	name     string
	password string
}

func init() {
	db, err = sql.Open("mysql", "root:secret@tcp(localhost:3000)/todos")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos.test(id INT UNSIGNED AUTO_INCREMENT, name varchar(50), password varchar(50),PRIMARY KEY (id))")
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
		password := c.PostForm("password")

		rows, err := db.Query("SELECT name,password FROM todos.test")
		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			var a account
			err = rows.Scan(&a.name, &a.password)
			if err != nil {
				log.Fatalln(err)
			}
			if name == a.name && password == a.password {
				c.String(200, "login successfully")
				return
			}
		}
		c.String(404, "account or password incorrect")
	})

	r.POST("/web", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")

		rows, err := db.Query("SELECT name FROM todos.test")
		if err != nil {
			log.Fatalln(err)
		}

		for rows.Next() {
			var s string

			err = rows.Scan(&s)
			if err != nil {
				log.Fatalln(err)
			}

			if s == name {
				c.String(404, "User name has been registereds")
				return
			}
		}
		rows.Close()

		rs, err := db.Exec("INSERT INTO test(name, password) VALUES (?,?)", name, password)
		if err != nil {
			log.Fatal(err)
		}
		c.String(200, "login successfully")

		rowCount, err := rs.RowsAffected()
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("inserted %d rows", rowCount)
	})

	r.Run(":8080")

	defer db.Close()
}
