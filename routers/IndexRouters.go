package routers

import (
	"gindemo02/controllers/index"
	"github.com/gin-gonic/gin"
)

func IndexRoutersInit(r *gin.Engine) {

	indexRouter := r.Group("index")
	{
		indexRouter.GET("/", index.IndexController{}.Index)
		indexRouter.GET("/news", index.IndexController{}.News)
		indexRouter.GET("/deleteCookie", index.IndexController{}.DeleteCookie)
	}

	userRouter := r.Group("user")
	{
		userRouter.GET("/", index.UserController{}.Index)
		userRouter.GET("/add", index.UserController{}.Add)
		userRouter.GET("/edit", index.UserController{}.Edit)
	}
}
