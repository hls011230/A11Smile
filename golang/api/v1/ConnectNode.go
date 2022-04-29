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


//征求者发布医疗信息
func Connect3_ReleaseMedicalInformation(gainer *model.Soliciter_solidity,gid int) error {
	DB := db.Get()
	var user model.Wallet
	DB.Table("gainers").First(&user,"id = ?",gid)

	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {
		log.Fatal("错误",err)
		return err
	}

	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(user.BlockAddress))
	if err != nil {

		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		log.Fatal("错误",err)
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))

	min := big.NewInt(int64(gainer.Min_))

	_, err = eth.Ins.GainerAddMedicalInformation(auth,"123",min,big.NewInt(int64(gainer.Max)),big.NewInt(int64(10)),gainer.MedicalName,gainer.MedicalNeed,gainer.RequirementDescription)
	if err != nil {
		log.Fatal("错误",err)
		return err
	}
	return err
}

//征求者审核与奖励
func Connect4_ReviewAndReward(gainer *model.Soliciter_solidity,gid int) error {
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


	_, err = eth.Ins.GainerWhether(auth,common.HexToAddress(gainer.Soliciter_),gainer.Proute, gainer.Whether,big.NewInt(int64(gainer.Ercnum_)))
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
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



