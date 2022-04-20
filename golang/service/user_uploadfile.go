package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"strings"
)

func user_uploadMedicalHistoryHandler(c *gin.Context)  {
	// 获取前端传递过来的文件
	f, err := c.FormFile("uploadMedicalHistory")
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

	// 上传用户的病历信息
	token, _ := v1.GetToken()
	srcFile, _ := f.Open()
	err = v1.UploadMedicalHistory(srcFile, token, uid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "上传病历信息成功")
}