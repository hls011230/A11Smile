package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func user_CheckTheAS(c *gin.Context)  {
	var seeAS model.AllPeople_solidity
	if err := c.ShouldBind(&seeAS); err != nil {
		serializer.RespError(c, err)
		return
	}

	err := v1.Connect6_CheckTheAS()
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	fmt.Print(err)


}

