package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Gainer_ViewMedicalName(gid int) ([]string,error)  {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)

	res,err := eth.Ins.ViewMedicalNames(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})
	return res,err
}

