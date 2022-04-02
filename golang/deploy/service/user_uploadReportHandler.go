package service

import (
	v1 "A11Smile/deploy/api/v1"
	"A11Smile/deploy/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
)

const (
	MAX_Report_SIZE = 1024 * 1024 * 20//50MB
)

//用户上传文件并保存文件
func user_uploadReportHandler(c *gin.Context){
	f,err:= c.FormFile("uploadMedical")
	//判断文件大小
	if err = c.ShouldBind(MAX_Report_SIZE);err != nil{
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
	//fileName := time.Now().Unix()
	//fildDir := fmt.Sprintf("%s%d%s/",time.Now().Year(),time.Now().Month().String())

	//判断文件夹是否存在
	_Reportdir := "./体检报告"
	exist, err := v1.UploadReport(_Reportdir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		fmt.Printf("has dir![%v]\n", _Reportdir)
	} else {
		fmt.Printf("no dir![%v]\n", _Reportdir)
		// 创建文件夹
		err := os.Mkdir(_Reportdir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			fmt.Printf("mkdir success!\n")
		}
	}
	//filePath := fmt.Sprintf("%s%s%s", fildDir, fileName, fileExt)

	//设置文件须要保存的指定位置并设置保存的文件名字
	dst := path.Join(_Reportdir,f.Filename)
	if err := c.SaveUploadedFile(f, dst); err != nil {
		serializer.RespError(c, "保存失败")
		return
	}


	serializer.RespOK(c,"上传成功")

}
