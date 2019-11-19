package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	webEng := gin.Default()


	webEng.GET("/get", HandleGeneric)
	webEng.POST("/post", HandleGeneric)
	webEng.HEAD("/head",HandleGeneric)

	webEng.Any("/",HandleAny)

	webEng.Handle("ECHO","/", HandleAny)

	webEng.Run(":8080")
}

func HandleGeneric(c *gin.Context){
	c.String(http.StatusOK,"generic method: %s", c.Request.Method)
}

func HandleAny(c *gin.Context){
	c.String(http.StatusOK,"method : %s", c.Request.Method)
}
