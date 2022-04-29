package service

import (
v1 "A11Smile/api/v1"
"A11Smile/serializer"
"github.com/gin-gonic/gin"
)

func gainer_displayHomepageHandler(c *gin.Context){
	gainermsg,err := v1.DisplayHomepage()
	if err != nil{
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,gainermsg)

}

