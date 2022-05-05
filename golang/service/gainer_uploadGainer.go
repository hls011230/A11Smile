package service

//func gainer_uploadGainer(c *gin.Context)() {
//	var upGainer model.Soliciter_solidity
//	if err := c.ShouldBind(&upGainer); err != nil {
//		serializer.RespError(c, err)
//		return
//	}
//
//	err, chainID, gasPrice, ins := v1.Connect1_uploadGainer(&upGainer)
//	if err != nil {
//		serializer.RespError(c, err)
//		return
//	}
//	serializer.RespOK(c, "征求者上传成功")
//	fmt.Print(err, chainID, gasPrice, ins)
//
//}