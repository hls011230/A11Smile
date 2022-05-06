package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func gainer_register_verifyBizlicenseHandler(c *gin.Context) {
	// 获取前端传递过来的文件
	f, err := c.FormFile("uploadBizlicense")
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

	//识别营业执照
	token, _ := v1.GetToken()
	srcFile, _ := f.Open()
	name,err := v1.PostBizlicense(srcFile, token)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, name)
}

func gainer_registerHandler(c *gin.Context) {

	gainer := model.Gainer{}
	if err := c.ShouldBindJSON(&gainer); err != nil {
		serializer.RespError(c,err)
		return
	}

	gainer.Gid,_ = strconv.Atoi(c.Request.Header.Get("gid"))
	err := v1.GainerRegister(&gainer)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,"注册成功")

}
