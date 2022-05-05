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

//征求者发布医疗信息
func ReleaseMedicalInformation(gid int,gainer *model.Soliciter_solidity) error {
	DB :=db.Get()
	var user model.Wallet
	DB.Table("gainers").First(&user,"id = ?",gid)

	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(user.BlockAddress))
	if err != nil {
		log.Fatal(err)
		return err
	}


	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
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

	_, err = eth.Ins.GainerAddMedicalInformation(auth,big.NewInt(gainer.Min_),big.NewInt(gainer.Max),big.NewInt(gainer.Account), gainer.MedicalName,gainer.MedicalNeed,gainer.RequirementDescription)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
