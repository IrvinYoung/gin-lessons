package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	//创建web实例
	webEng := gin.Default()

	//创建一个默认的路由
	webEng.GET("/", Home)

	//启动web服务
	webEng.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}

//Home
func Home(c *gin.Context){
	c.String(http.StatusOK, "gin running !")
}