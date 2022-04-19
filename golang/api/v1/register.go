package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 注册
func Register(user *model.User) error {

	addr, pk, _ := createAccount()

	err := addUserInToNode(addr)
	if err != nil {
		return err
	}

	err = sendUserETH(addr)
	if err != nil {
		return err
	}

	user.BlockAddress = addr
	user.PrivateKey = pk

	cli := db.Get()
	cli.Table("users").Save(user)

	return nil
}

// 创建私钥
func createAccount() (addr string, pk string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err)
		return "", "", nil
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	log.Println("address: ", address)

	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("privateKey: ", privateKey)
	return address, privateKey, nil
}

// 在合约中添加用户
func addUserInToNode(addr string) error {
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(model.Deployer.Address))
	if err != nil {
		log.Fatal(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(model.Deployer.PrivateKey)
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

	_, err = eth.Ins.UserAdduser(auth, common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 给新账户转ETH
func sendUserETH(addr string) error {
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(model.Deployer.Address))
	if err != nil {
		log.Fatal(err)
		return err
	}

	privateKey, err := crypto.HexToECDSA(model.Deployer.PrivateKey)
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
	auth.Value = big.NewInt(int64(1000000))

	_, err = eth.Ins.A11SmileGiveETH(auth, common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
