package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"fmt"

	"github.com/gin-gonic/gin"

	"path"
	"strings"
	"time"
)

func user_verifyIDCardHandler(c *gin.Context) {
	// 获取前端传递过来的文件
	f, err := c.FormFile("uploadIDCard")
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	// 处理文件
	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		serializer.RespError(c, "文件格式错误")
		return
	}

	// 识别身份证
	token, _ := v1.GetToken()
	srcFile, _ := f.Open()
	err = v1.PostIDCard(srcFile, token)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "认证成功")
}

func user_readMedicalInformation(c *gin.Context) {
	// 获取前端传递过来的文件
	f, err := c.FormFile("readMedicalInformation")
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	// 处理文件
	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		serializer.RespError(c, "文件格式错误")
		return
	}

	// 保存文件
	fileName := time.Now().Unix()
	fileDir := "static/img/"
	filePath := fmt.Sprintf("%s%d%s", fileDir, fileName, fileExt)
	if err := c.SaveUploadedFile(f, filePath); err != nil {
		serializer.RespError(c, err)
		return
	}

	// 识别医疗信息
	err = v1.GetMedicalInformation(filePath)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "认证成功")

}
