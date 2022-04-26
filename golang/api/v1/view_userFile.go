package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
	"strconv"
)

func ViewAllMedicalHistory(uid int) ([]string,error) {
	DB := db.Get()
	var user model.Wallet
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
	var user model.Wallet
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
	var user model.Wallet
	DB := db.Get()
	DB.Table("users").First(&user,"id = ?",uid)

	// 查询指定病历文件的云地址
	fileCloudPath, err := eth.Ins.UserViewMedicalInformation(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)},fileName)
	if err != nil {
		return "",err
	}

	// 根据云地址获取下载地址
	token,_ := GetToken()
	cloudFile := model.UserCloudLink{FileId: fileCloudPath,MaxAge: 720}
	fileList := []model.UserCloudLink{cloudFile}

	myReq := struct {
		Env  string `json:"env"`
		FileList []model.UserCloudLink `json:"file_list"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		FileList: fileList,
	}

	reqByte, err := json.Marshal(myReq)

	url := "https://api.weixin.qq.com/tcb/batchdownloadfile?access_token=%s"
	req, err := http.NewRequest("POST", fmt.Sprintf(url, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		return "",err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	// Json数据绑定
	var respWXLoadLink model.RespWXLoadLink
	err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
	if err != nil {
		return "",err
	}

	return respWXLoadLink.FileList[0].DownloadUrl, nil
}


func PreviewMedicalExaminationReport(uid int,fileName string) (string,error)  {
	// 获取对象钱包
	var user model.Wallet
	DB := db.Get()
	DB.Table("users").First(&user,"id = ?",uid)

	// 查询指定病历文件的云地址
	fileCloudPath, err := eth.Ins.UserViewMedicalExaminationReport(&bind.CallOpts{Context: context.Background(),From: common.HexToAddress(user.BlockAddress)},fileName)
	if err != nil {
		return "",err
	}

	// 根据云地址获取下载地址
	token,_ := GetToken()
	cloudFile := model.UserCloudLink{FileId: fileCloudPath,MaxAge: 720}
	fileList := []model.UserCloudLink{cloudFile}

	myReq := struct {
		Env  string `json:"env"`
		FileList []model.UserCloudLink `json:"file_list"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		FileList: fileList,
	}

	reqByte, err := json.Marshal(myReq)

	url := "https://api.weixin.qq.com/tcb/batchdownloadfile?access_token=%s"
	req, err := http.NewRequest("POST", fmt.Sprintf(url, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		return "",err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	// Json数据绑定
	var respWXLoadLink model.RespWXLoadLink
	err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
	if err != nil {
		return "",err
	}


	return respWXLoadLink.FileList[0].DownloadUrl, nil
}