package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_DisplayWarehouseHandler(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	var Warehouse model.PostWarehouse
	if err := c.ShouldBind(&Warehouse); err != nil {
		serializer.RespError(c, err)
		return
	}
	res,err:=v1.Gainer_DisplayWarehouse(gid,&Warehouse)
	if err!=nil {
		serializer.RespError(c,err)
	}
	serializer.RespOK(c,res)
}