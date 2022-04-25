package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
)

func ViewAllMedicalHistory(uid int) ([]string,error) {
	DB := db.Get()
	var user model.UserWallet
	DB.Table("users").First(&user,"id = ?",uid)

	var nameArray  []string
	// 从合约中获取用户的所有病历信息
	num, err := eth.Ins.ViewMedicalinformationCount(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)})
	if err != nil {
		return nameArray,err
	}

	numStr := num.String()
	count,_ := strconv.Atoi(numStr)

	for i := 0; i < count; i++ {
		medicalName, err := eth.Ins.UserMedicalinformationrName(nil,common.HexToAddress(user.BlockAddress),big.NewInt(int64(i)))
		if err != nil {
			return nameArray,err
		}
		nameArray = append(nameArray, medicalName)
	}

	return nameArray,err
}


func ViewAllMedicalExaminationReport(uid int) ([]string,error) {
	DB := db.Get()
	var user model.UserWallet
	DB.Table("users").First(&user,"id = ?",uid)

	var nameArray  []string
	// 从合约中获取用户的所有病历信息
	num, err := eth.Ins.ViewMedicalExaminationReportCount(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)})
	if err != nil {
		return nameArray,err
	}

	numStr := num.String()
	count,_ := strconv.Atoi(numStr)

	for i := 0; i < count; i++ {
		medicalName, err := eth.Ins.UserMedicalExaminationReportstrucrName(nil,common.HexToAddress(user.BlockAddress),big.NewInt(int64(i)))
		if err != nil {
			return nameArray,err
		}
		nameArray = append(nameArray, medicalName)
	}

	return nameArray,err
}

func PreviewMedicalHistory(uid int,fileName string) (string,error)  {
	// 获取对象钱包
	var user model.UserWallet
	DB := db.Get()
	DB.Table("users").First(&user,"id = ?",uid)

	// 查询指定病历文件的云地址
	fileCloudPath, err := eth.Ins.UserViewMedicalInformation(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)},fileName)
	if err != nil {
		return "",err
	}


	return fileCloudPath, nil
}


func PreviewMedicalExaminationReport(uid int,fileName string) (string,error)  {
	// 获取对象钱包
	var user model.UserWallet
	DB := db.Get()
	DB.Table("users").First(&user,"id = ?",uid)

	// 查询指定病历文件的云地址
	fileCloudPath, err := eth.Ins.UserViewMedicalExaminationReport(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)},fileName)
	if err != nil {
		return "",err
	}

	return fileCloudPath, nil
}