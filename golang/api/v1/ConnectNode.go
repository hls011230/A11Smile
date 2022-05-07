package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"

	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math"
	"math/big"
)



//查询ETH
func Connect5_CheckTheBalance(id ,genre int) (string,error) {

	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	if genre == 0{
		DB.Table("users").First(&w,"id = ?",id)
	}else {
		DB.Table("gainers").First(&w,"id = ?",id)
	}

	res, err := eth.Ins.GetUserETH(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})
	if err != nil {
		log.Fatal(err)
		return "",err
	}

	fBalance := new(big.Float)
	fBalance.SetString(res.String())
	balanceEther := new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))
	return balanceEther.String(),nil
}

//查询AS
func Connect6_CheckTheAS(id , genre int) (string,error) {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	if genre == 0{
		DB.Table("users").First(&w,"id = ?",id)
	}else {
		DB.Table("gainers").First(&w,"id = ?",id)
	}

	res, err := eth.Ins.GetUserAS(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)})

	if err != nil {
		log.Fatal(err)
		return "",err
	}

	return res.String(),nil
}





