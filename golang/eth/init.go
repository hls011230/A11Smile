package eth

import (
	"A11Smile/eth/gen"
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	GasPrice *big.Int
	ChainID  *big.Int
	Ins      *gen.Gen
	AS       *gen.AS
	Client   *ethclient.Client
	AS       *gen.A11Smile
)

func Init(contract_address []string) error {
	client, err := ethclient.Dial("http://47.106.124.34:8545")
	if err != nil {
		return err
	}

	Client = client

	// 设置gas的价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	GasPrice = gasPrice

	// 设置chainId
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	ChainID = chainID

<<<<<<< Updated upstream
	// 连接合约
	cAddr := common.HexToAddress(contract_address[0])
	ins, err := gen.NewGen(cAddr, client) // 创建合约实例
=======
	// 连接合约UploadMedicalrecords
	cAddr := common.HexToAddress(contract_address)

	// 创建合约实例
	ins, err := gen.NewGen(cAddr, client)
>>>>>>> Stashed changes
	if err != nil {
		return err
	}
	Ins = ins

	//
	cAddr = common.HexToAddress(contract_address[1])
	as, err := gen.NewAS(cAddr, client)
	if err != nil {
		return err
	}
	AS = as

	return nil

}
