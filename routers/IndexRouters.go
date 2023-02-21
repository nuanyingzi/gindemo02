package routers

import (
	"gindemo02/controllers/index"
	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {
	indexRouter := r.Group("index")
	{
		indexRouter.GET("/user", index.UserController{}.Index)
		indexRouter.GET("/user/add", index.UserController{}.Add)
		indexRouter.GET("/user/edit", index.UserController{}.Edit)
	}
}
