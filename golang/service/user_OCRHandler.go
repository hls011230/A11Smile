package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"

	"path"
	"strings"
)

func user_verifyIDCardHandler(c *gin.Context) {
	// 获取前端传递过来的文件
	f, err := c.FormFile("uploadIDCard")
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))

	// 处理文件
	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		serializer.RespError(c, "文件格式错误")
		return
	}

	// 识别身份证
	token, _ := v1.GetToken()
	srcFile, _ := f.Open()
	err = v1.PostIDCard(srcFile, token, uid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "认证成功")
}

func user_readMedicalInformationHandler(c *gin.Context) {
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

	// 识别医疗信息
	token, _ := v1.GetToken()
	srcFile, _ := f.Open()
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	err = v1.PostMedicalInformation(srcFile, token, uid)
	if err != nil {
		serializer.RespError(c, err.Error())
		return
	}

	serializer.RespOK(c, "识别成功")

}
