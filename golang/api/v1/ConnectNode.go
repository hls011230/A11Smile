package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"

	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"log"
	"math"
	"math/big"

)



//用户上传医疗信息
func Connect2_UploadMedicalInformation(user *model.User_solidity,uid int) error {
	cliadd :=db.Get()
	var add string
	cliadd.Select("select private_key from user where id = ? ",uid).Find(&add)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(add))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	clipk.Select("select address from user where id = ? ",uid).Find(&pk)

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))


	_, err = eth.Ins.UserAddMedicalInformation(auth, user.Proute,common.HexToAddress(user.Soliciter_))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}


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

//征求者查询ETH
func Connect5_GainCheckTheBalance(gid int) error {
	cliadd :=db.Get()
	var addr string
	cliadd.Select("select private_key from user where id = ? ",gid).Find(&addr)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	clipk.Select("select address from user where id = ? ",gid).Find(&pk)

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))

	ETH, err := eth.Ins.GetUserETH(nil)
	fmt.Println("剩余ETH",ETH)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}



