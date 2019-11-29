package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"net/textproto"
	"strings"
	"time"
)

var (
	webEng *gin.Engine
)

func main() {
	webEng = gin.Default()
	webEng.POST("/setAvatar", HandleUpload)
	webEng.POST("/multiple",HandleMultiple)
	webEng.Run()
}

func HandleMultiple(c *gin.Context){
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	pics := form.File["pics"]
	for _, pic := range pics {
		//todo: 检测文件大小
		//if fAvatar.Size > xxxx

		//获取存储路径
		name, err := localName(pic.Header)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		//保存
		c.SaveUploadedFile(pic, name)
	}
	c.String(http.StatusOK,"done")
}

func HandleUpload(c *gin.Context) {
	fAvatar, err := c.FormFile("avatar")
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//todo: 检测文件大小
	//if fAvatar.Size > xxxx

	//获取存储路径
	name, err := localName(fAvatar.Header)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	//保存
	c.SaveUploadedFile(fAvatar, name)

	c.String(http.StatusOK, "upload avatar done")
}

// localName 为图片分配本地存储路径，同时简单检测是否为图片文件
func localName(h textproto.MIMEHeader) (name string, err error) {
	ct := h.Get("Content-Type")
	if ct == "" {
		err = errors.New("Content-Type error")
		return
	}
	cts := strings.Split(ct, "/")
	if len(cts) != 2 {
		err = errors.New("wrong Content-Type format")
		return
	}
	if cts[0] != "image" {
		err = errors.New("we need an image file")
		return
	}
	//should get from configure
	base := "./"
	rand.Seed(time.Now().UnixNano())
	name = fmt.Sprintf("%s/%09d.%s", base, rand.Int63n(999999999), cts[1])
	return
}
