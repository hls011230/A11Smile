package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_eth/Test_register/depoly/db/model"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func VerifyIDCard(f string) error {

	// 小程序AppID和密钥
	appId := "wx1f6551a9b737be68"
	secret := "e8c9ec131ccbcc684b6cd0327404c869"

	// Get请求获取接口token
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	res, err := http.DefaultClient.Get(fmt.Sprintf(url, appId, secret))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	// Json绑定返回数据包
	var token model.RespWXSmall
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return err
	}

	// 执行身份证OCR识别(文件所在位置，小程序token)
	err = postIDCard(f, token)
	if err != nil {
		return err
	}

	return nil
}

func postIDCard(f string, token model.RespWXSmall) error {

	buf := new(bytes.Buffer)                    // 实例化一个结构体
	writer := multipart.NewWriter(buf)          // 返回一个writer指针
	ContentType := writer.FormDataContentType() // 表单提交的格式

	// 提供表单中的字段名<img>和文件名<new.jpg>,返回值是可写的接口io.Writer
	formFile, err1 := writer.CreateFormFile("img", "new.jpg")
	if err1 != nil {
		fmt.Println("创建form文件失败:", err1)
		return err1
	}

	// 从文件读取数据，写入表单
	srcFile, err := os.Open(f)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return err
	}

	defer srcFile.Close()
	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		fmt.Println("写入form文件失败", err)
		return err
	}

	writer.Close()

	// POST请求调用微信身份证OCR接口
	url := "https://api.weixin.qq.com/cv/ocr/idcard?type=photo&access_token=%s"
	res, err2 := http.Post(fmt.Sprintf(url, token.Access_token), ContentType, buf)
	if err2 != nil {
		fmt.Println("OCR接口请求失败:", err2)
		return err2
	}

	defer res.Body.Close()

	//Json数据绑定返回数据包
	var Front_IDCard model.RespWXIDOCRF
	err = json.NewDecoder(res.Body).Decode(&Front_IDCard)
	if err != nil {
		fmt.Println("数据绑定失败:", err)
		return err
	}
	return nil
}
