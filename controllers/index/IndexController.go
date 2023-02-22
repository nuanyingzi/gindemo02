package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true)
	c.SetCookie("hobby", "吃饭 睡觉 打豆豆", 10, "/", "localhost", false, true)
	c.String(http.StatusOK, "前台首页")
}

func (con IndexController) News(c *gin.Context) {
	userName, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusOK, "获取cookie失败")
	}
	hobby, err := c.Cookie("hobby")
	if err != nil {
		c.String(http.StatusOK, "获取cookie失败")
	}
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": userName,
		"hobby":    hobby,
	})
}

func (con IndexController) DeleteCookie(c *gin.Context) {
	// 删除Cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "删除Cookie成功")
}
