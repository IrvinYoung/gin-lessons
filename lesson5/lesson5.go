package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	webEng *gin.Engine
)

func main(){
	webEng = gin.Default()

	//处理 带必选参数的路由
	webEng.GET("/article/:searchBy", NecessaryParams)
	//处理 带可选参数的路由
	webEng.GET("/article/:searchBy/*condition", OptionalParams)
	//处理 其他没有定义的路由
	webEng.NoRoute(Others)

	webEng.Run(":8080")
}

/*
	处理带必选参数的路由：
	/article		no
	/article/		no
	/article/id		yes
	/article/id/	yes	(redirecting request 301: /article/id/ --> /article/id)
*/
func NecessaryParams(c *gin.Context){
	searchBy := c.Param("searchBy")
	c.String(http.StatusOK, "Necessary=> search article by %s", searchBy)
}

/*
	处理带可选参数的路由：
	/article/id				no
	/article/id/			yes
	/article/id/2159 		yes
	/article/id/2159/other 	yes
*/
func OptionalParams(c *gin.Context){
	condition := c.Param("condition")
	searchBy := c.Param("searchBy")
	c.String(http.StatusOK, "Optional=> search article by %s = %s",searchBy,condition)
}

// 处理其他没有定义的路由，自定义404
func Others(c *gin.Context){
	c.String(http.StatusNotFound, "No handler for => %s", c.Request.RequestURI)
}