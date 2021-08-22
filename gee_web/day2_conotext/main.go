package main

import (
	"gee"
	"net/http"
)

/*
 $Curl http://localhost:8080
 $Curl "http://localhost:8080/hello?name=leo"
 $Curl "http://localhost:8080/login" -X POST -d "username=leo&password=1234"
 $Curl "http://localhost:8080/hello12
*/
func main() {
	g := gee.New()
	g.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	g.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	g.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	g.Run(":8080")
}
