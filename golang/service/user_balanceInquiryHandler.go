package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func user_CheckTheBalanceHandler(c *gin.Context)  {

	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	balance,err := v1.Connect5_CheckTheBalance(uid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, balance)
}



func user_CheckTheAS(c *gin.Context)  {

	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	balance,err := v1.Connect6_CheckTheAS(uid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, balance)
}
