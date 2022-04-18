package v1

import (
	"A11Smile/db/model"
	"A11Smile/eth/gen"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

//设置用户
func Connect_uploadUser(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	opts, err := bind.NewTransactorWithChainID(reader, "123456",chainID)
	if err != nil {
		log.Fatal("NewTransactor", err)
	}
	opts.GasLimit = 3000000
	opts.GasPrice = gasPrice

	fmt.Println("opts:", opts)

	//上传用户
	tx, err := ins.UserAdduser(opts, user.User_)
	fmt.Println("上传用户:", tx)
	return err,chainID,gasPrice,ins
}

//设置征求者
func Connect1_uploadGainer(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	tx1, err := ins.GainerSetDoctor(opts, user.Soliciter_)
	fmt.Println("上传征求者:", tx1)
	return err, chainID, gasPrice, ins
}

//用户上传医疗信息
func Connect2_UploadMedicalInformation(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	//用户上传病历信息
	tx2, err := ins.UserAddMedicalInformation(opts, user.Proute,user.Soliciter_)
	fmt.Println("用户上传病历信息:", tx2)
	return err, chainID, gasPrice, ins
}

//征求者发布医疗信息
func Connect3_ReleaseMedicalInformation(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	//征求者发布病历信息
	tx3, err := ins.GainerAddMedicalInformation(opts,big.NewInt(user.Min_),big.NewInt(user.Max), user.MedicalName,user.MedicalNeed,user.RequirementDescription)
	fmt.Println("征求者发布病历信息:", tx3)
	return err, chainID, gasPrice, ins
}

//征求者审核与奖励
func Connect4_ReviewAndReward(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	//征求者审核与奖励
	tx4, err := ins.GainerWhether(opts,user.Soliciter_,user.Proute, user.Whether,big.NewInt(user.Ercnum_))
	fmt.Println("征求者审核与奖励:", tx4)
	return err, chainID, gasPrice, ins
}

//查询ETH
func Connect5_CheckTheBalance(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	//查询ETH
	tx5, err := ins.GetUserETH(nil,user.People_)
	fmt.Println("查询ETH:", tx5)
	return err, chainID, gasPrice, ins
}

//查询AS
func Connect6_CheckTheAS(user *model.User_AddMedicalInformation) (error,*big.Int,*big.Int,*gen.UploadMedicalrecords) {
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
	//查询ETH
	tx6, err := ins.BalanceOf(nil,user.People_)
	fmt.Println("查询AS:", tx6)
	return err, chainID, gasPrice, ins
}

