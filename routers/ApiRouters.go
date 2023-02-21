package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouter := r.Group("api")
	{
		apiRouter.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "接口首页")
		})
		apiRouter.GET("/news", func(c *gin.Context) {
			c.String(http.StatusOK, "接口新闻页")
		})
	}
}
