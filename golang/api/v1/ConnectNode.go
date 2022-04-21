package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"A11Smile/eth/gen"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"log"
	"math"
	"math/big"
	"os"
)



//设置征求者
func Connect1_uploadGainer(gainer *model.Soliciter_solidity) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
	//连接端口
	client, err := ethclient.Dial("http://127.0.0.1:8547")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	//获取ChainID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(chainID)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasPrice)

	//绑定合约
	ins, err := gen.NewUploadMedicalrecords(common.HexToAddress("0xd9145CCE52D386f254917e481eB44e9943F39138"), client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ins)

	//使用opts付费
	keyfile := "UTC--2022-04-08T08-31-50.357582000Z--14d1f23fbc43d02251f48ee70f43d3d4175bee8e"
	reader, _ := os.Open(keyfile)
	opts, err := bind.NewTransactorWithChainID(reader, "123456", chainID)
	if err != nil {
		log.Fatal("NewTransactor", err)
	}
	opts.GasLimit = 3000000
	opts.GasPrice = gasPrice

	fmt.Println("opts:", opts)
	//上传征求者
	tx1, err := ins.GainerSetDoctor(opts, common.HexToAddress(gainer .Soliciter_))
	fmt.Println("上传征求者:", tx1)
	return err, chainID, gasPrice, ins
}

//用户上传医疗信息
func Connect2_UploadMedicalInformation(user *model.User_solidity) error {
	cliadd :=db.Get()
	var add string
	var uidadd *gorm.DB
	cliadd.Select("select privatekey from user where id = ? ",uidadd).Find(&add)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(add))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	var uidpk *gorm.DB
	clipk.Select("select address from user where id = ? ",uidpk).Find(&pk)

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
	auth.Value = big.NewInt(int64(1000000))

	_, err = eth.Ins.UserAddMedicalInformation(auth, user.Proute,common.HexToAddress(user.Soliciter_))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}


//征求者发布医疗信息
func Connect3_ReleaseMedicalInformation(gainer *model.Soliciter_solidity) error {
	cliadd :=db.Get()
	var add string
	var uidadd *gorm.DB
	cliadd.Select("select privatekey from user where id = ? ",uidadd).Find(&add)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(add))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	var uidpk *gorm.DB
	clipk.Select("select address from user where id = ? ",uidpk).Find(&pk)

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
	auth.Value = big.NewInt(int64(1000000))

	_, err = eth.Ins.GainerAddMedicalInformation(auth, big.NewInt(gainer.Min_),big.NewInt(gainer.Max), gainer.MedicalName,gainer.MedicalNeed,gainer.RequirementDescription)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

//征求者审核与奖励
func Connect4_ReviewAndReward(gainer *model.Soliciter_solidity) error {
	cliadd :=db.Get()
	var add string
	var uidadd *gorm.DB
	cliadd.Select("select privatekey from user where id = ? ",uidadd).Find(&add)
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(add))
	if err != nil {
		log.Fatal(err)
		return err
	}

	clipk :=db.Get()
	var pk string
	var uidpk *gorm.DB
	clipk.Select("select address from user where id = ? ",uidpk).Find(&pk)

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
	auth.Value = big.NewInt(int64(1000000))

	_, err = eth.Ins.GainerWhether(auth,common.HexToAddress(gainer.Soliciter_),gainer.Proute, gainer.Whether,big.NewInt(gainer.Ercnum_))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

//查询ETH
func Connect5_CheckTheBalance(uid int) (string,error) {

	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var user model.UserWallet
	DB.Table("users").First(&user,"id = ?",uid)

	res, err := eth.Ins.GetUserETH(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)})
	if err != nil {
		log.Fatal(err)
		return "",err
	}

	fBalance := new(big.Float)
	fBalance.SetString(res.String())
	balanceEther := new(big.Float).Quo(fBalance,big.NewFloat(math.Pow10(18)))
	return balanceEther.String(),err
}

//查询AS
func Connect6_CheckTheAS(uid int) (string,error) {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var user model.UserWallet
	DB.Table("users").First(&user,"id = ?",uid)

	res, err := eth.Ins.GetUserAS(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)})

	if err != nil {
		log.Fatal(err)
		return "",err
	}

	return res.String(),err
}

