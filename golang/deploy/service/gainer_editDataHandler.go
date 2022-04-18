package service

import (
	v1 "A11Smile/deploy/api/v1"
	"A11Smile/deploy/db/model"
	"A11Smile/deploy/serializer"
	"github.com/gin-gonic/gin"
)

//编辑征求者资料
var gainer_id int
func gainer_editDataGetIdHandler(c *gin.Context){
	var gainer model.Gainer
	if err := c.ShouldBind(&gainer);err != nil{
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),
		})
	}

	gainer_id = gainer.Gid
	serializer.RespOK(c,nil)
}
func gainer_editDataGetDataHandler(c *gin.Context)  {
	todos,err := v1.GainerDataSeeUpdate(gainer_id)
	if err != nil{
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),
		})
	}
	serializer.RespOK(c,todos)

}

func gainer_editDataHandler(c *gin.Context){
	var gainer model.Gainer
	if err := c.ShouldBind(&gainer);err != nil {
		c.JSON(200,gin.H{
			"msg":"error",
			"err":err.Error(),

		})
	}

	if err := v1.GainerEditData(gainer.Gid,gainer.Enterprise_name,gainer.Introduce); err != nil{
		serializer.RespError(c,err)
		return

	}

	serializer.RespOK(c,"修改成功")


}
//查询资料
func gainer_SeeDataHandler(c *gin.Context)  {

	todos,err := v1.GainerSeeTodo()
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,todos)

}