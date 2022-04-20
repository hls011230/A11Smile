package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

//编辑用户资料
func user_editDataGetIdHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}
	serializer.RespOK(c, nil)
}

func user_editDataGetDataHandler(c *gin.Context) {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	user, err := v1.DataSeeUpdate(uid)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}
	serializer.RespOK(c, user)

}

func user_editDataHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	user.Id = uid

	if err := v1.EditData(user.Id, user.Uname, user.Resume); err != nil {
		serializer.RespError(c, err)
		return

	}

	serializer.RespOK(c, "修改成功!")

}

//查询资料
func user_SeeDataHandler(c *gin.Context) {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	user := v1.UserSeeTodo(uid)


	serializer.RespOK(c, user)

}
