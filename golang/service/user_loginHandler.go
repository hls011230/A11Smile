package service

import (
	v1 "A11Smile/deploy/api/v1"
	"A11Smile/deploy/db/model"
	"A11Smile/deploy/serializer"

	"github.com/gin-gonic/gin"
)

func user_loginHandler(c *gin.Context) {
	var user model.LoginUser
	if err := c.ShouldBind(&user); err != nil {
		serializer.RespError(c, err)
		return
	}
	count := v1.Login(&user)
	if count == 0 {
		serializer.RespError(c, "登录失败")
		return
	}

	res, err := v1.LoginWXSmall(c.PostForm("code"))
	if err != nil {
		serializer.RespError(c, "登录失败")
		return
	}

	serializer.RespOK(c, res)
}
