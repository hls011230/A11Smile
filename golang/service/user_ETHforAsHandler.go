package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func user_ETHforAsHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	var EthForAs model.PostETHforAS
	if err := c.ShouldBind(&EthForAs); err != nil {
		serializer.RespError(c, err)
		return
	}

	err := v1.User_ETHforAs(uid,gid,&EthForAs)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "兑换成功")
}
