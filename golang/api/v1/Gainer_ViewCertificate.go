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

func Gainer_ViewCertificate(gid int)([]gen.UploadMedicalrecords_useruser_MedicalCertificateState, error)  {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)
	res,err :=eth.Ins.GainerSeeuserUploadMedical(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})
	return res, err


}
