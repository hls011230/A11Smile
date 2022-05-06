package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

func User_submitCertificate(uid int,submitCertificate *model.PostSubmitCertificate) error  {
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

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))

	_,err = eth.Ins.UserAddMedicalInformation(auth,common.HexToHash(submitCertificate.Certificate),common.HexToAddress(submitCertificate.Soliciter),submitCertificate.MedicalName)
	if err != nil {
		return err
	}
	return nil
}
