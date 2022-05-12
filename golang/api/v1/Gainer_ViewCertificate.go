package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func Gainer_ViewCertificate(gid int)([]interface{}, error)  {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)
	res,err :=eth.Ins.GainerSeeuserUploadMedical(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})
	var r []interface{}
	for _, v := range res {
		if v.MedicalName != "" {

			r1 := struct{
				User         common.Address
				State        bool
				Soliciter    common.Address
				HospitalName string
				MedicalName  string
				Certificate  string
				Amount         *big.Int
			}{
				User: v.User,
				Soliciter: v.Soliciter,
				HospitalName: v.HospitalName,
				MedicalName: v.MedicalName,
				Certificate: fmt.Sprintf("0x%x",v.Certificate),
				Amount: v.Amount,
			}
			r = append(r, r1)
		}
	}
	return r, err


}
