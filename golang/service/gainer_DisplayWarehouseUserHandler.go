package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"github.com/gin-gonic/gin"
	"strconv"
)

func gainer_DisplayWarehouseUserHandler(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	var WarehouseUser model.PostWarehouse
	if err := c.ShouldBind(&WarehouseUser); err != nil {
		serializer.RespError(c, err)
		return
	}

	res, err := v1.Gainer_DisplayWarehouseUser(gid,&WarehouseUser)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c,res)
}

