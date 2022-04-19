package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
)

func gainer_ReviewAndReward(c *gin.Context)  {
	var upReviewAndReward model.Soliciter_solidity
	if err := c.ShouldBind(&upReviewAndReward); err != nil {
		serializer.RespError(c, err)
		return
	}

	err:= v1.Connect4_ReviewAndReward(&upReviewAndReward)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "征求者审核结束")
	fmt.Print(err)


}


