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


func ShowAllTransactionsHandler(id int)([]interface{},error)  {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w,"id = ?",id)
	res,err :=eth.Ins.UserNSeeCertificateState(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})

	var r []interface{}
	for _, v := range res {
		r1 := struct{
			User         common.Address
			State        bool
			Soliciter    common.Address
			HospitalName string
			MedicalName  string
			Certificate  string
			Erum         *big.Int
		}{
			User: v.User,
			State: v.State,
			Soliciter: v.Soliciter,
			HospitalName: v.HospitalName,
			MedicalName: v.MedicalName,
			Certificate: fmt.Sprintf("0x%x",v.Certificate),
			Erum: v.Erum,
		}
		r = append(r, r1)

	}

	return r,err
}

