package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func gainer_CheckTheAS(c *gin.Context)  {
	var seeAS model.AllPeople_solidity
	if err := c.ShouldBind(&seeAS); err != nil {
		serializer.RespError(c, err)
		return
	}

	err, chainID, gasPrice, ins := v1.Connect6_CheckTheAS(&seeAS)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "AS查询成功")
	fmt.Print(err, chainID, gasPrice, ins)


}
