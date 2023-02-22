package admin

import (
	"gindemo02/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strconv"
)

type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/useradd.html", gin.H{})
}

// DoUpload 单文件上传demo
func (con UserController) DoUpload(c *gin.Context) {
	//userName := c.PostForm("username")
	// 1 获取文件
	file, err := c.FormFile("face")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "文件上传失败",
			"err": err.Error(),
		})
	}
	// 2 获取后缀名
	extSuffix := path.Ext(file.Filename)
	// 3 判断后缀是否正确
	allowExtSuffix := map[string]bool{
		".jpg":  true,
		".png":  true,
		".jpeg": true,
		".gif":  true,
	}
	if _, ok := allowExtSuffix[extSuffix]; !ok {
		c.String(http.StatusBadRequest, "上传文件类型不合法")
	}
	// 4 创建图片保存目录
	day := models.GetDay() // 获取当天日期
	dir := "./static/upload/" + day
	err = os.MkdirAll(dir, 0666)
	if err != nil {
		c.String(http.StatusBadRequest, "MkdirAll失败")
		return
	}
	// 创建文件名
	fileName := strconv.FormatInt(models.GetUnix(), 10) + extSuffix
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
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
