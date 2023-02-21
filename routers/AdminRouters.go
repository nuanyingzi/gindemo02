package routers

import (
	"gindemo02/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminRoutersInit(r *gin.Engine) {
	adminRouter := r.Group("admin", middlewares.InitMiddleware)
	{
		adminRouter.GET("/", func(c *gin.Context) {
			/*userName, _ := c.Get("username")
			// 类型断言
			v, ok := userName.(string)
			if ok {
				c.String(http.StatusOK, "后台首页-"+v)
			} else {
				c.String(http.StatusOK, "获取用户失败")
			}*/
			c.String(http.StatusOK, "后台首页")

		})
		adminRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "后台新闻页")
		})

		adminUserRouter := adminRouter.Group("user")
		{
			adminUserRouter.GET("index", func(c *gin.Context) {
				c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
			})
		}
	}
}
