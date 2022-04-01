package service

import (
	"A11Smile/deploy/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
	"time"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 * 20//50MB
)

func user_uploadMedicalHandler(c *gin.Context){
	f,err:= c.FormFile("uploadMedical")
	//判断文件大小
	if err = c.ShouldBind(MAX_UPLOAD_SIZE);err != nil{
		serializer.RespError(c,"file is to big")
		return
	}
	//判断上传文件的类型
	fileExt := strings.ToLower(path.Ext(f.Filename))
	if fileExt != ".txt" && fileExt != ".doc" && fileExt != ".exe" {
		serializer.RespError(c, "文件格式错误")
		return
	}
	//保存文件
	fileName := time.Now().Unix()
	fildDir := fmt.Sprintf("%s%d%s/",time.Now().Year(),time.Now().Month().String())
	filePath := fmt.Sprintf("%s%s%s",fildDir,fileName,fileExt)
	if err := c.SaveUploadedFile(f, filePath); err != nil {
		serializer.RespError(c, "保存失败")
		return
	}


	serializer.RespOK(c,"上传成功")

}
