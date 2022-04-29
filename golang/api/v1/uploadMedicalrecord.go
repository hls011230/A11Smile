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
func ReleaseMedicalInformation(gid int) error {
	cliadd :=db.Get()
	var addr string
	cliadd.Select("select block_address from gainers where id = ? ",gid).Find(&addr)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	clipk.Select("select private_key from gainers where id = ? ",gid).Find(&pk)

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

	var gainer model.Soliciter_solidity
	_, err = eth.Ins.GainerAddMedicalInformation(auth,gainer.HospitalName,big.NewInt(gainer.Min_),big.NewInt(gainer.Max),big.NewInt(gainer.Account), gainer.MedicalName,gainer.MedicalNeed,gainer.RequirementDescription)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
