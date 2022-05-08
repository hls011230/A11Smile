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
	"math/big"
	"strconv"
)

func Ganiner_ETHforAs(uid int,gid int,as *model.PostETHforAS) error  {
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w,"id = ?",uid)

	// 在合约中存入用户病历信息
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(w.BlockAddress))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(w.PrivateKey)
	if err != nil {

		return err
	}

	value := as.Quantity
	valuef, err := strconv.ParseFloat(string(value), 64) //先转换为 float64

	if err != nil {
		log.Println("is not a number")
	}

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)
	if !isOk {
		log.Println("float to bigInt failed!")
	}


	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = valueWei

	var x model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)

	noncex, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(x.BlockAddress))
	if err != nil {

		return err
	}

	privateKeyx, err := crypto.HexToECDSA(x.PrivateKey)
	if err != nil {

		return err
	}

	authx, err := bind.NewKeyedTransactorWithChainID(privateKeyx, eth.ChainID)
	if err != nil {
		return err
	}


	authx.GasPrice = eth.GasPrice
	authx.GasLimit = uint64(6000000)
	authx.Nonce = big.NewInt(int64(noncex))

	_,err = eth.AS.EthGetAs(auth,common.HexToAddress(as.AddETH))
	_,err = eth.AS.AsgetETH(authx,common.HexToAddress(as.RedETH),big.NewInt(int64(as.Quantity)))
	if err!=nil{
		return err
	}

	return err

}

