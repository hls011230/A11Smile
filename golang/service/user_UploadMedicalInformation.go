package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func user_UploadMedicalInformation(c *gin.Context)  {
	var upMedical model.User_solidity
	if err := c.ShouldBind(&upMedical); err != nil {
		serializer.RespError(c, err)
		return
	}

	err:= v1.Connect2_UploadMedicalInformation(&upMedical)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "用户上传医疗信息成功")
	fmt.Print(err)

}
