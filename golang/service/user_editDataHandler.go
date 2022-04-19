package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
)

//编辑用户资料
var Id int

func user_editDataGetNameHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	Id = user.Id
	serializer.RespOK(c, nil)
}

func user_editDataGetDataHandler(c *gin.Context) {
	todos, err := v1.DataSeeUpdate(Id)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}
	serializer.RespOK(c, todos)

}

func user_editDataHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	if err := v1.EditData(user.Id, user.Uname, user.Resume); err != nil {
		serializer.RespError(c, err)
		return

	}

	serializer.RespOK(c, "修改成功!")

}

//查询资料
func user_SeeDataHandler(c *gin.Context) {

	todos, err := v1.UserSeeTodo()
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, todos)

}
