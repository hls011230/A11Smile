package service


import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_ReleaseMedicalInformation(c *gin.Context)  {
	var upMedical model.Soliciter_solidity
	if err := c.ShouldBind(&upMedical); err != nil {
		serializer.RespError(c, err)
		return
	}

	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	err := v1.Connect3_ReleaseMedicalInformation(&upMedical,gid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "征求者发布医疗信息成功")
	fmt.Print(err)

}

