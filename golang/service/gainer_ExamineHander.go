package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_ExamineHander(c *gin.Context)  {
	var examine model.PostExamine
	if err := c.ShouldBind(&examine); err != nil {
		serializer.RespError(c, err)
		return
	}
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	err := v1.Gainer_Examine(gid,&examine)
	if err != nil{
		serializer.RespError(c,err)
		return
	}
	serializer.RespOK(c, "审核成功")
}
