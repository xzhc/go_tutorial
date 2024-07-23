package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//路由无参数
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Who are you")
	})
	//路由有参数
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	//路由获取query参数
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	//获取POST参数
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") //可以设置默认值
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	//Query和POST混合参数
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000")

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	//Map参数
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.QueryMap("names")
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	//Redirect(重定向)
	//HTTP重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
	//路由重定向
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	//分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	//group:v1
	v1 := r.Group("v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	//group:v2
	v2 := r.Group("v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	//上传文件

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	r.MaxMultipartMemory = 8 << 20 //8Mib
	r.POST("/upload1", func(c *gin.Context) {
		//单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		dst := "./" + file.Filename
		//上传文件至完整的文件路径
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	r.POST("/upload2", func(c *gin.Context) {
		//多文件
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			//上传文件至指定目录
			dst := "./" + file.Filename
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	//HTML模板(Template)
	type student struct {
		Name string
		Age  int8
	}
	r.LoadHTMLGlob("templates/*")
	stu1 := &student{Name: "xzh", Age: 25}
	stu2 := &student{Name: "xsy", Age: 15}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	//使用中间件
	// 新建一个没有任何默认中间件的路由
	//r = gin.New()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())
	//你可以为每个路由添加任意数量的中间件(作用于单个路由)
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	//作用于某个组
	authorized := r.Group("/")
	authorized.Use(Authorized())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	r.Run(":9999") // listen and serve on 0.0.0.0:8080(default)
}

// 自定义中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//给Context实例设置一个值
		c.Set("xzh", "1111")
		//请求前
		c.Next()
		//请求后
		latency := time.Since(t)
		log.Print(latency)
	}
}
