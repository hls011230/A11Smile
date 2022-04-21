package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_CheckTheBalance(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	balance,err := v1.Connect5_CheckTheBalance(uid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, balance)


}

