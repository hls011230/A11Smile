package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
)

func gainer_loginHandler(c *gin.Context) {
	var user model.LoginUser
	if err := c.ShouldBind(&user); err != nil {
		serializer.RespError(c, err)
		return
	}
	UserId := v1.GainerLogin(&user)
	if UserId == 0 {
		serializer.RespError(c, "登录失败")
		return
	}

	serializer.RespOK(c, struct {
		Gid int `json:"gid"`
	}{
		Gid: UserId,
	})
}