package routers

import (
	"gindemo02/controllers/admin"
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
			// 用户页
			adminUserRouter.GET("index", admin.UserController{}.Index)
			adminUserRouter.GET("index2", admin.UserController{}.Index2)
			adminUserRouter.GET("index3", admin.UserController{}.Index3)

			// 单文件上传
			adminUserRouter.POST("doUpload", admin.UserController{}.DoUpload)

			// 不同名字多文件上传
			adminUserRouter.POST("doEdit", admin.UserController{}.DoEdit)

			// 相同名字多文件上传
			adminUserRouter.POST("doEdit2", admin.UserController{}.DoEdit2)

		}
	}
}
