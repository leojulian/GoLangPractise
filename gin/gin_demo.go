package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 无参数
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// 匹配 /user/leo
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 匹配users?name=xxx&role=xxx，role可选
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "Developer")
		c.String(http.StatusOK, "%s is %s", name, role)
	})

	// Post
	// $curl http://localhost:8080/form -X POST -d "username=leo&password=1234"
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// GET 和 POST 混合
	// curl "http://localhost:8080/posts?id=1224&page=12"  -X POST -d "username=leo&password=1234"
	//r.POST("/posts", func(c *gin.Context) {
	//	id := c.Query("id")
	//	page := c.DefaultQuery("page", "0")
	//	username := c.PostForm("username")
	//	password := c.DefaultPostForm("password", "000000")
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"id":       id,
	//		"page":     page,
	//		"username": username,
	//		"password": password,
	//	})
	//})

	//Map参数(字典参数)
	// $curl -g "http://localhost:8080/posts?ids[jack]=001&ids[leo]=002" -X POST -d "names[a]=sam&names[b]=micheal"
	r.POST("/posts", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 重定向(Redirect)
	// curl -i http://localhost:8080/redirect
	// curl http://localhost:8080/goindex
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 分组路由(Grouping Routes)
	// curl http://localhost:8080/v1/posts
	// curl http://localhost:8080/v2/posts
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	//group v1
	v1 := r.Group("v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	//group v2
	v2 := r.Group("v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	// 上传文件
	// 单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		//c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded", file.Filename)
	})
	//多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			//c.SaveUploadedFile(file, dst)
		}

		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	//HTML模板(Template)
	//$ curl http://localhost:9999/arr
	type student struct {
		Name string
		Age  int8
	}

	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "leo", Age: 30}
	stu2 := &student{Name: "Michael", Age: 4}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	//中间件(Middleware)
	//Logger := func() gin.HandlerFunc {
	//	return func(c *gin.Context) {
	//		t := time.Now()
	//		// 给Context实例设置一个值
	//		c.Set("geektutu", "1111")
	//		// 请求前
	//		c.Next()
	//		// 请求后
	//		latency := time.Since(t)
	//		log.Print(latency)
	//	}
	//}
	//
	//// 作用于全局
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//
	//// 作用于单个路由
	//r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//
	//// 作用于某个组
	//authorized := r.Group("/")
	//authorized.Use(AuthRequired())
	//{
	//	authorized.POST("/login", loginEndpoint)
	//	authorized.POST("/submit", submitEndpoint)
	//}

	go func() {
		_ = r.Run()
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	//Deal with some tasks when receive ctl+c
	fmt.Println("Server is closing...")
}
