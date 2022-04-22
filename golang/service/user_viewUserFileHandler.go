package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func user_viewAllMedicalHistory(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	nameArray,err := v1.ViewAllMedicalHistory(uid)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,nameArray)
}

func user_viewAllMedicalExaminationReport(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	nameArray,err := v1.ViewAllMedicalExaminationReport(uid)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,nameArray)
}