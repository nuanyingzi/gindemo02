package main

import (
	"fmt"
	"gindemo02/models"
	"gindemo02/routers"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

type New struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// Println 输出两个字符串拼接的函数
func Println(str1, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + str2
}

// 中间件1
func initMiddlewareOne(c *gin.Context) {
	fmt.Println("start-中间件-initMiddlewareOne")
	// 调用该请求的剩余处理程序
	c.Next()
	fmt.Println("end-中间件-initMiddlewareOne")
}

// 中间件2
func initMiddlewareTwo(c *gin.Context) {
	fmt.Println("start-中间件-initMiddlewareTwo")
	// 调用该请求的剩余处理程序
	c.Next()
	fmt.Println("end-中间件-initMiddlewareTwo")
}

func main() {
	r := gin.Default()

	// 注册模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Println":    Println,
	})
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态web目录
	r.Static("/static", "./static")

	// 全局中间件
	//r.Use(initMiddlewareOne, initMiddlewareTwo)

	r.GET("/", func(context *gin.Context) {
		fmt.Println("我是首页")
		context.String(http.StatusOK, "首页")
	})

	/*// GET请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})*/

	/*// POST演示
	r.GET("/index/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})
	r.POST("/userLogin", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "18")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	// GET POST获取的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		err := c.ShouldBind(&user)
		if err == nil {
			fmt.Printf("%#v", user)
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	r.POST("/userLogin2", func(c *gin.Context) {
		user2 := &UserInfo{}
		err := c.ShouldBind(&user2)
		if err == nil {
			c.JSON(http.StatusOK, user2)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		}
	})

	// 动态路由传值
	r.GET("/user/:uid", func(context *gin.Context) {
		uid := context.Param("uid")
		context.String(http.StatusOK, "UserId = %s", uid)
	})*/

	// 后台路由
	routers.AdminRoutersInit(r)
	// Api路由
	routers.ApiRoutersInit(r)
	// 前台路由
	routers.IndexRoutersInit(r)

	r.Run()
}
