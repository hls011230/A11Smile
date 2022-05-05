package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
)

func user_displayHomepageHandler(c *gin.Context){
	gainermsg,err := v1.DisplayHomepage()
	if err != nil{
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,gainermsg)

}

func user_showDetailsPageHandler(c *gin.Context) {
	var details model.PostDetails
	if err := c.ShouldBindJSON(&details); err != nil {
		serializer.RespError(c,err)
		return
	}

	res,err := v1.ShowDetailsPage(details)
	if err != nil {
		serializer.RespError(c,err)
		return
	}

	serializer.RespOK(c,res)

}
