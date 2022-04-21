package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

//编辑征求者资料
func gainer_editDataGetIdHandler(c *gin.Context){
	var gainer model.GainerAuthentication
	if err := c.ShouldBind(&gainer);err != nil{
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),
		})
	}
	serializer.RespOK(c,nil)
}
func gainer_editDataGetDataHandler(c *gin.Context)  {
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	gainer,err := v1.GainerDataSeeUpdate(gid)
	if err != nil{
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),
		})
	}
	serializer.RespOK(c,gainer)

}

func gainer_editDataHandler(c *gin.Context){
	var gainer model.GainerAuthentication
	if err := c.ShouldBind(&gainer);err != nil {
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),

		})
	}

	if err := v1.GainerEditData(gainer.Id,gainer.Resume); err != nil{
		serializer.RespError(c,err)
		return

	}

	serializer.RespOK(c,"修改成功")


}
//查询资料
func gainer_SeeDataHandler(c *gin.Context)  {
	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	gainer:= v1.GainerSeeTodo(gid)


	serializer.RespOK(c,gainer)

}