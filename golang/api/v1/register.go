package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// 注册
func UserRegister(user *model.User) error {

	addr, pk, _ := createAccount()

	err := addUserIntoNode(addr)
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

func GainerRegister(gainer *model.Gainer) error {
	DB := db.Get()
	var name string
	DB.Raw("select enterprise_name from gainer_authentication where id = ?",gainer.Gid).Find(&name)
	addr, pk, _ := createAccount()
	err := addGainerIntoNode(addr,name)
	if err != nil {
		return err
	}

	err = sendUserETH(addr)
	if err != nil {
		return err
	}

	err = sendGainerAS(addr)
	if err != nil {
		return err
	}

	gainer.BlockAddress = addr
	gainer.PrivateKey = pk


	DB.Table("gainers").Save(gainer)

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
func addUserIntoNode(addr string) error {
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

	value := "100"

	valuef, err := strconv.ParseFloat(value, 64) //先转换为 float64

	if err != nil {
		log.Println("is not a number")
	}

	valueWei, isOk := new(big.Int).SetString(fmt.Sprintf("%.0f", valuef*1000000000000000000), 10)

	if !isOk {
		log.Println("float to bigInt failed!")
	}

	auth.Value = valueWei

	_, err = eth.AS.A11SmileGiveETH(auth, common.HexToAddress(addr))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// 在合约中添加征求者
func addGainerIntoNode(addr string,name string) error {
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

	_, err = eth.Ins.GainerSetDoctor(auth, common.HexToAddress(addr),name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func sendGainerAS(addr string) error {
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

	amount := big.NewInt(int64(1000))
	_, err = eth.AS.Transfer(auth, common.HexToAddress(addr),amount)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
