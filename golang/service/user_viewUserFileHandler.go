package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func user_viewAllMedicalHistoryHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	nameArray,err := v1.ViewAllMedicalHistory(uid)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,nameArray)
}

func user_viewAllMedicalExaminationReportHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))

	nameArray,err := v1.ViewAllMedicalExaminationReport(uid)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,nameArray)
}

func user_previewMedicalHistoryHandler(c *gin.Context)  {
	uid,_ := strconv.Atoi(c.Request.Header.Get("uid"))
	var userFile model.PostUserFile
	if err := c.ShouldBindJSON(&userFile); err != nil {
		serializer.RespError(c,err)
		return
	}

	fileLoadPath,err := v1.PreviewMedicalHistory(uid,userFile.FileName)
	if err != nil {
		serializer.RespError(c,err)
		return
	}
	serializer.RespOK(c,fileLoadPath)
}

func user_previewMedicalExaminationReportHandler(c *gin.Context) {
	uid , _ := strconv.Atoi(c.Request.Header.Get("uid"))
	var userFile model.PostUserFile
	if err := c.ShouldBindJSON(&userFile); err != nil {
		serializer.RespError(c,err)
		return
	}

	fileLoadPath,err := v1.PreviewMedicalExaminationReport(uid,userFile.FileName)
	if err != nil {
		serializer.RespError(c,err)
		return
	}
	serializer.RespOK(c,fileLoadPath)

}