package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_ViewMedicalNameHandler(c *gin.Context)  {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	res,err := v1.Gainer_ViewMedicalName(gid)
	if err!= nil{
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c,res)
}

