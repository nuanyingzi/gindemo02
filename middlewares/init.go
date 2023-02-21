package middlewares

import (
	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	// 判断用户登录
	//fmt.Println("当前时间：", time.Now())
	//fmt.Println("当前请求路径：", c.Request.URL)

	// 协程下的context需要copy
	//cCopy := c.Copy()

	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("当前请求的URL：", c.Request.URL)
	//}()

	//c.Set("username", "Tom")
}
