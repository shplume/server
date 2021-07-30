package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3000)/todos")
	if err != nil {
		log.Fatal("create database failed", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS todos.test(
		id INT UNSIGNED AUTO_INCREMENT, 
		name varchar(50) unique, 
		password varchar(50),
		PRIMARY KEY (id)
	)`)

	if err != nil {
		log.Fatal("create table failed", err)
	}

	r := gin.Default()

	r.LoadHTMLGlob("html/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", "")
	})
	r.GET("/register", func(c *gin.Context) {
		c.HTML(200, "account.html", "")
	})

	r.POST("/login", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")

		var p string
		err := db.QueryRow("SELECT password FROM test WHERE name = ?", name).Scan(&p)
		if err != nil {
			log.Println("query failed,err:", err)
		}

		if password == p {
			c.String(200, "login successfully")
			return
		}

		c.String(400, "account or password incorrect")
	})

	r.POST("/register/result", func(c *gin.Context) {
		name := c.PostForm("name")
		password := c.PostForm("password")

		var s string

		err := db.QueryRow(`SELECT name FROM test WHERE name = ?`, name).Scan(&s)
		if err != nil {
			log.Println("query failed,err:", err)
		}

		if s == name {
			c.String(400, "User name has been registereds")
			return
		}

		rs, err := db.Exec("INSERT INTO test(name, password) VALUES (?,?)", name, password)
		if err != nil {
			log.Println("insert failed,err:", err)
			return
		}

		rowCount, err := rs.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("inserted %d rows", rowCount)

		c.String(200, "register successfully")
	})

	r.Run(":8080")
	db.Close()
}
