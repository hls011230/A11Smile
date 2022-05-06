package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func user_submitCertificateHandler(c *gin.Context)  {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	var submitcertificate  model.PostSubmitCertificate
	if err := c.ShouldBind(&submitcertificate); err != nil{
		serializer.RespError(c, err)
		return
	}

	err := v1.User_submitCertificate(uid,&submitcertificate)
	if err != nil{
		serializer.RespError(c,err)
		return
	}
	serializer.RespOK(c, "证书上传成功")

}