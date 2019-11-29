package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	webEng *gin.Engine
)

func main() {
	webEng = gin.Default()
	webEng.POST("/form", HandleForm)
	webEng.Run()
}

func HandleForm(c *gin.Context) {
	//获取表单值，不存在时，使用默认值
	name := c.DefaultPostForm("name", "guest")

	//获取表单值，如果不存在或者内容为空，都将得到空字符串
	mobile := c.PostForm("mobile")

	var (
		age string
		has bool
	)
	//获取表单值，同时检测是否有设置
	if age, has = c.GetPostForm("age"); !has {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "lost age",
		})
		return
	}

	//获取表单数组
	hobbies := c.PostFormArray("hobbies")
	//获取表单数组，同时检测是否设置
	if hobbies, has = c.GetPostFormArray("hobbies"); !has {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "lost hobbies",
		})
		return
	}

	//获取表单map
	marks := c.PostFormMap("marks")
	//获取表单map，同时检测是否设置
	if marks, has = c.GetPostFormMap("marks"); !has {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "lost marks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    name,
		"mobile":  mobile,
		"age":     age,
		"hobbies": hobbies,
		"marks":   marks,
	})
}
