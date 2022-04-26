package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_CheckTheAS(c *gin.Context)  {

	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))

	balance,err := v1.Connect6_CheckTheAS(gid,1)

	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, balance)

}
