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


func ShowAllTransactionsHandler(id int)([]gen.UploadMedicalrecords_useruser_MedicalCertificateState,error)  {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w,"id = ?",id)
	res,err :=eth.Ins.UserNSeeCertificateState(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})
	return res,err
}

