package service

import "C"
import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_ASforETHHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	var ASfroETH model.PostETHforAS
	if err := c.ShouldBind(&ASfroETH); err != nil {
		serializer.RespError(c, err)
		return
	}
	err := v1.Gainer_ASforETH(uid, gid,&ASfroETH)
	if err !=nil{
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,"兑换成功")

}

