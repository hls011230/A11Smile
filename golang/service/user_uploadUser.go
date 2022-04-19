package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func user_uploadUser(c *gin.Context)()  {
	var upUser model.User_solidity
	if err := c.ShouldBind(&upUser); err != nil {
		serializer.RespError(c, err)
		return
	}

	err,chainID,gasPrice,ins := v1.Connect_uploadUser(&upUser)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "用户上传成功")

	fmt.Print(err,chainID,gasPrice,ins)







}
