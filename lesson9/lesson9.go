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

	//account 路由组
	account := webEng.Group("account")
	{
		//path为相对路径
		account.POST("/login", Login)
		//path可以不以"/"开始
		account.POST("logout", Logout)
		//path可以包含子路径
		account.GET("profile/info", AccountInfo)
	}

	//article路由组
	article := webEng.Group("/article")
	{
		//嵌套路由组
		//按照版本划分路由组
		v1 := article.Group("v1")
		//path同样支持路由参数
		v1.GET("/:id", GetArticleById)

		v2 := article.Group("v2")
		v2.GET("/:tag",GetArticleByTag)
	}
	webEng.Run()
}

func Login(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func Logout(c *gin.Context) {
	c.String(http.StatusOK, "logout")
}

func AccountInfo(c *gin.Context) {
	c.String(http.StatusOK, "account info")
}

func GetArticleById(c *gin.Context){
	id := c.Param("id")
	c.String(http.StatusOK,"content of article: %s",id)
}

func GetArticleByTag(c *gin.Context){
	tag := c.Param("tag")
	c.String(http.StatusOK,"content of article: %s",tag)
}