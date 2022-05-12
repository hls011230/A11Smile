package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"A11Smile/eth/gen"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Gainer_DisplayWarehouse(gid int,Warehouse *model.PostWarehouse)(gen.UploadMedicalrecords_gainerData,error)  {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)
	res,err:=eth.Ins.ViewWareouse(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)},Warehouse.Medical,common.HexToAddress(Warehouse.User))
	if err!=nil{
		return res,err
	}
	return res,nil
}

