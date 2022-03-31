package service

import (
	v1 "A11Smile/deploy/api/v1"
	"A11Smile/deploy/serializer"
	"fmt"
	"path"
	"strings"
	"time"

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

	// 保存文件
	fileName := time.Now().Unix()
	fileDir := "static/img/"
	filePath := fmt.Sprintf("%s%d%s", fileDir, fileName, fileExt)
	if err := c.SaveUploadedFile(f, filePath); err != nil {
		serializer.RespError(c, err)
		return
	}

	// 识别营业执照
	err = v1.VerifyBizlicense(filePath)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "认证成功")
}
