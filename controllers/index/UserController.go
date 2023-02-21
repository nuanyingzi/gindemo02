package index

import (
	"fmt"
	"gindemo02/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	BaseController
}

func (con UserController) Index(c *gin.Context) {
	date := models.UnixToTime(1672712094)
	fmt.Println(date)
	c.HTML(http.StatusOK, "default/useradd.html", gin.H{
		"title": "用户首页",
		"time":  1672712094,
	})
}

func (con UserController) Add(c *gin.Context) {
	c.String(http.StatusOK, "前台首页添加用户")
}

func (con UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "前台首页修改用户")
}
