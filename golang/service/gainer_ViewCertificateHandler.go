package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_ViewCertificateHandler(c *gin.Context)  {
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	usercertificate,err :=v1.Gainer_ViewCertificate(gid)
	if err != nil{
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,usercertificate)

}
