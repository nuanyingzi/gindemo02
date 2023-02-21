package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

// DoUpload 单文件上传demo
func (con UserController) DoUpload(c *gin.Context) {
	userName := c.PostForm("username")
	file, err := c.FormFile("face")
	dst := path.Join("./static/upload", file.Filename)
	if err == nil {
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  userName + "文件上传成功",
			"dst":  dst,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  userName + "文件上传失败",
		})
	}
}

func (con UserController) Index2(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit.html", gin.H{})
}

// DoEdit 不同名字的多文件上传demo
func (con UserController) DoEdit(c *gin.Context) {
	userName := c.PostForm("username")

	// 文件1
	file1, err1 := c.FormFile("face1")
	dst1 := path.Join("./static/upload", file1.Filename)
	if err1 == nil {
		c.SaveUploadedFile(file1, dst1)
	}

	// 文件2
	file2, err2 := c.FormFile("face2")
	dst2 := path.Join("./static/upload", file2.Filename)
	if err2 == nil {
		c.SaveUploadedFile(file2, dst2)
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"username": userName,
		"dst1":     dst1,
		"dst2":     dst2,
	})
}

// Index3 相同名字多文件上传demo
func (con UserController) Index3(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useredit2.html", gin.H{})
}

func (con UserController) DoEdit2(c *gin.Context) {
	userName := c.PostForm("username")
	form, _ := c.MultipartForm()
	files := form.File["file[]"]
	for _, file := range files {
		dst := path.Join("./static/upload", file.Filename)
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     userName + "文件上传成功",
	})
}
