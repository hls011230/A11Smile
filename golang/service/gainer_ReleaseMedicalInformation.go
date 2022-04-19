package service

//import (
//	v1 "A11Smile/api/v1"
//	"A11Smile/db/model"
//	"A11Smile/serializer"
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//func gainer_ReleaseMedicalInformation(c *gin.Context)  {
//	var upMedical model.Soliciter_solidity
//func gainer_ReleaseMedicalInformation(c *gin.Context) {
//	var upMedical model.User_AddMedicalInformation
// 1cd5c200c1dd55df222463be19da3b87b821c6ee
//	if err := c.ShouldBind(&upMedical); err != nil {
//		serializer.RespError(c, err)
//		return
//	}
//
//	err, chainID, gasPrice, ins := v1.Connect3_ReleaseMedicalInformation(&upMedical)
//	if err != nil {
//		serializer.RespError(c, err)
//		return
//	}
//	serializer.RespOK(c, "征求者发布医疗信息成功")
//	fmt.Print(err, chainID, gasPrice, ins)
//
//}
