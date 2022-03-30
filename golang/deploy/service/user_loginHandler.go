package service

import (
	"github.com/gin-gonic/gin"
	v1 "go_eth/Test_register/depoly/api/v1"
	"go_eth/Test_register/depoly/db/model"
	"go_eth/Test_register/depoly/serializer"
)

func user_loginHandler(c *gin.Context)  {
	var user model.LoginUser
	if err := c.ShouldBind(&user);err != nil {
		serializer.RespError(c,err)
		return
	}
	count := v1.Login(&user)
	if count == 0 {
		serializer.RespError(c,"登录失败")
		return
	}

	res,err := v1.LoginWXSmall(c.PostForm("code"))
	if err != nil {
		serializer.RespError(c,"登录失败")
		return
	}

	serializer.RespOK(c,res)
}