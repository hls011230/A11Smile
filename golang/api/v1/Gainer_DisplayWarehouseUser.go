package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
)

func Gainer_DisplayWarehouseUser(gid int,warehouse *model.PostWarehouse)([]string, error) {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w,"id = ?",gid)

	reslen,err := eth.Ins.ViewWareHouseLength(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)},warehouse.Medical)
	if err != nil {
		return nil,err
	}
	number := reslen.String()//转成string
	num, err := strconv.Atoi(number)//string转int

	var Users []string
	for i := 0; i < num; i++ {
		resuser,err := eth.Ins.UserWareHouse(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(w.BlockAddress)},common.HexToAddress(w.BlockAddress),warehouse.Medical,big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		}
		Users = append(Users, fmt.Sprintf("0x%x",string(resuser[:])))
	}
	return Users,nil

}

