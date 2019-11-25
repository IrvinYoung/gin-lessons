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

	//获取文章列表
	webEng.GET("/articles", Articles)

	webEng.Run(":8080")
}

//查询文章列表，假设请求URL：host/articles?category=news&pageSize=10&pageNum=1
func Articles(c *gin.Context) {
	//如果参数不存在，或者为空，则返回空字符串
	pSize := c.Query("pageSize")
	//如果参数不存在，设置参数的默认值
	pNum := c.DefaultQuery("pageNum","1")

	var has bool
	//获取参数的同时，检测参数是否存在
	if pSize,has = c.GetQuery("pageSize");!has{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error": "lost page size",
		})
		return
	}

	//获取query数组
	// /articles?pageNum=1&pageSize=100&category=news&category=tech&category=music
	categories := c.QueryArray("category")
	//获取query数组,同时检测是否存在
	if categories,has = c.GetQueryArray("category");!has{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error": "lost category",
		})
		return
	}

	//获取query map数据
	// /articles?pageNum=1&pageSize=100&category=news&category=tech&category=music&date['start']=2019-07-06&date['end']=2019-11-10
	date := c.QueryMap("date")
	//获取query map数据,并检测是否设置
	if date,has = c.GetQueryMap("date");!has{
		c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
			"error": "lost date",
		})
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"pageSize":pSize,
		"pageNum":pNum,
		"category":categories,
		"categoryCount":len(categories),
		"date":date,
	})
}
