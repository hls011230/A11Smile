package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_CheckTheBalance(c *gin.Context)  {
	var seeETH model.AllPeople_solidity
	if err := c.ShouldBind(&seeETH ); err != nil {
		serializer.RespError(c, err)
		return
	}

	gid,_ := strconv.Atoi(c.Request.Header.Get("gid"))
	err:= v1.Connect5_GainCheckTheBalance(gid)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "ETH查询成功")
	fmt.Print(err)


}

