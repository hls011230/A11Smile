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
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	"time"
)

func UploadMedicalHistory(srcFile io.Reader, token model.RespWXToken, uid int , fileName string) error {

	// 获取用户的Address值
	DB := db.Get()
	var user model.Wallet
	DB.Table("users").First(&user,"id = ?",uid)

	// 获取文件上传地址
	path := fmt.Sprintf("a11smile/users/%v/MedicalHistory/%v.jpg",user.BlockAddress,fileName)
	myReq := struct {
		Env  string `json:"env"`
		Path string `json:"path"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		Path: path,
	}

	reqByte, err := json.Marshal(myReq)

	u := "https://api.weixin.qq.com/tcb/uploadfile?access_token=%s"

	req, err := http.NewRequest("POST", fmt.Sprintf(u, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	cli := http.Client{
		Timeout: 45 * time.Second,
	}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return err
	}

	// Json数据绑定
	var respUploadLink model.RespWXUploadLink
	err = json.NewDecoder(resp.Body).Decode(&respUploadLink)
	if err != nil {
		return err
	}

	fmt.Println(respUploadLink)
	// 在合约中存入用户病历信息
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(user.BlockAddress))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
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

	_, err = eth.Ins.UserUPMedicalinformation(auth,fileName,respUploadLink.FileId)
	if err != nil {
		return err
	}


	// 上传文件
	myUploadReq := struct {
		Key               string    `json:"key" form:"key"`
		Signature         string    `json:"Signature" form:"Signature"`
		XCosSecurityToken string    `json:"x_cos_security_token" form:"x_cos_security_token"`
		XCosMetaFileid    string    `json:"x_cos_meta_fileid" form:"x_cos_meta_fileid"`
		File              io.Reader `json:"file" form:"file"`
	}{
		Key:               path,
		Signature:         respUploadLink.Authorization,
		XCosSecurityToken: respUploadLink.Token,
		XCosMetaFileid:    respUploadLink.CosFileId,
		File:              srcFile,
	}

	//读出文本文件数据
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	content_type := w.FormDataContentType()


	_ = w.WriteField("key", myUploadReq.Key)
	_ = w.WriteField("Signature", myUploadReq.Signature)
	_ = w.WriteField("x-cos-security_token", myUploadReq.XCosSecurityToken)
	_ = w.WriteField("x-cos-meta-fileid", myUploadReq.XCosMetaFileid)

	//将文件数据写入
	formFile, _ := w.CreateFormFile("file", "new.jpg")
	_, err = io.Copy(formFile, myUploadReq.File)

	_ = w.Close()

	req, err = http.NewRequest("POST", respUploadLink.Url, buf)
	if err != nil {
		fmt.Println("req", err)
		return err
	}

	req.Header.Set("Content-Type", content_type)
	resp, _ = http.DefaultClient.Do(req)



	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return err
	}
	fmt.Println(string(b))

	return nil

}


func UploadMedicalExaminationReport(srcFile io.Reader, token model.RespWXToken, uid int , fileName string) error {

	// 获取用户的Address值
	DB := db.Get()
	var user model.Wallet
	DB.Table("users").First(&user,"id = ?",uid)

	// 获取文件上传地址
	path := fmt.Sprintf("a11smile/users/%v/MedicalExaminationReport/%v.jpg",user.BlockAddress,fileName)
	myReq := struct {
		Env  string `json:"env"`
		Path string `json:"path"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		Path: path,
	}

	reqByte, err := json.Marshal(myReq)

	u := "https://api.weixin.qq.com/tcb/uploadfile?access_token=%s"

	req, err := http.NewRequest("POST", fmt.Sprintf(u, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	cli := http.Client{
		Timeout: 45 * time.Second,
	}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	// Json数据绑定
	var respUploadLink model.RespWXUploadLink
	err = json.NewDecoder(resp.Body).Decode(&respUploadLink)
	if err != nil {
		return err
	}

	// 在合约中存入用户体检报告信息
	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(user.BlockAddress))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
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

	_, err = eth.Ins.UserUPMedicalExaminationReport(auth,fileName,respUploadLink.FileId)
	if err != nil {
		return err
	}


	// 上传文件
	myUploadReq := struct {
		Key               string    `json:"key" form:"key"`
		Signature         string    `json:"Signature" form:"Signature"`
		XCosSecurityToken string    `json:"x-cos-security_token" form:"x_cos_security_token"`
		XCosMetaFileid    string    `json:"x-cos-meta-fileid" form:"x_cos_meta_fileid"`
		File              io.Reader `json:"file" form:"file"`
	}{
		Key:               path,
		Signature:         respUploadLink.Authorization,
		XCosSecurityToken: respUploadLink.Token,
		XCosMetaFileid:    respUploadLink.CosFileId,
		File:              srcFile,
	}

	//读出文本文件数据

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	content_type := w.FormDataContentType()


	_ = w.WriteField("key", myUploadReq.Key)
	_ = w.WriteField("Signature", myUploadReq.Signature)
	_ = w.WriteField("x-cos-security-token", myUploadReq.XCosSecurityToken)
	_ = w.WriteField("x-cos-meta-fileid", myUploadReq.XCosMetaFileid)

	//将文件数据写入
	formFile, _ := w.CreateFormFile("file", "new.jpg")
	_, err = io.Copy(formFile, myUploadReq.File)

	_ = w.Close()

	req, err = http.NewRequest("POST", respUploadLink.Url, buf)
	if err != nil {
		fmt.Println("req", err)
		return err
	}

	req.Header.Set("Content-Type", content_type)
	resp, _ = http.DefaultClient.Do(req)



	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return err
	}
	fmt.Println(string(b))

	return nil

}


