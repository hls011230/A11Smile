package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_CheckTheBalance(c *gin.Context)  {
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	balance,err := v1.Connect5_CheckTheBalance(gid)

	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, balance)

}

