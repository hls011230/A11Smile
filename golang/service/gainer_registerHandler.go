package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

func gainer_register_verifyBizlicense(c *gin.Context) {
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
	err = v1.PostBizlicense(srcFile, token)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "认证成功")
}
