package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateCertificateHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	var pc model.PostCertificate
	if err := c.ShouldBindJSON(&pc);err != nil {
		serializer.RespError(c,err)
		return
	}

	err := v1.CreateCertificate(uid,pc)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,"证书生成成功")

}

func ShowAllCertificateHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	res,err := v1.ShowAllCertificate(uid)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,res)

}

func ShowDetailsCertificateHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	var serial model.PostCertificateHash
	if err := c.ShouldBindJSON(&serial);err != nil {
		serializer.RespError(c,err)
		return
	}
	res,err := v1.ShowDetailsCertificate(uid,serial.Serial)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,res)

}
