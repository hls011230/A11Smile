package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

//编辑征求者资料
func gainer_editDataHandler(c *gin.Context){
	var gainer model.Gainer
	if err := c.ShouldBind(&gainer);err != nil{
		serializer.RespError(c,err)
		return
	}

	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	gainer.Id = gid

	if err := v1.GainerEditData(gainer.Id,gainer.Resume); err != nil{
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,"修改成功")
}

//查询资料
func gainer_authenticationSeeHandler(c *gin.Context){
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	gainer := v1.GainerAuthenticationSee(gid)
	serializer.RespOK(c,gainer)
}