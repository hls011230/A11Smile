package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
)

func user_DisplayHomepageDepartmentHandler(c *gin.Context)  {
	var DepartmentMedical model.POSTDepartment
	if err := c.ShouldBind(&DepartmentMedical); err != nil{
		serializer.RespError(c, err)
		return
	}
	res,err := v1.DisplayHomepageDepartment(&DepartmentMedical)
	if err != nil{
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,res)

}


