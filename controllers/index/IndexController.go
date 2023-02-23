package index

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {

	session := sessions.Default(c)
	// 设置session的过期时间
	session.Options(sessions.Options{
		MaxAge: 3600 * 6, // 单位：秒
	})
	session.Set("username", "里斯")
	session.Save()

	c.SetCookie("username", "张三", 3600, "/", "localhost", false, true)
	c.SetCookie("hobby", "吃饭 睡觉 打豆豆", 10, "/", "localhost", false, true)
	c.String(http.StatusOK, "前台首页")
}

func (con IndexController) News(c *gin.Context) {

	session := sessions.Default(c)
	sessionUsername := session.Get("username")

	userName, err := c.Cookie("username")
	if err != nil {
		c.String(http.StatusOK, "获取cookie失败")
	}
	hobby, err := c.Cookie("hobby")
	if err != nil {
		c.String(http.StatusOK, "获取cookie失败")
	}
	c.JSON(http.StatusOK, gin.H{
		"success":          true,
		"username":         userName,
		"hobby":            hobby,
		"session_username": sessionUsername,
	})
}

func (con IndexController) DeleteCookie(c *gin.Context) {
	// 删除Cookie
	c.SetCookie("username", "张三", -1, "/", "localhost", false, true)
	c.String(http.StatusOK, "删除Cookie成功")
}
