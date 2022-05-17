package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

func Gainer_Examine(gid int,Examine *model.PostExamine) error {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)

	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(w.BlockAddress))
	if err != nil {
		log.Fatal(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(w.PrivateKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey,eth.ChainID)
	if err != nil {
		log.Fatal(err)
		return err
	}


	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	_,err = eth.Ins.GainerWhether(auth,common.HexToHash(Examine.Certificate),Examine.MedicalName,Examine.Whether,common.HexToAddress(Examine.Address),big.NewInt(Examine.Ercnum))


	nonce1, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(w.BlockAddress))
	if err != nil {
		log.Fatal(err)
		return err
	}


	auth1, err := bind.NewKeyedTransactorWithChainID(privateKey,eth.ChainID)
	if err != nil {
		log.Fatal(err)
		return err
	}


	auth1.GasPrice = eth.GasPrice
	auth1.GasLimit = uint64(3000000)
	auth1.Nonce = big.NewInt(int64(nonce1))

	_, err = eth.AS.Transfer(auth1, common.HexToAddress(Examine.Address), big.NewInt(Examine.Ercnum))
	if err != nil {
		return err
	}


	return nil
}
