package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ShowNoTransactionsHandler(id int)([]interface{},error)  {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w,"id = ?",id)
	r1,r2,err :=eth.Ins.UserYSeeCertificateState(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})

	var r []interface{}
	for k, v := range r1 {
		if r2[k] {
			r1 := struct{
				User         common.Address
				Soliciter    common.Address
				HospitalName string
				MedicalName  string
				Certificate  string
<<<<<<< HEAD
			}{
				User: v.User,

=======
				Amount         *big.Int
			}{
				User: v.User,
>>>>>>> bff38e53fce35580c1260ec9797be31e1f1e1d9f
				Soliciter: v.Soliciter,
				HospitalName: v.HospitalName,
				MedicalName: v.MedicalName,
				Certificate: fmt.Sprintf("0x%x",v.Certificate),
<<<<<<< HEAD

=======
				Amount: v.Amount,
>>>>>>> bff38e53fce35580c1260ec9797be31e1f1e1d9f
			}
			r = append(r, r1)
		}

	}

	return r,err
}

