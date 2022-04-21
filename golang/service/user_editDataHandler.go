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
//修改用户名
func user_editUserNameHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	if err := v1.EditUserName(user.Id, user.Uname); err != nil {
		serializer.RespError(c, err)
		return

	}

	serializer.RespOK(c, "修改成功!")

}

//修改用户简介
func user_editUserResumeHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	if err := v1.EditUserResume(user.Id, user.Resume); err != nil {
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
func User_AuthenticationSeeHandler(c *gin.Context){


	//serializer.RespOK(c,struct {
	//	Block_address string `json:"block_address"`
	//	Birthday      string `json:"birthday"`
	//	Resume        string `json:"resume"`
	//	Uname         string `json:"uname"`
	//	Gender        string `json:"gender"`
	//}{
	//	Block_address: block_address,
	//	Birthday: birthday,
	//	Resume: resume,
	//	Uname: uname,
	//	Gender: gender,
	//})
}
